package server

import (
	"net/http"

	"github.com/PietaTony/site-system/serverSide/db/step"
	"github.com/gin-gonic/gin"
)

func handleCheckErr(context *gin.Context, err error) {
	context.JSON(http.StatusOK, gin.H{
		"error": err.Error(),
	})
}

func handleCreateStep(context *gin.Context) {
	input := step.New()
	err := context.ShouldBindJSON(&input)
	if err != nil {
		handleCheckErr(context, err)
		return
	}

	step, err := step.Create(input)
	if err != nil {
		handleCheckErr(context, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"data": step,
	})
}

func handelSelectSteps(context *gin.Context) {
	steps, err := step.SelectAll()
	if err != nil {
		handleCheckErr(context, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": steps})
}
