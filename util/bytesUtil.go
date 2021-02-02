package util

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net"
)

func GetStringDataByMy(conn net.Conn) string {
	var result string
	for {
		buf := make([]byte, 1024)
		//conn.SetReadDeadline(time.Now().Add(time.Second * 1))
		//conn.SetReadDeadline(time.Now().Add(time.Second*1))//当在1秒类没有读到数据的时候就退出循
		n, err := conn.Read(buf)
		if n <= 0 || err != nil || err == io.EOF {
			break
		}
		isBreak := checkEnd(buf, n)
		result = result + string(buf[:n])
		if isBreak {
			break
		}
	}
	return result
}
func GetByteDateByMy(conn net.Conn) []byte {
	return []byte(GetByteDateByMy(conn))
}

/**
读取conn的start:end字节的字节流
目前无法读取大流
minReadSize为每次读取的最小
*/
func ReadByte(start int, end int, conn net.Conn, minReadSize int) []byte {
	var buf bytes.Buffer
	var cap int = 0
	for {
		bytes := make([]byte, minReadSize)
		if cap >= end {
			break
		}
		//conn.SetReadDeadline(time.Now().Add(time.Second*1))//当在1秒类没有读到数据的时候就退出循
		n, err := conn.Read(bytes)
		if n <= 0 || err != nil || err == io.EOF {
			break
		}
		buf.Write(bytes[:n])
		cap = cap + n
	}
	return buf.Bytes()[start:end]
}
func ReadData(conn net.Conn) []byte {
	//reader:=bufio.NewReader(conn)
	bytes, err := ioutil.ReadAll(conn)
	if err != nil && err == io.EOF {
		conn.Close()
		log.Fatal(err)
	}
	return bytes
}
func ReadStringData(conn net.Conn) string {
	return string(ReadData(conn))
}

//检查是否是文件消息末尾使用/r/n作为结尾的符号
func checkEnd(bytes []byte, ln int) bool {
	if bytes[ln-2] == 0xd && bytes[ln-1] == 0xa {
		return true
	}
	return false
}
