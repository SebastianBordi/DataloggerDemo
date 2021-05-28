package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sebastianbordi/DataloggerDemo/socket"
	"github.com/sebastianbordi/DataloggerDemo/view"
)

func GetRouter(baseURL string) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", redirectFunc)

	router.NotFoundHandler = http.HandlerFunc(redirectFunc)

	fileServer := http.FileServer(http.Dir("./www"))
	router.Handle("/webpage", http.StripPrefix("/webpage", fileServer)).Methods("GET")

	router.HandleFunc(baseURL+"/socket-subscribe/{isBroadcaster}/{mac}", socket.SocketEndpoint).Methods("GET")

	router.PathPrefix(baseURL).HandlerFunc(view.CorsOptionEndpoint).Methods("OPTIONS")

	router.HandleFunc(baseURL+"/login", view.Login).Methods("POST")
	router.HandleFunc(baseURL+"/sensor", view.JWTAuth(view.CreateSensor)).Methods("POST")
	router.HandleFunc(baseURL+"/sensor", view.GetSensors).Methods("GET")
	router.HandleFunc(baseURL+"/sensor/{id}", view.GetSensorById).Methods("GET")
	router.HandleFunc(baseURL+"/sensor", view.JWTAuth(view.UpdateSensor)).Methods("PUT")
	router.HandleFunc(baseURL+"/sensor/{id}", view.JWTAuth(view.DeleteSensor)).Methods("DELETE")

	router.HandleFunc(baseURL+"/measurement", view.CreateMeasurement).Methods("POST")
	router.HandleFunc(baseURL+"/measurement", view.GetMeasurements).Methods("GET")
	router.HandleFunc(baseURL+"/measurement/{id}", view.GetMeasurementById).Methods("GET")
	//router.HandleFunc(baseURL+"/temperature", view.UpdateMeasurement).Methods("PUT")
	router.HandleFunc(baseURL+"/measurement/{id}", view.JWTAuth(view.DeleteMeasurement)).Methods("DELETE")

	return router
}

func redirectFunc(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Location", "/webpage")
	rw.WriteHeader(301)
}
