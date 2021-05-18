package server

import (
	"github.com/gin-gonic/gin"
)

func gininit() {
	router := gin.Default()

	router.POST("/callback", post)

	router.Run()
}
