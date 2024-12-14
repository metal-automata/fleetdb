package fleetdbapi

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/metal-automata/fleetdb/internal/models"
)

// TODO: reference server credentials for Username, Password attributes and drop those columns from server_bmcs table

type ServerBMC struct {
	ID                 uuid.UUID `json:"id"`
	ServerID           uuid.UUID `json:"server_id" binding:"required,uuid"` // Note: binding attributes should not have spaces
	HardwareVendorName string    `json:"hardware_vendor_name" binding:"required"`
	HardwareVendorID   string    `json:"-"`
	HardwareModelName  string    `json:"hardware_model_name" binding:"required"`
	HardwareModelID    string    `json:"-"`
	Username           string    `json:"username" binding:"required"`
	Password           string    `json:"password" binding:"required"`
	IPAddress          string    `json:"ipaddress" binding:"required,ip"`
	MacAddress         string    `json:"macaddress" binding:"mac"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (t *ServerBMC) fromDBModel(dbT *models.ServerBMC) {
	t.ID = uuid.MustParse(dbT.ID)
	t.ServerID = uuid.MustParse(dbT.ServerID)
	t.Username = dbT.Username
	t.IPAddress = dbT.IPAddress
	t.MacAddress = dbT.MacAddress.String
	t.CreatedAt = dbT.CreatedAt.Time
	t.UpdatedAt = dbT.UpdatedAt.Time

	if dbT.R != nil && dbT.R.HardwareModel != nil {
		t.HardwareModelName = dbT.R.HardwareModel.Name
	}
	if dbT.R != nil && dbT.R.HardwareVendor != nil {
		t.HardwareVendorName = dbT.R.HardwareVendor.Name
	}
}

func (t *ServerBMC) toDBModel() *models.ServerBMC {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}

	return &models.ServerBMC{
		ID:               t.ID.String(),
		ServerID:         t.ServerID.String(),
		HardwareVendorID: t.HardwareVendorID,
		HardwareModelID:  t.HardwareModelID,
		Username:         t.Username,
		IPAddress:        t.IPAddress,
		MacAddress:       null.StringFrom(t.MacAddress),
	}
}

func (r *Router) serverBMCCreate(c *gin.Context) {
	var t ServerBMC
	if err := c.ShouldBindJSON(&t); err != nil {
		badRequestResponse(c, "invalid ServerBMC payload", err)
		return
	}

	ctx := c.Request.Context()
	// component data is written in a transaction along with versioned attributes
	tx, errTxBegin := r.DB.BeginTx(ctx, nil)
	if errTxBegin != nil {
		dbErrorResponse2(c, "ServerBMC insert error", errTxBegin)
	}

	defer loggedRollback(r, tx)

	id, errInsert := r.insertServerBMCWithTx(ctx, tx, t)
	if errInsert != nil {
		dbErrorResponse2(c, "ServerBMC insert error", errInsert)
		return
	}

	if errTxCommit := tx.Commit(); errTxCommit != nil {
		dbErrorResponse2(c, "ServerBMC insert error", errInsert)
		return
	}

	createdResponse(c, id)
}

func (r *Router) insertServerBMCWithTx(ctx context.Context, tx *sql.Tx, serverBMC ServerBMC) (string, error) {
	var credentialvalues *serverCredentialValues
	if serverBMC.Password != "" {
		credentialvalues = &serverCredentialValues{
			Username: serverBMC.Username,
			Password: serverBMC.Password,
		}
	}

	return r.insertServerBMCAndCredentialsWithTx(
		ctx,
		tx,
		serverBMC.HardwareVendorName,
		serverBMC.HardwareModelName,
		serverBMC.ServerID,
		serverBMC.toDBModel(),
		credentialvalues,
	)
}

func (r *Router) insertServerBMCAndCredentialsWithTx(ctx context.Context, tx *sql.Tx, hwVendor, hwModel string, serverID uuid.UUID, bmc *models.ServerBMC, creds *serverCredentialValues) (string, error) {
	// identify hardware vendor id
	dbHardwareVendor, err := r.hardwareVendorBySlug(ctx, hwVendor)
	if err != nil {
		return "", err
	}

	bmc.HardwareVendorID = dbHardwareVendor.ID

	dbHm, err := r.hardwareModelBySlug(ctx, hwModel)
	if err != nil {
		return "", err
	}

	bmc.HardwareModelID = dbHm.ID

	if err := bmc.Insert(ctx, tx, boil.Infer()); err != nil {
		return "", err
	}

	if creds != nil {
		if err := r.serverCredentialUpsert(ctx, tx, "bmc", serverID, *creds); err != nil {
			return "", err
		}
	}

	return bmc.ID, nil
}

func (r *Router) serverBMCGet(c *gin.Context) {
	serverID := c.Param("uuid")

	serverUUID, err := uuid.Parse(serverID)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "valid server UUID expected"))
		return
	}

	mods := []qm.QueryMod{
		qm.Where("server_id=?", serverUUID.String()),
		// join hardware vendor
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.HardwareVendors,
				models.HardwareVendorTableColumns.ID,
				models.ServerBMCTableColumns.HardwareVendorID,
			),
		),
		// Load N-1 relationship in db model struct field R
		qm.Load(models.ServerBMCRels.HardwareVendor),

		// join hardware model
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.HardwareModels,
				models.HardwareModelTableColumns.ID,
				models.ServerBMCTableColumns.HardwareModelID,
			),
		),
		// Load N-1 relationship in db model struct field R
		qm.Load(models.ServerBMCRels.HardwareModel),
	}

	dbServerBMC, err := models.ServerBMCS(mods...).One(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	serverBMC := &ServerBMC{}
	serverBMC.fromDBModel(dbServerBMC)
	itemResponse(c, serverBMC)
}

func (r *Router) serverBMCDelete(c *gin.Context) {
	serverID := c.Param("uuid")
	serverUUID, err := uuid.Parse(serverID)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "valid server UUID expected"))
		return
	}

	mod := qm.Where("server_id=?", serverUUID)

	_, err = models.ServerBMCS(mod).DeleteAll(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse2(c, "server BMC delete error", err)
		return
	}

	deletedResponse(c)
}
