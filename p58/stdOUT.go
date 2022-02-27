package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me on argument!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}

/*
이 경우 io.WriteString() 은 fmt.Print() 와 동일하다.
그러나, WriteString() 은 어디에 쓰길 원하는지에 대해 첫 번째 파라미터로 전달할 수 있고,
파일 입력에 사용이 가능할 것이다.
*/
