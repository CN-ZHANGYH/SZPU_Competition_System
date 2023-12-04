package model

type UserInfo struct {
	Id       int    `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:'主键ID'"`
	Username string `gorm:"column:username;NOT NULL;comment:'用户名称'"`
	Password string `gorm:"column:password;NOT NULL;comment:'用户密码'"`
	Email    string `gorm:"column:email;NOT NULL;comment:'邮箱地址'"`
	Phone    string `gorm:"column:phone;NOT NULL;comment:'手机号'"`
	Status   int    `gorm:"column:status;NOT NULL;comment:'状态'"`
	Role     int    `gorm:"column:role;NOT NULL;comment:'角色'"`
}
