export GEN_ROOT_DIRECTORY = './generated'

setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go mod vendor
generate:
	go generate ent/generate.go
generate-proto:
	protoc --proto_path=proto --go_out=build/proto --go_opt=paths=source_relative proto/brand.proto
%:
	@:
