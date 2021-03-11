package client

import (
	"context"
	"errors"
	"time"

	"github.com/golang/glog"
	"github.com/kolide/osquery-go"
)

// Client is the interface for the answering sql queries about the system.
type Client interface {
	// Query allows to take in a SQL query for the supported tables and
	// returns a list of maps which has the results.
	Query(ctx context.Context, sql string) ([]map[string]string, error)
}

// OsqueryClient implements the Client interface with osquery as the backend.
type OsqueryClient struct {
	client *osquery.ExtensionManagerClient
}

// NewOsqueryClient returns a OsqueryClient to perform queries.
func NewQsqueryClient(
	socket string, timeout time.Duration,
) (*OsqueryClient, error) {
	client, err := osquery.NewClient(socket, timeout)
	if err != nil {
		glog.Error("Failed to get an osquery client", err)
		return nil, err
	}

	return &OsqueryClient{client}, nil
}

func (c *OsqueryClient) Query(
	ctx context.Context,
	sql string,
) ([]map[string]string, error) {
	var (
		err = make(chan error)
		res = make(chan []map[string]string)
	)

	go func(r chan<- []map[string]string, e chan<- error) {
		res, err := c.client.QueryRows(sql)
		if err != nil {
			glog.V(2).Info("Failed to query osquery", err)
			e <- err
		}
		r <- res
	}(res, err)

	select {
	case <-ctx.Done():
		return nil, errors.New("query context cancelled")
	case e := <-err:
		return nil, e
	case r := <-res:
		return r, nil
	}
}
