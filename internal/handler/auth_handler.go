package handler

import (
	"net/http"
	"payment-system/internal/repository"
	"payment-system/pkg"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Repo *repository.UserRepository
}

func NewAuthHandler(repo *repository.UserRepository) *AuthHandler {
	return &AuthHandler{Repo: repo}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Repo.Create(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed register"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registered"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, password, err := h.Repo.FindByEmail(req.Email)
	if err != nil || password != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, _ := pkg.GenerateToken(id)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
