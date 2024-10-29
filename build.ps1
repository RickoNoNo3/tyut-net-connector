param (
  [string]$version
)

if ([string]::IsNullOrEmpty($version)) {
  echo 'Version need'
  exit -1
}

function Build {
  param (
    $GOARCH,
    $GOOS,
    $Name
  )
  mkdir ./build/$Name
  go env -w CGO_ENABLED=0
  go env -w GOARCH=$GOARCH
  go env -w GOOS=$GOOS
  if ($GOOS -eq 'windows') {
    go build -o ./build/$Name/tyut-net-connector.exe .
    cp ./startup.cmd ./build/$Name/tyut-net-connector-startup.cmd
  } else {
    go build -o ./build/$Name/tyut-net-connector .
    cp ./startup.sh ./build/$Name/tyut-net-connector-startup.sh
  }
  cp ./README.md ./build/$Name/
  cp ./encrypt.sh ./build/$Name/
  cp ./easy-connect.sh ./build/$Name/
  cd ./build
  Compress-Archive -Path ./$Name/* -DestinationPath ./tyut-net-connector-$Name-$version.zip
  rm -Recurse -Force ./$Name
  cd ../
}

rm -Recurse -Force ./build
mkdir ./build

Build -GOARCH amd64 -GOOS windows -Name win64
Build -GOARCH amd64 -GOOS linux -Name linux_amd64
Build -GOARCH arm64 -GOOS linux -Name linux_arm64
Build -GOARCH amd64 -GOOS darwin -Name osx_amd64
Build -GOARCH arm64 -GOOS darwin -Name osx_arm64

echo 'BUILD DONE'

go env -w CGO_ENABLED=1
go env -w GOARCH=amd64
go env -w GOOS=windows