package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	)

var db *gorm.DB
// 此处使用gorm关联到本地的数据库里（暂时） 后期更换成为链式储存的时候删去此部分

//初始化数据库
func InitSQLite() (err error) {

	db, err = gorm.Open("sqlite3", "shuibian.db")
	if err != nil {
		return
	}

	//db.AutoMigrate(&PatientPer{})
	db.AutoMigrate(&Audit{},&AuditMed{},&PatientInfo{},&PatientPer{},&BackPat{})
	return db.DB().Ping()
}

//关闭数据库
func Close() {
	err := db.Close()
	if err != nil {
		return
	}
}

