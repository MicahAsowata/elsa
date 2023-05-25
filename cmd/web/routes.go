package main

func TaskRoutes(a *application) {
	taskRoute := a.app.Group("/tarea")
	taskRoute.Get("/", TaskIndex)
	taskRoute.Get("/new", TaskNew)
	taskRoute.Post("/", a.TaskCreate)
	taskRoute.Get("/:id", TaskShow)
	taskRoute.Get("/:id/edit", TaskEdit)
	taskRoute.Post("/:id", TaskUpdate)
	taskRoute.Get("/:id/delete", TaskDestroy)
}
