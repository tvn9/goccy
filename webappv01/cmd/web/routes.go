package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (c *Config) routes() http.Handler {
	// create a router
	mux := chi.NewRouter()

	// set up middleware
	mux.Use(middleware.Recoverer)
	mux.Use(c.SessionLoad)

	// define application routes
	mux.Get("/", c.HomePage)
	mux.Get("/login", c.LoginPage)
	mux.Post("/login", c.PostLoginPage)
	mux.Get("/logout", c.Logout)
	mux.Get("/register", c.RegisterPage)
	mux.Post("/register", c.PostRegisterPage)
	mux.Get("/activateAccount", c.ActivateAccount)

	mux.Get("test-email", func(w http.ResponseWriter, r *http.Request) {

		m := Mail{
			Domain:      "127.0.0.1",
			Host:        "127.0.0.1",
			Port:        1025,
			Encryption:  "none",
			FromAddress: "info@company.com",
			FromName:    "info",
			ErrorChan:   make(chan error),
		}

		msg := Message{
			To:      "me@here.com",
			Subject: "Test email",
			Data:    "Hello, World!",
		}

		m.sendMail(msg, make(chan error))
	})

	return mux
}
