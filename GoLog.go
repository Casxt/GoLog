package GoLog

import (
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

var (
	logChannel chan string
	logCount   uint64
)

//StartLogger Initial Logger Processor
func StartLogger() {
	logChannel = make(chan string, 1024)
	logCount = 0
	go service()
}

//StopLogger Stop Logger Processor
func StopLogger() {
	close(logChannel)
}

//service is use to deal the log,
//consider with multi services are work
func service() {
	var sb strings.Builder
	for msg := range logChannel {
		sb.Grow(256)
		sb.WriteString("[")
		sb.WriteString(time.Now().Format("2006-01-02 15:04:05"))
		sb.WriteString("][")
		sb.WriteString(strconv.FormatUint(logCount, 10))
		sb.WriteString("]")
		sb.WriteString(msg)
		output(sb.String())
		sb.Reset()
		atomic.AddUint64(&logCount, 1)
	}
}

func output(log string) {
	//fmt.Println(log)
}
