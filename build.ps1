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
  SET CGO_ENABLED=0
  SET GOARCH=$GOARCH
  SET GOOS=$GOOS
  if ($GOOS -eq 'windows') {
    go build -o ./build/$Name/tyut-net-connector.exe .
    cp ./startup.cmd ./build/$Name/tyut-net-connector-startup.cmd
  } else {
    go build -o ./build/$Name/tyut-net-connector .
    cp ./startup.sh ./build/$Name/tyut-net-connector-startup.sh
    wsl -e chmod +x ./build/$Name/tyut-net-connector-startup.sh
  }
  cp ./README.md ./build/$Name/
  cd ./build
  wsl -e zip -r ./tyut-net-connector-$Name-$version.zip ./$Name
  rm -Recurse -Force ./$Name
  cd ../
}

rm -Recurse -Force ./build
mkdir ./build

Build -GOARCH amd64 -GOOS windows -Name win64
# Build -GOARCH amd64 -GOOS darwin -Name osx-intel64
# Build -GOARCH amd64 -GOOS linux -Name linux-amd64
# Build -GOARCH arm64 -GOOS linux -Name linux-arm64

echo 'BUILD DONE'