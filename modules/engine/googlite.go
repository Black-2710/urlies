// modules/engine/googlite.go
package engine

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"urlies/core"
	"github.com/PuerkitoBio/goquery"
)

func Run(domain string, outputDir string) {
	fmt.Println("[Engine] Running search engine recon...")

	query := fmt.Sprintf("site:%s", strings.TrimPrefix(domain, "https://"))
	searchURL := fmt.Sprintf("https://www.google.com/search?q=%s", query)

	req, _ := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("User-Agent", core.UserAgent)
	client := core.NewHTTPClient()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[!] Engine request failed:", err)
		return
	}
	defer resp.Body.Close()

	// DEBUG: Dump raw HTML body
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("====== RAW HTML DUMP START ======")
	os.WriteFile("debug_google_engine.html", body, 0644)
	fmt.Println("====== RAW HTML DUMP END ======")

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	os.MkdirAll(outputDir, 0755)
	out := []string{}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		fmt.Println("→ Anchor href:", href) // DEBUG

		if strings.Contains(href, "/url?q=") {
			raw := strings.Split(strings.TrimPrefix(href, "/url?q="), "&")[0]
			parsed, err := url.QueryUnescape(raw)
			if err != nil {
				fmt.Println("[-] Failed to decode:", raw)
				return
			}

			fmt.Println("[+] Parsed URL:", parsed)
			out = append(out, parsed)
		}
	})

	fname := fmt.Sprintf("%s/%s-engine.txt", outputDir, core.ExtractDomain(domain))
	core.WriteLines(out, fname)
}
