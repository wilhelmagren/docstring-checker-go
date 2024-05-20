//
//

package oswalk

import (
    "io/fs"
    "log"
    "path/filepath"
)

func find_go_files(root string) (files []string) {
    var files []string;
    filepath.WalkDir(root, func(p string, d fs.DirEntry, e error) error {
        if e != nil { return e; }
        if filepath.Ext(d.Name() == "go" {
            a = append(a, p);
        };
        return nil;
    }));

    if err != nil {
        log.Println("[ERROR] ", err);
    }

    return files;
}
