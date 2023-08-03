package walletbillctl

import (
	"backEndGo/models"
	userservice "backEndGo/service/userService"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewWalletController(router *gin.Engine) {
	wallet := router.Group("/wallet")
	{
		wallet.POST(":uid", Insertwallet)
		wallet.GET("", GetHistoryWallet)
	}

}
func GetHistoryWallet(ctx *gin.Context) {
	cusID := ctx.Query("uid")
	fmt.Println(cusID)
	uid, err := strconv.Atoi(cusID)
	wallet, err := userservice.NewGbprimeDataService().ServiceHistoryWallet(uid)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, wallet)
}
func Insertwallet(ctx *gin.Context) {
	cusID := ctx.Param("uid")

	uid, errs := strconv.Atoi(cusID)
	if errs != nil {
		panic(errs)
	}
	modelsWallet := models.Wallet{}
	err := ctx.ShouldBindJSON(&modelsWallet)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   "400",
			"result": err.Error(),
		})
	}
	rowsAffected, err := userservice.NewGbprimeDataService().ServiceInsertWallet(uid, &modelsWallet)
	if err != nil {
		if http.StatusBadRequest == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":   "400",
				"result": err.Error(),
			})
		}
	} else {
		if rowsAffected >= 1 {
			ctx.JSON(http.StatusOK, gin.H{
				"code":   "200",
				"result": strconv.Itoa(int(rowsAffected)),
			})

		} else {
			outputSoon(ctx)
		}
	}

}
func outputSoon(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"result": "0",
	})
}
