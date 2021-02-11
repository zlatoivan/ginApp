package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Создали маршрутизатор
	router = gin.Default()

	// Подгружаем шаблоны
	router.LoadHTMLGlob("static/templates/*")

	// Заполняем базу данных
	startDataInit()

	// Инициализируем маршруты
	routesInit()

	// Запуск сервера
	router.Run(":8080")
}