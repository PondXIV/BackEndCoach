package requestcontroller

import (
	"backEndGo/models"
	requestservice "backEndGo/service/RequestService"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var serviceShowRequest = requestservice.NewShowRequestDataService()

func NewRequestController(router *gin.Engine) {
	request := router.Group("/request")
	{
		request.GET("", showRequest)
		request.POST(":uid", Insertrequest)
	}

}
func Insertrequest(ctx *gin.Context) {
	cusID := ctx.Param("uid")

	uid, errs := strconv.Atoi(cusID)
	if errs != nil {
		panic(errs)
	}
	modelsRequest := models.Request{}
	err := ctx.ShouldBindJSON(&modelsRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   "400",
			"result": err.Error(),
		})
	}
	rowsAffected, err := requestservice.NewInsertRequestDataService().ServiceInsertRequest(uid, &modelsRequest)
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
func showRequest(ctx *gin.Context) {
	qrqID := ctx.Query("rqID")
	quid := ctx.Query("uid")
	qcid := ctx.Query("cid")

	rqID, err := strconv.Atoi(qrqID)
	uid, err := strconv.Atoi(quid)
	cid, err := strconv.Atoi(qcid)

	request, err := serviceShowRequest.ServiceGetRequest(rqID, uid, cid)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, request)

}
func outputSoon(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"result": "0",
	})
}
