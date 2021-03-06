package controllers

import (
	"github.com/astaxie/beego/config"
	. "github.com/bitly/go-simplejson"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/copier"

	"net/http"
	"net/url"

	"github.com/dremygit/xwindy-lite/models"
	"github.com/dremygit/xwindy-lite/utils"
)

// UserController handle /users
type UserController struct {
	BaseController
}

// GetBySno to get user info by sno
// @Title GetBySno
// @Description GetBySno
// @Param sno path string true 学号
// @Param access_token query string false access_token
// @Success 200 {object} models.UserInfo
// @Failure 400 Request Error
// @Failure 403 Forbidden
// @Failure 404 Not found
// @router /:sno [get]
func (c *UserController) GetBySno() {

	sno := c.Ctx.Input.Param(":sno")

	token, err := c.ParseToken()
	if err != nil {
		c.Failure(401, err.Error())
		return
	}
	tokenSno, ok := token["sno"].(string)
	if !ok || tokenSno != sno {
		c.Failure(403, "拒绝访问")
		return
	}

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

// UpdateInfo to update user info
// @Title UpdateUserInfo
// @Description Update user info
// @Param sno path string true 学号
// @Param body body models.UpdateUserPayload true 修改信息
// @router /:sno [put]
func (c *UserController) UpdateInfo() {
	// var payload models.UpdateUserPayload

	sno := c.Ctx.Input.Param(":sno")

	token, err := c.ParseToken()
	if err != nil {
		c.Failure(401, err.Error())
		return
	}
	tokenSno, ok := token["sno"].(string)
	if !ok || tokenSno != sno {
		c.Failure(403, "拒绝访问")
		return
	}

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

// ResetPassword to reset user's password
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

	newPassword, ok := payload["new_password"].(string)
	if !ok || len(newPassword) == 0 {
		c.Failure(400, "新密码不能为空")
		return
	}

	token, err := c.ParseToken()
	if err != nil {

		if err := userDB.GetBySno(sno); err != nil {
			c.Failure(404, "此用户不存在")
			return
		}

		easPassword, ok := payload["eas_password"].(string)
		if !ok || len(easPassword) == 0 {
			c.Failure(400, "教务系统密码不能为空")
			return
		}

		checked, err := checkEASAccount(sno, easPassword)
		if err != nil {
			c.Failure(500, "教务系统连接错误")
			return
		}

		if !checked {
			c.Failure(403, "教务系统验证失败")
			return
		}

	} else {

		oldPassword, ok := payload["old_password"].(string)
		if !ok || len(oldPassword) == 0 {
			c.Failure(400, "旧密码不能为空")
			return
		}

		tokenSno, ok := token["sno"].(string)
		if !ok || tokenSno != sno {
			c.Failure(403, "拒绝访问")
			return
		}

		if err := userDB.GetBySnoAndPassword(sno, oldPassword); err != nil {
			c.Failure(404, "旧密码错误")
			return
		}

	}

	if err := userDB.UpdatePassword(newPassword); err != nil {
		c.Failure(500, err.Error())
		return
	}

	c.Success(201, true)
}

// Authorize the login info and response token
// @Title Authorization
// @Description Authorization
// @Param body body models.AuthorizationPayload true Body
// @router /authorization [post]
func (c *UserController) Authorize() {

	payload, err := c.ParsePayload()
	if err != nil {
		c.Failure(400, err.Error())
		return
	}

	sno, ok := payload["sno"].(string)
	if !ok || len(sno) == 0 {
		c.Failure(400, "学号不能为空")
		return
	}

	password, ok := payload["password"].(string)
	if !ok || len(password) == 0 {
		c.Failure(400, "密码不能为空")
		return
	}

	var userDB models.User

	if err := userDB.GetBySnoAndPassword(sno, password); err != nil {
		c.Failure(404, "用户名或密码错误")
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["sno"] = sno
	tokenString, err := token.SignedString([]byte("test"))
	if err != nil {
		c.Failure(500, err.Error())
		return
	}

	c.Success(201, map[string]string{
		"token": tokenString,
	})
}

var easHost string

func checkEASAccount(sno, password string) (bool, error) {

	form := url.Values{
		"UserStyle": []string{"student"},
		"user":      []string{sno},
		"password":  []string{password},
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := client.PostForm(easHost+"/pass.asp", form)
	defer req.Body.Close()
	if err != nil {
		return false, err
	}

	if req.StatusCode == 302 {
		return true, nil
	}

	return false, nil
}

func init() {

	iniconf, err := config.NewConfig("ini", "./conf/app.conf")
	if err != nil {
		panic("Config file not found")
	}

	easHost = iniconf.String("app::eashost")
}
