# Go Error Handling Analyzer

## Overview

This Go project analyzes Go source code in a given repository or a specified list of files and measures a set of error-handling related skills. The tool checks for common error-handling patterns and best practices, such as proper error checking, the usage of `panic` and contextual error reporting.

## Features

- **Error Checking:** Verifies that errors are properly checked in the code.
- **Panic Usage:** Identifies inappropriate or excessive use of `panic()`, `log.Fatal()`, and `os.Exit()`.
- **Error Context:** Detects the use of `fmt.Errorf` or `errors.Wrap` to provide additional context in error messages.
- **Ignored Errors:** Monitors for ignored errors.
- **Error Logs:** Tracks and measures the frequency of error logs to ensure they aren't excessive.

## Prerequisites

- Go (version 1.17+)
- Git

## Skills Analyzed

The project measures the following skills related to error handling:

- `error_checking`: Ensures errors are correctly handled with `if err != nil`.
- `panic_usage`: Identifies unnecessary usage of `panic()`, `log.Fatal()`, and `os.Exit()`.
- `error_context`: Checks the addition of context to errors using `fmt.Errorf` or `errors.Wrap()`.
- `ignored_errors`: Tracks occurrences of ignored errors using `_`.
- `defer_usage`: Encourages the use of `defer` for cleaning up resources in case of errors.
- `log_error_count`: Monitors and minimizes frequent error logs during runtime.

## Installation

1. Clone the repository:
   ```
   git clone git@github.com:teamcubation/monitors.git
   ```
2. Navigate to the project directory:
   ```
   cd go/go-error-handling
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
    "metric_id": "error_checking",
    "score": 1,
    "evidence": [
      {
        "commit_id": "123456abcdef6789",
        "file": "cmd/api/handler/user_handler.go",
        "line": 15 
        },
      { 
        "commit_id": "123456abcdef6789",
        "file": "cmd/api/handler/file_handler.go",
        "line": 80
        }
    ]
  }
]
```
## How it works

1. The script opens the specified Git repository.
2. It identifies Go files to analyze.
3. For each file, it scans line by line, checking for patterns that indicate the use of specific testing skills.
4. It collects data on skill usage, associating it with the file's author.
5. Finally, it generates a JSON report of the findings.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

Copyright (c) 2024 Teamcubation. All rights reserved.