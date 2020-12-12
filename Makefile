ifneq (,$(wildcard ./.env))
	include .env
	export
endif

goose:
	goose -dir ./database/mysql/migrations mysql "$(DATABASE_MYSQL_DBSTRING)" $(filter-out goose,$(MAKECMDGOALS))
%:
	@:
