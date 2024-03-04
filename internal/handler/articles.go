package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getArticles(c *gin.Context) {
	articles, err := h.services.GetArticles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//дефолтное значение страницы это 1
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Page number"})
		return
	}
	// c.HTML(http.StatusOK, "index.html", gin.H{
	// 	"articles": articles,
	// })

	pageSize := 3

	totalPages := (len(articles) + pageSize - 1) / pageSize
	start := (page - 1) * pageSize
	end := start + pageSize
	if end > len(articles) {
		end = len(articles)
	}

	// Slice articles for the current page
	currentPageArticles := articles[start:end]
	prevPage := page - 1
	nextPage := page + 1
	// Render HTML template with articles and pagination data
	c.HTML(http.StatusOK, "index.html", gin.H{
		"articles":    currentPageArticles,
		"currentPage": page,
		"totalPages":  totalPages,
		"PrevPage":    prevPage,
		"NextPage":    nextPage,
	})
	//
	// fmt.Println(articles)
}

func (h *Handler) getArticleById(c *gin.Context) {
	articleID := c.Param("id")
	fmt.Println("pam pam pam", articleID)
	articleID = trimFirstRune(articleID)
	articleIDInt, err := strconv.Atoi(articleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This article does not exist, sorry"})
		return
	}
	article, err := h.services.GetArticleById(articleIDInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This article does not exist in this database"})
	}
	c.HTML(http.StatusOK, "single_article.html", gin.H{
		"article": article,
	})
}
func trimFirstRune(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return ""
}
