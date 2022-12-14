package routers

import (
	"github.com/Shteyd/notes-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	router   *gin.Engine
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		router:   gin.Default(),
		services: services,
	}
}

func (h *Handler) Run(port string) error {
	h.setupRouterHeaders()
	h.setupRouterGroups()
	if err := h.router.Run(port); err != nil {
		return err
	}

	return nil
}

func (h *Handler) setupRouterHeaders() {
	h.router.Use(func(c *gin.Context) {
		allowedHeaders := "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max"
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	})
}

func (h *Handler) setupRouterGroups() {
	v1 := h.router.Group("/api/v1")
	auth := v1.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	notes := v1.Group("/notes", h.userIdentify)
	{
		notes.POST("/", h.createNote)
		notes.GET("/", h.getNotesByUserID)
		notes.GET("/:id", h.getNoteByID)
		notes.PUT("/:id", h.updateNote)
		notes.DELETE("/:id", h.deleteNote)
	}
}
