package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mlhan1993/KongInterview/internal/service"
	"github.com/mlhan1993/KongInterview/internal/utils"
	"github.com/mlhan1993/KongInterview/pkg/errors"
)

type GetServicesRequest struct {
	SortOrder  string `json:"sortOrder,omitempty"`
	Filter     string `json:"filter,omitempty"`
	NumPerPage uint   `json:"numPerPage,omitempty"`
	PageNumber uint   `json:"pageNumber,omitempty"`
}

func (r *GetServicesRequest) Validate() error {
	if r.SortOrder != "" && r.SortOrder != "asc" && r.SortOrder != "desc" {
		return errors.NewBadRequestError("invalid sortOrder, sortOrder must be either asc or desc")
	}
	if (r.NumPerPage == 0 && r.PageNumber != 0) || (r.NumPerPage != 0 && r.PageNumber == 0) {
		a, b := "pageNumber", "numPerPage"
		if r.PageNumber == 0 {
			a, b = b, a
		}
		return errors.NewBadRequestError(fmt.Sprintf("numPerPage and pageNumber must be used together. cannot non-zero %s with zero %s", a, b))
	}
	return nil
}

type GetServicesResponse struct {
	Total    uint              `json:"total"`
	Services []service.Service `json:"services"`
}

type GetVersionsRequest struct {
	ServiceID  uint   `json:"serviceID"`
	SortOrder  string `json:"sortOrder,omitempty"`
	NumPerPage uint   `json:"numPerPage,omitempty"`
	PageNumber uint   `json:"pageNumber,omitempty"`
}

func (r *GetVersionsRequest) Validate() error {
	if r.SortOrder != "" && r.SortOrder != "asc" && r.SortOrder != "desc" {
		return errors.NewBadRequestError("invalid sortOrder, sortOrder must be either asc or desc")
	}

	if r.ServiceID == 0 {
		return errors.NewBadRequestError("invalid serviceID, serviceID must be a non-empty positive integer")
	}

	if (r.NumPerPage == 0 && r.PageNumber != 0) || (r.NumPerPage != 0 && r.PageNumber == 0) {
		a, b := "pageNumber", "numPerPage"
		if r.PageNumber == 0 {
			a, b = b, a
		}
		return errors.NewBadRequestError(fmt.Sprintf("numPerPage and pageNumber must be used together. cannot non-zero %s with zero %s", a, b))
	}

	return nil
}

type GetVersionResponse struct {
	Total    uint              `json:"total"`
	Versions []service.Version `json:"versions"`
}

type KongDB interface {
	GetServices(ctx context.Context, numPerPage, pageNumber uint, sortOrder, filter string) (uint, []service.Service, error)
	GetVersions(ctx context.Context, serviceId, numPerPage, pageNumber uint, sortOrder string) (uint, []service.Version, error)
}

type V1 struct {
	db KongDB
}

func NewV1(db KongDB) *V1 {
	return &V1{db: db}
}

func (h *V1) PostRetrieveServices(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body
	var req GetServicesRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ResponseFromError(errors.NewBadRequestError("invalid request format"), w)
		utils.LogRequestError(err, r)
		return
	}

	err = req.Validate()
	if err != nil {
		utils.ResponseFromError(err, w)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()
	totalServices, services, err := h.db.GetServices(ctx, req.NumPerPage, req.PageNumber, req.SortOrder, req.Filter)
	if err != nil {
		utils.ResponseFromError(err, w)
		utils.LogRequestError(err, r)
		return
	}

	// Send the response as JSON
	res := GetServicesResponse{
		Total:    totalServices,
		Services: services,
	}
	json.NewEncoder(w).Encode(res)
}

func (h *V1) PostRetrieveVersions(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body
	var req GetVersionsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ResponseFromError(errors.NewBadRequestError("invalid request format"), w)
		utils.LogRequestError(err, r)
		return
	}

	err = req.Validate()
	if err != nil {
		utils.ResponseFromError(err, w)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()
	total, versions, err := h.db.GetVersions(ctx, req.ServiceID, req.NumPerPage, req.PageNumber, req.SortOrder)
	if err != nil {
		utils.ResponseFromError(err, w)
		utils.LogRequestError(err, r)
		return
	}

	// Send the response as JSON
	res := GetVersionResponse{
		Total:    total,
		Versions: versions,
	}
	json.NewEncoder(w).Encode(res)
}
