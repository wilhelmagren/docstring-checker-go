# docstring-checker-go

Check that all docstrings in a project have been updated to match latest changes, written in Go. 

This program is structured according to the standard presented at [this](https://github.com/golang-standards/project-layout) link.

# Install go

if you dont have go installed on your system (check this by running `go version`) you can
use the script `/tools/go_install.sh` to install go on your system.

# Run program

Run everything with the command:
```
$ go run cmd/docstring-checker/main.go <path-to-search-from>
```

Or you can build it to binary and then run it using that:
```
$ go build cmd/docstring-checker/main.go
$ ./main <path-to-search-from>
```

