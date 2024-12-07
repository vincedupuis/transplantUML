
# transplantUML

**transplantUML** is a versatile tool built in Go, designed to convert SCXML (State Chart XML) files into any desired format using the Go templating engine. It empowers developers to easily transform state machine definitions into custom outputs such as code, documentation, or visualizations.

## Features

- **SCXML Parsing**: Reads and processes SCXML file.
- **Go Templating Engine**: Leverages the powerful Go templating system for flexible output generation.
- **Customizable Outputs**: Generate any desired format, such as source code, diagrams, or documentation, by providing the appropriate templates.
- **Lightweight and Fast**: Built with Go, ensuring speed and simplicity in deployment.

## Requirements

- Go (version 1.18 or later is recommended).
- A valid SCXML file as input.
- Custom templates written in Go's templating syntax.
    - See the [official documentation](https://pkg.go.dev/text/template) for more information.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/vincedupuis/transplantUML.git
   cd transplantUML
   ```

2. Build the project:
   ```bash
   go build cmd/tpuml.go
   ```

3. Run the tool:
   ```bash
   ./tpuml -i statechart.scxml -t template.tmpl -o output.ext
   ```

- `-i`: Path to the SCXML file to be converted.
- `-t`: Path to the Go template file defining the desired output format.
- `-o`: Path where the generated file will be saved.

## Example

To convert an SCXML file into a plantUML diagram:

```bash
./tpuml -i example/coffee-machine.scxml -t assets/plantuml.tmpl -o coffee.puml
```

The resulting `coffee.puml` file will contain a structured plantUML that can be visualized with [PlantUML Online](https://plantuml.online).

