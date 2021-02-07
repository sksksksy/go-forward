package core

import (
	"forward-router/logger"
	"forward-router/st"
)

func (*heart) do(zpro st.ZkPro) int {
	con := st.GetConn(zpro.Source)
	r := st.AssemblyProtocol(zpro, flag)
	i, err := (*con.Conn).Write([]byte(r))
	if err != nil {
		logger.R().Println(err)
		i = -1
	}
	Handle(con.Conn)
	return i
}

func (*forward) do(zpro st.ZkPro) int {
	con := st.GetConn(zpro.Target)
	r := st.AssemblyProtocol(zpro, flag)
	i, err := (*con.Conn).Write([]byte(r))
	if err != nil {
		i = -1
	}
	return i
}
func (*exit) do(zpro st.ZkPro) int {
	Stop()
	return 1
}
func (*iexit) do(zpro st.ZkPro) int {
	client := &st.Clients{
		Cname: zpro.Source,
	}
	st.RemoveConn(client)
	logger.R().Println("Conn has been removed.", zpro.Source)
	return 1
}
