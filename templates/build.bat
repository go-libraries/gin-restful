@echo off
set nowDir=%~dp0
set GO111MODULE=auto
set GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=auto
go env -w GOMOD=%nowDir%"/go.mod"
go env -w GOPROXY=https://goproxy.cn,direct

go mod init {{package}}
go mod vendor
go run main.go