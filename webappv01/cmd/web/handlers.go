package main

import "net/http"

func (c *Config) HomePage(w http.ResponseWriter, r *http.Request) {
	c.render(w, r, "home.page.html", nil)
}

func (c *Config) LoginPage(w http.ResponseWriter, r *http.Request) {
	c.render(w, r, "login.page.html", nil)
}

func (c *Config) PostLoginPage(w http.ResponseWriter, r *http.Request) {

}

func (c *Config) Logout(w http.ResponseWriter, r *http.Request) {

}

func (c *Config) RegisterPage(w http.ResponseWriter, r *http.Request) {

}

func (c *Config) PostRegisterPage(w http.ResponseWriter, r *http.Request) {

}

func (c *Config) ActivateAccount(w http.ResponseWriter, r *http.Request) {

}
