cd src/vfsgen/ && go test && cd ../..
GOOS=darwin GOARCH=amd64 cd src/ && go build . && mv src ../bin/cloud-events-generator-darwin-amd64 && cd ..
GOOS=darwin GOARCH=386 cd src/ && go build . && mv src ../bin/cloud-events-generator-darwin-386 && cd ..
GOOS=linux GOARCH=amd64 cd src/ && go build . && mv src ../bin/cloud-events-generator-linux-amd64 && cd ..
GOOS=linux GOARCH=386 cd src/ && go build . && mv src ../bin/cloud-events-generator-linux-386 && cd ..
GOOS=windows GOARCH=amd64 cd src/ && go build . && mv src ../bin/cloud-events-generator-windows-amd64.exe && cd ..
GOOS=windows GOARCH=386 cd src/ && go build . && mv src ../bin/cloud-events-generator-windows-386.exe && cd ..