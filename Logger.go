package GoLog

import (
	"os"
	"runtime"
	"strconv"
	"strings"
)

func getCallerInfo() (funcName, file, line string) {
	pc := make([]uintptr, 1)
	//skip 3 level
	if runtime.Callers(3, pc) == 0 {
		return "", "", ""
	}
	funcHandle := runtime.FuncForPC(pc[0])
	if funcHandle == nil {
		return strconv.FormatUint(uint64(pc[0]), 10), "", ""
	}
	funcName = funcHandle.Name()
	file, linenum := funcHandle.FileLine(pc[0])
	return funcName, file, strconv.Itoa(linenum)
}

//Info level offer caller name
func Info(msg string) {
	funcName, _, _ := getCallerInfo()
	funcNameGroup := strings.Split(funcName, "/")
	var sb strings.Builder
	sb.Grow(128)
	sb.WriteString("[INFO][")
	sb.WriteString(funcNameGroup[len(funcNameGroup)-1])
	sb.WriteString("]")
	sb.WriteString(msg)
	logChannel <- sb.String()
}

//WARN level offer caller detial
func Warn(msg string) {
	funcName, _, _ := getCallerInfo()
	var sb strings.Builder
	sb.Grow(128)
	sb.WriteString("[WARN][")
	sb.WriteString(funcName)
	sb.WriteString("]")
	sb.WriteString(msg)
	logChannel <- sb.String()
}

//Debug level offer file line caller name
func Debug(msg string) {
	funcName, file, line := getCallerInfo()
	funcNameGroup := strings.Split(funcName, "/")
	var sb strings.Builder
	sb.Grow(128)
	sb.WriteString("[DEBUG][")
	sb.WriteString(file)
	sb.WriteString(":")
	sb.WriteString(line)
	sb.WriteString("][")
	sb.WriteString(funcNameGroup[len(funcNameGroup)-1])
	sb.WriteString("]")
	sb.WriteString(msg)
	logChannel <- sb.String()
}

//Error level offer file line caller detial
func Error(msg string) {
	funcName, file, line := getCallerInfo()
	var sb strings.Builder
	sb.Grow(128)
	sb.WriteString("[ERROR][")
	sb.WriteString(file)
	sb.WriteString(":")
	sb.WriteString(line)
	sb.WriteString("][")
	sb.WriteString(funcName)
	sb.WriteString("]")
	sb.WriteString(msg)
	logChannel <- sb.String()
}

//Fatal level offer file line caller detial and exit with code 1
func Fatal(msg string) {
	funcName, file, line := getCallerInfo()
	var sb strings.Builder
	sb.Grow(128)
	sb.WriteString("[FATAL][")
	sb.WriteString(file)
	sb.WriteString(":")
	sb.WriteString(line)
	sb.WriteString("][")
	sb.WriteString(funcName)
	sb.WriteString("]")
	sb.WriteString(msg)
	logChannel <- sb.String()
	StopLogger()
	os.Exit(1)
}
