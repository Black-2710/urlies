// core/output.go
package core

import (
    "encoding/json"
    "fmt"
    "os"
)

// WriteLines writes lines to a .txt file
func WriteLines(lines []string, path string) error {
    f, err := os.Create(path)
    if err != nil {
        return err
    }
    defer f.Close()

    for _, line := range lines {
        fmt.Fprintln(f, line)
    }
    return nil
}

// WriteJSON writes data as JSON
func WriteJSON(data interface{}, path string) error {
    f, err := os.Create(path)
    if err != nil {
        return err
    }
    defer f.Close()

    enc := json.NewEncoder(f)
    enc.SetIndent("", "  ")
    return enc.Encode(data)
}
