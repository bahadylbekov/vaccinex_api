package apiserver

import (
	"net/http"
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
	"github.com/gin-gonic/gin"
)

// HandleVirusCreate ...
func (s *Server) HandleVirusCreate(c *gin.Context) {
	var v *model.Virus
	if err := c.BindJSON(&v); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}
	now := time.Now()
	v.CreatedAt = now
	v.CreatedBy = c.Value("userID").(string)

	if err := s.store.Virus().Create(v, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":          v.Name,
		"description":   v.Description,
		"photo_url":     v.PhotoUrl,
		"family":        v.Family,
		"fatality_rate": v.FatalityRate,
		"spread":        v.Spread,
		"is_active":     v.IsActive,
		"is_vaccine":    v.IsVaccine,
		"created_by":    v.CreatedBy,
	})
}

// HandleGetViruses returns all viruses from database
func (s *Server) HandleGetViruses(c *gin.Context) {
	viruses, err := s.store.Virus().GetViruses()
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, viruses)
}

//HandleGetVirus ...
func (s *Server) HandleGetVirusByID(c *gin.Context) {
	id := c.Param("id")
	viruses, err := s.store.Virus().GetVirusByID(id)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, viruses)
}

// HandleUpdateVirus allow user update virus information
func (s *Server) HandleUpdateVirus(c *gin.Context) {

	var v *model.Virus

	if err := c.BindJSON(&v); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	now := time.Now()
	v.UpdatedAt = now
	v.UpdatedBy = c.Value("userID").(string)

	if v.CreatedBy != v.UpdatedBy {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	if err := s.store.Virus().Update(v, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":            v.VirusID,
		"name":          v.Name,
		"description":   v.Description,
		"photo_url":     v.PhotoUrl,
		"family":        v.Family,
		"fatality_rate": v.FatalityRate,
		"spread":        v.Spread,
		"isActive":      v.IsActive,
		"isVaccine":     v.IsVaccine,
		"created_at":    v.CreatedBy,
		"created_by":    v.CreatedBy,
		"updated_at":    v.UpdatedAt,
		"updated_by":    v.UpdatedBy,
	})
}
