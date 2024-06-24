package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
)

func main() {
	buf := bytes.NewBufferString("hi,golang")
	reader := bytes.NewReader([]byte("oh,ops"))

	buf.Write([]byte("\r\n"))

	data := make([]byte, 2)

	reader.Read(data)
	reader.ReadByte()

	fmt.Printf("%s %s %d %d\n", buf.Next(3), data, reader.Len(), reader.Size())

	buf.WriteTo(os.Stdout)

	fmt.Println(ReadFile("system.log", 1024), runtime.GOMAXPROCS(1), runtime.NumCPU())
	fmt.Println(runtime.GOOS)

}
