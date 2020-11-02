set -x
set -e
$GOPATH="$(realpath ../../../..)"
curl -sSL https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
cat << EOF
{
  "envs": {
    "PATH": "$PATH:$GOPATH/bin",
    "GOPATH": "$GOPATH"
  }
}
EOF > $2