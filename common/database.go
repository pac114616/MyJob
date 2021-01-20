package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"text/model"
)

var DB *gorm.DB
func InitDB() *gorm.DB{
	//打开sql调试
	//DB.LogMode(true)
	port :="3306"
	host :="localhost"
	database :="test"
	password :="chen360219"
	username :="root"
	charset :="utf8"
	args :=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",//相当于connString
		username,
		password,
		host,
		port,
		database,
		charset)
	/*constr := fmt.Sprintf("username=%s;password=%s;sever=%s;port=%s;database=%s;charset=%s",
	username,
	password,
	server,
	port,
	database,
	charset)*/
	//constr :=fmt.Sprintf("Provider=SQLOLEDB;Data Source=DESKTOP-6TJBT2R;Initial Catalog=testOne;user id=sa;password=chen360219")
	/*connString := fmt.Sprintf("@server=%s:port=%s/database=%s;username=%s;password=%s;charset=%s,driver={sql server}",
	server, port, database, username, password,charset,mssql_data_source)*/
	db,err := gorm.Open("mysql",args)
	//db, err := gorm.Open(mssql.Open("test.db"), &gorm.Config{})
	//db, err := gorm.Op
	//db, err := sql.Open("mssql", constr)
	if err != nil{
		panic("failed to connect database, err:"+err.Error())
	}
	db.AutoMigrate(&model.User{})//自动创建数据表
	DB = db
	return db
}
func GetDB()*gorm.DB{
	return DB
}