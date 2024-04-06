# Reading From standard input

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)
```

위 코드를 보면, 처음 보는 `bufio` package 가 있을 것이다.  
`bufio` 는 파일 입/출력에 가장 많이 사용되는 package 이지만, 보통은 `os` package 를 보게 될텐데,  
그 이유는 보통 Command-line arguments 를 다루는데 사용되기 때문이다 (`os.Args`)

```go
// stdIN.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()
	
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}
```

먼저, `os.Stdin` 을 파라미터로 `bufio.NewScanner()` 를 호출하는데, 이는 `bufio.Scanner` 를 리턴해준다.  
이는 `Scan()` 을 통해 `os.Stdin` 을 라인마다 실행한다.  
각 라인은 텍스트를 읽어들이고, 출력해낸다.  
