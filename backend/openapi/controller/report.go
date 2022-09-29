package controller

import (
	"bot-huan/core"
	"bot-huan/core/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lonelyevil/khl"
	"github.com/lonelyevil/khl/log_adapter/plog"
)

var (
	khlSession *khl.Session
)

func init() {
	khlSession = khl.New(core.GetKaiheilaBotToken(), plog.NewLogger(log.GetLogger()), khl.SessionWithEncryptKey([]byte(core.GetKaiheilaBotEncryptKey())), khl.SessionWithVerifyToken(core.GetKaiheilaBotVerifyToken()))
}

type Reporter struct {
	WorkIP string
}

func Report(c *gin.Context) {
	reportInfo := Reporter{}
	err := c.BindJSON(&reportInfo)
	if err != nil {
		// TODO log
		fmt.Println(err)
	}
	if reportInfo.WorkIP != "" {
		khlSession.MessageCreate(&khl.MessageCreate{
			MessageCreateBase: khl.MessageCreateBase{
				TargetID: "1655514156713483",
				Content:  "(met)1712811410(met) " + reportInfo.WorkIP,
				Quote:    "d83c799e-76c9-4a21-94c2-0c1bb2e88635",
				Type:     khl.MessageTypeText,
			},
		})
	}
}
