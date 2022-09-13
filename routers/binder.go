package routers

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/tapiaw38/irrigation-api/server"
)

func BinderRoutes(s server.Server, router *mux.Router) {
	// mount the routers
	mount(router, "/users", UserRoutes(s))
	mount(router, "/producers", ProducerRoutes(s))
	mount(router, "/productions", ProductionRoutes(s))
	mount(router, "/sections", SectionRoutes())
	mount(router, "/intakes", IntakeRoutes())
	mount(router, "/turns", TurnRoutes())
	mount(router, "/ws", WebSocketRoutes(s))
}

// mount is a helper function to mount a router to a path
func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
