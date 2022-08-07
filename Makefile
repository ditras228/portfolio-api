.PHONY: dev
dev:
	$(call _info, "\nНакатывание миграций");
	migrate -database postgresql://postgres:pM4*fJoS@localhost:5432/postgres?sslmode=disable -path ./schema/ up
	$(call _info, "\nСтарт GraphQL сервера");
	go run server.go
.PHONY: drop
drop:
	$(call _warning, "\nДроп миграций");
	migrate -database postgresql://postgres:pM4*fJoS@localhost:5432/postgres?sslmode=disable -path ./schema/ drop
	go run migrate.go
.PHONY: gen
gen:
	$(call _info, "\nГенерация GraphQL типов");
	go run github.com/99designs/gqlgen generate.

.PHONY: img
img:
	$(call _info, "\nСтарт сервера раздачи картинок");
	go run server-images.go

define _info
	$(call _echoColor, $1, 6)
endef

define _warning
	$(call _echoColor, $1, 1)
endef

define _echoColor
	@tput setaf $2
	@echo $1
	@tput sgr0
endef

