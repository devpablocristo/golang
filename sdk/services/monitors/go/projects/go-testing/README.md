# Go Test Analyzer

This Go script analyzes Go test files in a Git repository to identify various testing skills and practices used by different authors.

## Features

The analyzer checks for the following skills:

1. Writing unit tests using Go's testing package
2. Using Testify for enhanced assertions
3. Implementing table-driven tests
4. Using mock objects for testing

## Prerequisites

- Go (version 1.17+)
- Git

## Installation

1. Clone the repository:
   ```
   git clone git@github.com:teamcubation/monitors.git
   ```
2. Navigate to the project directory:
   ```
   cd go/go-testing
   ```
3. Install dependencies:
   ```
   go mod tidy
   ```

## Usage

To run the tool, use the following command:

```bash
go run main.go <repo_path> [file1] [file2] ...
```

- `<repo_path>`: The path to the Git repository you want to analyze.
- `[file1] [file2] ...`: (Optional) Specific files to analyze. If not provided, all Go test files in the repository will be analyzed.

## Output

The tool outputs a JSON-formatted report detailing the metrics for each analyzed file, including:

- Skills measured.
- Evidence (lines where patterns were detected).
- **Score:** Indicates if the skill was found in the file. A score of `0` means no cases were found, and a score of `1` means at least one case was detected.

### Example of output:

```json
[
  {
    "git_author": "JohnDoe",
    "metric_id": "go_unit_tests",
    "score": 1,
    "evidence": [
      {
        "commit_id": "123456abcdef6789",
        "file": "cmd/api/handler/user_handler_test.go",
        "line": 15 
        },
      { 
        "commit_id": "123456abcdef6789",
        "file": "cmd/api/handler/file_handler_test.go",
        "line": 80
        }
    ]
  },
  {
    "git_author": "JohnDoe",
    "metric_id": "testify_assertions",
    "score": 1,
    "evidence": [
      {
        "commit_id": "123456abcdef6789",
        "file": "cmd/api/handler/user_handler_test.go",
        "line": 18 
        }
    ]
  },
  {
    "git_author": "JaneSmith",
    "metric_id": "table_driven_tests",
    "score": 0,
    "evidence": []
  }
]
```

## How it works

1. The script opens the specified Git repository.
2. It identifies Go test files (ending with `_test.go`) to analyze.
3. For each file, it scans line by line, checking for patterns that indicate the use of specific testing skills.
4. It collects data on skill usage, associating it with the file's author.
5. Finally, it generates a JSON report of the findings.

## Customization

You can easily add new skills to detect by modifying the `skills` slice in the script. Each skill is defined with an ID, name, and either a regex pattern or a set of check functions.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

Copyright (c) 2024 Teamcubation. All rights reserved.