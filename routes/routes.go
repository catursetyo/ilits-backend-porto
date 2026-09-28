package routes

import (
	"net/http"

	"ilits-porto-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hi! selamat datang lee",
			"status":  "running",
		})
	})

	v1 := r.Group("/api/v1")
	{
		portfolios := v1.Group("/portfolios")
		{
			portfolios.GET("", controllers.GetAllPortfolios)
			portfolios.GET("/:id", controllers.GetPortfolioByID)
			portfolios.POST("", controllers.CreatePortfolio)
			portfolios.PUT("/:id", controllers.UpdatePortfolio)
			portfolios.DELETE("/:id", controllers.DeletePortfolio)
		}
	}

	return r
}
