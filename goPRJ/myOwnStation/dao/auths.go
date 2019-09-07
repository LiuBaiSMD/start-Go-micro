package dao

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"
)

var sqlDB *gorm.DB

type Auth struct {
	Id       string `gorm:"default:'peter'"`
	Password string
}

func InitMysql(driverName, mysqlURL string) (*gorm.DB) {
	var err error
	sqlDB, err = gorm.Open(driverName, mysqlURL)
	if err != nil {
		log.Log(err)
	}
	createMyTable()
	return sqlDB
}

func createMyTable() {
	//初始化Auth,并且插入一个成员
	auth := Auth{
		Id:       "123",
		Password: "345",
	}
	res := Auth{}
	sqlDB.First(&auth).Scan(&res)
	if res.Id == ""{
		sqlDB.CreateTable(&Auth{})
	}
}

func QueryUserIdPass(userId, password string) error {
	userAuth := Auth{
		Id: userId,
	}
	res := Auth{}
	sqlDB.Find(&userAuth).Scan(&res)
	if res.Password == "" {
		return errors.New("无此用户，请进行注册")
	} else if res.Password != password {
		return errors.New("密码错误，请重试")
	}
	return nil
}

func RegisterUserIdPass(userId, password string) error {
	//0 操作成功  1 密码不合格 2 账号已被注册
	//var users []Auth
	userAuth := Auth{
		Id: userId,
	}
	res := Auth{}
	sqlDB.Find(&userAuth).Select("id").Scan(&res)
	if err := checkPassword(password); err != nil {
		return err
	}
	if res.Id != "" {
		return errors.New("此账户已注册，请重新输入")
	}
	userAuth.Password = password
	sqlDB.Create(&userAuth)
	return nil
}

func ChangePWD(userId, password, newPassword string) error {
	if password == newPassword {
		return errors.New("新老密码一致，请重新填写！")
	}
	//检查用户是否存在，不存在不让改密码
	userAuth := Auth{
		Id: userId,
	}
	res := Auth{}
	sqlDB.Find(&userAuth).Select("id").Scan(&res)
	if err := checkPassword(password); err != nil {
		return err
	}
	if res.Id == "" {
		return errors.New("此账户未注册，请重新输入")
	}
	if err := QueryUserIdPass(userId, password); err != nil {
		return err
	}
	res.Password = newPassword
	sqlDB.Model(&res).Update("password", newPassword)
	return nil
}

func checkPassword(password string) error {
	//检查密码的工具
	if len(password) < 8 {
		fmt.Println("密码过于短小，请重设密码!（8个字符起！）")
		return errors.New("密码过于短小，请重设密码!（8个字符起！）")
	}
	return nil
}

func checkOKs(oks ...bool) bool{
	//检查oks是否全为true
	for _, v := range oks{
		if !v{
			return false
		}
	}
	return true
}
