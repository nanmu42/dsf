.PHONY: config clean dir all dsf cross

all: dsf

dir:
	@mkdir -p bin

clean:
	@rm -rf bin

dsf: dir
	@go build -trimpath -ldflags "-s -w" -o bin/

cross: dir
	@GOOS=linux GOARCH=amd64 go build -o bin/dsf-linux-amd64 -trimpath -ldflags "-s -w" && \
	GOOS=linux GOARCH=arm64 go build -o bin/dsf-linux-arm64 -trimpath -ldflags "-s -w"  && \
	GOOS=windows GOARCH=amd64 go build -o bin/dsf-windows-amd64.exe -trimpath -ldflags "-s -w" && \
	GOOS=darwin GOARCH=amd64 go build -o bin/dsf-darwin-amd64 -trimpath -ldflags "-s -w" && \
	GOOS=darwin GOARCH=arm64 go build -o bin/dsf-darwin-arm64 -trimpath -ldflags "-s -w"