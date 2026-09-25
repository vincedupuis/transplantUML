BINARY := fsm

# PlantUML jar used by TestPlantUMLSyntax (internal/render/plantuml_test.go).
# The test skips when neither `plantuml` is on PATH nor PLANTUML_JAR is set.
PLANTUML_VERSION := 1.2025.4
PLANTUML_JAR ?= bin/plantuml-$(PLANTUML_VERSION).jar
PLANTUML_URL := https://github.com/plantuml/plantuml/releases/download/v$(PLANTUML_VERSION)/plantuml-$(PLANTUML_VERSION).jar

# Boost.SML header used by TestSMLCompiles (internal/render/sml_test.go).
# The test skips when SML_INCLUDE is not set or no C++ compiler is found.
SML_VERSION := 1.2.0
SML_INCLUDE ?= bin/sml-$(SML_VERSION)
SML_URL := https://raw.githubusercontent.com/boost-ext/sml/v$(SML_VERSION)/include/boost/sml.hpp

# ANTLR tool used to regenerate internal/fsm/parser from internal/fsm/fsm.g4
# (`make generate`). Its version must match the Go runtime in go.mod.
ANTLR_VERSION := 4.13.2
ANTLR_JAR ?= bin/antlr-$(ANTLR_VERSION)-complete.jar
ANTLR_URL := https://www.antlr.org/download/antlr-$(ANTLR_VERSION)-complete.jar

.PHONY: build run test fmt vet clean plantuml sml antlr generate

build:
	go build -o bin/$(BINARY) ./cmd/$(BINARY)

run:
	go run ./cmd/$(BINARY) $(ARGS)

# Exports PLANTUML_JAR and SML_INCLUDE only when they exist, so `make test` works without them.
test:
	$(if $(wildcard $(PLANTUML_JAR)),PLANTUML_JAR=$(abspath $(PLANTUML_JAR))) \
	$(if $(wildcard $(SML_INCLUDE)/boost/sml.hpp),SML_INCLUDE=$(abspath $(SML_INCLUDE))) \
	go test ./...

# Download the PlantUML jar so TestPlantUMLSyntax runs. Needs java on PATH.
plantuml: $(PLANTUML_JAR)

$(PLANTUML_JAR):
	mkdir -p $(dir $@)
	curl -fsSL -o $@ $(PLANTUML_URL)

# Download the Boost.SML header so TestSMLCompiles runs. Needs a C++20 compiler.
sml: $(SML_INCLUDE)/boost/sml.hpp

$(SML_INCLUDE)/boost/sml.hpp:
	mkdir -p $(dir $@)
	curl -fsSL -o $@ $(SML_URL)

# Download the ANTLR tool and regenerate the DSL parser. Needs java on PATH.
antlr: $(ANTLR_JAR)

$(ANTLR_JAR):
	mkdir -p $(dir $@)
	curl -fsSL -o $@ $(ANTLR_URL)

generate: antlr
	ANTLR_JAR=$(abspath $(ANTLR_JAR)) go generate ./...

fmt:
	gofmt -l -w .

# The generated ANTLR parser trips vet's unreachable-code check; it is not ours to fix.
vet:
	go vet -unreachable=false ./...

clean:
	rm -rf bin
