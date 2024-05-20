// File created: 2024-05-20
// Last updated: 2024-05-20

package oswalk

import (
    "io/fs"
    "log"
    "path/filepath"
)

func FindGoFiles(root string) []string {
    var files []string;
    err := filepath.WalkDir(root, func(p string, d fs.DirEntry, e error) error {
        if e != nil { return e; }
        if filepath.Ext(d.Name()) == ".go" {
            files = append(files, p);
        };
        return nil;
    });

    if err != nil {
        log.Fatal(err);
    }

    return files;
}
