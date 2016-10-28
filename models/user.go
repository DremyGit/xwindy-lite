package models

import "github.com/dremygit/xwindy-lite/utils"

// User 用户表模型
type User struct {
	Sno       string `orm:"size(10);pk" json:"sno"`
	Nickname  string `orm:"size(32)" json:"nickname"`
	Password  string `orm:"size(32)" json:"password"`
	Phone     string `orm:"size(11)" json:"phone"`
	Email     string `orm:"size(50)" json:"email"`
	AvatarURL string `orm:"size(50);column(avatar_url)" json:"avatar_url"`
}

type UserInfo struct {
	Sno       string `json:"sno"`
	Nickname  string `json:"nickname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

type CreateUserPayload struct {
	Sno         string `json:"sno" required:"true"`
	Nickname    string `json:"nickname" required:"true"`
	Password    string `json:"password" required:"true"`
	EASPassword string `json:"eas_password"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	AvatarURL   string `json:"avatar_url"`
}

type UpdateUserPayload struct {
	Nickname  string `json:"nickname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

type ResetPasswordPayload struct {
	NewPassword string `json:"new_password" required:"true"`
	OldPassword string `json:"old_password"`
	EASPassword string `json:"eas_password"`
}

func (user *User) GetBySno(sno string) error {
	err := o.QueryTable("user").Filter("sno", sno).One(user)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) GetBySnoAndPassword(sno string, password string) error {
	err := o.QueryTable("user").Filter("sno", sno).Filter("password", password).One(user)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) CreateFrom(payload map[string]interface{}) error {
	if err := utils.CopyFromMap(user, payload); err != nil {
		return err
	}

	_, err := o.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) UpdateBy(payload map[string]interface{}) error {
	err := utils.CopyFromMap(user, payload)
	if err != nil {
		return err
	}

	if _, err := o.Update(user); err != nil {
		return err
	}
	return nil
}

func (user *User) UpdatePassword(password string) error {
	user.Password = password
	if _, err := o.Update(user, "password"); err != nil {
		return err
	}
	return nil
}
