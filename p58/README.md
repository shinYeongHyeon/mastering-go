# Using standard output

Standard out 은 더도 말고 덜도 말고 화면에 출력하는 것이다.  
그러나, standard output 을 사용하는 것은 `fmt` 패키지에 포함되지 않는 것이 있을 수도 있기에 이 섹션이 존재한다.  

세 가지 부분으로 설명될 예정이며 해당 파일은 `stdOUT.go` 이다.  
첫번째 파트는 아래와 같다.

```go
// stdOUT.go
package main

import (
	"io"
	"os"
)
```

`fmt` 패키지 대신 `io` 패키지를 사용할 것이다.  
`os` 패키지는 command-line 의 인자를 읽늗네 사용된다.  

두 번째 부분은 하기와 같다.

```go
func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me on argument!"
	} else {
		myString = arguments[1]
	}
```

`myString` 변수는 화면에 출력될 command-line 인자의 첫번째를 가지고 있거나, 인자가 없을 경우 하드코딩된 메시지를 출력한다.  

세 번째 부분은 하기와 같다.

```go
    io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
```

이 경우 io.WriteString() 은 fmt.Print() 와 동일하다.
그러나, WriteString() 은 어디에 쓰길 원하는지에 대해 첫 번째 파라미터로 전달할 수 있고,
파일 입력에 사용이 가능할 것이다.

파일을 실행하면 아래와 같은 화면들을 볼 수 있다.

```shell
$ go run stdOUT.go
Please give me one argument!
$ go run stdOUT.go 123 12
123
```
