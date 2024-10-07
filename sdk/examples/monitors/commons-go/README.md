# commons-go

Shared utilities module for the Monitors project, providing common functionality for metrics generation and Git operations. This module serves as a foundation for all Go-based analyzers in the project.

## Package Structure

```
commons-go/
└── pkg/
    ├── common/     # Common utilities for metrics and analysis
    └── gitutils/   # Git-related operations and utilities
```

## Packages

### common

The `common` package provides core structures and utilities for metrics generation and analysis.

```go
import "github.com/teamcubation/monitors/commons-go/pkg/common"
```

#### Key Features

- Standard metric structures
- Evidence tracking
- Score calculation utilities

#### Main Types

```go
type Metric struct {
	MetricID  string     `json:"metric_id"`
	GitAuthor string     `json:"git_author"`
	Score     string     `json:"score"`
	Evidence  []Evidence `json:"evidence"`
}

type Evidence struct {
	CommitID string `json:"commit_id"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}
```

#### Usage Example

```go
metric := common.Metric{
    MetricID: "architecture_compliance",
    Score: "3",
    Evidence: []common.Evidence{
        {
            CommitID: "abc123",
            File: "internal/core/domain/entity.go",
            Line: 15,
        },
    },
}
```

### gitutils

The `gitutils` package provides Git-related utilities for repository analysis and metadata extraction.

```go
import "github.com/teamcubation/monitors/commons-go/pkg/gitutils"
```

#### Key Features

- Repository access and navigation
- Author information retrieval
- Commit history analysis
- File filtering utilities

#### Main Functions

```go
// Opens and returns a Git repository instance
GetRepo(path string) (*git.Repository, error)

// Get author and commit information for a specific file
GetAuthorAndCommit(repo *git.Repository, filePath string) (string, string, error)

// Filter files based on specific criteria
GetFilesToAnalyze(repo *git.Repository, modifiedFiles []string, filter func(string) bool) ([]string, error)
```

#### Usage Example

```go
repo, err := gitutils.GetRepo("path/to/repo")
if err != nil {
    log.Fatal(err)
}

author, commitID, err := gitutils.GetAuthorAndCommit(repo, "path/to/file")
if err != nil {
    log.Fatal(err)
}
```

## Installation

To use this module in your analyzer, add it to your `go.mod`:

```bash
go get github.com/teamcubation/monitors/commons-go
```

## Requirements

- Go 1.17 or higher

## Contributing

When contributing to this module:

1. Ensure all new utilities are properly documented
2. Add appropriate tests for new functionality
3. Maintain backward compatibility when possible
4. Update this README with any new major features

### Testing

Run the tests with:

```bash
go test ./...
```

## Usage in Monitors

This module is used by various analyzers in the Monitors project. When creating a new analyzer:

1. Import the required packages:
```go
import (
    "github.com/teamcubation/monitors/commons-go/pkg/common"
    "github.com/teamcubation/monitors/commons-go/pkg/gitutils"
)
```

2. Use the standard metric structures for consistency:
```go
metrics := []common.Metric{
    {
        MetricID: "your_metric",
        Score: "3",
        // ...
    },
}
```

3. Utilize Git utilities for repository analysis:
```go
repo, _ := gitutils.GetRepo(repoPath)
files, _ := gitutils.GetFilesToAnalyze(repo, modifiedFiles, filterFunc)
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

Copyright (c) 2024 Teamcubation. All rights reserved.