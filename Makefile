BIN := tommy-grpc


.PHONY: build test lint proto clean-proto generate


build:
	go build -ldflags=$(BUILD_LDFLAGS) -o build/$(BIN) ./...

test: build
	go test -v ./...
	rm $(BIN)

lint:
	go vet ./...
	golint -set_exit_status ./...

.PHONY: clean
clean:
	rm -rf build
	go clean

protoc:
	protoc --go_out=plugins=grpc:. ./proto/*.proto

clean-proto:
	rm ./proto *.pb.go

generate:
