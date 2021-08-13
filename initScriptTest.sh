set -x
set -e
export GOPATH="$(realpath ../../../..)"
if [ ! -d "$GOPATH/bin" ]
then
  mkdir $GOPATH/bin
fi
PATH=$PATH:$GOPATH/bin
export GO111MODULE=on
cd tools/generator && go build && cp generator $GOPATH/bin && cd ../..
ln -s /usr/bin/pwsh $GOPATH/bin/pwsh.exe
pwsh -Command Get-ChildItem -recurse -path ./sdk -filter go.mod
ls sdk
cat > $2 << EOF
{
  "envs": {
    "PATH": "$PATH:$GOPATH",
    "GOPATH": "$GOPATH"
  }
}
EOF