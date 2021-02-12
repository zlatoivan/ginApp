package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Включаем авторизацию при любом запросе
func authUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		fmt.Println("   * MW: ", token)
		// Basic 9a75e347-3d31-11eb-829b-54bf6414a1a2
		/*a := string([]rune(token)[6:])*/
		if err != nil {
			return
		}
		username, ok := tokens[token] // [6:]
		if ok {
			c.Set("author", username) // храниться в контексте (что-то не хранится пока дальше)
			c.Set("isAuthIn", true)
			fmt.Println("   *", username, true)
		} else {
			c.SetCookie("token", "no token", 10000, "", "", false, true) // Удаляем старые куки
			c.Set("isAuthIn", false)
			fmt.Println("   * Guest", false)
		}
	}
}

// Требуем авторизацию при каждом следующем запросе
func requireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.Status(http.StatusForbidden)
			return
		}
		if _, ok := tokens[token]; !ok {
			c.AbortWithStatus(http.StatusUnauthorized) // а c.Status не прерывает, поэтому так лучше
		}
	}
}