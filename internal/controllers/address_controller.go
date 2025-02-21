package controllers

import (
	"encoding/json"
	"net/http"

	customhttp "test/internal/infrastructure/http"
	"test/internal/service"
)

type AddressController struct {
	geoProvider service.GeoProvider
	responder   customhttp.Responder
}

func NewAddressController(geoProvider service.GeoProvider, responder customhttp.Responder) *AddressController {
	return &AddressController{
		geoProvider: geoProvider,
		responder:   responder,
	}
}

type SearchRequest struct {
	Query string `json:"query"`
}

type GeocodeRequest struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

func (c *AddressController) Search(w http.ResponseWriter, r *http.Request) {
	var req SearchRequest

	if err := r.ParseForm(); err != nil {
		c.responder.Error(w, http.StatusBadRequest, "Неверные параметры запроса")
		return
	}

	req.Query = r.Form.Get("query")
	if req.Query == "" {
		c.responder.Error(w, http.StatusBadRequest, "Параметр 'query' не предоставлен")
		return
	}

	result, err := c.geoProvider.AddressSearch(req.Query)
	if err != nil {
		c.responder.Error(w, http.StatusInternalServerError, "Ошибка поиска адреса")
		return
	}

	c.responder.JSON(w, http.StatusOK, map[string]interface{}{
		"addresses": result,
	})
}

func (c *AddressController) Geocode(w http.ResponseWriter, r *http.Request) {
	var req GeocodeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.responder.Error(w, http.StatusBadRequest, "Неверный формат запроса")
		return
	}

	if req.Lat == "" || req.Lng == "" {
		c.responder.Error(w, http.StatusBadRequest, "Параметры 'lat' и 'lng' обязательны")
		return
	}

	result, err := c.geoProvider.GeoCode(req.Lat, req.Lng)
	if err != nil {
		c.responder.Error(w, http.StatusInternalServerError, "Ошибка геокодирования")
		return
	}

	c.responder.JSON(w, http.StatusOK, map[string]interface{}{
		"addresses": result,
	})
}
