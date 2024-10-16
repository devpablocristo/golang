# Monitors

A comprehensive collection of static code analysis tools for evaluating software development best practices across multiple programming languages and architectural patterns.

## Overview

This repository hosts a suite of specialized analyzers designed to assess various aspects of code quality, architectural patterns, and development practices in Go, Java, and Python projects. Each analyzer is specifically crafted to evaluate different aspects of software development, from architectural compliance to testing practices.

## Repository Structure

```
monitors/
├── commons-go           # Shared utilities and common code in Go (legacy)
├── go                   # Analyzers written in Go
│   ├── commons          # Common Go utilities
│   │   └── repo-tools
│   │       ├── ast
│   │       └── go-git
│   └── projects         # Specific Go projects
│       ├── go-dip
│       ├── go-error-handling
│       └── ...          # Other Go projects
├── java                 # Analyzers written in Java
│   ├── commons          # Common Java utilities
│   │   └── repo-tools
│   │       └── ...      # Java-specific repo tools
│   └── projects         # Specific Java projects
│       ├── java-solid
│       └── ...          # Other Java projects
├── python               # Analyzers written in Python
│   ├── commons          # Common Python utilities
│   │   └── repo-tools
│   │       └── ...      # Python-specific repo tools
│   └── projects         # Specific Python projects
│       ├── go-error-handling
│       ├── go-hexagonal-arch
│       └── ...          # Other Python projects
└── README.md            # Project documentation



```

## Available Analyzers

### Analyzers Written in Go
Located in the `/go` directory:
- **go-error-handling**: Evaluates for common error-handling patterns and best practices
- **go-hexagonal**: Evaluates the implementation of hexagonal architecture patterns in Go projects
- **go-testing**: Analyzes testing coverage and practices in Go codebases
- **meli**: Collection of analyzers specifically designed to evaluate Mercado Libre's internal tools and practices
  - Monitors compliance with company-specific patterns
  - Validates usage of internal frameworks and libraries
  - Ensures adherence to Mercado Libre's development standards

### Analyzers Written in Java
Located in the `/java` directory:
- **java-client**: Assesses client implementation patterns in Java
- [Additional Java analyzers...]

### Analyzers Written in Python
Located in the `/python` directory:
- [Python analyzers list...]

## Common Features

All analyzers share these core characteristics:
- Generate standardized JSON output
- Provide numerical scoring (typically 1-3) for each evaluated metric
- Include evidence and specific examples in results
- Support for analyzing specific files or entire repositories
- Git integration for tracking changes and attributions

## Getting Started

Each analyzer is contained in its own directory and can be used independently. To use a specific analyzer:

1. Navigate to the desired analyzer's directory
2. Follow the analyzer-specific README for installation and usage instructions

Example:
```bash
cd go/go-hexagonal
# Follow specific analyzer instructions
```

## Common Dependencies

### Go Projects
- Go 1.17 or higher
- Git

### Java Projects
- Java JDK 11 or higher
- Maven/Gradle

### Python Projects
- Python 3.x
- pip

## Usage Pattern

While each analyzer has its specific usage instructions, they generally follow this pattern:

```bash
# Go analyzers
go run main.go <repo_path> [optional_files]

# Java analyzers
mvn exec:java -Dexec.args="<repo_path> [optional_files]"

# Python analyzers
python analyzer.py <repo_path> [optional_files]
```

## Output Format

All analyzers produce JSON output following this general structure:
```json
[
  {
    "metricId": "metric_name",
    "score": "3",
    "gitAuthor": "author@example.com",
    "evidence": [
      {
        "commitId": "abc123",
        "file": "path/to/file",
        "line": 10
      }
    ]
  }
]
```

## Contributing

We welcome contributions to any of the analyzers! To contribute:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

Please ensure your analyzer follows our standard output format and includes appropriate documentation.

## Development Guidelines

When creating new analyzers:
- Follow the existing naming convention: `{target-language}-{analysis-type}`
- Include a comprehensive README in the analyzer's directory
- Implement the standard JSON output format
- Add appropriate tests
- Document all dependencies and requirements

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

Copyright (c) 2024 Teamcubation. All rights reserved.