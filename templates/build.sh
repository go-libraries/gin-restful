#/bin/bash
current_path=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

export GO111MODULE="auto"
export GOMOD=${current_path}"/go.mod"

go mod init {{package}}
go mod vendor
go run main.go