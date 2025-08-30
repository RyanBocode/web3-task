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

type CommentController struct{}

func NewCommentController() *CommentController {
	return &CommentController{}
}

func (cc *CommentController) ListByPost(c *gin.Context) {
	postID := c.Param("id")
	// 检查文章是否存在
	var dummy models.Post
	if err := database.DB.First(&dummy, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			responses.JSONError(c, http.StatusNotFound, "not_found", "post not found")
			return
		}
		responses.JSONError(c, http.StatusInternalServerError, "server_error", err.Error())
		return
	}
	var comments []models.Comment
	if err := database.DB.Where("post_id = ?", postID).Preload("User").Order("id ASC").Find(&comments).Error; err != nil {
		responses.JSONError(c, http.StatusInternalServerError, "server_error", err.Error())
		return
	}
	responses.JSONOK(c, comments)
}

func (cc *CommentController) Create(c *gin.Context) {
	postID := c.Param("id")
	// 确保文章存在
	var post models.Post
	if err := database.DB.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			responses.JSONError(c, http.StatusNotFound, "not_found", "post not found")
			return
		}
		responses.JSONError(c, http.StatusInternalServerError, "server_error", err.Error())
		return
	}
	var req dto.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.JSONError(c, http.StatusBadRequest, "bad_request", err.Error())
		return
	}
	uid := middleware.MustGetUserID(c)
	comment := models.Comment{Content: req.Content, UserID: uid, PostID: post.ID}
	if err := database.DB.Create(&comment).Error; err != nil {
		responses.JSONError(c, http.StatusBadRequest, "create_failed", err.Error())
		return
	}
	// 带作者返回
	database.DB.Preload("User").First(&comment, comment.ID)
	c.JSON(http.StatusCreated, comment)
}
