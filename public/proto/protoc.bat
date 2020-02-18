@echo off
protoc -I=./ --go_out=../pb ./*.proto

pause

