lib.wasm: main.go
	GOARCH=wasm GOOS=js go build -o lib.wasm main.go

.PHONY: serve
serve:
	go run server/server.go

wasm_exec.js:
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" .

build: lib.wasm wasm_exec.js