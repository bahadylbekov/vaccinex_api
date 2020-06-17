package apiserver

import (
	"net/http"
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
	"github.com/gin-gonic/gin"
)

// HandleProfileCreate ...
func (s *Server) HandleProfileCreate(c *gin.Context) {
	var organization *model.Organization
	if err := c.BindJSON(&organization); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}
	now := time.Now()
	organization.CreatedAt = now
	organization.CreatedBy = c.Value("userID").(string)

	if err := s.store.Organization().Create(organization, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":       organization.Name,
		"email":      organization.Email,
		"phone":      organization.Phone,
		"website":    organization.Website,
		"country":    organization.Country,
		"city":       organization.City,
		"street":     organization.Street,
		"postcode":   organization.Postcode,
		"is_active":  organization.IsActive,
		"created_by": organization.CreatedBy,
		"created_at": organization.CreatedAt,
	})
}

// HandleGetMyProfile ...
func (s *Server) HandleGetMyProfile(c *gin.Context) {
	createdBy := c.Value("userID").(string)
	organization, err := s.store.Organization().GetMyOrganization(createdBy)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, organization)
}

// HandleUpdateProfile changes account data in database
func (s *Server) HandleUpdateProfile(c *gin.Context) {
	var organization *model.Organization

	if err := c.BindJSON(&organization); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	now := time.Now()
	organization.UpdatedAt = now
	organization.UpdatedBy = c.Value("userID").(string)

	if organization.CreatedBy != organization.UpdatedBy {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	if err := s.store.Organization().Update(organization, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"organization_id": organization.OrganizationID,
		"name":            organization.Name,
		"email":           organization.Email,
		"phone":           organization.Phone,
		"website":         organization.Website,
		"country":         organization.Country,
		"city":            organization.City,
		"street":          organization.Street,
		"postcode":        organization.Postcode,
		"is_active":       organization.IsActive,
		"created_by":      organization.CreatedBy,
		"created_at":      organization.CreatedAt,
		"updated_by":      organization.UpdatedBy,
		"updated_at":      organization.UpdatedAt,
	})
}
