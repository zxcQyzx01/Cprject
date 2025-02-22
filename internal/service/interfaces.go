package service

import "test/internal/domain"

// GeoServicer определяет интерфейс сервиса для работы с геоданными
type GeoServicer interface {
	// Search ищет адреса по строке запроса
	Search(query string) ([]*domain.Address, error)
	// Geocode получает адрес по координатам
	Geocode(lat, lng string) ([]*domain.Address, error)
}

// GeoService реализует интерфейс GeoServicer
type GeoService struct {
	provider domain.GeoProvider
}

// NewGeoService создает новый экземпляр GeoService
func NewGeoService(provider domain.GeoProvider) GeoServicer {
	return &GeoService{
		provider: provider,
	}
}

// Search реализует поиск адресов через провайдер
func (s *GeoService) Search(query string) ([]*domain.Address, error) {
	return s.provider.AddressSearch(query)
}

// Geocode реализует геокодирование через провайдер
func (s *GeoService) Geocode(lat, lng string) ([]*domain.Address, error) {
	return s.provider.GeoCode(lat, lng)
}
