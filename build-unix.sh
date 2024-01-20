go get
mkdir build
go build -o build/tyut-net-connector .
cp startup.sh build/tyut-net-connector-startup.sh
chmod +x build/*
cp README.md build/
