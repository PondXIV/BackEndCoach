package walletctl

import (
	"backEndGo/models"
	userservice "backEndGo/service/userService"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewWalletGBController(router *gin.Engine) {
	walletgb := router.Group("/gbcallback")
	{
		walletgb.POST("", getgbprime)
	}

}

func getgbprime(ctx *gin.Context) {
	jsonDto := models.Gbprimpay{}
	err := ctx.ShouldBindJSON(&jsonDto)

	if err != nil {

		panic(err)
	}

	if jsonDto.ResultCode == "00" {
		rowsAffected, _ := userservice.NewGbprimeDataService().ServiceWallet(jsonDto.ReferenceNo, &jsonDto)
		ctx.JSON(http.StatusOK, gin.H{
			"code":   "200",
			"result": strconv.Itoa(int(rowsAffected)),
		})

	}

}
