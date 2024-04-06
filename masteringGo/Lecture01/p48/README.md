# You either use a Go package or you do not include it

Go 는 package 사용에 있어서 엄격한 규칙을 가지고 있다.  
그러므로, 나중에 필요할 것 같다라던지, 사용하지 않는 package 를 그냥 include 할 필요가 없다.

아래 순진한 코드를 보자.

```go
// packageNotUsed.go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello there !")
}
```

해당 파일을 실행하면 아마 `imported not used: "os"` 와 같은 메시지를 받게 될 것이다.  
해당 패키지를 지우게 되면 정상적으로 동작한다.  

그러나, 이것에 대해 우회할 수 있는 방법이 있긴하다.

```go
// packageNotUsedUnderScore.go
package main

import (
	"fmt"
	_ "os"
)

func main() {
	fmt.Println("Hello there !")
}
```

Underscore 를 통해서 에러 메시지가 출력되지 않게 할 수 있고, 정상적으로 동작이 가능하다.
