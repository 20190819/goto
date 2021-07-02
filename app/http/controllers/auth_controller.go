package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"goto/app/models/user"
	"goto/database/mysql"
	"goto/pkg/utils"
	"goto/pkg/validate"
)

type AuthController struct {
	BaseController
}

func (*AuthController) Register(ctx *gin.Context) {
	userModel := user.User{}
	validateMap, validateNameMap := validate.MapDataForStruct(userModel)
	validateMap["password_confirm"] = []string{"required", "min:6", "max:20"}
	if ok, msg := validate.MapValidate(ctx.Request, validateMap, validateNameMap); !ok {
		utils.ErrorResponse(ctx, 402, -1, msg)
	}
	whereMap := make(map[string]interface{})
	whereMap["email"] = ctx.PostForm("email")
	mysql.DB.Where(whereMap).First(&userModel)
	if userModel.Id == 0 {
		if err := mysql.DB.Create(&userModel).Error; err != nil {
			utils.ErrorResponse(ctx, 200, -1, "注册失败")
			fmt.Println("register err:", err)
			return
		}
		// 注册成功
		utils.Success(ctx)
	} else {
		utils.ErrorResponse(ctx, 200, -1, "邮箱已被占用")
		return
	}
}

func (*AuthController) Login(ctx *gin.Context) {
	userModel := user.User{}
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	validateMap, validateNameMap := validate.MapDataForStruct(userModel)
	delete(validateMap, "nickname")
	if ok, msg := validate.MapValidate(ctx.Request, validateMap, validateNameMap); !ok {
		utils.ErrorResponse(ctx, 402, -1, msg)
	}
	where := make(map[string]interface{})
	where["email"] = email
	mysql.DB.Where(where).First(&userModel)
	if userModel.Id == 0 {
		utils.ErrorResponse(ctx, 200, -1, "用户名或密码错误")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(password)); err != nil {
		utils.ErrorResponse(ctx, 200, -1, "密码错误")
		return
	}

	utils.SuccessWithData(ctx, &userModel)
}

func (*AuthController) Logout(ctx *gin.Context) {
	// todo
}