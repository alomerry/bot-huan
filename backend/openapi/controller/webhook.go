package controller

import (
	"bot-huan/core"
	component2 "bot-huan/core/component"
	"bot-huan/core/log"
	"github.com/gin-gonic/gin"
	"github.com/lonelyevil/khl"
	"github.com/lonelyevil/khl/log_adapter/plog"
	"math/rand"
	"strconv"
	"strings"
)

func Webhook() func(ctx *gin.Context) {
	s := khl.New(core.GetKaiheilaBotToken(), plog.NewLogger(log.GetLogger()), khl.SessionWithEncryptKey([]byte(core.GetKaiheilaBotEncryptKey())), khl.SessionWithVerifyToken(core.GetKaiheilaBotVerifyToken()))
	s.AddHandler(messageHan)
	return gin.WrapF(s.WebhookHandler())
}

func messageHan(ctx *khl.KmarkdownMessageContext) {
	if ctx.Common.Type != khl.MessageTypeKMarkdown || ctx.Extra.Author.Bot {
		return
	}
	switch {
	case strings.HasPrefix(ctx.Common.Content, "-ping"):
		component2.SubmitTask(func() {
			ctx.Session.MessageCreate(&khl.MessageCreate{
				MessageCreateBase: khl.MessageCreateBase{
					TargetID: ctx.Common.TargetID,
					Content:  "[{\"type\":\"card\",\"theme\":\"none\",\"color\":\"#569185\",\"size\":\"lg\",\"modules\":[{\"type\":\"section\",\"text\":{\"type\":\"kmarkdown\",\"content\":\"`pong`\"}},{\"type\":\"context\",\"elements\":[{\"type\":\"kmarkdown\",\"content\":\"BOT - 服务来源 [清欢](https://blog.alomerry.com) | [源码](https://github.com/Alomerry)\"}]}]}]",
					Quote:    ctx.Common.MsgID,
					Type:     khl.MessageTypeCard,
				},
			})
		})
	case strings.HasPrefix(ctx.Common.Content, "-roll"):
		component2.SubmitTask(func() {
			ctx.Session.MessageCreate(&khl.MessageCreate{
				MessageCreateBase: khl.MessageCreateBase{
					TargetID: ctx.Common.TargetID,
					Content:  "[{\"type\":\"card\",\"theme\":\"none\",\"color\":\"#569185\",\"size\":\"lg\",\"modules\":[{\"type\":\"section\",\"text\":{\"type\":\"kmarkdown\",\"content\":\"`" + strconv.FormatInt(int64(rand.Float64()*100), 10) + "`\"}},{\"type\":\"context\",\"elements\":[{\"type\":\"kmarkdown\",\"content\":\"BOT - 服务来源 [清欢](https://blog.alomerry.com) | [源码](https://github.com/Alomerry)\"}]}]}]",
					Quote:    ctx.Common.MsgID,
					Type:     khl.MessageTypeCard,
				},
			})
		})
	case strings.HasPrefix(ctx.Common.Content, "-music"):
		component2.SubmitTask(func() {
			page := uint8(1)
			// TODO 并发
		RETRY:
			searKey := strings.TrimPrefix(ctx.Common.Content, "-music ")
			musicId, musicPicUrl, musicName := component2.NeteaseAPI.Search(searKey, page)
			musicUrl := component2.NeteaseAPI.GetSongUrl(musicId)
			if musicUrl != "" {
				ctx.Session.MessageCreate(&khl.MessageCreate{
					MessageCreateBase: khl.MessageCreateBase{
						Type:     khl.MessageTypeCard,
						TargetID: ctx.Common.TargetID,
						Content: `[
							  {
								"type": "card","theme": "none","color": "#569185","size": "lg",
								"modules": [
								  {
									"type": "audio",
									"title": "` + musicName + `",
									"src": "` + musicUrl + `",
									"cover": "` + musicPicUrl + `"
								  },
								  {"type": "context","elements": [{"type": "kmarkdown","content": "BOT - 服务来源 [清欢](https://blog.alomerry.com) | [源码](https://github.com/Alomerry)"}]}
								]
							  }
							]`,
						Quote: ctx.Common.MsgID,
					},
				})
			} else {
				if page >= 5 {
					ctx.Session.MessageCreate(&khl.MessageCreate{
						MessageCreateBase: khl.MessageCreateBase{
							Type:     khl.MessageTypeCard,
							TargetID: ctx.Common.TargetID,
							Content:  `[{"type":"card","theme":"none","color":"#569185","size":"lg","modules":[{"type":"section","text":{"type":"kmarkdown","content":"暂无版权"}},{"type":"context","elements":[{"type":"kmarkdown","content":"BOT - 服务来源 [清欢](https://blog.alomerry.com) | [源码](https://github.com/Alomerry)"}]}]}]`,
							Quote:    ctx.Common.MsgID,
						},
					})
				} else {
					page++
					goto RETRY
				}
			}
		})
	case strings.HasPrefix(ctx.Common.Content, "-help"):
		component2.SubmitTask(func() {
			ctx.Session.MessageCreate(&khl.MessageCreate{
				MessageCreateBase: khl.MessageCreateBase{
					Type:     khl.MessageTypeCard,
					TargetID: ctx.Common.TargetID,
					Content:  "[{\"type\":\"card\",\"theme\":\"none\",\"color\":\"#569185\",\"size\":\"lg\",\"modules\":[{\"type\":\"header\",\"text\":{\"type\":\"plain-text\",\"content\":\"指令帮助：\"}},{\"type\":\"divider\"},{\"type\":\"section\",\"text\":{\"type\":\"kmarkdown\",\"content\":\"`-help` 帮助指令\"}},{\"type\":\"section\",\"text\":{\"type\":\"kmarkdown\",\"content\":\"`-ping` 戳一戳我\"}},{\"type\":\"section\",\"text\":{\"type\":\"kmarkdown\",\"content\":\"`-roll` 掷色子 1-100\"}},{\"type\":\"section\",\"text\":{\"type\":\"kmarkdown\",\"content\":\"`-music` 点播网易云音乐\"}},{\"type\":\"divider\"},{\"type\":\"section\",\"text\":{\"type\":\"kmarkdown\",\"content\":\"**[BOT - 最新动态](https://kaihei.co/r7Gva2)**\"}},{\"type\":\"action-group\",\"elements\":[{\"type\":\"button\",\"theme\":\"warning\",\"value\":\"ok\",\"text\":{\"type\":\"plain-text\",\"content\":\"赞助名单\"}},{\"type\":\"button\",\"theme\":\"danger\",\"value\":\"cancel\",\"text\":{\"type\":\"plain-text\",\"content\":\"我要赞助\"}}]},{\"type\":\"context\",\"elements\":[{\"type\":\"kmarkdown\",\"content\":\"BOT - 服务来源 [清欢](https://blog.alomerry.com) | [源码](https://github.com/Alomerry)\"}]}]}]",
					Quote:    ctx.Common.MsgID,
				},
			})
		})
	default:
		// empty
	}
}
