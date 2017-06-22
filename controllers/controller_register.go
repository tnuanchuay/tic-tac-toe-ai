package controllers

import "net/http"

type Controller_register struct {
	GET	map[string]func(w http.ResponseWriter, r *http.Request)
	POST	map[string]func(w http.ResponseWriter, r *http.Request)
	PUT	map[string]func(w http.ResponseWriter, r *http.Request)
	DELETE map[string]func(w http.ResponseWriter, r *http.Request)
}

func (c * Controller_register)Get(p string, f func(w http.ResponseWriter, r *http.Request)){
	c.GET[p] = f
}

func (c * Controller_register)Post(p string, f func(w http.ResponseWriter, r *http.Request)){
	c.POST[p] = f
}

func (c * Controller_register)Put(p string, f func(w http.ResponseWriter, r *http.Request)){
	c.PUT[p] = f
}

func (c * Controller_register)Delete(p string, f func(w http.ResponseWriter, r *http.Request)){
	c.DELETE[p] = f
}

func (c *Controller_register) New() Controller_register{

	c.GET = make(map[string]func(w http.ResponseWriter, r *http.Request))
	c.POST = make(map[string]func(w http.ResponseWriter, r *http.Request))
	c.PUT =make(map[string]func(w http.ResponseWriter, r *http.Request))
	c.DELETE = make(map[string]func(w http.ResponseWriter, r *http.Request))
	return *c
}