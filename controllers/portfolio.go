package controllers

import (
	"errors"
	"net/http"

	"ilits-porto-backend/config"
	"ilits-porto-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllPortfolios(c *gin.Context) {
	var portfolios []models.Portfolio
	query := config.DB.Model(&models.Portfolio{})

	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}

	if q := c.Query("q"); q != "" {
		searchTerm := "%" + q + "%"
		query = query.Where("title LIKE ? OR tech_stack LIKE ?", searchTerm, searchTerm)
	}

	if err := query.Order("created_at desc").Find(&portfolios).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseFormat{
			Success: false,
			Message: "gagal ambil data",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseFormat{
		Success: true,
		Message: "jago banget, bisa ambil data",
		Data:    portfolios,
	})
}

func GetPortfolioByID(c *gin.Context) {
	id := c.Param("id")
	var portfolio models.Portfolio

	if err := config.DB.First(&portfolio, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, models.ResponseFormat{
				Success: false,
				Message: "porto gaada",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, models.ResponseFormat{
			Success: false,
			Message: "gagal ambil detail porto",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseFormat{
		Success: true,
		Message: "berhasil ambil detail porto",
		Data:    portfolio,
	})
}

func CreatePortfolio(c *gin.Context) {
	var input models.CreatePortfolioInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseFormat{
			Success: false,
			Message: "input data ngga valid",
			Error:   err.Error(),
		})
		return
	}

	status := input.Status
	if status == "" {
		status = "Completed"
	}

	portfolio := models.Portfolio{
		Title:       input.Title,
		Description: input.Description,
		Category:    input.Category,
		TechStack:   input.TechStack,
		ProjectURL:  input.ProjectURL,
		Status:      status,
	}

	if err := config.DB.Create(&portfolio).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseFormat{
			Success: false,
			Message: "gagal nyimpen porto",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.ResponseFormat{
		Success: true,
		Message: "porto berhasil ditambahkan",
		Data:    portfolio,
	})
}

func UpdatePortfolio(c *gin.Context) {
	id := c.Param("id")
	var portfolio models.Portfolio

	if err := config.DB.First(&portfolio, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, models.ResponseFormat{
				Success: false,
				Message: "porto ngga ada",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, models.ResponseFormat{
			Success: false,
			Message: "gagal mencari data porto",
			Error:   err.Error(),
		})
		return
	}

	var input models.UpdatePortfolioInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseFormat{
			Success: false,
			Message: "input update tidak valid",
			Error:   err.Error(),
		})
		return
	}

	updates := map[string]interface{}{}
	if input.Title != "" {
		updates["title"] = input.Title
	}
	if input.Description != "" {
		updates["description"] = input.Description
	}
	if input.Category != "" {
		updates["category"] = input.Category
	}
	if input.TechStack != "" {
		updates["tech_stack"] = input.TechStack
	}
	if input.ProjectURL != "" {
		updates["project_url"] = input.ProjectURL
	}
	if input.Status != "" {
		updates["status"] = input.Status
	}

	if err := config.DB.Model(&portfolio).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseFormat{
			Success: false,
			Message: "gagal update portofolio",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseFormat{
		Success: true,
		Message: "berhasil update portofolio",
		Data:    portfolio,
	})
}

func DeletePortfolio(c *gin.Context) {
	id := c.Param("id")
	var portfolio models.Portfolio

	if err := config.DB.First(&portfolio, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, models.ResponseFormat{
				Success: false,
				Message: "porto tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, models.ResponseFormat{
			Success: false,
			Message: "gagal mencari data portofolio",
			Error:   err.Error(),
		})
		return
	}

	if err := config.DB.Delete(&portfolio).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseFormat{
			Success: false,
			Message: "gagal hapus portofolio",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseFormat{
		Success: true,
		Message: "porto udah berhasil dihapus",
	})
}
