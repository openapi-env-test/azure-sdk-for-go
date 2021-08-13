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
pwsh ./eng/scripts/build.ps1 -skipBuild -format -tidy -generate armpostgresql
cat > $2 << EOF
{
  "envs": {
    "PATH": "$PATH:$GOPATH",
    "GOPATH": "$GOPATH"
  }
}
EOF