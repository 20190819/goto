package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
		utils.Success(ctx)
	} else {
		utils.ErrorResponse(ctx, 200, -1, "邮箱已被占用")
		return
	}
}

func (*AuthController) Login(ctx *gin.Context) {
	// todo
}
