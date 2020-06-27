package apiserver

import (
	"net/http"
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
	"github.com/gin-gonic/gin"
)

// HandleCreateVaccine creates vaccine using Vaccine model
func (s *Server) HandleCreateVaccine(c *gin.Context) {
	var v *model.Vaccine
	if err := c.BindJSON(&v); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}
	now := time.Now()
	v.CreatedAt = now
	v.CreatedBy = c.Value("userID").(string)

	if err := s.store.Vaccine().Create(v, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":             v.Name,
		"virus_id":         v.VirusID,
		"virus_name":       v.VirusName,
		"description":      v.Description,
		"requested_amount": v.RequestedAmount,
		"funded_amount":    v.FundedAmount,
		"is_active":        v.IsActive,
		"created_by":       v.CreatedBy,
		"created_at":       v.CreatedAt,
	})
}

// HandleGetVaccines returns all vaccines from database
func (s *Server) HandleGetVaccines(c *gin.Context) {

	vaccines, err := s.store.Vaccine().GetVaccines()
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, vaccines)
}

//HandleGetVirus ...
func (s *Server) HandleGetVaccineByID(c *gin.Context) {
	id := c.Param("id")
	vaccines, err := s.store.Vaccine().GetVaccineByID(id)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, vaccines)
}

// HandleUpdateVaccineAmount updates vaccine's funded amount
func (s *Server) HandleUpdateVaccineAmount(c *gin.Context) {
	type repsonse struct {
		VaccineID    int       `json:"vaccine_id"`
		FundedAmount string    `json:"funded_amount"`
		UpdatedBy    string    `json:"updated_by"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
	var v *repsonse

	if err := c.BindJSON(&v); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	now := time.Now()
	v.UpdatedAt = now
	v.UpdatedBy = c.Value("userID").(string)

	if err := s.store.Vaccine().UpdateAmount(v.FundedAmount, v.UpdatedBy, v.VaccineID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"vaccine_id":    v.VaccineID,
		"funded_amount": v.FundedAmount,
		"updated_by":    v.UpdatedBy,
		"updated_at":    v.UpdatedAt,
	})
}

// HandleUpdateVaccineName updates vaccine's name
func (s *Server) HandleUpdateVaccineName(c *gin.Context) {
	type repsonse struct {
		VaccineID int       `json:"vaccine_id"`
		Name      string    `json:"vaccine_name"`
		UpdatedBy string    `json:"updated_by"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	var v *repsonse

	if err := c.BindJSON(&v); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	now := time.Now()
	v.UpdatedAt = now
	v.UpdatedBy = c.Value("userID").(string)

	if err := s.store.Vaccine().UpdateName(v.Name, v.UpdatedBy, v.VaccineID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"vaccine_id":   v.VaccineID,
		"vaccine_name": v.Name,
		"updated_by":   v.UpdatedBy,
		"updated_at":   v.UpdatedAt,
	})
}

// HandleUpdateVaccineDescription updates vaccine's description
func (s *Server) HandleUpdateVaccineDescription(c *gin.Context) {
	type repsonse struct {
		VaccineID   int       `json:"vaccine_id"`
		Description string    `json:"description"`
		UpdatedBy   string    `json:"updated_by"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
	var v *repsonse

	if err := c.BindJSON(&v); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	now := time.Now()
	v.UpdatedAt = now
	v.UpdatedBy = c.Value("userID").(string)

	if err := s.store.Vaccine().UpdateDescription(v.Description, v.UpdatedBy, v.VaccineID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"vaccine_id":  v.VaccineID,
		"description": v.Description,
		"updated_by":  v.UpdatedBy,
		"updated_at":  v.UpdatedAt,
	})
}
