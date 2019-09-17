cd src/vfsgen/ && go test && cd ../..
cd src/ && gox -os="linux darwin windows" -arch="amd64 386"
mv src_darwin_386 ../bin/cloudevents-generator-darwin-386
mv src_darwin_amd64 ../bin/cloudevents-generator-darwin-amd64
mv src_linux_386 ../bin/cloudevents-generator-linux-386
mv src_linux_amd64 ../bin/cloudevents-generator-linux-amd64
mv src_windows_386.exe ../bin/cloudevents-generator-windows-386.exe
mv src_windows_amd64.exe ../bin/cloudevents-generator-windows-amd64.exe
cd ..
