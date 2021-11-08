#!/bin/sh
rm -rf ./dist
rm -rf ./static
tar -xvf dist.tar.xz
mv ./dist/static ./static
mv ./dist/index.html ./templates/index.html
sed -i 's/^/{{define "index"}}/' ./templates/index.html
sed -i 's/$/{{end}}/' ./templates/index.html