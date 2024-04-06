# Hello Go World

## 코드 뜯어보기

```go
package main

import "fmt"

func main() {
	// Hello Go World 출력
	fmt.Println("Hello Go World")
}
```

가장 기본이 되는 코드를 뜯어보자.

### 1/ package main

패키지 선언으로, 이 코드가 어떤 패키지에 속하는지 알려준다.  
패키지는 코드 묶음으로 여러 기능을 제공하고, Go 언어의 모든 코드는 반드시 패키지 선언으로 시작해야 한다.  

그 중에서 `main` 패키지는 프로그램의 시작점을 포함하는 특별한 패키지이다.  
`main()` 함수가 없는 패키지는 패키지 이름으로 `main` 을 쓸 수 없다.  

⭐️ Go 언어는 패키지 선언으로 시작되고, `package main` 은 프로그램의 시작점이다.

### 2/ import "fmt"

fmt 패키지를 가져오는 것이다.
`fmt` 패키지는 표준 입출력을 다루는 내장 패키지이다.  


