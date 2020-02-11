#/bin/bash
current_path=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

export GO111MODULE="auto"
export GOMOD=${current_path}"/go.mod"

rm go.sum

go get github.com/swaggo/swag/cmd/swag
# ${GOPATH}/bin/swag init
go mod init {{package}}
go mod vendor
go run main.go