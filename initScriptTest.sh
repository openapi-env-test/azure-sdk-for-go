set -x
set -e
export GOPATH="$(realpath ../../../..)"
if [ ! -d "$GOPATH/bin" ]
then
  mkdir $GOPATH/bin
fi
PATH=$PATH:$GOPATH/bin
export GO111MODULE=on
go get github.com/openapi-env-test/azure-sdk-for-go/tools/generator@track2_automation_and_release
cat > $2 << EOF
{
  "envs": {
    "PATH": "$PATH:$GOPATH",
    "GOPATH": "$GOPATH"
  }
}
EOF