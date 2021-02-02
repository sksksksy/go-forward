package logger

import (
	"io"
	"log"
	"os"
)

var oss = make(map[string]*os.File)
var k string = "logfile"

func init() {
	f, err := os.OpenFile("program-log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	oss[k] = f
}
func R() *log.Logger {
	f := oss[k]
	writes := []io.Writer{
		f,
		os.Stdout}
	fsw := io.MultiWriter(writes...)
	logger := log.New(fsw, "[*]", log.Ldate|log.Ltime|log.Lshortfile)
	// logger.Fatalln("")
	return logger
}
