# go
### !!! md 사진 붙여넣기 단축키 cmd + shift + option !!!
## commend 모음
``` go 
go env -w GO111MODULE=auto\n  // 모듈 관련 환경 설정
go build main.go // 실행 전 빌드 명령어
go run main.go // 실행 명령어
go get github.com/hangyuCho/game-api // 레포지토리로부터 소스 다운로드
go mod init github.com/hangyuCho/game-api // 모듈 등록
go mod tidy // 모듈 등록2
```

## bugfix
1. swagger의 명령어를 못찾는 문제
``` shell
# 실행 결과
$ go get github.com/swaggo/swag/cmd/swag
$ go mod tidy
$ swag
zsh: command not found: swag

# 해결법
# GOROOT의 위치에서 명령어를 실행 해줘야함..ㅜ
~/go/bin/swag init
```

2. bad.go:2:43: expected 'package', found 'EOF' 에러 발생

``` shell
# 실행 결과
$ pwd
~/go/bin/swag
$ ~/go/bin/swag init
$ 블라블라블라블라............bad.go:2:43: expected 'package', found 'EOF'

# 해결법
$ pwd
/Users/1g1204/go/src/github.com/hangyuCho/game-api/hello
$ ~/go/bin/swag init
2021/08/13 00:27:06 Generate swagger docs....
2021/08/13 00:27:06 Generate general API Info
2021/08/13 00:27:06 create docs.go at  docs/docs.go
```

3. swagger Fetch errorInternal Server Error doc.json 발생
``` shell
# 실행 결과
$ go run main.go # 실행 자체는 문제 없음..
# http://{host}:{port}/swagger/index.html 접속 시 아래 에러 발생
# Fetch errorInternal Server Error doc.json
# Parser error on line 14 end of the stream or a document separator is expected
```

4. swagger api 정보가 swagger.uaml에 등록안됨..
* 해결법.. 이랄까 문제는 코멘트 주석의 문제..
![](2021-08-14-21-01-05.png)

``` go
// 해결법
package main

import (
	_ "github.com/hangyuCho/game-api/docs" // << 도큐먼트 파일 출력 패스를 메인 패키지에 등록 해줘야 함..!
)
``` 

 ## 주석 사용 규칙 참고
 https://github.com/swaggo/swag#declarative-comments-format
 
``` shell


# 실행 결과

# 해결법

```