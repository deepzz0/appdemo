#!/usr/bin/env sh

set -e

# ensuer new app?
read -p "Are you sure new app? [y/n] " ensure
if [ "$ensure" != "y" ]; then
  exit 0
fi

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

_contains() {
  _str="$1"
  _sub="$2"
  echo "$_str" | grep -- "$_sub" >/dev/null 2>&1
}

printf 'Project [\33[1;32m%b\33[0m], initializing...\n' "$appname"

# rename appname
_sed_i "1s/demo/$appname/" "conf/app.yml"
echo "Clean conf/app.yml"

# rename module
_sed_i "1s/$old/$new/" "go.mod"
echo "Clean go.mod"

# rename package ref
find . -name "*.go" | while read fname; do
  _sed_i "s/$old/$new/g" "$fname"
  if _contains "$fname" "pkg/config"; then
    # special file
    _path=$(echo $new | sed 's/\\\//", "/')
    _sed_i "s/\"github.com\", \"deepzz0\", \"appdemo\"/\"$_path\"/g" \
      "$fname"
  fi
done
echo "Clean *.go"

# clean git
rm -rf .git
echo "Clean git repo"

# init empty repo
git init

printf '\33[1;32m%b\33[0m' "Successful initialization."
