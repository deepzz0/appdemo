#!/usr/bin/env sh

set -e

os="$(uname -s)"
old=$(echo "github.com/deepzz0/appdemo" | sed "s/\//\\\\\//g")

tmp="$(pwd)"
new=$(echo "${tmp/$GOPATH\/src\//}" | sed "s/\//\\\\\//g")
appname="${tmp##*/}"

_sed_i() {
    option="$1"
    file="$2"
    if [ "$os" = "Darwin" ]; then
        sed -i "" "$option" "$file"
    else
        sed -i "$option" "$file"
    fi
}

printf 'Project [\33[1;32m%b\33[0m], initializing...\n' "$appname"

_sed_i "1s/demo/$appname/" "conf/app.yml"
echo "clean conf/app.yml"

_sed_i "1s/$old/$new/" "go.mod"
echo "clean go.mod"

find . -name "*.go" | while read fname; do
    _sed_i "s/$old/$new/g" "$fname"
done
echo "clean *.go"

rm -rf .git
echo "clean git repo"

git init

printf '\33[1;32m%b\33[0m' "Successful initialization."
