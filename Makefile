setup:
	go mod vendor
	go get github.com/facebook/ent/cmd/ent
ent-generate:
	ent generate --target ent/generated ./ent/schema
%:
	@:
