package = github.com/fcingolani/fibertel-stats

.PHONY: release

release:
	go get
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/fibertel-stats-linux-amd64 $(package)
	GOOS=linux GOARCH=386 go build -o release/fibertel-stats-linux-386 $(package)
	GOOS=darwin GOARCH=amd64 go build -o release/fibertel-stats-darwin-amd64 $(package)
	GOOS=darwin GOARCH=386 go build -o release/fibertel-stats-darwin-386 $(package)
	GOOS=windows GOARCH=amd64 go build -o release/fibertel-stats-windows-amd64.exe $(package)
	GOOS=windows GOARCH=386 go build -o release/fibertel-stats-windows-386.exe $(package)