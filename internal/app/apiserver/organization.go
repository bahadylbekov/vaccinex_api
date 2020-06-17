package apiserver

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//HandleFindOrganizations ...
func (s *Server) HandleFindOrganizations(c *gin.Context) {
	search := c.Query("search")
	organizations, err := s.store.Organization().FindOrganizationsByEmail(search)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, organizations)
}

//HandleAddOrganizationToMyList ...
func (s *Server) HandleAddOrganizationToMyList(c *gin.Context) {
	organizationID := c.Query("id")
	createdBy := c.Value("userID").(string)
	now := time.Now()

	myID, err := s.store.Organization().ID(createdBy)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}

	if err := s.store.Organization().AddOrganizationToMyList(*myID, organizationID, createdBy, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":              myID,
		"organization_id": organizationID,
		"created_by":      createdBy,
		"created_at":      now,
	})
}

//HandleGetConnectedOrganizations ...
func (s *Server) HandleGetConnectedOrganizations(c *gin.Context) {
	createdBy := c.Value("userID").(string)
	organizations, err := s.store.Organization().GetConnectedOrganizations(createdBy)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, organizations)
}

//HandleGetOrganization ...
func (s *Server) HandleGetOrganization(c *gin.Context) {
	id := c.Query("id")
	organization, err := s.store.Organization().GetOrganization(id)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, organization)
}

// HandleDeleteOrganization deletes organization from connected organizations list
func (s *Server) HandleDeleteOrganization(c *gin.Context) {
	organizationID := c.Query("organization_id")

	createdBy := c.Value("userID").(string)

	myID, err := s.store.Organization().ID(createdBy)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}

	err = s.store.Organization().Delete(*myID, organizationID)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"my_id":           myID,
		"organization_id": organizationID,
	})
}
