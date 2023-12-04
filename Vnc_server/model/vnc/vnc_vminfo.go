package model

type VmInfo struct {
	Id       int    `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:'主键ID'"`
	Name     string `gorm:"column:name;NOT NULL;comment:'主机名称'"`
	Host     string `gorm:"column:host;NOT NULL;comment:'主机地址'"`
	Password string `gorm:"column:password;NOT NULL;comment:'主机VNC密码'"`
	Port     int    `gorm:"column:port;NOT NULL;comment:'端口'"`
	Url      string `gorm:"column:url;NOT NULL;comment:'VNC的URL'"`
}
