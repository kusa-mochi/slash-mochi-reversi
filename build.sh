#!/bin/bash
# Run this shell and upload the zip file to AWS lambda.
# 2021/5/22 mochi

srcDir=./
distDir=./dist
handlerName=reversi_nextmove_handler

[[ -d $distDir ]] && rm -R $distDir
mkdir $distDir
GOOS=linux GOARCH=amd64 go build -o $handlerName
zip $distDir/$handlerName.zip $handlerName
rm $handlerName

exit 0
