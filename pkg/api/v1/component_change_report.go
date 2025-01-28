package fleetdbapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/metal-automata/fleetdb/internal/models"
)

var (
	errComponentChangeReportPayload = errors.New("server component report payload error")
)

// ComponentChangeReport represents addition/removal of server components to or from an existing set.
//
// An inventory client submits component changes when a server has existing component records,
// the changes submitted here are stored in the component_change_review table and once reviewed by an operator
// or automation, they can be merged.
type ComponentChangeReport struct {
	CollectionMethod string             `json:"collection_method" binding:"required"` // inband OR outofband
	Creates          []*ServerComponent `json:"creates,omitempty"`                    // Components to be added
	Deletes          []*ServerComponent `json:"deletes,omitempty"`                    // Components to be removed
}

type ComponentChangeReportResponse struct {
	ReportID        string   // identifies all changes part of a report
	ChangeIDCreates []string // identifiers for the changes that adds components
	ChangeIDDeletes []string // identifiers for the changes that removes components
}

func (t *ComponentChangeReport) toDBModel(slugMap map[string]string, collectionMethod, serverID, reportID string) (creates, deletes []*models.ComponentChangeReport, err error) {
	// uniq component serials
	serials := make(map[string]bool)

	errAdd := errors.New("error in component add change")
	for _, add := range t.Creates {
		key := ComponentKey(add.Name, add.Serial)
		if serials[key] {
			return nil, nil, errors.Wrap(errAdd, fmt.Sprintf("duplicate component, name: %s, serial: %s", add.Name, add.Serial))
		}

		serials[key] = true

		if serverID != add.ServerUUID.String() {
			return nil, nil, errors.Wrap(errAdd, fmt.Sprintf("serverID mismatch, name: %s, serial: %s", add.Name, add.Serial))
		}

		add.ServerComponentTypeID = slugMap[add.Name]
		data, err := json.Marshal(add)
		if err != nil {
			return nil, nil, errors.Wrap(errAdd, fmt.Sprintf("name: %s, serial: %s", add.Name, add.Serial))
		}

		if slugMap[add.Name] == "" {
			return nil, nil, errors.Wrap(errAdd, fmt.Sprintf("unknown component slug, name: %s", add.Name))
		}

		creates = append(creates, &models.ComponentChangeReport{
			CollectionMethod:      collectionMethod,
			ReportID:              reportID,
			ServerID:              add.ServerUUID.String(),
			ServerComponentTypeID: slugMap[add.Name],
			Serial:                add.Serial,
			ServerComponentName:   add.Name,
			Data:                  data,
		})
	}

	errRemove := errors.New("error in component remove change")
	for _, remove := range t.Deletes {
		key := ComponentKey(remove.Name, remove.Serial)
		if serials[key] {
			return nil, nil, errors.Wrap(errAdd, fmt.Sprintf("duplicate component, name: %s, serial: %s", remove.Name, remove.Serial))
		}

		serials[key] = true

		if remove.UUID == uuid.Nil {
			return nil, nil, errors.Wrap(errRemove, "nil component ID")
		}

		if serverID != remove.ServerUUID.String() {
			return nil, nil, errors.Wrap(errRemove, fmt.Sprintf("serverID mismatch, name: %s, serial: %s", remove.Name, remove.Serial))
		}

		remove.ServerComponentTypeID = slugMap[remove.Name]
		data, err := json.Marshal(remove)
		if err != nil {
			return nil, nil, errors.Wrap(err, fmt.Sprintf("name: %s, serial: %s", remove.Name, remove.Serial))
		}

		if slugMap[remove.Name] == "" {
			return nil, nil, errors.Wrap(errRemove, fmt.Sprintf("unknown component slug, name: %s", remove.Name))
		}

		remove.ServerComponentTypeID = slugMap[remove.Name]
		deletes = append(deletes, &models.ComponentChangeReport{
			CollectionMethod:      collectionMethod,
			ReportID:              reportID,
			ServerID:              remove.ServerUUID.String(),
			ServerComponentID:     null.StringFrom(remove.UUID.String()),
			ServerComponentTypeID: slugMap[remove.Name],
			RemoveComponent:       null.BoolFrom(true),
			Serial:                remove.Serial,
			ServerComponentName:   remove.Name,
			Data:                  data,
		})
	}

	return creates, deletes, nil
}

