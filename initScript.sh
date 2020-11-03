set -x
set -e
export GOPATH="$(realpath ../../../..)"
if [ ! -d "$GOPATH/bin" ]
then
  mkdir $GOPATH/bin
fi
curl -sSL https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
dep ensure
cat > $2 << EOF
{
  "envs": {
    "PATH": "$PATH:$GOPATH/bin:$GOPATH",
    "GOPATH": "$GOPATH"
  }
}
EOF