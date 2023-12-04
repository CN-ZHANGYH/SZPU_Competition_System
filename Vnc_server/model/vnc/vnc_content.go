package model

type ContentInfo struct {
	Id      int    `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:'主键ID'"`
	Name    string `gorm:"column:name;NOT NULL;comment:'主机名称'"`
	Type    string `gorm:"column:type;NOT NULL;comment:'类型'"`
	VmList  string `gorm:"column:vmList;NOT NULL;comment:'主机列表'"`
	Content string `gorm:"column:content;NOT NULL;comment:'内容'"`
}
