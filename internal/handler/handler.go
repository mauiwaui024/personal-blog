package handler

import (
	"day_06/internal/credentials"
	"day_06/internal/service"
	"log"
	"net/http"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"

	"github.com/gin-gonic/gin"
)

// наш хэндлер будет обращаться к сервисам
type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(setupRateLimiter())
	router.LoadHTMLGlob("./internal/templates/*")
	router.Static("/css", "./css")
	router.Static("/img", "./images")
	router.GET("/", h.getArticles)

	router.GET("/article/:id", h.getArticleById)
	///написать функцию считывания из файла
	creds, err := credentials.GetCredentials("./internal/credentials/credentials.txt")
	if err != nil {
		log.Fatal("error reading from credentials file")
	}
	//ахуенно
	adminRoute := router.Group("/admin", AuthMiddleware(&creds))
	{
		adminRoute.GET("", h.GetAdminPanel)
		//post new article
		adminRoute.POST("/post-article", h.AddArticle)

	}
	router.NoRoute(func(c *gin.Context) {
		// Render the 404 page HTML template
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	return router
}

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}
func setupRateLimiter() gin.HandlerFunc {
	// Create a rate limiter store (e.g., in-memory or Redis)
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 100,
	})

	// Create the rate limiter middleware
	limiter := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	// Return the middleware handler
	return limiter
}
