package handlers

import (
	"github.com/JustZyx/wechat-assistant/gtp"
	"github.com/eatmoreapple/openwechat"
	"log"
	"strings"
	"time"
)

var _ MessageHandlerInterface = (*GroupMessageHandler)(nil)

// GroupMessageHandler 群消息处理
type GroupMessageHandler struct {
}

// handle 处理消息
func (g *GroupMessageHandler) handle(msg *openwechat.Message) error {
	if msg.IsText() {
		return g.ReplyText(msg)
	}
	return nil
}

// NewGroupMessageHandler 创建群消息处理器
func NewGroupMessageHandler() MessageHandlerInterface {
	return &GroupMessageHandler{}
}

// ReplyText 发送文本消息到群
var m = make(map[string]int)

func (g *GroupMessageHandler) ReplyText(msg *openwechat.Message) error {
	// 接收群消息
	sender, err := msg.Sender()
	group := openwechat.Group{sender}
	log.Printf("Received Group %v Text Msg : %v ChatRoomId: %v", group.NickName, msg.Content, group.ChatRoomId)

	// 三点钟以后,不是这两个群的不提供服务
	if group.NickName != "王姐农药开black群5th（substitute）" && group.NickName != "" && time.Now().Unix() > 1676703600 {
		if _, ok := m[group.NickName]; !ok {
			msg.ReplyText("2月18日15点起不再提供群聊服务,感谢大家的厚爱,江湖再见")
			m[group.NickName]++
			return nil
		}
		if m[group.NickName] == 1 {
			msg.ReplyText("再见啦👋")
			m[group.NickName]++
			return nil
		}
		return nil
	}

	// 不是@的不处理
	if !msg.IsAt() {
		return nil
	}

	// 替换掉@文本，然后向GPT发起请求
	replaceText := "@" + sender.Self.NickName
	requestText := strings.TrimSpace(strings.ReplaceAll(msg.Content, replaceText, ""))
	reply, err := gtp.Completions(requestText)
	if err != nil {
		log.Printf("gtp request error: %v \n", err)
		msg.ReplyText("对不起，我累了")
		return err
	}
	if reply == "" {
		return nil
	}

	// 获取@我的用户
	groupSender, err := msg.SenderInGroup()
	if err != nil {
		log.Printf("get sender in group error :%v \n", err)
		return err
	}

	// 回复@我的用户
	reply = strings.TrimSpace(reply)
	reply = strings.Trim(reply, "\n")
	atText := "@" + groupSender.NickName
	replyText := atText + reply
	_, err = msg.ReplyText(replyText)
	if err != nil {
		log.Printf("response group error: %v \n", err)
	}
	return err
}
