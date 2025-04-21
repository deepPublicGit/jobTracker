package main

import (
	"encoding/csv"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"log"
	"os"
	"strings"
	"time"
)

type Job struct {
	Title    string
	Location string
	Date     time.Time
	URL      string
}

func main() {

	fName := "ghtest.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Title", "Location", "Date", "URL"})

	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}),
		// Visit only domains: coursera.org, www.coursera.org
		colly.AllowedDomains("job-boards.greenhouse.io"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./gh_cache"),
	)

	jobs := make([]Job, 0, 200)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if !strings.Contains(link, "/jobs/") {
			return
		}
		//
		//if e.DOM.Find("td[class=cell]").Length() == 0 {
		//	return
		//}
		//title := strings.Split(e.ChildText(".job-title"), "\n")[0]
		//course_id := e.ChildAttr("p[class=body body--medium]", "value")
		////start_date, _ := time.Parse(DATE_FORMAT, e.ChildText("span.start-date"))
		////end_date, _ := time.Parse(DATE_FORMAT, e.ChildText("span.final-date"))
		//var run string
		//if len(strings.Split(course_id, "_")) > 1 {
		//	run = strings.Split(course_id, "_")[1]
		//}
		/*		e.ForEach("p", func(i int, ch *colly.HTMLElement) {
				log.Printf("Raw: %s, %s, %s", ch.Name, ch.DOM, ch.Index)
				log.Printf(ch.Text)
			})*/
		//log.Printf("Checking %s\n", )
		//log.Printf("Title %s\n", e.ChildText("p.body body--medium"))
		//log.Printf("Location %s\n", e.ChildText("p.body body__secondary body--metadata"))

		job := Job{
			Title:    e.DOM.Find("p").Eq(0).Text(),
			Location: e.DOM.Find("p").Eq(1).Text(),
			Date:     time.Now(),
			URL:      link,
		}
		countries := []string{"India", "Hyderabad", "Chennai", "Gurgaon", "Gurugram", "Delhi", "APAC", "Bengaluru", "Bangalore", "Remote"}

		for _, country := range countries {
			if strings.Contains(strings.ToLower(job.Location), strings.ToLower(country)) {
				jobs = append(jobs, job)
				writer.Write([]string{job.Title, job.Location, job.Date.Format("2006-01-02"), job.URL})
				break
			}
		}

		// start scaping the page under the link found
		//e.Request.Visit(link)
	})
	log.Printf("Collector starting...")

	c.Visit("https://job-boards.greenhouse.io/workatbackbase/")

	log.Printf("Scraping finished, check file %q for results\n", fName)
	log.Printf("Done. Jobs: %d\n, %s", len(jobs), jobs)
}
