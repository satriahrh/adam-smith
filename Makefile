ifneq (,$(wildcard ./.env))
	include .env
	export
endif

setup:
	GO111MODULE=off go get -u github.com/pressly/goose/cmd/goose

goose:
	docker run --rm \
	-v "$(pwd)/host-directory-with-migrations:/migrations" \
	artsafin/goose-migrations -dir ./database/mysql/migrations \
	mysql "$(DATABASE_MYSQL_DBSTRING)" \
	$(filter-out goose,$(MAKECMDGOALS))
%:
	@:
