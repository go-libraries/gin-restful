#/bin/bash
current_path=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

cd $current_path;

export GO111MODULE="auto"
export GOMOD=${current_path}"/go.mod"

if [ -f "go.mod" ]; then
  rm -f go.mod
fi

if [ -f "go.sum" ]; then
  rm -f go.sum
fi

if [ ! -d "vendor" ];then
  rm -rf vendor
fi

go mod init {{package}}
go mod vendor
go run main.go