setup:
	go get github.com/facebook/ent/cmd/ent
ent-generate:
	ent generate ./ent/schema
%:
	@:
