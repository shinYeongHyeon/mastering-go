# 변수

변수란, 값을 저장하는 메모리 공간을 가리키는 `이름` 이다.  

```go
package main

import "fmt"

func main() {
	var a int = 10
	var msg string = "Hello variable"
	
	a = 20
	msg = "Good morning"
	fmt.Println(msg, a)
}
```

변수를 사용하려면 변수를 선언해야 한다.  
변수 선언은 컴퓨터에게 값을 저장할 공간을 마련하라고 명령하는 것과 같다. 이것을 메모리 할당이라고 부른다.  

`var a int = 10`  
여기에서 `var` 는 변수 선원 키워드로 `variable` 의 약자이다.  
이어서 `a` 라는 변수`명` 을 적는다.  
그 다음 `int` 라는 타입을 적습니다.  

## 또 다른 형태의 변수선언 

이 외에도 `var b int` 와 같이 초깃값을 생략하는 선언  
`var c = 4` 와 같이 변수 타입을 생략하는 선언 (타입은 할당 된 값과 같은 타입이 된다)  
` d := 5` 와 같이 선언 대입문 `:=` 을 통해 `var` 과 타입을 생략하는 선언방식이 있다.  
