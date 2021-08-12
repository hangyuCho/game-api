# go

go env -w GO111MODULE=auto\n
go build main.go
go run main.go

go mod init github.com/hangyuCho/game-api 
go mod tidy

error:
    in:
       swag i
   out:
       zsh: command not found: swag

solution:
    step1:
       open ~/.zshrc

    step2:
       # enter the next line
       export "/Users/{username}/go/bin:$PATH"

~/go/bin/swag init
2021/08/13 00:27:06 Generate swagger docs....
2021/08/13 00:27:06 Generate general API Info
2021/08/13 00:27:06 create docs.go at  docs/docs.go
 1g1204@1g1204ui-MacBook-Pro  ~/go/src/github.com/hangyuCho/game-api   develop  pwd
/Users/1g1204/go/src/github.com/hangyuCho/game-api
 1g1204@1g1204ui-MacBook-Pro  ~/go/src/github.com/hangyuCho/game-api   develop  