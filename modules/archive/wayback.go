// modules/archive/wayback.go
package archive

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"

	"urlies/core"
)

func Run(domain string, outputDir string, maxPages int) {
	fmt.Println("[Archive] Fetching Wayback history (concurrent)...")

	domainOnly := core.ExtractDomain(domain)
	baseURL := fmt.Sprintf("http://web.archive.org/cdx/search/cdx?url=*.%s/*&output=json&fl=original&collapse=urlkey&limit=1000&page=", domainOnly)

	var allURLs []string
	var wg sync.WaitGroup
	var mu sync.Mutex

	client := core.NewHTTPClient()

	for page := 0; page < maxPages; page++ {
		wg.Add(1)
		go func(pg int) {
			defer wg.Done()
			pageURL := baseURL + strconv.Itoa(pg)
			req, _ := http.NewRequest("GET", pageURL, nil)
			req.Header.Set("User-Agent", core.UserAgent)

			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("[!] Page %d failed: %v\n", pg, err)
				return
			}
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()

			var results [][]string
			if err := json.Unmarshal(body, &results); err != nil {
				fmt.Printf("[!] JSON parse failed on page %d: %v\n", pg, err)
				return
			}

			if len(results) <= 1 {
				fmt.Printf("[~] Page %d empty.\n", pg)
				return
			}

			local := []string{}
			for i := 1; i < len(results); i++ {
				local = append(local, results[i][0])
			}

			mu.Lock()
			allURLs = append(allURLs, local...)
			mu.Unlock()

			fmt.Printf("[+] Fetched %d URLs from page %d\n", len(local), pg)
		}(page)
	}

	wg.Wait()

	if len(allURLs) == 0 {
		fmt.Println("[!] No archived URLs found after concurrent fetch.")
		return
	}

	os.MkdirAll(outputDir, 0755)
	fname := fmt.Sprintf("%s/%s-wayback.txt", outputDir, domainOnly)
	if err := core.WriteLines(allURLs, fname); err != nil {
		fmt.Println("[!] Failed to write file:", err)
		return
	}

	for i := 0; i < len(allURLs) && i < 5; i++ {
		fmt.Println("[+] Wayback URL:", allURLs[i])
	}

	fmt.Printf("[+] Collected %d total URLs from Wayback (concurrent mode)\n", len(allURLs))
}
