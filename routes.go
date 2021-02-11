package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// Вывести всех пользователей
func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// Вывести информацию о конкрутном пользователе
func GetUserId(c *gin.Context)  {
	id := c.Param("id")
	c.JSON(http.StatusOK, users[id])
}

// Страница регистрации
func showRegistrPage(c *gin.Context) {
	c.HTML(http.StatusOK,
		"5_registr.html",
		gin.H{"Author": getAuthor(c)})
}

// Регистрация
func PostUser(c *gin.Context) {
	jsUser := user {
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
		Email: c.PostForm("email"),
	}
	if _, ok := users[jsUser.Username]; ok {
		c.HTML(http.StatusOK,
			"6_registr.html",
			gin.H{"Error": true})
		c.Status(http.StatusConflict) // чтб не перетереть пользователя
		return
	}

	newToken, _ := uuid.NewUUID()
	tokens[newToken.String()] = jsUser.Username
	users[jsUser.Username] = &jsUser

	//c.String(http.StatusOK, newToken.String())

	c.SetCookie("token", newToken.String(), 10000, "", "", false, true) // потом проверить

	c.HTML(http.StatusOK,
		"6_registr-success.html",
		nil)
}

// Страница успешного завершения регистрации
func showRegistrSuccessPage(c *gin.Context) {
	fmt.Println("REG SUCC")
	//fmt.Println("getAuthor() -> ", getAuthor(c))
	//fmt.Println("getIsAuthIn() -> ", getIsAuthIn(c))
	c.HTML(http.StatusOK,
		"6_registr-success.html",
		gin.H{
			"Author": getAuthor(c),
			"IsAuthIn": getIsAuthIn(c),
		})
}

// Страница авторизации
func showAuthPage(c *gin.Context) {
	c.HTML(http.StatusOK,
		"7_auth.html",
		gin.H{"Author": getAuthor(c)})
}

// Авторизация
func PostUserAuth(c *gin.Context)  {
	var jsUser user
	err := c.BindJSON(&jsUser)
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		if u.Username == jsUser.Username &&
			u.Password == jsUser.Password {
			t, err := u.getToken()
			if err != nil {
				c.Status(http.StatusForbidden)
				return
			}
			c.String(http.StatusOK, t)
			return
		}
	}
	c.Status(http.StatusForbidden) // пользователь не найден (403 - запрещено)
}

func showAuthSuccessPage(c *gin.Context) {
	c.HTML(http.StatusOK,
		"8_auth-success.html",
		gin.H{"Author": getAuthor(c)})
}

func PutUser(c *gin.Context) {
	var jsUser user
	err := c.BindJSON(&jsUser)
	if err != nil {
		log.Fatal(err)
	}
	users[jsUser.Username].Password = jsUser.Password
}

func DeleteUser(c *gin.Context) {

}

func showMainPage(c *gin.Context) {
	c.HTML(http.StatusOK,
		"4_index.html",
		gin.H{"Author": getAuthor(c)})
}

func logout(c *gin.Context) {

}

func showChPassPage(c *gin.Context) {

}

func getAuthor(c *gin.Context) string {
	a, ok := c.Get("author")
	if ok {
		return a.(string)
	}
	return "Guest"
}

func getIsAuthIn(c *gin.Context) bool {
	a, ok := c.Get("isAuthIn")
	if ok {
		return a.(bool)
	}
	return false
}

func authVars(c *gin.Context) interface{} {
	return gin.H{
		"Author": getAuthor(c),
		"IsAuthIn": getIsAuthIn(c),
	}
}


func routesInit()  {
		// СТАТИКА

	router.Static("/static", "./static")

	// Сессия
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.Use(authUser())


		// ПОЛЬЗОВАТЕЛЬ

	// Показать главную страницу
	// GET /
	router.GET("/", showMainPage)

	// Показать страницу регистрации
	// GET /registr
	router.GET("/registr", showRegistrPage)
	router.GET("/registr-success", showRegistrSuccessPage)

	// Показать страницу авторизации
	// GET /auth
	router.GET("/auth", showAuthPage)
	router.GET("/auth-success", showAuthSuccessPage)

	// Выйти из учетной записи
	// GET /logout
	router.GET("/logout", logout)

	// Показать страницу изменения пароля
	// GET /chPass
	router.GET("/chPass", showChPassPage)


		// REST API

	// Вывести всех пользователей
	// GET /user
	router.GET("/user", requireAuth(), GetUser)

	// Вывести одного пользователя
	// GET /user/:id
	router.GET("/user/:id", GetUserId)

	// Регистрация
	// POST /user
	router.POST("/user", PostUser)

	// Авторизация
	// POST /user/auth
	router.POST("/user/auth", PostUserAuth)

	// PUT /user
	router.PUT("user", PutUser)

	// DELETE /user
	//router.DELETE("user", DeleteUser)
}