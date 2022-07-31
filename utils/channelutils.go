package utils

import (
	"context"
)

type RetChannelStuct struct {
	RetVal    uint
	ReplyMsg  *ChannelMsg
	ReqType   ReqType
	MessageId uint64
}

type ChannelMsg struct {
	Text  string
	At    bool
	Image string
	File  string
}

type ChannelPlugin interface {
	ChannelDo(ctx *context.Context, botId, botChannelId int64, guildId, channelId, userId uint64, rawMsg, card string, super, userRole bool, rs, rd, rf int) (retStuct RetChannelStuct)
}

var ChannelPluginSet map[string]ChannelPlugin

func init() {
	ChannelPluginSet = make(map[string]ChannelPlugin)
}

func ChannelRegister(k string, v ChannelPlugin) {
	ChannelPluginSet[k] = v
}
