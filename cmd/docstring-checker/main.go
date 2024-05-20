// File created: 2024-05-18
// Last updated: 2024-05-20

package main

import (
    "log"
    "os"
    "github.com/wilhelmagren/docstring-checker-go/pkg/oswalk/oswalk"
)

func main() {
    if len(os.Args) != 1 {
        log.Println("[ERROR] you need to provide at least one argument, the path to search from!");
    }

    root := os.Args[0];

    files := oswalk.find_go_files(root);
    log.Println(files);
}
