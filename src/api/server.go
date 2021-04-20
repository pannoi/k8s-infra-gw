package api

import (
	"context"
	"net/http"

	c "infra-gw/src/cont"
	"infra-gw/src/api/resources"
	"infra-gw/src/api/apps"

	"github.com/gorilla/mux"
)

// Web is http api implementation
type Web struct {
	httpServer *http.Server
}

// NewWeb creates a new web api server
func NewWeb(ctx context.Context, apiListenAddr string) *Web {
	w := &Web{
		httpServer: &http.Server{
			Handler: mux.NewRouter(),
			Addr:    apiListenAddr,
		},
	}
	return w
}

// Start starts the web server
func (w *Web) Start(ctx context.Context, appCtx *c.AppContext) chan error {
	log := appCtx.Log
	doneChan := make(chan error)

	log.With("addr", w.httpServer.Addr).Debug("starting web api")
	go func() {
		err := w.httpServer.ListenAndServe()
		doneChan <- err
	}()
	return doneChan
}

// SetupRoutes defines web api routes
func (w *Web) SetupRoutes(ctx context.Context, appCtx *c.AppContext) {
	log := appCtx.Log

	w.httpServer.Handler.(*mux.Router).NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugw("NotFoundHandler called", "path", r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
	})

	w.httpServer.Handler.(*mux.Router).HandleFunc("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		log.Debug("return application status")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("1"))
	}))
	// GET Requests
	w.httpServer.Handler.(*mux.Router).Path("/pods").Methods("GET").HandlerFunc(resources.ListPods(ctx, appCtx))
	w.httpServer.Handler.(*mux.Router).Path("/services").Methods("GET").HandlerFunc(resources.ListServices(ctx, appCtx))
	w.httpServer.Handler.(*mux.Router).Path("/deployments").Methods("GET").HandlerFunc(resources.ListDeployments(ctx, appCtx))
	w.httpServer.Handler.(*mux.Router).Path("/ingresses").Methods("GET").HandlerFunc(resources.ListIngresses(ctx, appCtx))
	w.httpServer.Handler.(*mux.Router).Path("/namespaces").Methods("GET").HandlerFunc(resources.ListNamespaces(ctx, appCtx))
	// POST Requests
	w.httpServer.Handler.(*mux.Router).Path("/namespaces").Methods("POST").HandlerFunc(resources.CreateNamespaces(ctx, appCtx))
	w.httpServer.Handler.(*mux.Router).Path("/apps/mysql").Methods("POST").HandlerFunc(apps.CreateMySQL(ctx, appCtx))
	// DELETE Requests
	w.httpServer.Handler.(*mux.Router).Path("/namespaces").Methods("DELETE").HandlerFunc(resources.DeleteNamespaces(ctx, appCtx))
}
