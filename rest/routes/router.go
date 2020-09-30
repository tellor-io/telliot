package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/tellor-io/TellorMiner/common"
	"github.com/tellor-io/TellorMiner/db"
)

//Handler handle incoming http requests and produces a response code and JSON payload string
type Handler interface {
	Incoming(ctx context.Context, req *http.Request) (int, string)
}

//Router holds all url to handler mappings and routes incoming http calls to handlers
type Router struct {
	DB     db.DB
	routes map[string]Handler
}

//NewRouter creates a new router instance using the given DB instance
func NewRouter(DB db.DB) *Router {
	return &Router{DB, make(map[string]Handler)}
}

//Default http handler callback which will route to appropriate handler internally
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	url := strings.ToLower(req.URL.String())
	h := r.routes[url]
	if h == nil {
		w.WriteHeader(404)
		return
	}
	log.Println("Incoming request", url)
	if e := recover(); e != nil {
		fmt.Printf("Problem with controller %s: %v\n", url, e)
		fmt.Fprintf(w, "Problem handling request")
		return
	}
	ctx := context.WithValue(context.Background(), common.DBContextKey, r.DB)
	code, payload := h.Incoming(ctx, req)
	w.WriteHeader(code)
	fmt.Fprintf(w, payload)
}

//AddRoute adds handler for the given url
func (r *Router) AddRoute(path string, handler Handler) {
	r.routes[strings.ToLower(path)] = handler
}
