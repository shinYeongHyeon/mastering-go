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

보다 시피 `stdERR.go` 는 `stdOUT.go` 와 많이 비슷하다.

두 번째 부분을 보자

```go
    io.WriteString(os.Stdout, "This is Standard output\n")
    io.WriteString(os.Stderr, myString)
    io.WriteString(os.Stderr, "\n")
}
```
`io.WriteString()` 을 standard error 를 작성하기 위해 두번 불렀고, standard output 을 작성하기 위해 한번 불렀다.
이제 이를 실행해보면 아래와 같은 결과를 본다.

```shell
This is Standard output
Please give me one argument!
```

해당 결과는 이전에 작성된 standard output 과 차이점을 찾는데 어렵다.  
그러나, bash shell 을 이용하면 standard output 과의 차이점을 알 수 있다.  
거의 모든 UNIX shell 은 그들의 방법으로 이 기능을 제공한다.  

bash 을 사용할 때, standard error output 을 파일로 전송할 수 있다.

```shell
$ go run stdERR.go 2>/tmp/stdError                                                                                                                                      ─╯
This is Standard output

$ cat /tmp/stdError                                                                                                                                                     ─╯
Please give me one argument!
```

> 이 뒤 내용은 이해하지 못하였다.. 나중에 다시 보자