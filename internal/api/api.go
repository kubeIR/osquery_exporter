package api

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/prateeknischal/osqueryexporter/internal/client"
)

func Server(ctx context.Context, addr string, cl client.Client) {
	handler := QueryHandler{cl}
	r := mux.NewRouter()
	r.HandleFunc("/query", handler.handle).Methods(http.MethodPost)

	server := http.Server{
		Addr:    addr,
		Handler: r,
	}
	glog.Fatal(server.ListenAndServe())
}
