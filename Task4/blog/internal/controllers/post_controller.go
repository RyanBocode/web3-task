package controllers

import (
	"blog/internal/database"
	"blog/internal/dto"
	"blog/internal/middleware"
	"blog/internal/models"
	"blog/internal/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct{}

func NewPostController() *PostController {
	return &PostController{}
}

func (p *PostController) List(c *gin.Context) {
	var posts []models.Post
	if err := database.DB.Preload("User").Find(&posts).Error; err != nil {
		responses.JSONError(c, http.StatusInternalServerError, "server_error", err.Error())
		return
	}
	responses.JSONOK(c, posts)
}

func (p *PostController) Get(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := database.DB.Preload("User").First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			responses.JSONError(c, http.StatusNotFound, "not_found", "post not found")
			return
		}
		responses.JSONError(c, http.StatusInternalServerError, "server_error", err.Error())
		return
	}
	responses.JSONOK(c, post)
}

func (p *PostController) Create(c *gin.Context) {
	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.JSONError(c, http.StatusBadRequest, "bad_request", err.Error())
		return
	}

	uid := middleware.MustGetUserID(c)
	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  uid,
	}

	if err := database.DB.Create(&post).Error; err != nil {
		responses.JSONError(c, http.StatusInternalServerError, "server_error", err.Error())
		return
	}

	// 返回作者信息
	database.DB.Preload("User").First(&post, post.ID)
	responses.JSONOK(c, post)
}

func (p *PostController) Update(c *gin.Context) {
	id := c.Param("id")
	uid := middleware.MustGetUserID(c)
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			responses.JSONError(c, http.StatusNotFound, "not_found", "post not found")
			return
		}
		responses.JSONError(c, http.StatusInternalServerError, "server_error", err.Error())
		return
	}
	if post.UserID != uid {
		responses.JSONError(c, http.StatusForbidden, "forbidden", "only author can modify")
		return
	}
	var req dto.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.JSONError(c, http.StatusBadRequest, "bad_request", err.Error())
		return
	}
	post.Title = req.Title
	post.Content = req.Content
	if err := database.DB.Save(&post).Error; err != nil {
		responses.JSONError(c, http.StatusInternalServerError, "server_error", err.Error())
		return
	}
	// 返回作者信息
	database.DB.Preload("User").First(&post, post.ID)
	responses.JSONOK(c, post)
}

func (p *PostController) Delete(c *gin.Context) {
	id := c.Param("id")
	uid := middleware.MustGetUserID(c)
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			responses.JSONError(c, http.StatusNotFound, "not_found", "post not found")
			return
		}
		responses.JSONError(c, http.StatusInternalServerError, "server_error", err.Error())
		return
	}
	if post.UserID != uid {
		responses.JSONError(c, http.StatusForbidden, "forbidden", "only author can delete")
		return
	}
	if err := database.DB.Delete(&post).Error; err != nil {
		responses.JSONError(c, http.StatusInternalServerError, "server_error", err.Error())
		return
	}
	c.Status(http.StatusNoContent)
}
