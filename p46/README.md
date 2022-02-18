# Executing Go code

실행 파일을 만들지 않고 Go 코드를 실행할 수 있는 다른 방법이 존재한다. - *사실은 실행 파일을 만들고 바로 사라지는 것이다.*  
> 이 방식이 Python, Ruby, Perl 과 같은 스크립트 프로그래밍 언어처럼 Go를 사용할 수 있게 해준다.

그래서, `a_source_file.go` 를 실행파일 없이 실행하기 위해서는, 아래의 command 를 실행해야 한다.  

```shell
$ go run a_source_file.go
This is a sample Go Program!
```

보다 시피, command 에 출력물이 이전과 동일하게 나오고 있다.  

> 하지만 분명히 기억해야 할 것은, `go run` 도 실행파일은 필요하다는 것이다. 볼 순 없어도 자동으로 실행이 됐고, 그 후 자동으로 삭제되는 것이다 

더 간단하기 때문에, 주로 `go run` 을 통해 예제를 진행해나갈 것이다.  
더해, `go run` 은 어떤 프로그램도 여러분의 하드디스크에 남기지 않습니다. 
