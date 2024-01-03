package config

import (
	model "Vnc_Server/model/vnc"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

var DB *gorm.DB

func InitMySQL() {
	username := "root"           //账号
	password := "Wsczkmucdis251" //密码
	host := "fallingcreams.top"  //数据库地址，可以是Ip或者域名
	port := 3306                 //数据库端口
	Dbname := "vnc_traning"      //数据库名
	timeout := "10s"             //连接超时，10秒

	// root:root@tcp(127.0.0.1:3306)/insure?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 单数表名
			NoLowerCase:   false, // 关闭小写转换
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功 初始化全局变量DB
	log.Println("连接数据库成功")
	DB = db
	// 注册数据库表
	RegisterTables()
}

// RegisterTables 注册数据库表专用
func RegisterTables() {
	err := DB.AutoMigrate(
		// 系统平台所有的数据库表
		model.UserInfo{},
		model.ContentInfo{},
		model.VmInfo{},
	)
	if err != nil {
		log.Println("注册数据库表失败")
		os.Exit(0)
	}
	log.Println("注册数据库表成功")

}
