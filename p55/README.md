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

--

* `format specifier`: %s, %d 와 같은 형식 지정자  
