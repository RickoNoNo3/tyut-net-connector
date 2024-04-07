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
  go build -o ./build/$Name/tyut-net-connector.exe .
  cp ./startup.cmd ./build/$Name/tyut-net-connector-startup.cmd
  cp ./README.md ./build/$Name/
  cd ./build
  Compress-Archive -Path ./$Name/* -DestinationPath ./tyut-net-connector-$Name-$version.zip
  rm -Recurse -Force ./$Name
  cd ../
}

rm -Recurse -Force ./build
mkdir ./build

Build -GOARCH amd64 -GOOS windows -Name win64

echo 'BUILD DONE'