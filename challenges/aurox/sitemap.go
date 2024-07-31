/*
	Implement simple sitemap (https://www.sitemaps.org) generator as command line tool.
	Please implement this test task in the same way as you would do it for production
	code, which means pay attention to edge cases and details.

	It should:
	1. Accept start url as argument.
	2. Recursively navigate by site pages in parallel.
	3. Should not use any external dependencies, only standard golang library.
	4. Extract page urls only from <a> elements and take in account <base> element if declared.
	5. Should be well tested (automated testing).

	Suggested program options:
	-parallel= number of parallel workers to navigate through site.
	-output-file= output file path.
	-max-depth=  max depth of url navigation recursion.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	nu "github.com/devpablocristo/sitemap-generator/navigate-url"
	xml "github.com/devpablocristo/sitemap-generator/xml-file"
)

// type URL struct {
// 	url   string
// 	depth int
// }

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nSitemap Generator commands:")
	fmt.Println("-u starting url.")
	fmt.Println("-p number of parallel workers to navigate through site.")
	fmt.Println("-of output file path.")
	fmt.Println("-md  max depth of url navigation recursion.")
	fmt.Println("example: sitemapgen -u http://localhost:8081/index.html -p 4 -of sitemap.xml -md 0")
	fmt.Println("")
	fmt.Print("sitemapgen ")
	cmd, _ := reader.ReadString('\n')
	cmd = strings.Replace(cmd, "\n", "", 1)

	cmds := strings.Split(cmd, " ")
	//fmt.Println(cmds)

	var startURL string
	var parallel int
	var outputFile string
	var maxDepth int
	for i := 0; i < len(cmds); i++ {
		switch cmds[i] {
		case "-u":
			startURL = cmds[i+1]
		case "-p":
			parallel, _ = strconv.Atoi(cmds[i+1])
		case "-of":
			outputFile = cmds[i+1]
		case "-md":
			maxDepth, _ = strconv.Atoi(cmds[i+1])
		}
	}

	maxDepth++
	nu.ListURLs = make(map[string]bool)
	nu.ListURLs[startURL] = true

	var wg sync.WaitGroup

	fmt.Println("parallel workers:", parallel)

	if parallel > 0 {
		for i := 0; i < parallel; i++ {
			wg.Add(1)
			go nu.Worker(startURL, maxDepth, &wg)
		}
	} else {
		nu.NavigateUrl(startURL, maxDepth)
	}

	wg.Wait()

	err := xml.CreateXML(outputFile)
	if err != nil {
		log.Fatal(err)
	}

	for u := range nu.ListURLs {
		fmt.Println(u)
	}

}
