export GEN_ROOT_DIRECTORY = './generated'

setup:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get github.com/facebook/ent/cmd/ent
generate:
	make generate-ent
	make generate-proto
generate-ent:
	go generate ent/generate.go
generate-proto:
	protoc --proto_path=proto --go_out=build/proto --go_opt=paths=source_relative proto/*.proto
test:
	go run github.com/ory/go-acc --ignore build ./... -- -v
%:
	@:
