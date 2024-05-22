/usr/sbin/go test -timeout 30s -run TestMockAll github.com/117503445/markdown-translate/test -v > 1.out

/usr/sbin/go test -run TestOpenAI github.com/117503445/markdown-translate/test -v

GOOS=js GOARCH=wasm go build -o markdown-translate.wasm ./cmd/wasm

go install github.com/spf13/cobra-cli@latest
cobra-cli init

/usr/sbin/go test -run TestUniAll github.com/117503445/markdown-translate/test -v

go run . 

docker build -t markdown-translate-builder -f Dockerfile.builder .
docker run --rm -v $(pwd):/workspace markdown-translate-builder