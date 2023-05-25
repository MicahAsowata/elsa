package main

func TaskRoutes(a *application) {
	taskRoute := a.app.Group("/tarea")
	taskRoute.Get("/", a.TaskIndex)
	taskRoute.Get("/new", a.TaskNew)
	taskRoute.Post("/", a.TaskCreate)
	taskRoute.Get("/:id", a.TaskShow)
	taskRoute.Get("/:id/edit", a.TaskEdit)
	taskRoute.Post("/:id", a.TaskUpdate)
	taskRoute.Get("/:id/delete", a.TaskDestroy)
}
