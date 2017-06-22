package controllers

func Init() Controller_register{
	route := (&Controller_register{}).New()
	staticInit()

	route.Get("/assets/{file}",  static_controller)

	return route
}
