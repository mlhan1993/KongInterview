package v1

import (
	"encoding/json"
	"github.com/mlhan1993/KongInterview/internal/service"
	"github.com/mlhan1993/KongInterview/internal/utils"
	"github.com/mlhan1993/KongInterview/pkg/errors"
	"net/http"
)

type GetServiceOverviewRequest struct {
	SortOrder  string `json:"sortOrder,omitempty"`
	Filter     string `json:"filter,omitempty"`
	PageSize   uint   `json:"pageSize,omitempty"`
	PageNumber uint   `json:"pageNumber,omitempty"`
}

func (r *GetServiceOverviewRequest) Validate() error {
	if r.SortOrder != "" && r.SortOrder != "asc" && r.SortOrder != "desc" {
		return errors.NewBadRequestError("Invalid sortOrder, sortOrder must be either asc or desc")
	}
	return nil
}

type GetServiceOverviewResponse struct {
	Total            uint               `json:"total"`
	ServiceOverviews []service.Overview `json:"serviceOverviews"`
}

type GetServiceDetailsRequest struct {
	ServiceID  string `json:"serviceId"`
	SortOrder  string `json:"sortOrder,omitempty"`
	PageSize   uint   `json:"pageSize,omitempty"`
	PageNumber uint   `json:"pageNumber,omitempty"`
}

func (r *GetServiceDetailsRequest) Validate() error {
	if r.SortOrder != "" && r.SortOrder != "asc" && r.SortOrder != "desc" {
		return errors.NewBadRequestError("Invalid sortOrder, sortOrder must be either asc or desc")
	}

	return nil
}

type GetServiceDetailsResponse struct {
	Total          uint             `json:"total"`
	ServiceDetails []service.Detail `json:"serviceDetails"`
}

type ServiceDB interface {
	GetServiceOverview(sortOrder string, filter string, numPerPage, pageNumber uint) (uint, []service.Overview, error)
	GetServiceDetails(serviceId string, sortOrder string, numPerPage, pageNumber uint) (uint, []service.Detail, error)
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
		return
	}

	totalServices, serviceOverview, err := h.db.GetServiceOverview(req.SortOrder, req.Filter, req.PageSize, req.PageNumber)
	if err != nil {
		utils.ResponseFromError(err, w)
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
		return
	}

	total, serviceDetails, err := h.db.GetServiceDetails(req.ServiceID, req.SortOrder, req.PageSize, req.PageNumber)
	if err != nil {
		utils.ResponseFromError(err, w)
		return
	}

	// Send the response as JSON
	res := GetServiceDetailsResponse{
		Total:          total,
		ServiceDetails: serviceDetails,
	}
	json.NewEncoder(w).Encode(res)
}
