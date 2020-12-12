ifneq (,$(wildcard ./.env))
	include .env
	export
endif

setup:
	GO111MODULE=off go get -u github.com/pressly/goose/cmd/goose

goose:
	goose -dir ./database/mysql/migrations mysql "$(DATABASE_MYSQL_DBSTRING)" $(filter-out goose,$(MAKECMDGOALS))
%:
	@:
