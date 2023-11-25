package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	controller "api/src/controllers"
)

func RegisterRoutes() http.Handler {
	router := mux.NewRouter()

	PlacesController := controller.PlacesController{}

	router.HandleFunc("/place-details/{placeId}", PlacesController.GetPlaceDetails()).Methods(http.MethodGet)

	corsHandler := cors.Default().Handler(router)

	return corsHandler
}
