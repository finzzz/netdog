build:
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o dist/nd
	GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o dist/nd.exe

compress:
	ls dist/ | xargs -n 1 -I {} upx dist/{}