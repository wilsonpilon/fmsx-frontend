set shell := ["bash", "-cu"]

app := "fmsx-frontend"

fmt:
	go fmt ./...

test:
	go test ./...

build:
	mkdir -p bin
	go build -o bin/{{app}} ./cmd/fmsx-frontend

build-tui:
	mkdir -p bin
	go build -tags tview -o bin/{{app}}-tui ./cmd/fmsx-frontend

build-gui:
	mkdir -p bin
	go build -tags fyne -o bin/{{app}}-gui ./cmd/fmsx-frontend

run:
	go run ./cmd/fmsx-frontend --mode tui

run-tui:
	go run -tags tview ./cmd/fmsx-frontend --mode tui

run-gui:
	go run -tags fyne ./cmd/fmsx-frontend --mode gui
