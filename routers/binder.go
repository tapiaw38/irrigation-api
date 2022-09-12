package routers

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func BinderRoutes(router *mux.Router) {

	// Mount the routers
	mount(router, "/users", UserRoutes())
	mount(router, "/producers", ProducerRoutes())
	mount(router, "/productions", ProductionRoutes())
	mount(router, "/sections", SectionRoutes())
	mount(router, "/intakes", IntakeRoutes())
	mount(router, "/turns", TurnRoutes())
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
