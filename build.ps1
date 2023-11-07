rm -fo ./build
mkdir ./build
mkdir ./build/osx_amd64
mkdir ./build/linux_amd64
mkdir ./build/win_amd64

SET CGO_ENABLED=0
SET GOARCH=amd64
SET GOOS=windows
go build -o ./build/win_amd64/tyut-net-connector.exe .

SET GOOS=darwin
go build -o ./build/osx_amd64/tyut-net-connector .

SET GOOS=linux
go build -o ./build/linux_amd64/tyut-net-connector .
