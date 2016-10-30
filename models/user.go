package models

import "github.com/dremygit/xwindy-lite/utils"

// User is the model of database
type User struct {
	Sno       string `orm:"size(10);pk" json:"sno"`
	Nickname  string `orm:"size(32)" json:"nickname"`
	Password  string `orm:"size(32)" json:"password"`
	Phone     string `orm:"size(11)" json:"phone"`
	Email     string `orm:"size(50)" json:"email"`
	AvatarURL string `orm:"size(50);column(avatar_url)" json:"avatar_url"`
}

// UserInfo is the resource of User
type UserInfo struct {
	Sno       string `json:"sno"`
	Nickname  string `json:"nickname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

// CreateUserPayload is the request payload when create a new user
type CreateUserPayload struct {
	Sno         string `json:"sno" required:"true"`
	Nickname    string `json:"nickname" required:"true"`
	Password    string `json:"password" required:"true"`
	EASPassword string `json:"eas_password"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	AvatarURL   string `json:"avatar_url"`
}

// UpdateUserPayload is the request payload when update user info
type UpdateUserPayload struct {
	Nickname  string `json:"nickname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

// ResetPasswordPayload is the request payload when reset password
type ResetPasswordPayload struct {
	NewPassword string `json:"new_password" required:"true"`
	OldPassword string `json:"old_password"`
	EASPassword string `json:"eas_password"`
}

// AuthorizationPayload is the request payload when authorize
type AuthorizationPayload struct {
	Sno      string `json:"sno"`
	Password string `json:"password"`
}

// GetBySno to get user by sno
func (user *User) GetBySno(sno string) error {
	err := o.QueryTable("user").Filter("sno", sno).One(user)
	if err != nil {
		return err
	}
	return nil
}

// GetBySnoAndPassword to get User by sno and password
func (user *User) GetBySnoAndPassword(sno string, password string) error {
	err := o.QueryTable("user").Filter("sno", sno).Filter("password", password).One(user)
	if err != nil {
		return err
	}
	return nil
}

// CreateFrom to create user to database from payload
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

// UpdateBy to update user info by payload
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

// UpdatePassword to update user password
func (user *User) UpdatePassword(password string) error {
	user.Password = password
	if _, err := o.Update(user, "password"); err != nil {
		return err
	}
	return nil
}
