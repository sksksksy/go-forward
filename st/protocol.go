package st

import (
	"fmt"
	"forward-router/iconst"
	"forward-router/logger"
	"net"
	"strconv"
	"strings"
	"sync"
)

// zk protocol
type ZkPro struct {
	Header int    //协议头
	Source string //原地址
	Target string //目的地址
	Cmd    int    //命令
	Msg    string //信息
}
type Clients struct {
	Conn   *net.Conn  //连接
	Cname  string     //连接名
	IsKeep bool       //是否保持连接
	Mu     sync.Mutex //锁
}
type pool struct {
	sync.RWMutex //读写锁
	conns        map[string]*Clients
}

var p *pool

func init() {
	p = &pool{
		conns: make(map[string]*Clients),
	}
}

//添加连接
func AddConn(c *Clients) {
	p.Lock()
	cons := p.conns
	cons[c.Cname] = c
	defer p.Unlock()
}

// 删除连接
func RemoveConn(c *Clients) {
	p.Lock()
	cons := p.conns
	if cons[c.Cname] != nil {
		delete(cons, c.Cname)
	}
	defer p.Unlock()
}

// 获取连接
func GetConn(cname string) *Clients {
	p.RLock()
	defer p.RUnlock()
	cons := p.conns
	c := cons[cname]
	return c
}
func ParserProtocol(ostr string, seq string) *ZkPro {
	var zpro ZkPro
	ostrs := strings.Split(ostr, seq)
	zpro.Header, _ = strconv.Atoi(ostrs[0])
	if zpro.Header != iconst.PRO_TOCOL {
		logger.R().Println("it's not zk protocol.", zpro.Header, "is not ", iconst.PRO_TOCOL)
		return &zpro
	}
	l := len(ostrs)
	zpro.Source = getI(ostrs, l, 1)
	zpro.Target = getI(ostrs, l, 2)
	zpro.Cmd, _ = strconv.Atoi(getI(ostrs, l, 3))
	zpro.Msg = getI(ostrs, l, 4)
	return &zpro
}
func AssemblyProtocol(zpro ZkPro, seq string) string {
	var s string
	s = fmt.Sprintf("%d%s", zpro.Header, seq)
	s = fmt.Sprintf("%s%s%s", s, zpro.Source, seq)
	s = fmt.Sprintf("%s%s%s", s, zpro.Target, seq)
	s = fmt.Sprintf("%s%d%s", s, zpro.Cmd, seq)
	s = fmt.Sprintf("%s%s\r\n", s, zpro.Msg)
	return s
}
func getI(s []string, l int, i int) string {
	if i >= l {
		return ""
	}
	return s[i]
}
