TOOLS_SOURCE := $(wildcard ./tools/*.go)
tools:
	@for file in $(TOOLS_SOURCE); do \
    	filename=$$(basename $$file .go); \
    	echo "generating tools $$filename"; \
    	go build -o ./tools/$$filename $$file; \
	done
serve:
	./tools/build && ./tools/serve

.PHONY: tools
