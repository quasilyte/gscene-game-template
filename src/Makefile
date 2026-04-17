.PHONY: run-game
run-game:
	go run --tags gmath ./cmd/game --data ../game_data

wasm:
	GOARCH=wasm GOOS=js go build --tags gmath -ldflags="-s -w" -trimpath -o ../_web/main.wasm ./cmd/game

itchio-wasm: wasm
	cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" ../_web/wasm_exec.js
	cd ../_web && \
		mkdir -p ../bin && \
		rm -f ../bin/game.zip && \
		zip ../bin/game.zip -r main.wasm index.html wasm_exec.js game_data
