package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5"
	"net/http"
	"simple-crud/config"
)

type ClientService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type Server struct {
	router      *chi.Mux
	db          *pgx.Conn
	userService ClientService
}

func NewServer(db *pgx.Conn, userService ClientService) *Server {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	return &Server{
		router:      router,
		db:          db,
		userService: userService,
	}
}

func (s *Server) Start(cfg *config.Config) error {
	s.router.Route("/api", func(r chi.Router) {
		r.Put("/create", s.userService.Create)
		r.Post("/read", s.userService.Read)
		r.Patch("/update", s.userService.Update)
		r.Delete("/delete", s.userService.Delete)
	})

	addr := getApiAddr(cfg)
	return http.ListenAndServe(addr, s.router)
}

// Format http server address
func getApiAddr(cfg *config.Config) string {
	return fmt.Sprintf("%s:%s", cfg.ApiPort, cfg.ApiPort)
}
