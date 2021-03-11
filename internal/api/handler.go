package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
	"github.com/prateeknischal/osqueryexporter/internal/client"
)

type QueryHandler struct {
	c client.Client
}

type query struct {
	Query string `json:"query"`
}

func (qh QueryHandler) handle(w http.ResponseWriter, r *http.Request) {
	q := query{}
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		glog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	glog.V(2).Infof("Query: %s", q.Query)
	res, err := qh.c.Query(context.Background(), q.Query)
	if err != nil {
		glog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}
