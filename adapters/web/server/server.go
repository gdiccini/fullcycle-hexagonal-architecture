package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codeedu/go-hexagonal/adapters/web/handler"
	"github.com/codeedu/go-hexagonal/application"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	router := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(router, n, w.Service)
	http.Handle("/", router)


	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr: ":8080",
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}