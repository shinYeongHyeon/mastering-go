# Writing to log files

`log` package 는 로그 메세지를 UNIX 의 시스템 로그 서비스로 전달할수 있게 해주고, `log` 패키지의 일부인 `syslog` 는 **logging level** 과 **logging facility** 를 정의하게 해준다.

UNIX 운영체제 대부분의 로그 파일 시스템은 `/var/log` 에 위치한다.  
그러나, 많은 서비스의 로그 파일 - Apache, Nginx 등 - 은 그들의 설정에 따라 어디든 있을 수 있다.

로그 파일은 화면에 같은 결과를 출력하기보다 어떤 정보들을 작성할때 사용되는데, 이에는 두가지 이유정도가 있다.  
첫째, 파일에 저장되면 해당 출력을 잃을 일이 없다.  
둘째, UNIX 도구 - grep, awk, sed 등 - 로 검색할 수 있다. 

`log` package 는 UNIX 에 출력을 전달하는데 도움을 주는 많은 함수들이 있다.  
가령, `log.Printf()`, `log.Fatalf()`, `log.Print()`, `log.Panicf()` 등등.
