package handler

import (
	"day_06/internal/credentials"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(credentials *credentials.LogPass) gin.HandlerFunc {
	return func(c *gin.Context) {
		// if c.Request.URL.Path == "/admin" || c.Request.URL.Path == "/admin/post-article" {
		user, pass, ok := c.Request.BasicAuth()

		// Проверяем, были ли получены данные аутентификации из заголовка запроса
		if !ok || user != credentials.AdminLogin || pass != credentials.AdminPassword {
			// Если аутентификационные данные не были предоставлены или они неверные, отправляем заголовок WWW-Authenticate с сообщением об ошибке
			c.Header("WWW-Authenticate", `Basic realm="Please enter your username and password for access"`)
			// Отправляем статус HTTP 401 Unauthorized и завершаем выполнение обработчика запроса
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Если аутентификация прошла успешно, передаем управление следующему middleware или обработчику запроса
		c.Next()
	}
	// }
}
