// File created: 2024-05-18
// Last updated: 2024-05-20

package main

import (
    "log"
    "fmt"
    "os"
    "github.com/wilhelmagren/docstring-checker-go/pkg/checker"
    "github.com/wilhelmagren/docstring-checker-go/pkg/oswalk"
)

func main() {
    if len(os.Args) != 2 {
        log.Fatal("You need to provide exactly one argument which is the path to search from!");
    }

    root := os.Args[1];

    files := oswalk.FindGoFiles(root);
    invalidFiles := checker.CheckDocstrings(files);

    numInvalidFiles := len(invalidFiles);

    if numInvalidFiles > 0 {
        fmt.Printf("\n 💥 You had %d file(s) with outdated docstring(s)!\n", numInvalidFiles);
        for _, file := range invalidFiles {
            fmt.Printf("    - %s\n", file);
        }
    } else {
        fmt.Println("\n 🎉 You had no files with outdated docstrings, good job!");
    }
}
