package pratice

import (
	"log"
	"net/http"
	"strconv"

	"github.com/line/line-bot-sdk-go/v7/linebot"

	"github.com/gin-gonic/gin"
)

func gintest() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	/*router.POST("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})*/

	router.POST("/callback", post)

	router.Run()
}

func post(c *gin.Context) {
	//msg := context.DefaultPostForm("input", "表單沒有input。") // 沒有輸入參數時 可設定預設值

	//context.String(http.StatusOK, "您輸入的文字為: \n%s", msg)

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.JSON(http.StatusBadRequest, nil)
		} else {
			c.JSON(http.StatusInternalServerError, nil)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				quota, err := bot.GetMessageQuota().Do()
				if err != nil {
					log.Println("Quota err:", err)
				}
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK! remain message:"+strconv.FormatInt(quota.Value, 10))).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
