package main

import "net/http"

func (c *Config) SessionLoad(next http.Handler) http.Handler {
	return c.Session.LoadAndSave(next)
}
