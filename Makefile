export GEN_ROOT_DIRECTORY = './generated'

setup:
	go mod vendor
	make generate
generate-setup:
	rm -Rf $(GEN_ROOT_DIRECTORY)
	mkdir $(GEN_ROOT_DIRECTORY)
	touch $(GEN_ROOT_DIRECTORY)/generated.go
	echo package generated > $(GEN_ROOT_DIRECTORY)/generated.go
generate:
	make generate-setup
	ent generate --target $(GEN_ROOT_DIRECTORY)/ent ./ent/schema
%:
	@:
