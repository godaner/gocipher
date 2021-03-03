pgk:pkglinuxamd64 pkglinux386 pkgwindowsamd64 pkgwindows386
# build
build:
	CGO_ENABLED=0 go build

# linux amd64 build and package
pkglinuxamd64:buildlinuxamd64 upxlinuxamd64
	tar -czf ./bin/gocipher_linux-amd64.tar.gz ./bin/gocipher
buildlinuxamd64:
	-rm ./bin/gocipher
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/gocipher
upxlinuxamd64:
	-upx -9 ./bin/gocipher

# linux 386 build and package
pkglinux386:buildlinux386 upxlinux386
	tar -czf ./bin/gocipher_linux-386.tar.gz ./bin/gocipher
buildlinux386:
	-rm ./bin/gocipher
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o ./bin/gocipher
upxlinux386:
	-upx -9 ./bin/gocipher

# windwos amd64 build and package
pkgwindowsamd64:buildwindowsamd64 upxwindowsamd64
	tar -czf ./bin/gocipher_windows-amd64.tar.gz ./bin/gocipher.exe
buildwindowsamd64:
	-rm ./bin/gocipher.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/gocipher.exe
upxwindowsamd64:
	-upx -9 ./bin/gocipher.exe

# windwos 386 build and package
pkgwindows386:buildwindows386 upxwindows386
	tar -czf ./bin/gocipher_windows-386.tar.gz ./bin/gocipher.exe
buildwindows386:
	-rm ./bin/gocipher.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ./bin/gocipher.exe
upxwindows386:
	-upx -9 ./bin/gocipher.exe