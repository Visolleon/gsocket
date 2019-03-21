@ECHO OFF
for /f "tokens=2 delims==" %%a in ('wmic OS Get localdatetime /value') do set "dt=%%a"
set "YY=%dt:~2,2%" & set "YYYY=%dt:~0,4%" & set "MM=%dt:~4,2%" & set "DD=%dt:~6,2%"
set "HH=%dt:~8,2%" & set "Min=%dt:~10,2%" & set "Sec=%dt:~12,2%"
REM set datestamp=%YYYY%%MM%%DD%" & set "timestamp=%HH%%Min%%Sec%
REM set fullstamp=%YYYY%%MM%%DD%%HH%%Min%%Sec%
set fullstamp=%YYYY%%MM%%DD%%HH%%Min%

ECHO  ====Windows cross-build to Linux 64====
ECHO ------------go env---------------
set GOARCH=amd64
set GOBIN=
set GOEXE=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOOS=linux
set GOPATH=%GOPATH%;%CD%
set GORACE=
set GO15VENDOREXPERIMENT=1
set CC=gcc
set GOGCCFLAGS=-fPIC -m64 -fmessage-length=0
set CXX=g++
set CGO_ENABLED=0

go env
ECHO -----------go get----------------
go get -v
ECHO -----------go build----------------
go build -v -ldflags "-s -w -X main.Version='%fullstamp%'"