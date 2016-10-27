package controllers

type UserController struct {
	BaseController
}

func (c *UserController) Get() {
	c.Ctx.WriteString("Hello")
}
