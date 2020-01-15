@echo off
set nowDir=%~dp0
set GO111MODULE=on
set GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=on
go env -w GOMOD=%nowDir%"/go.mod"
go env -w GOPROXY=https://goproxy.cn,direct
if exist go.sum (
    del go.sum
)

:: go get github.com/swaggo/swag/cmd/swag
:: %GOPATH%/bin/swag init

go mod vendor
go run main.go