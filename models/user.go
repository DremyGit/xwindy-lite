package models

// User 用户表模型
type User struct {
	Sno       string `orm:"size(10);pk"`
	Nickname  string `orm:"size(32)"`
	Password  string `orm:"size(32)"`
	Phone     string `orm:"size(11)"`
	Email     string `orm:"size(50)"`
	AvatarURL string `orm:"size(50)"`
}

func (user *User) GetBySno(sno string) error {
	err := o.QueryTable("user").Filter("sno", sno).One(user)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) Create() error {
	_, err := o.Insert(&User{
		Sno:      user.Sno,
		Password: user.Password,
		Phone:    user.Phone,
	})
	if err != nil {
		return err
	}
	return nil
}
