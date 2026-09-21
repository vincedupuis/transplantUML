BINARY := tpuml

# PlantUML jar used by TestPlantUMLSyntax (internal/render/plantuml_test.go).
# The test skips when neither `plantuml` is on PATH nor PLANTUML_JAR is set.
PLANTUML_VERSION := 1.2025.4
PLANTUML_JAR ?= bin/plantuml-$(PLANTUML_VERSION).jar
PLANTUML_URL := https://github.com/plantuml/plantuml/releases/download/v$(PLANTUML_VERSION)/plantuml-$(PLANTUML_VERSION).jar

.PHONY: build run test fmt vet clean plantuml

build:
	go build -o bin/$(BINARY) ./cmd/$(BINARY)

run:
	go run ./cmd/$(BINARY) $(ARGS)

# Exports PLANTUML_JAR only when the jar exists, so `make test` works without it.
test:
	$(if $(wildcard $(PLANTUML_JAR)),PLANTUML_JAR=$(abspath $(PLANTUML_JAR))) go test ./...

# Download the PlantUML jar so TestPlantUMLSyntax runs. Needs java on PATH.
plantuml: $(PLANTUML_JAR)

$(PLANTUML_JAR):
	mkdir -p $(dir $@)
	curl -fsSL -o $@ $(PLANTUML_URL)

fmt:
	gofmt -l -w .

vet:
	go vet ./...

clean:
	rm -rf bin
