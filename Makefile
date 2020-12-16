export GEN_ROOT_DIRECTORY = './generated'

setup:
	go mod vendor
generate:
	go generate ent/generate.go
%:
	@:
