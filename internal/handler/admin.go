package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAdminPanel(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_panel.html", nil)
}

func (h *Handler) AddArticle(c *gin.Context) {
	articleName := c.PostForm("article-name")
	articleContent := c.PostForm("article-content")
	fmt.Println(articleContent)
	fmt.Println(articleContent)
	if err := h.services.AddArticle(articleName, articleContent); err != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error adding article"})
			return
		}
	}
	c.Redirect(http.StatusSeeOther, "/")
}