// upsert change report and return change IDs
func (r *Router) upsertComponentChangeReports(ctx context.Context, tx boil.ContextExecutor, changes []*models.ComponentChangeReport) ([]string, error) {
	changeIDs := make([]string, 0, len(changes))
	for _, change := range changes {
		if err := change.Upsert(
			ctx,
			tx,
			true, // update on conflict
			// conflict columns
			[]string{
				models.ComponentChangeReportColumns.ServerID,
				models.ComponentChangeReportColumns.Serial,
				models.ComponentChangeReportColumns.ServerComponentName,
				models.ComponentChangeReportColumns.RemoveComponent,
			},
			// Columns to update when its an UPDATE
			boil.Whitelist(
				models.ComponentChangeReportColumns.ReportID,
				models.ComponentChangeReportColumns.RemoveComponent,
				models.ComponentChangeReportColumns.Data,
				models.ComponentChangeReportColumns.ServerComponentID,
			),
			// Columns to insert when its an INSERT
			boil.Infer(),
		); err != nil {
			return nil, err
		}

		changeIDs = append(changeIDs, change.ID)
	}

	return changeIDs, nil
}

// componentChangeReport sets up component inventory changes for review
func (r *Router) componentChangeReport(c *gin.Context) {
	// load server based on the UUID parameter
	server, err := r.loadServerFromParams(c.Request.Context(), c.Param("uuid"))
	if err != nil {
		if errors.Is(err, ErrUUIDParse) {
			badRequestResponse(c, "", err)
			return
		}

		dbErrorResponse2(c, "server query error", err)
		return
	}

	// components payload
	var incoming ComponentChangeReport
	if errBind := c.ShouldBindJSON(&incoming); errBind != nil {
		badRequestResponse(
			c,
			"",
			errors.Wrap(
				errComponentChangeReportPayload, errBind.Error()),
		)
		return
	}

	// validate
	if len(incoming.Creates) == 0 && len(incoming.Deletes) == 0 {
		badRequestResponse(
			c,
			"",
			errors.Wrap(errComponentChangeReportPayload, "expected components to Creates/Deletes"),
		)
		return
	}

	if !slices.Contains(collectionMethods, incoming.CollectionMethod) {
		badRequestResponse(
			c,
			"",
			errors.Wrap(errComponentChangeReportPayload, "unexpected CollectionMethod: "+incoming.CollectionMethod),
		)
		return
	}

	slugMap, err := r.serverComponentTypeSlugMap(c.Request.Context())
	if err != nil {
		dbErrorResponse2(c, "server component types query", err)
		return
	}

	// reportID associates all changes part of this request together for lookups
	reportID := uuid.New()

	// prepare change reports
	creates, deletes, err := incoming.toDBModel(slugMap, incoming.CollectionMethod, server.ID, reportID.String())
	if err != nil {
		badRequestResponse(
			c,
			"",
			err,
		)
		return
	}

	ctx := c.Request.Context()
	// component data is written in a transaction along with versioned attributes
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	defer loggedRollback(r, tx)

	// create record for server component adds
	changeIDCreates, err := r.upsertComponentChangeReports(ctx, tx, creates)
	if err != nil {
		dbErrorResponse(c, errors.Wrap(err, "error creating component change review records - create records"))
		return
	}

	// create record for server component deletes
	changeIDDeletes, err := r.upsertComponentChangeReports(ctx, tx, deletes)
	if err != nil {
		dbErrorResponse(c, errors.Wrap(err, "error creating component change review records - removal records"))
		return
	}

	if err := tx.Commit(); err != nil {
		dbErrorResponse(c, err)
		return
	}

	response := &ServerResponse{
		Message: "resource created",
		Data: &ComponentChangeReportResponse{
			ReportID:        reportID.String(),
			ChangeIDCreates: changeIDCreates,
			ChangeIDDeletes: changeIDDeletes,
		},
		Links: ServerResponseLinks{
			Self: &Link{Href: c.Request.URL.String()},
		},
	}

	c.JSON(http.StatusOK, response)
}

// ComponentChangeAccept is the payload sent by the client to merge component add/deletes records
type ComponentChangeAccept struct {
	ChangeIDs []string `json:"change_ids" binding:"required"` // ChangeIDs for which component changes are to be merged
}

