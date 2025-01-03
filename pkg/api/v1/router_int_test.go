package fleetdbapi_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/metal-automata/rivets/ginjwt"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"gopkg.in/go-jose/go-jose.v2"
	"gopkg.in/go-jose/go-jose.v2/jwt"

	"github.com/metal-automata/fleetdb/internal/dbtools"
	"github.com/metal-automata/fleetdb/internal/httpsrv"
	fleetdbapi "github.com/metal-automata/fleetdb/pkg/api/v1"
)

type integrationServer struct {
	h      http.Handler
	Client *fleetdbapi.Client
}

func serverTest(t *testing.T) *integrationServer {
	jwksURI := ginjwt.TestHelperJWKSProvider(ginjwt.TestPrivRSAKey1ID, ginjwt.TestPrivRSAKey2ID)

	db := dbtools.DatabaseTest(t)

	zcfg := zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zap.ErrorLevel),
	}

	l, _ := zcfg.Build()

	hs := httpsrv.Server{
		Logger:      l,
		DB:          db,
		OIDCEnabled: true,
		AuthConfigs: []ginjwt.AuthConfig{
			{
				Enabled:    true,
				Audience:   "hollow.test",
				Issuer:     "hollow.test.issuer",
				JWKSURI:    jwksURI,
				RolesClaim: "userPerms",
			},
		},
		SecretsKeeper: dbtools.TestSecretKeeper(t),
	}
	s := hs.NewServer()

	ts := &integrationServer{
		h: s.Handler,
	}

	c, err := fleetdbapi.NewClientWithToken("testToken", "http://test.hollow.com", ts)
	require.NoError(t, err)

	ts.Client = c

	return ts
}

func (s *integrationServer) Do(req *http.Request) (*http.Response, error) {
	// if the context is expired return the error, used for timeout tests
	if err := req.Context().Err(); err != nil {
		return nil, err
	}

	w := httptest.NewRecorder()
	s.h.ServeHTTP(w, req)

	return w.Result(), nil
}

func validToken(scopes []string) string {
	claims := jwt.Claims{
		Subject:   "test-user",
		Issuer:    "hollow.test.issuer",
		NotBefore: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)),
		Expiry:    jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Audience:  jwt.Audience{"hollow.test", "another.test.service"},
	}
	signer := ginjwt.TestHelperMustMakeSigner(jose.RS256, ginjwt.TestPrivRSAKey1ID, ginjwt.TestPrivRSAKey1)

	return ginjwt.TestHelperGetToken(signer, claims, "userPerms", scopes)
}
