# golang

How to run gofmt recursively
gofmt -w -s .
This will run gofmt recursively and update the files directly

-s simplifies the code
-w writes results directly
How to run gofmt recursively as dry-run
gofmt -s -d .
This will write the changes to standard output instead of writing directly to the files

-s simplifies the code
-d display diffs instead of rewriting files
