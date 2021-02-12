package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

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

// Главная страица
func showMainPage(c *gin.Context) {
	c.HTML(http.StatusOK,
		"4_index.html",
		authVars(c))
}

// Страница регистрации
func showRegistrPage(c *gin.Context) {
	c.HTML(http.StatusOK,
		"5_registr.html",
		authVars(c))
}

// Регистрация
func PostUser(c *gin.Context) {
	fmt.Println(c.PostForm("username"))
	newUser := user {
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
		Email: c.PostForm("email"),
	}
	if _, ok := users[newUser.Username]; ok {
		c.HTML(http.StatusOK,
			"5_registr.html",
			gin.H{
				"Author": getAuthor(c),
				"IsAuthIn": getIsAuthIn(c),
				"Error": true})
		c.Status(http.StatusConflict) // чтоб не перетереть пользователя
		return
	}

	newToken, _ := uuid.NewUUID()
	tokens[newToken.String()] = newUser.Username
	users[newUser.Username] = &newUser

	c.SetCookie("token", newToken.String(), 10000, "", "", false, true) // потом проверить
	cok, _ := c.Cookie("token")
	fmt.Println("---", cok, "---")

	c.Set("author", newUser.Username)
	c.Set("isAuthIn", true)
	//c.SetCookie("isAuthin", "yes", 10000, "", "", false, true)

	c.HTML(http.StatusOK,
		"6_registr-success.html",
		authVars(c))
}

// Страница авторизации
func showAuthPage(c *gin.Context) {
	c.HTML(http.StatusOK,
		"7_auth.html",
		authVars(c))
}

// Авторизация
func PostUserAuth(c *gin.Context) {
	newUser := user {
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
		Email: c.PostForm("email"),
	}
	for _, u := range users {
		if u.Username == newUser.Username &&
			u.Password == newUser.Password {
			t, err := u.getToken()
			if err != nil {
				c.Status(http.StatusForbidden) // у пользователя не найден токен
				return
			}
			c.Set("author", newUser.Username)
			c.Set("isAuthIn", true)
			c.SetCookie("token", t, 10000, "", "", false, true)
			// authUser() - почему не работает?
			c.HTML(http.StatusOK,
				"8_auth-success.html",
				authVars(c))
			return
		}
	}

	c.Status(http.StatusForbidden) // пользователь не найден (403 - запрещено)
	c.HTML(http.StatusOK,
		"7_auth.html",
		gin.H{
			"Author": getAuthor(c),
			"IsAuthIn": getIsAuthIn(c),
			"Error": true})
}

func showLogoutPage(c *gin.Context) {
	c.HTML(http.StatusOK,
		"9_logout.html",
		authVars(c))
}

func PostUserLogout(c *gin.Context) {
	c.Set("author", "Guest")
	c.Set("isAuthIn", false)
	c.SetCookie("token", "no token", 10000, "", "", false, true)
	showMainPage(c)
}

func showChPassPage(c *gin.Context) {
	c.HTML(http.StatusOK,
		"10_chPass.html",
		authVars(c))
}

func PostUserChPass(c *gin.Context) {
	/*type pass struct {
		PrevPass string // и без `json:"prevPass"` работает
		NewPass  string
	}
	var p pass
	err := c.BindJSON(&p)
	if err != nil {
		log.Fatal(err)
	}*/
	prevPass := c.PostForm("prevPass")
	newPass := c.PostForm("newPass")
	for _, u := range users {
		if u.Password == prevPass {
			u.Password = newPass
			fmt.Println("good job")
			c.HTML(http.StatusOK,
				"11_chPass-success.html",
				authVars(c))
			return
		}
	}

	c.Status(http.StatusForbidden)
	c.HTML(http.StatusOK,
		"10_chPass.html",
		gin.H{
			"Author": getAuthor(c),
			"IsAuthIn": getIsAuthIn(c),
			"Error": true})
}

// Вывести всех пользователей
func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// Вывести информацию о конкрутном пользователе
func GetUserId(c *gin.Context)  {
	id := c.Param("id")
	c.JSON(http.StatusOK, users[id])
}

/*func PutUser(c *gin.Context) {
	var jsUser user
	err := c.BindJSON(&jsUser)
	if err != nil {
		log.Fatal(err)
	}
	users[jsUser.Username].Password = jsUser.Password
}*/

func showDeleteAccPage(c *gin.Context) {
	c.HTML(http.StatusOK,
		"12_deleteAcc.html",
		authVars(c))
}

func DeleteUser(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}
	username, _ := tokens[token]
	fmt.Println(username)
	delete(users, username)

	c.Set("author", "Guest")
	c.Set("isAuthIn", false)
	c.SetCookie("token", "no token", 10000, "", "", false, true)
}


func routesInit()  {
		// СТАТИКА

	router.Static("/static", "./static")

	// Сессия
	//store := sessions.NewCookieStore([]byte("secret"))
	//router.Use(sessions.Sessions("mysession", store))

	router.Use(authUser())


		// ПОЛЬЗОВАТЕЛЬ

	// Показать главную страницу
	// GET /
	router.GET("/", showMainPage)

	// Показать страницу регистрации
	// GET /registr
	router.GET("/registr", showRegistrPage)
	//router.GET("/registr-success", showRegistrSuccessPage)

	// Показать страницу авторизации
	// GET /auth
	router.GET("/auth", showAuthPage)
	//router.GET("/auth-success", showAuthSuccessPage)

	// Выйти из учетной записи
	// GET /showLogoutPage
	router.GET("/logout", showLogoutPage)

	// Показать страницу изменения пароля
	// GET /chPass
	router.GET("/chPass", showChPassPage)

	// Удалить аккаунт
	//
	router.GET("/deleteAcc", showDeleteAccPage)


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

	// Выйти из аккаунта
	// POST /user/logout
	router.POST("/user/logout", PostUserLogout)

	// Изменить пароль
	// POST /user/chPass
	router.POST("/user/chPass", PostUserChPass)

	// Удаление аккаунта
	// DELETE /user
	router.DELETE("/user", DeleteUser)
}