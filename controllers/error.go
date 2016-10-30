package controllers

// ErrorController response the system error
type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.Failure(404, "资源不存在或者请求方法错误")
}

func (c *ErrorController) Error500() {
	c.Failure(500, "系统错误")
}
func (c *ErrorController) Error501() {
	c.Failure(501, "系统错误")
}
