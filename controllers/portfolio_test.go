package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ilits-porto-backend/config"
	"ilits-porto-backend/models"
	"ilits-porto-backend/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDBAndRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	var err error
	config.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	_ = config.DB.AutoMigrate(&models.Portfolio{})

	return routes.SetupRouter()
}

func TestHealthCheck(t *testing.T) {
	router := setupTestDBAndRouter()

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}
}

func TestPortfolioCRUD(t *testing.T) {
	router := setupTestDBAndRouter()

	payload := models.CreatePortfolioInput{
		Title:       "Portofolio",
		Description: "REST API Portofolio Go, Gin, dan GORM",
		Category:    "Backend",
		TechStack:   "Golang, Gin, GORM, SQLite",
		ProjectURL:  "https://github.com/catursetyo/ilits-porto-backend",
		Status:      "Completed",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/portfolios", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status 201, got %d: %s", w.Code, w.Body.String())
	}

	var createRes struct {
		Success bool             `json:"success"`
		Data    models.Portfolio `json:"data"`
	}
	_ = json.Unmarshal(w.Body.Bytes(), &createRes)
	createdID := createRes.Data.ID

	req, _ = http.NewRequest(http.MethodGet, "/api/v1/portfolios", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	req, _ = http.NewRequest(http.MethodGet, "/api/v1/portfolios/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	updatePayload := models.UpdatePortfolioInput{
		Title: "Portofolio v2",
	}
	updateBody, _ := json.Marshal(updatePayload)
	req, _ = http.NewRequest(http.MethodPut, "/api/v1/portfolios/1", bytes.NewBuffer(updateBody))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	req, _ = http.NewRequest(http.MethodDelete, "/api/v1/portfolios/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	req, _ = http.NewRequest(http.MethodGet, "/api/v1/portfolios/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("Expected status 404, got %d", w.Code)
	}

	_ = createdID
}
