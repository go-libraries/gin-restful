#/bin/bash
current_path=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

export GO111MODULE="on"
export GOMOD=${current_path}"/go.mod"

rm go.sum

swag init

go run main.go