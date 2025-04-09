package backend

import (
	"inventory-service/config"
	"inventory-service/internal/entities"
	"inventory-service/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	router  *gin.Engine
	usecase usecase.ProductUsecase
	opts    *Options
}

func NewHTTPServer(cfg *config.Config, usecase usecase.ProductUsecase) *HTTPServer {
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
	// api.Use(s.authMiddleware()) // Remove or comment out this line

	api.POST("/products", s.createProduct)
	api.GET("/products/:id", s.getProduct)
	api.PATCH("/products/:id", s.updateProduct)
	api.DELETE("/products/:id", s.deleteProduct)
	api.GET("/products", s.listProducts)
}

func (s *HTTPServer) createProduct(c *gin.Context) {
	var product entities.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := s.usecase.CreateProduct(c.Request.Context(), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (s *HTTPServer) getProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := s.usecase.GetProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (s *HTTPServer) updateProduct(c *gin.Context) {
	id := c.Param("id")
	var product entities.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product.ID = id
	if err := s.usecase.UpdateProduct(c.Request.Context(), &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (s *HTTPServer) deleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := s.usecase.DeleteProduct(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (s *HTTPServer) listProducts(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	products, err := s.usecase.ListProducts(c.Request.Context(), offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (s *HTTPServer) Run(port string) error {
	return s.router.Run(port)
}
