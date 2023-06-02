package requestcontroller

import (
	"backEndGo/models"
	requestservice "backEndGo/service/RequestService"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var modelsRequest = models.Request{}
var serviceShowRequest = requestservice.NewShowRequestDataService()

func NewBuyCourseController(router *gin.Engine) {
	request := router.Group("/request")
	{
		request.GET("", showRequest)
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
