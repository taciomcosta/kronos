package rest

import "github.com/julienschmidt/httprouter"

// NewRouter creates a new router with endpoints set
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.POST("/jobs", CreateJob)
	router.GET("/jobs", FindJobs)
	router.GET("/jobs/:name", DescribeJob)
	router.PUT("/jobs/:name", UpdateJobStatus)
	router.DELETE("/jobs/:name", DeleteJob)
	router.GET("/executions", FindExecutions)
	router.POST("/notifiers", CreateNotifier)
	router.GET("/notifiers", FindNotifiers)
	router.GET("/notifiers/:name", DescribeNotifier)
	router.DELETE("/notifiers/:name", DeleteNotifier)
	router.POST("/assignments", CreateAssignment)
	router.DELETE("/assignments", DeleteAssignment)
	return router
}
