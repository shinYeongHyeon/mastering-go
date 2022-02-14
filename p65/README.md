# Working with command-line arguments

이번에 작성해볼 프로그램은 커맨드 라인의 arguments 중 최소값과 최대값을 구하는 것이다.  

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)
```

여기서 중요한 것은 커맨드라인 arguments 를 가져오는데 필요한 것이 `os` package 라는 것이고, 여기에 더해 `strconv` 라는 package 가 필요하다는 것이다.
커맨드라인 argument 를 변환하기 위해서 필요하기 때문이다.  

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)
}
```

위 프로그램을 보면 `os.Args` 를 통해 command-line arguments 의 길이를 구하고, 검사한다.    
여기서 짚고 넘어가야할 것은 `os.Args` 는 `string` 값들의 Go slice 라는 것이다 (slice 는 후에 자세히 알아본다. 배열이라고 생각하자 일단)  
slice 의 첫 번째 요소는 실행가능한 프로그램의 이름이다. 그렇기 때문에, `min`, `max` 의 값을 초기화 하기위해서는 두 번째 요소를 가져와야 한다.  

또한 여기서 더 알 수 있는 것은, argument 의 요소를 가져와 `ParseFloat` 을 한 리턴 값 중 min/max 와 더불어 `_` (underscore, called *blank identifier*) 가 있는데, 원래 이 자리는 `ParseFloat` 의 결과로 에러가 올 수 있다.  
즉, `ParseFloat` 은 Float 으로 Parse 한 값을 첫 번째 리턴 값으로, 두 번째로는 `Error` 를 리턴하게 되는데, 여기서 _ 는 이를 무시 한다는 것이다.  
Go 에서는 사용되지 않는 변수가 있으면  컴파일이 되지 않는데, 이를 위한 변수라고 보면 된다.  
실제 사용할 프로그램을 작성한다면 `Error` 가 왔을때에 대한 처리가 필요하게 될 것이다.

그러면 이제 나머지 코드를 작성해보자

```go
// cla.go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}
```