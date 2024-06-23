package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

var ErrType = fmt.Errorf("")

type any = interface{}
type bytes []byte
type str string
type StringSlice []string

func (b bytes) Decode() string {
	return string(b)
}

func (s str) PadStart(targetLength int, padString string) string {
	format := fmt.Sprintf("%%%s%ds", padString, targetLength)
	return fmt.Sprintf(format, s)
}

func (s StringSlice) Contain(value string) bool {
	for _, val := range s {
		if val == value {
			return true
		}
	}
	return false
}

type Logger struct {
	Scope    string
	Level    string
	Filename string
}

func (l *Logger) Info(format, message string) {
	var filename string

	if l.Scope == "global" {
		filename = path.Join(os.TempDir(), l.Filename)
	} else if l.Scope == "local" {
		pwd, _ := os.Getwd()
		filename = path.Join(pwd, l.Filename)
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	fmt.Fprintf(file, "%s\r\n", message)
}

func Log(values ...any) {
	fmt.Fprintln(os.Stdout, values...)
}

func Bin(value int) string {
	return fmt.Sprintf("%#b", value)
}

func Chr(code int) string {
	return fmt.Sprintf("%c", code)
}

func Input(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')

	return strings.TrimSpace(data)
}

func Tree(dirname string, excludes StringSlice) {
	entries, err := os.ReadDir(dirname)

	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if !excludes.Contain(entry.Name()) {
				Tree(path.Join(dirname, entry.Name()), excludes)
			}
		} else {
			info, _ := entry.Info()
			fmt.Printf("|-- %s %d\n", info.Name(), info.Size())
		}
	}
}

func Exec(script string) {
	process, err := os.StartProcess("D:/Git/bin/bash.exe", []string{"deploy.sh"}, &os.ProcAttr{})

	if err != nil {
		return
	}

	time.AfterFunc(time.Second*3, func() {
		process.Signal(os.Kill)
	})
}

func GetGoPath() string {
	return os.Getenv("GOPATH")
}

func Now() string {
	return time.Now().Format("2006年01月02日 15时04分05秒")
}

func SetTimeout(callback func(), seconds time.Duration) {
	time.Sleep(time.Second * seconds)
	callback()
}

func SetInterval(callback func(time.Time), seconds time.Duration) {
	for current := range time.Tick(time.Second * seconds) {
		callback(current)
	}
}
