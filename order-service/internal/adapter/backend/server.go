package backend

import (
	"net/http"
	"order-service/config"
	"order-service/internal/entities"
	"order-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	router  *gin.Engine
	usecase usecase.OrderUsecase
	opts    *Options
}

func NewHTTPServer(cfg *config.Config, usecase usecase.OrderUsecase) *HTTPServer {
	router := gin.Default()
	opts := NewOptions(cfg.Port)

	server := &HTTPServer{
		router:  router,
		usecase: usecase,
		opts:    opts,
	}
	server.setupRoutes()
	return server
}

func (s *HTTPServer) setupRoutes() {
	api := s.router.Group("/api")

	api.POST("/orders", s.createOrder)
	api.GET("/orders/:id", s.getOrder)
	api.PATCH("/orders/:id", s.updateOrder)
	api.GET("/orders", s.listOrders)
}

func (s *HTTPServer) createOrder(c *gin.Context) {
	var order entities.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.usecase.CreateOrder(c.Request.Context(), &order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}

func (s *HTTPServer) getOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := s.usecase.GetOrder(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (s *HTTPServer) updateOrder(c *gin.Context) {
	id := c.Param("id")
	var order entities.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.ID = id
	if err := s.usecase.UpdateOrder(c.Request.Context(), &order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (s *HTTPServer) listOrders(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	orders, err := s.usecase.ListOrders(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (s *HTTPServer) Run(port string) error {
	return s.router.Run(port)
}
