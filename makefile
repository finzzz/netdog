build:
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o dist/nd
	GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o dist/nd.exe

compress:
	cp dist/nd dist/nd_upx
	cp dist/nd.exe dist/nd_upx.exe
	ls dist/nd_upx* | xargs -n 1 upx --ultra-brute