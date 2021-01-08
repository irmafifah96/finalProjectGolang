// Package controller loads the routes for each of the controllers.
package controller

import (
	// "finalProject/controller/about"
	"finalProject/controller/debug"
	"finalProject/controller/home"
	"finalProject/controller/login"
	//"finalProject/controller/notepad"
	"finalProject/controller/register"
	"finalProject/controller/static"
	"finalProject/controller/status"
	"finalProject/controller/article"
)

// LoadRoutes loads the routes for each of the controllers.
func LoadRoutes() {
	// about.Load()
	debug.Load()
	register.Load()
	login.Load()
	home.Load()
	static.Load()
	status.Load()
	// notepad.Load()
	article.Load()
}
