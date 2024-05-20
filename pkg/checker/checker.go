// File created: 2024-05-20
// Last updated: 2024-05-20

package checker

import (
    "log"
    "os"
    "strings"
)

func CheckDocstrings(files []string) []string {
    var invalidFiles []string;
    for _, file := range files {
        contents, err := os.ReadFile(file);
        if err != nil {
            log.Fatal(err);
        }
        body := string(contents);

        for _, line := range strings.Split(body, "\n") {
            // Ugly hack to find only the docstring, we know its longer than 20 runes.
            if len(line) > 20 {
                if strings.Contains(line[:20], "// Last updated: ") {
                    docstringDate := strings.TrimSpace(line[16:])
                    fileInfo, err := os.Stat(file);

                    if err != nil {
                        log.Fatal(err);
                    }

                    modificationTime := fileInfo.ModTime().Format("2006-01-02");
                    if modificationTime != docstringDate {
                        invalidFiles = append(invalidFiles, file);
                    }
                    break;
                }
            }
        }
    }

    return invalidFiles;
}

