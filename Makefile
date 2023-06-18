define HELP_TEXT
Welcome!

Targets:
help				provides help text
test				run tests
legal				prepends legal header to source code

endef
export HELP_TEXT

help:
	@echo "$$HELP_TEXT"

doc:
	bash ./scripts/gendoc.sh

test:
	go test ./... -v -race -covermode=atomic -coverprofile=.coverprofile -coverpkg=./...

legal: .legal
.legal:
	git ls-files | xargs -I{} addlicense -f ./templates/legal/header.txt -skip yaml -skip yml {}
