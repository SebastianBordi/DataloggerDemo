package middleware

import (
	"log"
	"net/http"

	"github.com/sebastianbordi/DataloggerDemo/configuration"
	"github.com/sebastianbordi/DataloggerDemo/view"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		config, err := configuration.GetInstance()
		if err != nil {
			log.Println(err)
			view.BasicResponse(&w, 500, "internal server error")
			return
		}
		view.SetHeaders(w, r, *config)
		next.ServeHTTP(w, r)
	})
}
