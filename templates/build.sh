#/bin/bash
current_path=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)
cd $current_path;
export GO111MODULE="auto"
export GOMOD=${current_path}"/go.mod"

if [ ! -f "go.mod" ]; then
  go mod init {{package}}
fi

if [ ! -d "vendor" ];then
  go mod vendor
fi

go run main.go