# The godoc utility

Go 는 프로그래머로서의 삶을 보다 쉽게 많들 수 있는 도구가 함께 제공됩니다.  
이 것들 중 하나가 `godoc` utility 인데, 이는 Go 함수들과 package 에 대한 Document 를 인터넷 연결이 없이도 볼 수 있게 해줍니다.  

`godoc` utility 는 일반적인 command-line application 에서 실행할 수 있고, 웹서버를 실행한 command-line application 에서도 실행할 수 있습니다.  
후자의 경우엔 Go document 를 보기위해서는 웹 브라우저가 필요합니다.  

```shell
$ go install golang.org/x/tools/cmd/godoc@latest
```

`fmt.Printf` 에 대해서 알아보자.  
shell 에 하기와 같이 입력한다.
```shell
$ go doc fmt.Printf 

package fmt // import "fmt"

func Printf(format string, a ...interface{}) (n int, err error)
    Printf formats according to a format specifier and writes to standard
    output. It returns the number of bytes written and any write error
    encountered.
```
