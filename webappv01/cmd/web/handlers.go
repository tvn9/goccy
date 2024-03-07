package main

import "net/http"

func (c *Config) HomePage(w http.ResponseWriter, r *http.Request) {
	c.render(w, r, "home.page.html", nil)
}

func (c *Config) LoginPage(w http.ResponseWriter, r *http.Request) {
	c.render(w, r, "login.page.html", nil)
}

func (c *Config) PostLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = c.Session.RenewToken(r.Context())

	// parse from post
	err := r.ParseForm()
	if err != nil {
		c.ErrorLog.Println(err)
	}

	// get email and password from form post
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := c.Models.User.GetByEmail(email)
	if err != nil {
		c.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// check password
	validPassword, err := user.PasswordMatches(password)
	if err != nil {
		c.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if !validPassword {
		c.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// okay, so log user in
	c.Session.Put(r.Context(), "userID", user.ID)
	c.Session.Put(r.Context(), "user", user)

	c.Session.Put(r.Context(), "flash", "Successful login!")

	// redirect the user
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (c *Config) Logout(w http.ResponseWriter, r *http.Request) {
	// // clean up session
	_ = c.Session.Destroy(r.Context())
	_ = c.Session.RenewToken(r.Context())

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (c *Config) RegisterPage(w http.ResponseWriter, r *http.Request) {
	c.render(w, r, "register.page.html", nil)
}

func (c *Config) PostRegisterPage(w http.ResponseWriter, r *http.Request) {
	// create a user

	// send and activation email

	// subscribe the user to an account

}

func (c *Config) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	// validate url

	// generate an invoice

	// send an email with attachments

	// send an email with the invoice attached
}
