package api

import (
	"context"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prateeknischal/osqueryexporter/internal/client"
)

func Server(ctx context.Context, addr string, cl client.Client) {
	handler := QueryHandler{cl}
	r := mux.NewRouter()
	r.HandleFunc("/query", handler.handle).Methods(http.MethodPost)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// health check endpoint for kubernetes
		w.WriteHeader(http.StatusOK)
	})

	server := http.Server{
		Addr:    addr,
		Handler: handlers.CombinedLoggingHandler(os.Stdout, r),
	}
	glog.V(2).Infof("Starting server at %s", addr)
	glog.Fatal(server.ListenAndServe())
}
