cd src/vfsgen/ && go test && cd ../..
cd src/ && gox -os="linux darwin windows" -arch="amd64 386"
mv src_darwin_386 ../bin/cloud-events-generator-darwin-386
mv src_darwin_amd64 ../bin/cloud-events-generator-darwin-amd64
mv src_linux_386 ../bin/cloud-events-generator-linux-386
mv src_linux_amd64 ../bin/cloud-events-generator-linux-amd64
mv src_windows_386.exe ../bin/cloud-events-generator-windows-386.exe
mv src_windows_amd64.exe ../bin/cloud-events-generator-windows-amd64.exe
cd ..
