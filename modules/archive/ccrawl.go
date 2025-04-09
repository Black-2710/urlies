// modules/archive/ccrawl.go
package archive

import (
    "fmt"
    //"net/http"
    "os"
    "strings"
    "urlies/core"
)

func FetchFromCommonCrawl(domain string, outputDir string) {
    fmt.Println("[Archive] Fetching from Common Crawl (placeholder)...")

    fake := fmt.Sprintf("https://data.commoncrawl.org/archive-%s/fake-sample.txt", strings.TrimPrefix(domain, "https://"))
    os.MkdirAll(outputDir, 0755)
    core.WriteLines([]string{fake}, fmt.Sprintf("%s/commoncrawl.txt", outputDir))
}
