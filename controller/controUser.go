package controller

import (
	dto "backEndGo/DTO"
	"backEndGo/service"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userDateService = service.NewUserDataService()

func NewUserController(router *gin.Engine) {
	ping := router.Group("/user")
	{
		ping.POST("/login", loginPostBody)
		ping.POST("/registerCus", registerCus)
		ping.POST("/registerCoach", registerCoach)
	}

}
func loginPostBody(ctx *gin.Context) {
	// input json
	jsonDto := dto.LoginDTO{}
	err := ctx.ShouldBindJSON(&jsonDto)
	fmt.Printf(jsonDto.Email, jsonDto.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	// process
	coach, cus, err := userDateService.ServiceLogin(jsonDto.Email, jsonDto.Password, jsonDto.Type)

	// output json
	// println("===============")
	if err != nil {
		panic(err)
	}

	if len(*coach) == 1 {
		ctx.JSON(http.StatusOK, coach)
	} else if len(*cus) == 1 {
		ctx.JSON(http.StatusOK, cus)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	// println("=================")
}
func registerCus(ctx *gin.Context) {
	jsonDto := dto.RegisterCusDTO{}
	err := ctx.ShouldBindJSON(&jsonDto)
	fmt.Printf(jsonDto.AliasName, jsonDto.Password,
		jsonDto.Email, jsonDto.FullName,
		jsonDto.Birthday, jsonDto.Gender,
		jsonDto.Phone, jsonDto.Image,
		jsonDto.Weight, jsonDto.Height, jsonDto.Price)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, jsonDto)
	// cus, err := userDateService.ServiceRegisterCus(jsonDto.AliasName, jsonDto.Password,
	// 	jsonDto.Email, jsonDto.FullName,
	// 	jsonDto.Birthday, jsonDto.Gender,
	// 	jsonDto.Phone, jsonDto.Image,
	// 	jsonDto.Weight, jsonDto.Height, jsonDto.Price)
	// if err != nil {
	// 	panic(err)
	// }
	// if len(*cus) == 1 {
	// 	ctx.JSON(http.StatusOK, cus)
	// } else {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err,
	// 	})
	// }

}
func registerCoach(ctx *gin.Context) {

}
