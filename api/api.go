package api

import (
	"net/http"
	"new-proto/api/handlers"
	"new-proto/api/repository"
	"new-proto/api/services"

	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New() *http.Server {
	r := chi.NewRouter()

	// global middlware
	r.Use(Tracer)
	r.Use(middleware.Recoverer)

	// repositories
	taskRepo := repository.NewTaskRepository()

	//services
	taskSvc := services.NewTaskService(taskRepo)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})

	// routes
	r.Route("/api/v1", func(v1 chi.Router) {
		v1.Mount("/tasks", handlers.NewTaskHandler(taskSvc))
		v1.Mount("/env", handlers.NewEnvHandler())
	})

	return &http.Server{
		Handler: r,
		Addr:    ":8888",
	}

}

func Tracer(next http.Handler) http.Handler {
	return xray.Handler(xray.NewFixedSegmentNamer("mgtApp"), next)
}
