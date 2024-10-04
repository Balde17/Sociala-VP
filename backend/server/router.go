package server

import (
	"fmt"
	"net/http"
)

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Use(middleware Middleware) {
	r.middlewares = append(r.middlewares, middleware)
}

func (r *Router) AddRoute(method, pattern string, handler RouteHandler) {
	r.routes = append(r.routes, route{method, pattern, handler})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var pathMatch bool

	for _, route := range r.routes {
		if req.URL.Path == route.pattern {
			pathMatch = true
			if req.Method == route.method {
				handler := route.handler
				for i := len(r.middlewares) - 1; i >= 0; i-- {
					handler = r.middlewares[i](handler)
				}
				handler(w, req)
				return
			}
		}
	}

	if pathMatch {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	http.NotFound(w, req)
}

func (r *Router) ServeStatic(path, directory string) {

	fileServer := http.FileServer(http.Dir(directory))
	handler := http.StripPrefix(path, fileServer)

	r.AddRoute("GET", path, func(w http.ResponseWriter, req *http.Request) {
		handler.ServeHTTP(w, req)
	})
}

func (r *Router) ConfigureRoutes() {
	r.Use(LoggerMiddleware)
	r.Use(CORSMiddleware)

	for _, endpoint := range Endpoints {
		for _, method := range endpoint.Method {
			r.AddRoute(method, endpoint.Path, endpoint.Handler)
		}
	}

	for _, endpoint := range Protected {
		for _, method := range endpoint.Method {
			r.AddRoute(method, endpoint.Path, AuthenticationMiddleware(endpoint.Handler))
		}
	}

}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "user.html")
}

func (r *Router) StartServer(port string) error {
	// Configure routes and middlewares
	r.ConfigureRoutes()

	// Start the server
	fmt.Printf("Server listening on http://localhost:%v\n", port)
	return http.ListenAndServe(":"+port, r)
}
