# URLies

**URLies** is a powerful, minimalistic passive recon tool that combines the strengths of tools like `waybackurls`, `gauplus`, and `katana` — all under one CLI interface. Built with concurrency, speed, and modular recon in mind.

> Made by nobody...

---

##  Features

-  **Wayback Machine Recon (paged & concurrent)**
-  **Search Engine scraping (Google-style)**
-  **JS-based link extraction (headless recon)**
-  Fully customizable: `--depth`, output dir, and more

---

##  Installation

```bash
git clone https://github.com/black/urlies.git
cd urlies
go mod tidy
go build -o urlies main.go
```

---

##  Usage

```bash
./urlies -u https://example.com -m archive -o output/
./urlies -u https://target.com -m all --depth 10 -o output/
```

- `-u`: Target domain (required)
- `-m`: Mode — `archive`, `engine`, `headless`, or `all`
- `-o`: Output directory
- `--depth`: Page depth for archive crawling (default: 20)

---

##  Output Structure

Results are saved in:

```
output/
├── example.com-wayback.txt
├── example.com-engine.txt
├── example.com-headless.txt
```

---

## 🛠 Modules Overview

- `archive/wayback.go` – Wayback Machine, paged API
- `engine/googlite.go` – Google scraping with regex
- `headless/jsfinder.go` – JS + DOM link detection
- `core/*` – Shared tools (client, helpers, output)

---

##  Dependencies (A Friendly Guide)

>  **If you're new to Go** or never installed a CLI tool before, don’t worry — follow this and you’ll be running `urlies` in no time.

###  1. Make sure you have Go installed
```bash
go version
```
You should see something like `go version go1.20.6 linux/amd64`
→ [Install Go](https://golang.org/dl/) if needed

###  2. Clone the Repo
```bash
git clone https://github.com/black/urlies.git
cd urlies
```

###  3. Download the Dependencies
```bash
go mod tidy
```
This fetches the required packages:
- [goquery](https://github.com/PuerkitoBio/goquery) – DOM parsing
- [urfave/cli/v2](https://github.com/urfave/cli) – CLI flag parsing
- [cascadia](https://github.com/andybalholm/cascadia) – CSS selectors
- [golang.org/x/net](https://pkg.go.dev/golang.org/x/net) – HTTP extensions

###  4. Build the tool
```bash
go build -o urlies main.go
```

###  5. Run It!
```bash
./urlies -u https://example.com -m all --depth 15 -o output/
```

---

##  License

MIT License. See `LICENSE` for details.

---

## Updates and Patches...

Well Dont Worry I'll take care...

---
