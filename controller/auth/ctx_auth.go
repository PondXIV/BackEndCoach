package auth

import (
	dto "backEndGo/DTO"
	"backEndGo/models"
	"backEndGo/service"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

var userDateService = service.NewUserDataService()
var modelsCus = models.Customer{}
var modelsCoach = models.Coach{}

func NewAuthController(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", loginPostBody)
		auth.POST("/loginfb", loginFBPostBody)
		auth.POST("/registerCus", registerCus)
		auth.POST("/registerCoach", registerCoach)
	}

}
func loginFBPostBody(ctx *gin.Context) {
	// input json
	jsonDto := dto.LoginFBDTO{}
	err := ctx.ShouldBindJSON(&jsonDto)
	fmt.Printf(jsonDto.FacebookID)
	if err != nil {
		panic(err)
	}

	coach, cus, err := userDateService.ServiceLoginFB(jsonDto.FacebookID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"cid": nil,
			"uid": nil,
		})
	}
	if coach.Cid > 0 && cus.Uid > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"cid": coach.Cid,
			"uid": cus.Uid,
		})
	} else if coach.Cid > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"cid": coach.Cid,
			"uid": nil,
		})
	} else if cus.Uid > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"cid": nil,
			"uid": cus.Uid,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"cid": nil,
			"uid": nil,
		})
	}
}
func loginPostBody(ctx *gin.Context) {
	// input json
	jsonDto := dto.LoginDTO{}
	err := ctx.ShouldBindJSON(&jsonDto)
	fmt.Printf(jsonDto.Email, jsonDto.Password)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, err)
		panic(err)
	}
	// process
	//coach, cus, err := userDateService.ServiceLogin(jsonDto.Email, jsonDto.Password, jsonDto.Type)
	coach, cus, err := userDateService.ServiceLogin(jsonDto.Email, jsonDto.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"cid": nil,
			"uid": nil,
		})
	}
	if coach.Cid > 0 && cus.Uid > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"cid": coach.Cid,
			"uid": cus.Uid,
		})
	} else if coach.Cid > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"cid": coach.Cid,
			"uid": nil,
		})
	} else if cus.Uid > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"cid": nil,
			"uid": cus.Uid,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"cid": nil,
			"uid": nil,
		})
	}
}

func registerCus(ctx *gin.Context) {
	err := ctx.ShouldBindJSON(&modelsCus)
	// fmt.Printf("%v", cus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	RowsAffected := userDateService.ServiceRegisterCus(&modelsCus)

	if RowsAffected > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"cid": nil,
			"uid": modelsCus.Uid,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"cid": nil,
			"uid": nil,
		})
	}
}
func registerCoach(ctx *gin.Context) {
	err := ctx.ShouldBindJSON(&modelsCoach)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	RowsAffected := userDateService.ServiceRegisterCoach(&modelsCoach)
	if RowsAffected > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"cid": nil,
			"uid": modelsCus.Uid,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"cid": nil,
			"uid": nil,
		})
	}
}
