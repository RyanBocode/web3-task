package controllers

import (
	"blog/internal/config"
	"blog/internal/database"
	"blog/internal/dto"
	"blog/internal/middleware"
	"blog/internal/models"
	"blog/internal/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	Cfg *config.Config
}

func NewAuthController(cfg *config.Config) *AuthController {
	return &AuthController{Cfg: cfg}
}

func (a *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.JSONError(c, http.StatusBadRequest, "bad_request", err.Error())
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		responses.JSONError(c, http.StatusInternalServerError, "server_error", "failed to hash password")
		return
	}
	user := models.User{Username: req.Username, Email: req.Email, Password: string(hash)}
	if err := database.DB.Create(&user).Error; err != nil {
		responses.JSONError(c, http.StatusBadRequest, "create_failed", err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

func (a *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.JSONError(c, http.StatusBadRequest, "bad_request", err.Error())
		return
	}
	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		responses.JSONError(c, http.StatusUnauthorized, "invalid_credentials", "username or password is wrong")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		responses.JSONError(c, http.StatusUnauthorized, "invalid_credentials", "username or password is wrong")
		return
	}
	token, err := middleware.NewToken(a.Cfg, user.ID, user.Username)
	if err != nil {
		responses.JSONError(c, http.StatusInternalServerError, "server_error", "failed to create token")
		return
	}
	c.JSON(http.StatusOK, dto.LoginResponse{Token: token})
}

func (a *AuthController) Me(c *gin.Context) {
	uid := middleware.MustGetUserID(c)
	var user models.User
	if err := database.DB.First(&user, uid).Error; err != nil {
		responses.JSONError(c, http.StatusNotFound, "not_found", "user not found")
		return
	}
	// 不返回密码
	user.Password = ""
	responses.JSONOK(c, user)
}
