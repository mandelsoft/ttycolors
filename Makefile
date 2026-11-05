build:
	go build ./...

test:
	go test ./...

# All examples (folders below examples containing a main.go).
EXAMPLES := $(patsubst examples/%/main.go,%,$(shell find examples -name main.go -not -path 'examples/internal/*'))
# Examples without demo gif: the terminal emulation of terminalizer
# supports neither underline styles nor underline colors.
GIF_SKIP := ttycolors/underline
GIFS     := $(patsubst %,examples/%/demo.gif,$(filter-out $(GIF_SKIP),$(EXAMPLES)))
SOURCES  := $(shell find . -name '*.go' -not -name '*_test.go' -not -path './examples/*' -not -path './local/*')

# Generates the demo.gif files for the examples (requires bash,
# jq, script and terminalizer, see hack/demo-gif.sh).
# Only outdated gifs are generated, use make -B gifs to generate all.
.PHONY: gifs
gifs: $(GIFS)

examples/%/demo.gif: examples/%/main.go $(SOURCES) hack/demo-gif.sh hack/terminalizer.json
	hack/demo-gif.sh $*
