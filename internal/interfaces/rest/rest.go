package rest

import "github.com/julienschmidt/httprouter"

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.POST("/jobs", CreateJob)
	router.GET("/jobs", FindJobs)
	return router
}
