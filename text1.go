package main
import (
	"fmt"
	"github.com/gin-gonic/gin" //引入gin
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"reflect"

	//"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)
type User struct {
	//User内嵌了gorm.Model，内置了ID、CreatedAt、UpdatedAt、DeletedAt属性，
	//同时Create的时候会自动设置CreatedAt、UpdatedAt，Update的时候会自动更新UpdatedAt
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null;unique"`
	Password string `gorm:"size:255;not null"`
}
func main() {
	db := InitDB()
	defer db.Close()
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		//获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password :=ctx.PostForm("password")
		//数据验证
		if len(telephone) != 11{
			// gin.H 是 map[string]interface{} 的一种快捷方式
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code": 422,"msg": "手机号必须为11位"})
			return
		}
		if len(password) < 6{
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422, "msg":"密码不能少于6位"})
			return
		}

		//如果名称没有传，给一个10位的随机字符串
		if len(name) == 0{
			name = RandomString(10)
		}
		//判断手机号是否存在
		log.Println(name,telephone,password)
		if isTelephoneExist(db,telephone){
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422, "msg":"用户已经存在"})
			return
		}
		//创建用户
		newUser :=User{
			Name: name,
			Telephone: telephone,
			Password: password,
		}
		db.Create(&newUser)
		//返回结果
		ctx.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	// SELECT * FROM diversion_card WHERE state = 1 limit 1;
	db.Where("telephone = ?",telephone).First(&user)
	fmt.Print(reflect.TypeOf(db.Where("telephone = ?",telephone).First(&user)))
	if user.ID !=0 {
		return true
	}
	return false
}
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte,n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
func InitDB() *gorm.DB{
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
	db.AutoMigrate(&User{})//自动创建数据表
	return db
}