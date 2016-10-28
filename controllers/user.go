package controllers

import (
	. "github.com/bitly/go-simplejson"
	"github.com/jinzhu/copier"

	"github.com/dremygit/xwindy-lite/models"
	"github.com/dremygit/xwindy-lite/utils"
)

// UserController handle /users
type UserController struct {
	BaseController
}

// GetBySno get user info by sno
// @Title GetBySno
// @Description GetBySno (need Admin)
// @Param sno path string true 学号
// @Success 200 {object} models.UserInfo
// @Failure 400 Request Error
// @Failure 403 Forbidden
// @Failure 404 Not found
// @router /:sno [get]
func (c *UserController) GetBySno() {

	sno := c.Ctx.Input.Param(":sno")
	var userDB models.User
	if err := userDB.GetBySno(sno); err != nil {
		c.Failure(404, "用户不存在")
		return
	}

	var userInfo models.UserInfo
	copier.Copy(&userInfo, &userDB)
	c.Success(200, userInfo)
}

// CreateUser create new user
// @Title CreateUser
// @Description Create new user
// @Param body body models.CreateUserPayload true 用户
// @Success 201 {object} models.UserInfo
// @Failure 400 Request Error
// @router / [post]
func (c *UserController) CreateUser() {

	js, err := NewJson(c.Ctx.Input.RequestBody)
	if err != nil {
		c.Failure(400, "请求格式错误")
		return
	}

	payload, err := js.Map()
	if err != nil {
		c.Failure(400, "请求格式错误")
		return
	}

	if len(payload["sno"].(string)) != 10 {
		c.Failure(400, "学号错误")
		return
	}

	if len(payload["nickname"].(string)) == 0 || len(payload["password"].(string)) == 0 {
		c.Failure(400, "请将信息填写完整")
		return
	}

	var userDB models.User
	if err := userDB.CreateFrom(payload); err != nil {
		c.Failure(500, err.Error())
		return
	}

	var userInfo models.UserInfo
	copier.Copy(&userInfo, &userDB)
	c.Success(201, userInfo)
}

// UpdateInfo update user info
// @Title UpdateUserInfo
// @Description Update user info
// @Param sno path string true 学号
// @Param body body models.UpdateUserPayload true 修改信息
// @router /:sno [put]
func (c *UserController) UpdateInfo() {
	// var payload models.UpdateUserPayload

	sno := c.Ctx.Input.Param(":sno")
	var userDB models.User
	if err := userDB.GetBySno(sno); err != nil {
		c.Failure(404, "用户不存在")
		return
	}

	payload, err := c.ParsePayload()
	if err != nil {
		c.Failure(400, err.Error())
		return
	}

	filter := []string{"nickname", "phone", "email", "avatar_url"}
	payload = utils.FilterMap(payload, filter)

	if err := userDB.UpdateBy(payload); err != nil {
		c.Failure(500, err.Error())
		return
	}

	var userInfo models.UserInfo
	copier.Copy(&userInfo, &userDB)
	c.Success(201, userInfo)
}

// ResetPassword reset user's password
// @Title ResetPassword
// @Description Reset user's password
// @Param sno path string true 学号
// @Param body body models.ResetPasswordPayload true Body
// @router /:sno/password [put]
func (c *UserController) ResetPassword() {

	payload, err := c.ParsePayload()
	if err != nil {
		c.Failure(400, err.Error())
		return
	}

	var userDB models.User
	sno := c.Ctx.Input.Param(":sno")
	oldPassword, ok := payload["old_password"].(string)
	if !ok || len(oldPassword) == 0 {
		c.Failure(400, "密码不能为空")
		return
	}

	if err := userDB.GetBySnoAndPassword(sno, oldPassword); err != nil {
		c.Failure(404, "旧密码错误")
		return
	}

	newPassword, ok := payload["new_password"].(string)
	if !ok || len(newPassword) == 0 {
		c.Failure(400, "新密码不能为空")
		return
	}

	if err := userDB.UpdatePassword(newPassword); err != nil {
		c.Failure(500, err.Error())
		return
	}

	c.Success(201, true)
}
