package fleetdbapi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/metal-automata/fleetdb/internal/dbtools"
	"github.com/metal-automata/fleetdb/internal/models"
)

// Server represents a server in a facility
type Server struct {
	UUID         uuid.UUID          `json:"uuid"`
	Name         string             `json:"name"`
	FacilityCode string             `json:"facility" binding:"required"`
	Vendor       string             `json:"vendor" binding:"required"`
	Model        string             `json:"model"`
	Serial       string             `json:"serial"`
	BMC          *ServerBMC         `json:"bmc,omitempty"`
	Components   []*ServerComponent `json:"components"`
	Status       *ServerStatus      `json:"status"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	// DeletedAt is a pointer to a Time in order to be able to support checks for nil time
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// InventoryRefreshedAt indicates the last time the server inventory was collected
	InventoryRefreshedAt time.Time `json:"inventory_refreshed_at"`
}

func (s *Server) fromDBModel(dbS *models.Server) error {
	var err error

	s.UUID, err = uuid.Parse(dbS.ID)
	if err != nil {
		return err
	}

	s.Name = dbS.Name.String
	s.FacilityCode = dbS.FacilityCode.String
	s.Vendor = dbS.VendorID.String
	s.CreatedAt = dbS.CreatedAt.Time
	s.UpdatedAt = dbS.UpdatedAt.Time

	if dbS.R != nil {
		if dbS.R.ServerBMC != nil {
			bmc := ServerBMC{}
			bmc.fromDBModel(dbS.R.ServerBMC)
			s.BMC = &bmc
		}

		if dbS.R.ServerStatus != nil {
			status := ServerStatus{}
			status.fromDBModel(dbS.R.ServerStatus)
			s.Status = &status
		}

		if dbS.R.ServerComponents != nil {
			scl := ServerComponentSlice{}
			scl.fromDBModel(dbS.R.ServerComponents)
			s.Components = make([]*ServerComponent, 0, len(scl))
			for _, component := range scl {
				s.Components = append(s.Components, component)
			}
		}
	}

	if !dbS.DeletedAt.IsZero() {
		s.DeletedAt = &dbS.DeletedAt.Time
	}

	if !dbS.InventoryRefreshedAt.IsZero() {
		s.InventoryRefreshedAt = dbS.InventoryRefreshedAt.Time
	}

	return nil
}

func (s *Server) toDBModel(vendorID, modelID string) (*models.Server, error) {
	dbS := &models.Server{
		Name:         null.StringFrom(s.Name),
		FacilityCode: null.StringFrom(s.FacilityCode),
		SerialNumber: null.StringFrom(s.Serial),
	}

	if vendorID != "" {
		id, err := uuid.Parse(vendorID)
		if err != nil {
			return nil, errors.Wrap(err, "vendor ID")
		}

		dbS.VendorID = null.StringFrom(id.String())
	}

	if modelID != "" {
		id, err := uuid.Parse(modelID)
		if err != nil {
			return nil, errors.Wrap(err, "model ID")
		}

		dbS.ModelID = null.StringFrom(id.String())
	}

	if s.UUID.String() != uuid.Nil.String() {
		dbS.ID = s.UUID.String()
	}

	return dbS, nil
}

func (r *Router) serverGet(c *gin.Context) {
	serverUUID, err := r.parseUUID(c.Param("uuid"))
	if err != nil {
		badRequestResponse(c, "invalid server UUID", err)
		return
	}

	params := &ServerGetParams{}
	params.decode(c.Request.URL.Query())
	// Bind basic query parameters
	mods := []qm.QueryMod{}
	mods = append(mods, params.queryMods(serverUUID.String())...)

	dbSrv, err := models.Servers(mods...).One(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	var srv Server
	if err = srv.fromDBModel(dbSrv); err != nil {
		failedConvertingToVersioned(c, err)
		return
	}

	// TODO: reference the BMC username, password in the server credentials table
	// decrypt the BMC password
	if dbSrv.R != nil && srv.BMC != nil {
		// if we have more than one credential returned - the query mod is incorrect
		if len(dbSrv.R.ServerCredentials) > 0 {
			value, err := dbtools.Decrypt(c.Request.Context(), r.SecretsKeeper, dbSrv.R.ServerCredentials[0].Password)
			if err != nil {
				c.JSON(http.StatusInternalServerError, &ServerResponse{Message: "error decrypting value", Error: err.Error()})
				return
			}
			srv.BMC.Password = value
		}
	}

	itemResponse(c, srv)
}

func (r *Router) serverCreate(c *gin.Context) {
	var srv Server

	if err := c.ShouldBindJSON(&srv); err != nil {
		badRequestResponse(c, "invalid server", err)
		return
	}

	ctx := c.Request.Context()

	hwVendor, err := r.hardwareVendorBySlug(ctx, srv.Vendor)
	if err != nil {
		dbErrorResponse2(c, "hardware vendor query error: "+srv.Vendor, err)
		return
	}

	// hw model is optional at server create, its populated at inventory collection
	var hwModelID string
	if srv.Model != "" {
		hwModel, errHwModel := r.hardwareModelBySlug(ctx, srv.Model)
		if errHwModel != nil {
			dbErrorResponse2(c, "hardware model query error: "+srv.Model, errHwModel)
			return
		}

		hwModelID = hwModel.ID
	}

	dbSRV, err := srv.toDBModel(hwVendor.ID, hwModelID)
	if err != nil {
		failedConvertingToVersioned(c, errors.Wrap(err, "invalid server"))
		return
	}

	// tx
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		dbErrorResponse2(c, "server insert error", err)
		return
	}

	defer loggedRollback(r, tx)

	if err := dbSRV.Insert(c.Request.Context(), tx, boil.Infer()); err != nil {
		dbErrorResponse2(c, "", err)
		return
	}

	if srv.BMC != nil {
		srv.BMC.ServerID = uuid.MustParse(dbSRV.ID)
		dbBMC := srv.BMC.toDBModel()

		if _, err := r.insertServerBMC(ctx, tx, srv.Vendor, srv.Model, dbBMC); err != nil {
			dbErrorResponse2(c, "server insert error", err)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		dbErrorResponse(c, err)
		return
	}

	createdResponse(c, dbSRV.ID)
}

func (r *Router) serverDelete(c *gin.Context) {
	dbSRV, err := r.loadServerFromParams(c.Request.Context(), c.Param("uuid"))
	if err != nil {
		r.Logger.Error(fmt.Sprintf("failed to load server %v, err %v", dbSRV, err))

		if errors.Is(err, ErrUUIDParse) {
			badRequestResponse(c, "", err)
			return
		}

		dbErrorResponse(c, err)

		return
	}

	if _, err = dbSRV.Delete(c.Request.Context(), r.DB, false); err != nil {
		r.Logger.Error(fmt.Sprintf("failed to delete server %v, err %v", dbSRV.ID, err))
		dbErrorResponse(c, err)

		return
	}

	deletedResponse(c)
}

func (r *Router) serverUpdate(c *gin.Context) {
	srv, err := r.loadServerFromParams(c.Request.Context(), c.Param("uuid"))
	if err != nil {
		if errors.Is(err, ErrUUIDParse) {
			badRequestResponse(c, "", err)
			return
		}

		dbErrorResponse(c, err)

		return
	}

	var newValues Server
	if err := c.ShouldBindJSON(&newValues); err != nil {
		badRequestResponse(c, "invalid server", err)
		return
	}

	srv.Name = null.StringFrom(newValues.Name)
	srv.FacilityCode = null.StringFrom(newValues.FacilityCode)

	cols := boil.Infer()

	if _, err := srv.Update(c.Request.Context(), r.DB, cols); err != nil {
		dbErrorResponse(c, err)
		return
	}

	updatedResponse(c, srv.ID)
}

func (r *Router) serverInventoryRefreshed(ctx context.Context, tx boil.ContextExecutor, server *models.Server) error {
	server.InventoryRefreshedAt = null.TimeFrom(time.Now())
	_, err := server.Update(ctx, tx, boil.Whitelist(models.ServerColumns.InventoryRefreshedAt))
	return err
}

func (r *Router) loadServerFromParams(ctx context.Context, serverID string) (*models.Server, error) {
	u, err := r.parseUUID(serverID)
	if err != nil {
		return nil, errors.Wrap(ErrUUIDParse, err.Error())
	}

	srv, err := models.FindServer(ctx, r.DB, u.String())
	if err != nil {
		return nil, err
	}

	return srv, nil
}
