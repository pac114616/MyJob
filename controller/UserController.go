package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"text/common"
	"text/model"
	"text/util"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
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
		name = util.RandomString(10)
	}
	//判断手机号是否存在
	log.Println(name,telephone,password)
	if isTelephoneExist(DB,telephone){
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422, "msg":"用户已经存在"})
		return
	}
	//创建用户加密密码
	hasedPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"code":500,"msg":"加密错误"})
		return
	}
	newUser :=model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hasedPassword),//加密后的密码
	}
	DB.Create(&newUser)
	//返回结果
	ctx.JSON(200, gin.H{
		"code":200,
		"msg": "注册成功",
	})
}

func Login(ctx *gin.Context){
	DB := common.GetDB()
	//获取参数
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

	//判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?",telephone).First(&user)
	if user.ID == 0{
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"用户不存在"})
		return
	}

	//判断密码是否正确 第一个参数是原始的加密后的，第二个是对比的
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"code":400,"msg":"密码错误"})
		return
	}
	//密码验证通过 发放token
	token :="11"
	//返回结果
	ctx.JSON(200,gin.H{
		"code":200,
		"data":gin.H{"token":token},
		"msg":"登陆成功",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	// SELECT * FROM diversion_card WHERE state = 1 limit 1;
	db.Where("telephone = ?",telephone).First(&user)
	//fmt.Print(reflect.TypeOf(db.Where("telephone = ?",telephone).First(&user)))
	if user.ID !=0 {
		return true
	}
	return false
}