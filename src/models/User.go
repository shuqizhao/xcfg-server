package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id            int
	Name          string
	Pwd           string
	IsDeleted     int
	IsEnabled     int
	CreateTime    string
	CreateBy      int
	UpdateTime    string
	UpdateBy      int
	LastLoginTime string
}

type UserInputViewModel struct {
	Id            int
	Name          string
	Pwd           string
	IsDeleted     int
	IsEnabled     string
	CreateTime    string
	CreateTime1   string
	CreateBy      int
	UpdateTime    string
	UpdateBy      string
	LastLoginTime string
}

type UserLoginViewModel struct {
	Un string `json:"un"`
	Pwd  string `json:"pwd"`
}

type UserModViewModel struct {
	OldPwd string
	Pwd    string
}

type UserRoleViewModel struct {
	Id        string
	Name      string
	Roles     string
	IsEnabled int
}

func Login(un string, pwd string) User {
	user := User{}
	o := orm.NewOrm()
	qs := o.QueryTable("User").Filter("is_deleted", "0").Filter("is_enabled", "1").Filter("name", un).Filter("pwd", pwd)
	qs.One(&user)
	return user
}

func IsExistsUserName(un string, id int) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("User").Exclude("id", id).Filter("is_deleted", "0").Filter("is_enabled", "1").Filter("name__iexact", un)
	return qs.Exist()
}

func GetUsersByPage(offset int64, length int64, usr UserInputViewModel) ([]*User, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable("User").Filter("is_deleted", "0")
	if usr.Name != "" {
		qs = qs.Filter("name__icontains", usr.Name)
	}
	if usr.IsEnabled != "" && usr.IsEnabled != "2" {
		qs = qs.Filter("is_enabled", usr.IsEnabled)
	}
	if usr.CreateTime != "" {
		qs = qs.Filter("create_time__gte", usr.CreateTime)
	}
	if usr.CreateTime1 != "" {
		qs = qs.Filter("create_time__lte", usr.CreateTime1)
	}

	var users []*User
	qs.Limit(length, offset).All(&users)
	count, _ := qs.Count()
	return users, count
}

func EnableUser(ids []string) {
	o := orm.NewOrm()
	for _, v := range ids {
		i, _ := strconv.Atoi(v)
		user := User{Id: i, IsEnabled: 1}
		o.Update(&user, "IsEnabled")
	}

}
func DisableUser(ids []string) {
	o := orm.NewOrm()
	for _, v := range ids {
		i, _ := strconv.Atoi(v)
		user := User{Id: i, IsEnabled: 0}
		o.Update(&user, "IsEnabled")
	}
}

func AddUser(name string, pwd string) int64 {
	o := orm.NewOrm()
	var user User
	user.Name = name
	user.Pwd = pwd
	user.IsDeleted = 0
	user.IsEnabled = 1
	user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	user.UpdateTime = user.CreateTime
	user.LastLoginTime = user.CreateTime
	fmt.Println(user.CreateTime)
	id, err := o.Insert(&user)
	if err == nil {
		return id
	} else {
		return 0
	}
}

func GetUser(id int) User {
	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Id, user.Name)
	}
	return user
}

func UpdatePwd(id int, pwd string) bool {
	o := orm.NewOrm()
	user := User{Id: id}
	//err := o.Read(&user)
	if o.Read(&user) == nil {
		user.Pwd = pwd
		if _, err := o.Update(&user, "pwd"); err == nil {
			return true
		}
	}
	return false
}
