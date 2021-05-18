package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func lineinit() {
	if err87 := godotenv.Load(); err87 != nil {
		log.Print("No .env file found")
	}
	var err error
	bot, err = linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_ACCESS_TOKEN"))
	log.Println("Bot:", bot, " err:", err)
	//http.HandleFunc("/ping", callbackHandler)
	//http.HandleFunc("/callback", callbackHandler)

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func post(c *gin.Context) {
	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.JSON(http.StatusBadRequest, nil)
		} else {
			c.JSON(http.StatusInternalServerError, nil)
		}
		return
	}
	/*
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
	*/
	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeMessage:
			profile, err := bot.GetProfile(event.Source.UserID).Do()
			if err != nil {
				log.Print(err)
			}

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				matchText := message.Text
				buyProductName := ""

				if strings.Contains(matchText, "存錢") {
					buyProductName = strings.Split(matchText, ",")[1]
					matchText = strings.Split(matchText, ",")[0]
				} else if strings.Contains(matchText, "領錢") {
					buyProductName = strings.Split(matchText, ",")[1]
					matchText = strings.Split(matchText, ",")[0]
				}

				switch matchText {
				case "存錢":
					replyMsg := "成功! "
					money, err := strconv.Atoi(buyProductName)
					if err != nil {
						replyMsg = "請輸入整數"
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMsg)).Do(); err != nil {
							log.Print(err)
						}
					} else {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMsg+profile.DisplayName+" 存入 "+strconv.Itoa(money)+"元")).Do(); err != nil {
							log.Print(err)
						}
						deposit(profile.DisplayName, "存錢", money)
					}

				case "領錢":
					replyMsg := "成功! "
					money, err := strconv.Atoi(buyProductName)
					if err != nil {
						replyMsg = "請輸入整數"
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMsg)).Do(); err != nil {
							log.Print(err)
						}
					} else {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMsg+profile.DisplayName+" 領出 "+strconv.Itoa(money)+"元")).Do(); err != nil {
							log.Print(err)
						}
						withdraw(profile.DisplayName, "領錢", money)
					}

				case "餘額查詢":
					replyMsg := "成功! "
					bal := balance(profile.DisplayName)
					if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMsg+profile.DisplayName+"的餘額為"+strconv.Itoa(bal))).Do(); err != nil {
						log.Print(err)
					}

				default:
					if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("輸入格式:\n存錢,數字\n領錢,數字\n餘額查詢")).Do(); err != nil {
						log.Print(err)
					}
				}
			case *linebot.ImageMessage:
				log.Print(message)
			case *linebot.VideoMessage:
				log.Print(message)
			case *linebot.AudioMessage:
				log.Print(message)
			case *linebot.FileMessage:
				log.Print(message)
			case *linebot.LocationMessage:
				log.Print(message)
			case *linebot.StickerMessage:
				log.Print(message)
			default:
				log.Printf("Unknown message: %v", message)
			}
		default:
			log.Printf("Unknown event: %v", event)
		}
	}
}
