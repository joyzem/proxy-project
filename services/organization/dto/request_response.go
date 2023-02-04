package dto

import "github.com/joyzem/proxy-project/services/organization/domain"

type GetOrganizationsRequest struct {
}

type GetOrganizationsResponse struct {
	Organizations []*domain.Organization `json:"organizations,omitempty"`
	Err           string                 `json:"error,omitempty"`
}
