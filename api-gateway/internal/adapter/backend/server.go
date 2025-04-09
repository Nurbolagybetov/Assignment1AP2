package backend

import (
	"api-gateway/config"
	"api-gateway/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTPServer struct {
	router  *gin.Engine
	usecase usecase.GatewayUsecase
	opts    *Options
}

func NewHTTPServer(cfg *config.Config, usecase usecase.GatewayUsecase) *HTTPServer {
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

	s.router.POST("/api/products", s.forwardToInventory)
	s.router.GET("/api/products/:id", s.forwardToInventory)
	s.router.PATCH("/api/products/:id", s.forwardToInventory)
	s.router.DELETE("/api/products/:id", s.forwardToInventory)
	s.router.GET("/api/products", s.forwardToInventory)

	s.router.POST("/api/orders", s.forwardToOrder)
	s.router.GET("/api/orders/:id", s.forwardToOrder)
	s.router.PATCH("/api/orders/:id", s.forwardToOrder)
	s.router.GET("/api/orders", s.forwardToOrder)
}

func (s *HTTPServer) forwardToInventory(c *gin.Context) {
	body, _ := c.GetRawData()
	endpoint := c.Request.URL.Path
	resp, err := s.usecase.ForwardToInventory(c.Request.Context(), endpoint, c.Request.Method, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", resp)
}

func (s *HTTPServer) forwardToOrder(c *gin.Context) {
	body, _ := c.GetRawData()
	endpoint := c.Request.URL.Path
	resp, err := s.usecase.ForwardToOrder(c.Request.Context(), endpoint, c.Request.Method, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", resp)
}

func (s *HTTPServer) Run(port string) error {
	return s.router.Run(port)
}
