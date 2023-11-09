rm -fo ./build
mkdir ./build
mkdir ./build/win_amd64
mkdir ./build/osx_amd64
mkdir ./build/linux_amd64
mkdir ./build/linux_arm64

SET CGO_ENABLED=0
SET GOARCH=amd64
SET GOOS=windows
go build -o ./build/win_amd64/tyut-net-connector.exe .
cp ./README.md ./build/win_amd64/
cp ./startup.cmd ./build/win_amd64/
wsl -e zip -r ./build/tyut-net-connector-win64.zip ./build/win_amd64

SET GOOS=darwin
go build -o ./build/osx_amd64/tyut-net-connector .
cp ./README.md ./build/osx_amd64/
cp ./startup.sh ./build/osx_amd64/
wsl -e chmod +x ./build/osx_amd64/startup.sh
wsl -e zip -r ./build/tyut-net-connector-osx-intel64.zip ./build/osx_amd64

SET GOOS=linux
go build -o ./build/linux_amd64/tyut-net-connector .
cp ./README.md ./build/linux_amd64/
cp ./startup.sh ./build/linux_amd64/
wsl -e chmod +x ./build/linux_amd64/startup.sh
wsl -e zip -r ./build/tyut-net-connector-linux-amd64.zip ./build/linux_amd64

SET GOARCH=arm64
SET GOOS=linux
go build -o ./build/linux_arm64/tyut-net-connector .
cp ./README.md ./build/linux_arm64/
cp ./startup.sh ./build/linux_arm64/
wsl -e chmod +x ./build/linux_arm64/startup.sh
wsl -e zip -r ./build/tyut-net-connector-linux-arm64.zip ./build/linux_arm64