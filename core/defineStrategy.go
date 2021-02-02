package core

import (
	"forward-router/logger"
	"forward-router/st"
)

// type Command struct {
// 	Heart   int `enum:"heart"`
// 	Forward int `enum:"forward"`
// }
type HandStrategy interface {
	do(zpro st.ZkPro) int
}

//0x1 心跳
type heart struct{}

//0x2 转发
type forward struct{}

//0x3 转发程序退出
type exit struct{}

//0x4 清除自己的信息
type iexit struct{}

var strategys = []HandStrategy{&heart{}, &forward{}, &exit{}, &iexit{}}

type MessageOperator struct {
	strategy HandStrategy
}

func (operator *MessageOperator) setStrategy(strategy HandStrategy) {
	operator.strategy = strategy
}
func (operator *MessageOperator) doHand(zpro st.ZkPro) {
	i := operator.strategy.do(zpro)
	if i == -1 {
		logger.R().Panic("do hand error.", i)
	}
}

func DoStrategy(zpro st.ZkPro) {
	l := len(strategys)
	if l <= zpro.Cmd {
		logger.R().Panic("command in message is invalid. command is ", zpro.Cmd)
	}
	var msgOp = MessageOperator{}
	msgOp.setStrategy(strategys[zpro.Cmd-1])
	msgOp.doHand(zpro)
}
