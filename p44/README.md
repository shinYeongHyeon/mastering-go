# Compiling Go code

Go code 를 어떻게 컴파일하는지 배워보자.  
Graphical application 없이 command line 에서 Go code 를 컴파일 할 수 있다.  
게다가, `main` 이라는 이름을 가진 package 안에 `main() function` 만 있다면 파일의 이름에 제한이 없다.  
왜냐하면, `main() function` 이 프로그램의 시작점이기 때문이다.  
그 결과, 여러개의 `main() function` 을 한 파일 혹은 한 프로젝트에 지닐 수 없다.

`aSourceFile.go` 라는 이름을 가진 파일을 시작으로 첫 Go program 을 작성해보자.

```go
package main

import "fmt"

func main() {
	fmt.Println("This is a sample Go Program!")
}
```

참고로, Go community 는 `aSourceFile.go` 보다는 `source_file.go` 의 이름을 더 선호한다.  
나는 `a_source_file.go` 로 작성했다.  
`a_source_file.go` 를 컴파일 하기 위해서, 그리고 `statically linked (정적으로 연결된)` 실행 파일을 만들기 위해서는 아래 command 를 실행해야 한다.  

```shell
$ go build a_source_file.go
```

그러고 나면, `a_source_file` 이라는 이름을 가진 실행 파일이 생성될 것이다.

```shell
$ file a_source_file
a_source_file: Mach-0 64-bit executable x86_64
$ ls -l a_source_file
-rwxr-xr-x  1 sin-yeonghyeon  staff  1869504 Feb 19 01:42 a_source_file
$ ./a_source_file
This is a sample Go Program!
```

`a_source_file` 의 용량이 큰 주된 이유는 정적으로 연결됐기 때문인데, 이는 파일을 실행하기 위한 외부 라이브러리가 전혀 필요없다는 것을 의미한다.
