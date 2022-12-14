package router

import (
	"FieldProgramming/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ConvertRequest struct {
	Input string `json:"input" binding:"required"`
}

type ConvertResponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

func ConvertService(ctx *gin.Context) {
	req := &ConvertRequest{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(400, gin.H{})
		return
	}
	tpe := util.JudgeInputType(req.Input)
	if tpe == util.UNKNOWN {
		ctx.JSON(200, ConvertResponse{
			Code: 0,
			Data: "",
		})
		return
	}
	if tpe == util.CN {
		i := util.CNConvertToDigit(req.Input)
		if i == -1 || i == 0 {
			ctx.JSON(200, ConvertResponse{
				Code: 0,
				Data: "",
			})
			return
		}
		if i == -2 {
			i = 0
		}
		ctx.JSON(200, ConvertResponse{
			Code: 1,
			Data: strconv.Itoa(i),
		})
		return
	}
	if tpe == util.DIGITS {
		digit, err := strconv.Atoi(req.Input)
		if err != nil {
			ctx.JSON(200, ConvertResponse{
				Code: 0,
				Data: "",
			})
			return
		}
		data := util.DigitConvertToCN(strconv.FormatInt(int64(digit), 10))
		if data == "输入错误" {
			ctx.JSON(200, ConvertResponse{
				Code: 0,
				Data: "",
			})
			return
		}
		ctx.JSON(200, ConvertResponse{
			Code: 1,
			Data: data,
		})
		return
	}
	ctx.JSON(500, gin.H{})
	return
}
