# About printing output

UNIX, C, Go 는 화면에 출력하기 위한 여러가지 방법을 지원한다.  
이번에 설명하는 모든 것은 `printing.go` 에 작성될것이고 두가지 part 로 구분되며 Go 의 표준 `fmt` 패키지로 출력할 것이다.  

Go 에서 무언가를 출력하는 가장 간단한 방법은 `fmt.Println()` 과 `fmt.Printf()` 함수가 있다.  
`fmt.Printf()` 함수는 C 의 `printf(3)` 과 많은 유사성을 지니고 있다.  
또한 `fmt.Println()` 대신에 `fmt.Print()` 를 사용할 수 있다.  
`fmt.Print()` 와 `fmt.Println()` 의 주요 차이점은 각각 호출된 이후에 **자동적으로 개행문자를 추가하느냐** 이다.  

반면에 `fmt.Println()`과 `fmt.Printf()` 의 가장 큰 차이점은 원하는 것을 출력하려고 할때 **format specifier** 같은 것들이 필요하느냐 이다.  
C의 `printf(3)` 과 마찬가지로 수행하려는 작업을 더 잘 제어할 수 있지만, 더 많은 코드를 작성해야 한다.  
Go 는 format specifier 를 동사(verbs) 라고 부른다.  
이에 대해서는 [https://golang.org/pkg/fmt](https://golang.org/pkg/fmt) 에서 더 많은 verbs 정보를 확인할 수 있다.  

만약 어떤 것을 출력하기 전에 formatting 을 할 생각이라면, 여러 변수를 정렬할 생각이라면 `fmt.Printf()` 를 쓰는 것이 좋은 선택일 수 있다.  
그러나, 한 변수만 출력할 예정이라면 `fmt.Print()` 혹은 `fmt.Println()` 중 개행 여부에 따라서 선택하면 된다.  

`printing.go` 의 첫 번째 부분부터 보자.

```go
// printing.go

package main

import "fmt"

func main() {
	v1 := "123"
	v2 := 123
	v3 := "Have a nice day\n"
	v4 := "abc"
```

이 부분에서, `fmt` 패키지를 import 한 것을 볼 수 있다.  
v3 에 사용된 `\n` 은 개행 문자이다.  
그러나, 별다른 인자 없이 출력에 개행을 원하면 `fmt.Println()` 을 쓰면 된다.  

두 번째 부분을 보자.
```go
// printing.go SecondPart

fmt.Print(v1, v2, v3, v4)
fmt.Println()
fmt.Println(v1, v2, v3, v4)
fmt.Print(v1, " ", v2, " ", v3, " ", v4, "\n")
fmt.Printf("%s%d %s %s\n", v1, v2, v3, v4)
}
```

이 부분에서는 4 변수를 각종 Print 함수로 사용한다.  
`printing.go` 를 사용하면 아래의 결과를 볼 수 있다.  

```shell
123123Have a nice day
abc
123 123 Have a nice day
 abc
123 123 Have a nice day
 abc
123123 Have a nice day
 abc
```

--

* `format specifier`: %s, %d 와 같은 형식 지정자  
* [형식 지정자 목록](../../appendix/fmt/FormatOfPrintf.md)