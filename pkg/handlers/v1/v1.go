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

type GetServiceOverviewRequest struct {
	SortOrder  string `json:"sortOrder,omitempty"`
	Filter     string `json:"filter,omitempty"`
	NumPerPage uint   `json:"numPerPage,omitempty"`
	PageNumber uint   `json:"pageNumber,omitempty"`
}

func (r *GetServiceOverviewRequest) Validate() error {
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

type GetServiceOverviewResponse struct {
	Total            uint               `json:"total"`
	ServiceOverviews []service.Overview `json:"serviceOverviews"`
}

type GetServiceDetailsRequest struct {
	ServiceID  uint   `json:"serviceID"`
	SortOrder  string `json:"sortOrder,omitempty"`
	NumPerPage uint   `json:"numPerPage,omitempty"`
	PageNumber uint   `json:"pageNumber,omitempty"`
}

func (r *GetServiceDetailsRequest) Validate() error {
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

type GetServiceDetailsResponse struct {
	Total          uint             `json:"total"`
	ServiceDetails []service.Detail `json:"serviceDetails"`
}

type ServiceDB interface {
	GetServiceOverview(ctx context.Context, numPerPage, pageNumber uint, sortOrder, filter string) (uint, []service.Overview, error)
	GetServiceDetails(ctx context.Context, serviceId, numPerPage, pageNumber uint, sortOrder string) (uint, []service.Detail, error)
}

type V1 struct {
	db ServiceDB
}

func NewV1(db ServiceDB) *V1 {
	return &V1{db: db}
}

func (h *V1) PostRetrieveServiceOverview(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body
	var req GetServiceOverviewRequest
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
	totalServices, serviceOverview, err := h.db.GetServiceOverview(ctx, req.NumPerPage, req.PageNumber, req.SortOrder, req.Filter)
	if err != nil {
		utils.ResponseFromError(err, w)
		utils.LogRequestError(err, r)
		return
	}

	// Send the response as JSON
	res := GetServiceOverviewResponse{
		Total:            totalServices,
		ServiceOverviews: serviceOverview,
	}
	json.NewEncoder(w).Encode(res)
}

func (h *V1) PostRetrieveServiceDetails(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body
	var req GetServiceDetailsRequest
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
	total, serviceDetails, err := h.db.GetServiceDetails(ctx, req.ServiceID, req.NumPerPage, req.PageNumber, req.SortOrder)
	if err != nil {
		utils.ResponseFromError(err, w)
		utils.LogRequestError(err, r)
		return
	}

	// Send the response as JSON
	res := GetServiceDetailsResponse{
		Total:          total,
		ServiceDetails: serviceDetails,
	}
	json.NewEncoder(w).Encode(res)
}
