package models

// User 用户表模型
type User struct {
	Sno       string `orm:"size(10)"`
	Username  string `orm:"size(16)"`
	Nickname  string `orm:"size(32)"`
	Password  string `orm:"size(32)"`
	Phone     string `orm:"size(11)"`
	Email     string `orm:"size(50)"`
	AvatarURL string `orm:"size(50)"`
}

func (u *User) GetBySno(sno string) (user *User, err error) {

	err = o.QueryTable("user").Filter("sno", sno).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil

}
