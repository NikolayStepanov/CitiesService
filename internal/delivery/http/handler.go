package http

import (
	_ "CitiesService/docs" // docs is generated by Swag CLI, you have to import it.
	"CitiesService/internal/service"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	citiesService    service.Cities
	regionsService   service.Regions
	districtsService service.Districts
}

func NewHandler(citiesService service.Cities, regionsService service.Regions, districtsService service.Districts) *Handler {
	return &Handler{citiesService: citiesService, regionsService: regionsService, districtsService: districtsService}
}

func (h *Handler) Init(host, port string) *chi.Mux {
	router := GetRouter()
	h.MountRoutes(router)
	router.GetMux().Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))
	return router.GetMux()
}
