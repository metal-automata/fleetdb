package fleetdbapi

import (
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

// https://www.postgresql.org/docs/current/errcodes-appendix.html
const pgUniqueViolationErrorCode = "23505"

var (
	errBadRequest = errors.New("bad request")
)

// ServerResponse represents the data that the server will return on any given call
type ServerResponse struct {
	PageSize         int                 `json:"page_size,omitempty"`
	Page             int                 `json:"page,omitempty"`
	PageCount        int                 `json:"page_count,omitempty"`
	TotalPages       int                 `json:"total_pages,omitempty"`
	TotalRecordCount int64               `json:"total_record_count,omitempty"`
	Links            ServerResponseLinks `json:"_links,omitempty"`
	Message          string              `json:"message,omitempty"`
	Error            string              `json:"error,omitempty"`
	Slug             string              `json:"slug,omitempty"`
	Data             interface{}         `json:"data,omitempty"` // data is not a DB record/records, but is structured
	Record           interface{}         `json:"record,omitempty"`
	Records          interface{}         `json:"records,omitempty"`
}

// ServerResponseLinks represent links that could be returned on a page
type ServerResponseLinks struct {
	Self     *Link `json:"self,omitempty"`
	First    *Link `json:"first,omitempty"`
	Previous *Link `json:"previous,omitempty"`
	Next     *Link `json:"next,omitempty"`
	Last     *Link `json:"last,omitempty"`
}

// Link represents an address to a page
type Link struct {
	Href string `json:"href,omitempty"`
}

// HasNextPage will return if there are additional resources to load on additional pages
func (r *ServerResponse) HasNextPage() bool {
	return r.Records != nil && r.Links.Next != nil
}

// notFoundResponse writes a 404 response with the given message
func notFoundResponse(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, &ServerResponse{Message: message})
}

func badRequestResponse(c *gin.Context, message string, err error) {
	if err == nil {
		err = errBadRequest
	}

	c.JSON(http.StatusBadRequest, &ServerResponse{Message: message, Error: err.Error()})
}

func createdResponse(c *gin.Context, slug string) {
	uri := fmt.Sprintf("%s/%s", uriWithoutQueryParams(c), slug)
	r := &ServerResponse{
		Message: "resource created",
		Slug:    slug,
		Links: ServerResponseLinks{
			Self: &Link{Href: uri},
		},
	}

	c.Header("Location", uri)
	c.JSON(http.StatusCreated, r)
}

// DEPRECATED; Replace with deletedResponse2
func deletedResponse(c *gin.Context) {
	c.JSON(http.StatusOK, &ServerResponse{Message: "resource deleted"})
}

func deletedResponse2(c *gin.Context, count int64) {
	if count <= 0 {
		c.JSON(http.StatusNotFound, &ServerResponse{Message: "resource not found", Error: "Unable to delete resource", Slug: "0"})
	} else {
		c.JSON(http.StatusOK, &ServerResponse{Message: "resource deleted", Slug: fmt.Sprintf("%d", count)})
	}
}

func updatedResponse(c *gin.Context, slug string) {
	r := &ServerResponse{
		Message: "resource updated",
		Slug:    slug,
		Links: ServerResponseLinks{
			Self: &Link{Href: uriWithoutQueryParams(c)},
		},
	}

	c.JSON(http.StatusOK, r)
}

// DEPRECATED; Replace with deletedResponse2
func dbErrorResponse(c *gin.Context, err error) {
	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		badRequestResponse(c, "", err)
		return
	}

	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, &ServerResponse{Message: "resource not found", Error: err.Error()})
	} else {
		c.JSON(http.StatusInternalServerError, &ServerResponse{Message: "datastore error", Error: err.Error()})
	}
}

func dbErrorResponse2(c *gin.Context, message string, err error) {
	if pgErr, ok := err.(*pq.Error); ok { // nolint:errorlint // TODO fixme - use errors.As
		if pgErr.Code == pgUniqueViolationErrorCode { // Unique violation error code in PostgreSQL
			err = errors.Wrapf(err, "duplicate key value violates unique constraint %s: %s", pgErr.Constraint, pgErr.Detail)
			badRequestResponse(c, pgErr.Detail, err)
			return
		}

		// TODO: fix this to log or return the error detail
		badRequestResponse(c, pgErr.Detail, err)
		return
	}

	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, &ServerResponse{Message: fmt.Sprintf("resource not found; %s", message), Error: err.Error()})
	} else {
		c.JSON(http.StatusInternalServerError, &ServerResponse{Message: fmt.Sprintf("datastore error; %s", message), Error: err.Error()})
	}
}

func failedConvertingToVersioned(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, &ServerResponse{Message: "failed parsing the datastore results", Error: err.Error()})
}

func listResponse(c *gin.Context, i interface{}, p paginationData) {
	uri := c.Request.URL

	r := &ServerResponse{
		PageSize:  p.pager.limitUsed(),
		PageCount: p.pageCount,
		Records:   i,
		Links: ServerResponseLinks{
			Self: &Link{Href: uri.String()},
		},
	}

	d := float64(p.totalCount) / float64(p.pager.limitUsed())
	r.TotalPages = int(math.Ceil(d))
	r.Page = p.pager.Page
	r.TotalRecordCount = p.totalCount

	r.Links.First = &Link{Href: getURIWithQuerySet(*uri, "page", "1")}
	r.Links.Last = &Link{Href: getURIWithQuerySet(*uri, "page", strconv.Itoa(r.TotalPages))}

	if r.Page < r.TotalPages {
		r.Links.Next = &Link{Href: getURIWithQuerySet(*uri, "page", strconv.Itoa(r.Page+1))}
	}

	if r.Page != 1 {
		r.Links.Previous = &Link{Href: getURIWithQuerySet(*uri, "page", strconv.Itoa(r.Page-1))}
	}

	c.JSON(http.StatusOK, r)
}

func itemResponse(c *gin.Context, i interface{}) {
	r := &ServerResponse{
		Message: "resource retrieved",
		Record:  i,
		Links: ServerResponseLinks{
			Self: &Link{Href: c.Request.URL.String()},
		},
	}
	c.JSON(http.StatusOK, r)
}

func getURIWithQuerySet(uri url.URL, key, value string) string {
	q := uri.Query()
	q.Del(key)
	q.Add(key, value)
	uri.RawQuery = q.Encode()

	return uri.String()
}

func uriWithoutQueryParams(c *gin.Context) string {
	uri := c.Request.URL
	uri.RawQuery = ""

	return uri.String()
}
