package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"log"
)

func main() {
	//bot := openwechat.DefaultBot()
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	if err := bot.Login(); err != nil {
		log.Println("login error:", err)
		return
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		log.Println("get current user", err)
		return
	}

	// 获取所有的好友
	friends, err := self.Friends(true)
	if err != nil {
		log.Println("get friends error:", err)
		return
	}
	for _, f := range friends {
		msg := fmt.Sprintf("%s，祝你新年快乐", f.RemarkName)
		r := []rune(f.RemarkName)
		if len(r) > 0 && len(r) <= 4 {
			_, err := f.SendText(msg)
			if err != nil {
				log.Println("send error :", r, err)
				continue
			}
			log.Println("send successful:", msg)
		}
	}
	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
