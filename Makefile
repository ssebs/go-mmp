build:
	@rm main.exe
	go build main.go
run: build
	./main.exe
test:
	go test ./...  -coverpkg=./... -coverprofile ./coverage.out
	go tool cover -func ./coverage.out
pkg:
	fyne package -os windows
pkg-mac:
	fyne package -os darwin