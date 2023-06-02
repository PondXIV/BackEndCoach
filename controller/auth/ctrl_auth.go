package auth

import (
	dto "backEndGo/DTO"
	"backEndGo/models"
	"backEndGo/service"
	coachsv "backEndGo/service/CoachService/CoachSV"
	"fmt"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

var userDateService = service.NewUserDataService()
var updateCoachService = coachsv.NewUpdateCoachDataService()
var modelsCus = models.Customer{}
var modelsCoach = models.Coach{}

func NewAuthController(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", loginPostBody)
		auth.POST("/loginfb", loginFBPostBody)
		auth.POST("/Cus", registerCus)
		auth.POST("/Coach", registerCoach)
		auth.PUT("/Coach/:cid", updateCoach)
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
		if http.StatusBadRequest == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": err.Error(),
				"cid":  "0",
				"uid":  "0",
			})
		}
	}
	if coach.Cid > 0 && cus.Uid > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "พบสมาชิก",
			"cid":  strconv.Itoa(int(coach.Cid)),
			"uid":  strconv.Itoa(int(cus.Uid)),
		})
	} else if coach.Cid > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "พบสมาชิก",
			"cid":  strconv.Itoa(int(coach.Cid)),
			"uid":  "0",
		})
	} else if cus.Uid > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "พบสมาชิก",
			"cid":  "0",
			"uid":  strconv.Itoa(int(cus.Uid)),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "ไม่พบสมาชิก",
			"cid":  "0",
			"uid":  "0",
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
		if http.StatusBadRequest == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": err.Error(),
				"cid":  "0",
				"uid":  "0",
			})
		}
	}
	if coach.Cid > 0 && cus.Uid > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "พบสมาชิก",
			"cid":  strconv.Itoa(int(coach.Cid)),
			"uid":  strconv.Itoa(int(cus.Uid)),
		})
	} else if coach.Cid > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "พบสมาชิก",
			"cid":  strconv.Itoa(int(coach.Cid)),
			"uid":  "0",
		})
	} else if cus.Uid > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "พบสมาชิก",
			"cid":  "0",
			"uid":  strconv.Itoa(int(cus.Uid)),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "ไม่พบสมาชิก",
			"cid":  "0",
			"uid":  "0",
		})
	}
}

func registerCus(ctx *gin.Context) {
	err := ctx.ShouldBindJSON(&modelsCus)
	// fmt.Printf("%v", cus)
	if err != nil {
		if http.StatusBadRequest == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": err.Error(),
				"cid":  "0",
				"uid":  "0",
			})
		}
	}
	RowsAffected, err := userDateService.ServiceRegisterCus(&modelsCus)
	if err != nil {
		if http.StatusBadRequest == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": err.Error(),
				"cid":  "0",
				"uid":  "0",
			})
		}
	}
	if RowsAffected > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "สมัครสมาชิกสำเร็จ",
			"cid":  "0",
			"uid":  strconv.Itoa(int(modelsCus.Uid)),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "สมัครสมาชิกไม่สำเร็จ",
			"cid":  "0",
			"uid":  "0",
		})
	}
}
func registerCoach(ctx *gin.Context) {
	err := ctx.ShouldBindJSON(&modelsCoach)
	if err != nil {
		if http.StatusBadRequest == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": err.Error(),
				"cid":  "0",
				"uid":  "0",
			})
		}
	}
	RowsAffected, err := userDateService.ServiceRegisterCoach(&modelsCoach)
	if err != nil {
		if http.StatusBadRequest == 400 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": err.Error(),
				"cid":  "0",
				"uid":  "0",
			})
		}
	}
	if RowsAffected > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "สมัครสมาชิกสำเร็จ",
			"cid":  strconv.Itoa(int(modelsCoach.Cid)),
			"uid":  "0",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "สมัครสมาชิกไม่สำเร็จ",
			"cid":  "0",
			"uid":  "0",
		})
	}
}

func updateCoach(ctx *gin.Context) {
	coachID := ctx.Param("cid")

	cid, errs := strconv.Atoi(coachID)
	if errs != nil {
		panic(errs)
	}
	err := ctx.ShouldBindJSON(&modelsCoach)
	// fmt.Printf("%v", cus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   "400",
			"result": err.Error(),
		})
	} else {
		rowsAffected, err := updateCoachService.ServiceUpdateCoach(cid, &modelsCoach)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":   "400",
				"result": err.Error(),
			})

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

}
func outputSoon(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"result": "0",
	})
}
