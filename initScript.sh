set -x
set -e
export GOPATH="$(realpath ../../../..)"
mkdir $GOPATH/bin
curl -sSL https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
cat << EOF
{
  "envs": {
    "PATH": "$PATH:$GOPATH/bin",
    "GOPATH": "$GOPATH"
  }
}
EOF > $2