// componentChangeAccept merges the changes referenced by the changeID into the server_components table and its relations
//
// nolint:gocyclo // TODO: split method
func (r *Router) componentChangeAccept(c *gin.Context) {
	// load server based on the UUID parameter
	server, err := r.loadServerFromParams(c.Request.Context(), c.Param("uuid"))
	if err != nil {
		if errors.Is(err, ErrUUIDParse) {
			badRequestResponse(c, "", err)
			return
		}

		dbErrorResponse(c, err)
		return
	}

	var incoming ComponentChangeAccept
	if errBind := c.ShouldBindJSON(&incoming); errBind != nil {
		badRequestResponse(
			c,
			"",
			errors.Wrap(
				errComponentChangeReportPayload, errBind.Error()),
		)

		return
	}

	if len(incoming.ChangeIDs) == 0 {
		badRequestResponse(
			c,
			"",
			errors.Wrap(
				errComponentChangeReportPayload, "mergeIDs empty"),
		)
		return
	}

	// https://github.com/volatiletech/sqlboiler/issues/227#issuecomment-348053252
	convertedIDs := make([]interface{}, 0, len(incoming.ChangeIDs))
	for _, id := range incoming.ChangeIDs {
		convertedIDs = append(convertedIDs, id)
	}

	mods := []qm.QueryMod{
		qm.WhereIn(
			fmt.Sprintf(
				"%s in ?",
				models.ComponentChangeReportColumns.ID,
			),
			convertedIDs...,
		),
	}

	ctx := c.Request.Context()
	changeReviews, err := models.ComponentChangeReports(mods...).All(ctx, r.DB)
	if err != nil {
		dbErrorResponse2(c, "current records query error", err)
		return
	}

	if len(changeReviews) == 0 {
		badRequestResponse(
			c,
			"",
			errors.Wrap(
				errComponentChangeReportPayload,
				"no changes identified for review",
			),
		)
		return
	}

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		dbErrorResponse2(c, "tx begin error", err)
		return
	}

	defer loggedRollback(r, tx)

	//	models.ServerComponent
	for _, changeReview := range changeReviews {
		if changeReview.RemoveComponent.Bool {
			// remove component
			mod := qm.Where(models.ServerComponentColumns.ID+"= ?", changeReview.ServerComponentID)
			if _, err := models.ServerComponents(mod).DeleteAll(c.Request.Context(), r.DB); err != nil {
				dbErrorResponse(c, err)
				return
			}
		} else {
			// add component
			addComponent := &ServerComponent{}
			if err := json.Unmarshal(changeReview.Data, addComponent); err != nil {
				badRequestResponse(
					c,
					"",
					errors.Wrap(
						errComponentChangeReportPayload,
						err.Error()+" unable to unmarshal component data from change record",
					),
				)
				return
			}

			addComponent.ServerComponentTypeID = changeReview.ServerComponentTypeID
			if err := r.componentAndRelationsUpsert(ctx, tx, changeReview.ServerID, *addComponent); err != nil {
				dbErrorResponse(c, err)
				return
			}
		}

		// purge change report
		if _, err := changeReview.Delete(ctx, tx); err != nil {
			dbErrorResponse(c, err)
			return
		}
	}

	// mark server inventory refreshed
	if err := r.serverInventoryRefreshed(ctx, tx, server); err != nil {
		dbErrorResponse(c, err)
		return
	}

	if err := tx.Commit(); err != nil {
		dbErrorResponse(c, err)
		return
	}

	createdResponse(c, "changes merged")
}

// serverComponentChangeDeleteAll deletes all server component change reviews
func (r *Router) componentChangeReportDeleteAll(c *gin.Context) {
	// load server based on the UUID parameter
	server, err := r.loadServerFromParams(c.Request.Context(), c.Param("uuid"))
	if err != nil {
		if errors.Is(err, ErrUUIDParse) {
			badRequestResponse(c, "", err)
			return
		}

		dbErrorResponse(c, err)

		return
	}

	if _, err := server.ComponentChangeReports().DeleteAll(c.Request.Context(), r.DB); err != nil {
		dbErrorResponse(c, err)

		return
	}

	deletedResponse(c)
}
