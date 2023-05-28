package main

func routes(b *base) error {
	b.base.Static("/static", "./ui/static")
	b.base.Get("/", b.Index)
	b.base.Get("/new", b.New)
	b.base.Post("/", b.Create)
	b.base.Get("/:id", b.Show)
	b.base.Get("/:id/edit", b.Edit)
	b.base.Post("/:id", b.Update)
	b.base.Get("/:id/delete", b.Destroy)
	return nil
}
