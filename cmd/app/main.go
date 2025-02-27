package main

import (
	"log"
	"net/http"

	"test/internal/controllers"
	"test/internal/infrastructure/dadata"
	customhttp "test/internal/infrastructure/http"
	"test/internal/infrastructure/http/server"
	"test/internal/service"
)

func main() {
	// Инициализация зависимостей
	responder := customhttp.NewHTTPResponder()

	// Создание провайдера DaData
	dadataProvider := dadata.NewProvider(
		"627de73a10855ebb80eb0191f2bbb55cc72eef89",
		"7886bc85cac2562af90304564e7f04078d18dc4b",
	)

	// Создание сервиса геокодирования
	geoService := service.NewGeoService(dadataProvider)

	// Создание контроллера
	addressController := controllers.NewAddressController(geoService, responder)

	// Настройка и запуск сервера
	srv := server.New(addressController)
	if err := http.ListenAndServe(":8080", srv.Router()); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
