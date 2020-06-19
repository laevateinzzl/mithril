package model

import (
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	gorm.Model
	Account        string
	PasswordDigest string
	Username       string
	Status         string
	Avatar         string `gorm:"size:1000"`
}
type Claims struct {
	ID     int    `gorm:"primary_key" json:"claim_id"`
	AuthID int    `json:"user_id"`
	Type   string `json:"type"`
	Value  string `json:"value"`
}

type UserCheck struct {
	Account    string
	UserClaims []Claims
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	//Inactive string = "inactive"
	//// Suspend 被封禁用户
	//Suspend string = "suspend"
)

func (user *User) AvatarURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(user.Avatar, oss.HTTPGet, 600)
	return signedGetURL
}

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

func GetUserClaims(Account string) (claims []Claims) {
	var user User
	DB.Where("account = ?", Account).First(&user)
	DB.Where("user_id = ?", user.Account).Find(&claims)
	return
}

func CheckAuth(account, password string) bool {
	var user User

	if err := DB.Where("account = ?", account).First(&user).Error; err != nil {
		return false
	}

	if user.CheckPassword(password) == false {
		return false
	}

	return true
}
func GetUserInfo(account string) (user User) {
	DB.Where("account = ?", account).First(&user)
	return
}
