# There is only one way to format curly braces

```go
// curly.go
package main 
 
import ( 
    "fmt" 
) 
 
func main()
{
    fmt.Println("Go has strict rules for curly braces!")
}
```

위 코드를 보면 괜찮아 보이지만, 실행하고 나면 이내 출력되는 문법 에러에 실망하게 될 것이다. 
```shell
./curly.go:7:6: missing function body
./curly.go:8:1: syntax error: unexpected semicolon or newline before {
```

해당 오류에 대한 공식적인 설명은 Go가 여러 context 에서 세미콜론을 종결자로 사용해야 한다는 것이고 컴파일러는 세미콜론이 필요하다고 생각될 때 **자동으로** 세미콜론을 삽입한다는 것이다.  
따라서 중괄호을 열면 Go 컴파일러 이전 라인 끝에 (`func main()`) 세미콜론을 넣게되며 이는 메시지의 원인이 된다.
