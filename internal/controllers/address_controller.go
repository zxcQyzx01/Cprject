package controllers

import (
	"encoding/json"
	"net/http"

	customhttp "test/internal/infrastructure/http"
	"test/internal/service"
)

// AddressController обрабатывает HTTP-запросы для работы с адресами
type AddressController struct {
	geoService service.GeoServicer  // Сервис для работы с геоданными
	responder  customhttp.Responder // Компонент для отправки HTTP-ответов
}

// NewAddressController создает новый экземпляр контроллера
func NewAddressController(geoService service.GeoServicer, responder customhttp.Responder) *AddressController {
	return &AddressController{
		geoService: geoService,
		responder:  responder,
	}
}

// Search обрабатывает запросы на поиск адресов
func (c *AddressController) Search(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.responder.Error(w, http.StatusBadRequest, "Неверные параметры запроса")
		return
	}

	query := r.Form.Get("query")
	if query == "" {
		c.responder.Error(w, http.StatusBadRequest, "Параметр 'query' не предоставлен")
		return
	}

	result, err := c.geoService.Search(query)
	if err != nil {
		c.responder.Error(w, http.StatusInternalServerError, "Ошибка поиска адреса")
		return
	}

	c.responder.JSON(w, http.StatusOK, map[string]interface{}{
		"addresses": result,
	})
}

// Geocode обрабатывает запросы на геокодирование
func (c *AddressController) Geocode(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Lat string `json:"lat"`
		Lng string `json:"lng"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.responder.Error(w, http.StatusBadRequest, "Неверный формат запроса")
		return
	}

	if req.Lat == "" || req.Lng == "" {
		c.responder.Error(w, http.StatusBadRequest, "Параметры 'lat' и 'lng' обязательны")
		return
	}

	result, err := c.geoService.Geocode(req.Lat, req.Lng)
	if err != nil {
		c.responder.Error(w, http.StatusInternalServerError, "Ошибка геокодирования")
		return
	}

	c.responder.JSON(w, http.StatusOK, map[string]interface{}{
		"addresses": result,
	})
}
