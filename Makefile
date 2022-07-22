GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

run:
	@echo ''
	@echo '${YELLOW}Initializing...${RESET}'
	@echo '${WHITE}'
	@go run .
	@echo '${RESET}'
	@echo '${GREEN}Finished.'