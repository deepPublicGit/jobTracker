package scrapers

import (
	"awesomeProject1/models"
	"github.com/gocolly/colly"
	"log"
	"strings"
	"time"
)

func ScrapePlainHTML(url string) ([]models.Job, error) {
	var jobs []models.Job
	c := colly.NewCollector(
		//colly.AllowedDomains("headout.com"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./generic_cache"),
	)

	c.OnHTML("a", func(e *colly.HTMLElement) {

		title := e.Text
		link := e.Attr("href")
		roles := []string{"software", "engineer", "developer", "sde", "swe", "java", "backend"}
		log.Printf("Scraping link: %s, %q\n, ", title, link)
		//log.Printf("Raw: %s", e.Text)

		for _, role := range roles {
			if strings.Contains(strings.ToLower(title), role) {
				jobs = append(jobs, models.Job{
					Title:    strings.TrimSpace(title),
					URL:      e.Request.AbsoluteURL(link),
					Location: "",
					Date:     time.Now(),
				})
				break
			}
		}

	})

	err := c.Visit(url)
	return jobs, err
}
