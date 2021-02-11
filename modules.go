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
		// Basic 9a75e347-3d31-11eb-829b-54bf6414a1a2
		//fmt.Println("authUser() -> ", token)

		/*a := string([]rune(token)[6:])
		fmt.Println("a = ", a)*/
		if err != nil {
			return
		}
		username, ok := tokens[token] // [6:]
		if ok {
			c.Set("author", username) // храниться в контексте (что-то не хранится пока дальше)
			c.Set("isAuthIn", true)
			fmt.Println(username, true)
		} else {
			c.Set("isAuthIn", false)
			fmt.Println("Guest", false)
		}
	}
}

// Требуем авторизацию при каждом следующем запросе
func requireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		a, _ := c.Get("author")
		fmt.Println("requireAuth() -> ", a)
		if _, ok := c.Get("author"); !ok {
			c.AbortWithStatus(http.StatusUnauthorized) // а c.Status не прерывает, поэтому так лучше
		}
	}
}