vpath %.sql sql/query ./sql/schema
vpath query/%.sql sql/query
vpath %.sql sql/query ./sql/schema
SHELL=/bin/bash
GO := $(shell command -v go 2> /dev/null)
TERN := $(shell command -v tern 2> /dev/null)

.DEFAULT_GOAL := run

check-go:
ifndef GO
	$(error "go is not installed! Aborting")
endif

check-tern:
ifndef TERN
	$(shell go install github.com/jackc/tern)
endif

run:
	go run -ldflags "-X main.buildTime=$(shell date -u '+%Y-%m-%dT%T%z')" main.go

migrate: check-go check-tern
	tern migrate -m "sql/schema"

schema/%.sql:
	tern new "$(basename $*)" -m "sql/schema"

query/%.sql:
	touch sql/query/$*.sql
	$$EDITOR sql/query/$*.sql
