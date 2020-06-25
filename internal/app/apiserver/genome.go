package apiserver

import (
	"net/http"
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
	"github.com/gin-gonic/gin"
)

// HandleGenomeCreate ...
func (s *Server) HandleGenomeCreate(c *gin.Context) {
	var g *model.Genome
	if err := c.BindJSON(&g); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}
	now := time.Now()
	g.CreatedAt = now
	g.CreatedBy = c.Value("userID").(string)

	if err := s.store.Genome().Create(g, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":              g.Name,
		"organization_name": g.OrganizationName,
		"file_url":          g.FileUrl,
		"price":             g.Price,
		"virus_name":        g.VirusName,
		"simularity_rate":   g.SimularityRate,
		"origin":            g.Origin,
		"is_active":         g.IsActive,
		"is_sold":           g.IsSold,
		"created_by":        g.CreatedBy,
	})
}

// HandleGetGenomes returns all viruses from database
func (s *Server) HandleGetGenomes(c *gin.Context) {
	genomes, err := s.store.Genome().GetGenomes()
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, genomes)
}

// HandleGetMyGenomes returns all viruses created by this user
func (s *Server) HandleGetMyGenomes(c *gin.Context) {
	createdBy := c.Value("userID").(string)
	genomes, err := s.store.Genome().GetMyGenomes(createdBy)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, genomes)
}

//HandleGetGenomesByVirus returns all genomes for specific virus
func (s *Server) HandleGetGenomesByVirus(c *gin.Context) {
	id := c.Param("id")
	genomes, err := s.store.Genome().GetGenomesByVirus(id)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, genomes)
}

//HandleGetGenomesByOrganization returns all genomes for specific virus
func (s *Server) HandleGetGenomesByOrganization(c *gin.Context) {
	id := c.Param("id")
	genomes, err := s.store.Genome().GetGenomesByOrganization(id)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, genomes)
}
