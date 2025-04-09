// modules/headless/jsfinder.go
package headless

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"urlies/core"
	"github.com/PuerkitoBio/goquery"
)

func Run(domain string, outputDir string) {
	fmt.Println("[Headless] Scanning for JS URLs...")

	client := core.NewHTTPClient()
	req, _ := http.NewRequest("GET", domain, nil)
	req.Header.Set("User-Agent", core.UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[!] Failed to fetch:", err)
		return
	}
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	os.MkdirAll(outputDir, 0755)
	found := map[string]bool{}
	out := []string{}

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists {
			if !strings.HasPrefix(src, "http") {
				full := domain + src
				fmt.Println("→ Relative JS:", src)
				if !found[full] {
					out = append(out, full)
					found[full] = true
				}
			} else {
				fmt.Println("[+] Full JS Script:", src)
				if !found[src] {
					out = append(out, src)
					found[src] = true
				}
			}
		}
	})

	jsRegex := regexp.MustCompile(`https?://[^"'>]+\.js`)
	bodyText := doc.Text()

	// DEBUG: Dump raw body text from HTML
	fmt.Println("====== HEADLESS RAW TEXT ======")
	os.WriteFile("debug_headless_text.txt", []byte(bodyText), 0644)

	matches := jsRegex.FindAllString(bodyText, -1)
	for _, js := range matches {
		if !found[js] {
			found[js] = true
			out = append(out, js)
			fmt.Println("[+] JS URL:", js)
		}
	}

	fname := fmt.Sprintf("%s/%s-headless.txt", outputDir, core.ExtractDomain(domain))
	core.WriteLines(out, fname)
}
