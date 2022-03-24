# About error output

이번에는 **UNIX standard error** 를 보여주는 방법에 대해서 알아볼 것이다.  

`stdERR.go` 에서 Go 에서의 standard error 의 사용에 대해 설명할 것이고 두 가지 부분으로 나누어 설명할 것이다.  
standard error 를 사용하기 위해 파일 설명자를 이용할 것이고 이는 `stdOUT.go` 에 기반해 있다.  

첫 번째 부분을 보자.

```go
package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
    } else {
		myString = arguments[1]
    }
}
```
