package view

import (
	"log"
	"net/http"

	"github.com/sebastianbordi/DataloggerDemo/configuration"
	"github.com/sebastianbordi/DataloggerDemo/util"
)

func CorsOptionEndpoint(w http.ResponseWriter, r *http.Request) {
	config, err := configuration.GetInstance()
	if err != nil {
		log.Println(err)
		BasicResponse(&w, 500, "internal server error")
		return
	}
	SetHeaders(w, r, *config)
	w.WriteHeader(204)
}

func SetHeaders(w http.ResponseWriter, r *http.Request, config configuration.Config) {
	origin := r.Header.Get("Origin")
	method := r.Method

	if config.GetEnvironment() == configuration.DEVELOPMENT {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Header", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
	} else {
		if util.StringContains(config.GetAllowedOrigins(), origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Header", "*")
			w.Header().Set("Access-Control-Allow-Method", method)
		}
	}

}
