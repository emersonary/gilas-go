package main

import (
	"fmt"
	"net/http"

	"gitgub.com/emersonary/gilasw/go/configs"
	"gitgub.com/emersonary/gilasw/go/internal/seeders"
	"gitgub.com/emersonary/gilasw/go/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func defaultHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
func starthttpserver(db *gorm.DB, cfg *(configs.Conf)) {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(defaultHeaders)
	r.Use(middleware.Recoverer)

	dbhandler := handlers.NewDBHandler(db)

	r.Get("/users/", handlers.GetUsers)
	r.Get("/categories/", handlers.GetCategories)
	r.Get("/channels/", handlers.GetChannels)

	r.Route("/messages/", func(r chi.Router) {

		r.Options("/*", handlers.GetOptions)
		r.Post("/create", dbhandler.CreateMessage)
		r.Get("/lastmessages", dbhandler.GetLastMessages)
		r.Get("/lastnotifications", dbhandler.GetLastMessagesNotifications)

	})

	fmt.Printf("Starting Web Server at port %v\n", cfg.WebServerPort)

	http.ListenAndServe(":8000", r)

}

func main() {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("gilasw.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	seeders.LoadOrSeed(db)
	starthttpserver(db, configs)

}
