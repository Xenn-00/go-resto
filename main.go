package main

import (
	"os"

	"github.com/Xenn-00/go-resto/middleware"
	"github.com/Xenn-00/go-resto/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRouter(router)
	routes.TableRoutes(router)
	routes.OrderItemRoutes(router)
	router.Run("localhost:" + port)
}
