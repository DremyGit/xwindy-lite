package controllers

import (
	"encoding/json"

	"github.com/dremygit/xwindy-lite/models"
	"github.com/jinzhu/copier"
)

// UserController handle /users
type UserController struct {
	BaseController
}

type UserInfo struct {
	Sno       string `json:"sno"`
	Nickname  string `json:"nickname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

// GetBySno get user info by sno
// @Title GetBySno
// @Description Add user (need Admin)
// @Param body body models.User true 用户
// @Success 201 {object} models.User
// @Failure 400 Request Error
// @Failure 403 Forbidden
// @router /:userId [get]
func (c *UserController) GetBySno() {
	var userDB models.User
	if err := userDB.GetBySno("123"); err != nil {
		c.Ctx.WriteString("Not Found")
		return
	}

	var userInfo UserInfo
	copier.Copy(&userInfo, &userDB)
	c.Success(200, userInfo)
}

type CreateUserPayload struct {
	Sno         string `json:"sno"`
	Nickname    string `json:"nickname"`
	Password    string `json:"password"`
	EASPassword string `json:"eas_password"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	AvatarURL   string `json:"avatar_url"`
}

// CreateUser create new user
// @Title CreateUser
// @router / [post]
func (c *UserController) CreateUser() {

	var payload CreateUserPayload
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload); err != nil {
		c.Failure(400, err.Error())
		return
	}

	var userDB models.User
	copier.Copy(&userDB, &payload)
	if err := userDB.Create(); err != nil {
		c.Failure(500, err.Error())
		return
	}

	c.Success(201, userDB)
}
