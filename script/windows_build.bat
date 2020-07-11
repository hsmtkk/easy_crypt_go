cd cmd\easycrypt

set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=0

go build -o easycrypt.exe

cd ..\..
