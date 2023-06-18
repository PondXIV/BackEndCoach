package walletgbctl

import (
	"backEndGo/models"
	userservice "backEndGo/service/userService"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ServiceMoney = userservice.NewGbprimeDataService()

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
		rowsAffected, _ := ServiceMoney.ServiceWallet(jsonDto.ReferenceNo, &jsonDto)
		//rowAff,_ := ServiceMoney.
		ctx.JSON(http.StatusOK, gin.H{
			"code":   "200",
			"result": strconv.Itoa(int(rowsAffected)),
		})

	}

}
