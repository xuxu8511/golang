set GOPATH=E:\my\golang
go install -v base/common
go install -v base/network
go install -v base/log4go

cd bin
go build -v socket
go build -v client

pause