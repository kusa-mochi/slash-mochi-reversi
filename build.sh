#!/bin/bash
# Run this shell and upload the zip file to AWS lambda.
# 2021/5/22 mochi

distDir=./dist

[[ -d $distDir ]] && rm -R $distDir
mkdir $distDir
GOOS=linux GOARCH=amd64 go build -o $distDir/NextReversiMove
zip $distDir/NextReversiMove.zip $distDir/NextReversiMove

exit 0
