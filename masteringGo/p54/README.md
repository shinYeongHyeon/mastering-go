# UNIX stdin, stdout, and stderr

모든 UNIX OS 는 프로세스를 위해 항상 열려 있는 3가지 파일이 있습니다.  
UNIX 는 프린터나 마우스 등 모든 것을 파일로 간주한다.  

UNIX는 열려 있는 모든 파일에 액세스하기 위한 내부 표현으로 양의 정수 값인 파일 설명자(file descriptor)를 사용합니다. 이는 긴 경로를 사용하는 것보다 훨씬 좋습니다.  

따라서 기본적으로 모든 UNIX 시스템은 `/dev/stdin`, `/dev/stdout`, `/dev/stderr` 의 세 가지 특수 및 표준 파일 이름을 지원하며 각각 파일 설명자 0, 1, 2를 사용하여 액세스할 수도 있습니다.
위 세가지 파일 설명자는 **standard input**, **standard output**, **standard error** 라고 불린다.  
추가적으로, MacOS 에서 파일 설명자 0은 `/as/fd/0` 으로 불리고 Debian Linux 에서는 `/dev/fd/0` 와 `/dev/pts/0` 이라고 불린다.

Go 는 `os.Stdin` 을 standard input 으로 사용하고, `os.Stdout` 과 `os.Stderr` 로 standard output 과 standard error 로 사용한다.  

동일한 디바이스에 액세스하는 데 /dev/stdin, /dev/stdout 및 /dev/stderr 또는 관련 파일 설명자 값을 사용할 수 있지만 Go 에서 제안한 `os.Stdin/Stdout/Stderr` 를 사용하는 것이 더 낫고 안전하며 휴대성이 더 낫다.  
