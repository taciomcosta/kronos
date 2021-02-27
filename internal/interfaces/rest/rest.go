package rest

import "github.com/julienschmidt/httprouter"

// NewRouter creates a new router with endpoints set
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.POST("/jobs", CreateJob)
	router.GET("/jobs", FindJobs)
	router.DELETE("/jobs/:name", DeleteJob)
	router.GET("/executions", FindExecutions)
	return router
}
