# Downloading Go packages

Go 의 표준라이브러리들이 굉장히 풍부하긴 하지만, 외부 Go package 들을 다운받아서 이용하는게 도움이 될 때가 많다.  
`getPackage.go` 로 된 하기 프로그래밍을 보자.   

```go
// getPackage.go
package main

import (
	"fmt"
	"github.com/mactsouk/go/simpleGitHub"
)

func main() {
	fmt.Println(simpleGitHub.AddTwo(5, 6))
}
```

위 프로그램은 `import` 를 인터넷 주소처럼 사용해, 외부 package 를 이용한다. `simpleGiHub` 이라 불리는 package 를 사용하고, 이는 `github.com/mactsouk/go/simpleGitHub` 에 위치해있다.  

만약 바로 실행하려고 한다면, 패키지를 찾을 수 없다는 오류를 보게 될 것이다.  
그래서, 패키지를 컴퓨터에 다운받아야 한다.  다운 받기 위해서는 하기 명령어를 입력한다.  

```shell
$ go get -v github.com/mactsouk/go/simpleGitHub
```

이제 프로그램을 실행하는데 문제가 없다.  
하기 명령어를 통해 package 삭제가 가능하다.

```shell
$ go clean -i -v -x github.com/mactsouk/go/simpleGitHub
```
