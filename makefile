.PHONY: web-tinygo web-go

web-tinygo:
	GOOS=js GOARCH=wasm tinygo build -o ./web/repl-lib.wasm ./cmd/repl_wasm_lib.go
	cp `tinygo env TINYGOROOT`/targets/wasm_exec.js ./web/

web-go: 
	GOOS=js GOARCH=wasm go build -o ./web/repl-lib.wasm ./cmd/repl_wasm_lib.go
	cp `go env GOROOT`/lib/wasm/wasm_exec.js ./web/

