package core

import (
	"forward-router/iconst"
	"forward-router/logger"
	"forward-router/st"
	"forward-router/util"
	"net"
)

var flag = ":"
var stop bool

// 开启监听
func IStart() {
	stop = false
	listenIp := "0.0.0.0:10081"
	nl, _ := net.Listen("tcp4", listenIp)
	logger.R().Println("listening addr -->" + listenIp)
	for {
		nconn, err := nl.Accept()
		logger.R().Println("connecion received.")
		if err != nil {
			logger.R().Fatal(err)
		}
		go handle(&nconn)
		if stop {
			break
		}
	}
}
func Stop() {
	stop = true
}
func handle(conn *net.Conn) {
	connection := *conn
	remoteAddr := connection.RemoteAddr().String()
	logger.R().Println("remote connection：" + remoteAddr)
	msg := util.GetStringDataByMy(connection)
	logger.R().Println("orginal message is>EOF:\n", msg, "EOF<<")
	zpro := st.ParserProtocol(msg, flag)
	if zpro.Header != iconst.PRO_TOCOL {
		logger.R().Println("it's not zk protocol.", zpro.Header, "is not ", iconst.PRO_TOCOL)
		defer connection.Close()
		return
	}
	c := &st.Clients{
		Conn:  conn,
		Cname: zpro.Source,
	}
	st.AddConn(c)
	logger.R().Println("zk protocol message is :" + msg)
	dispatch(*zpro)
}
func dispatch(zpro st.ZkPro) {
	DoStrategy(zpro)
}
