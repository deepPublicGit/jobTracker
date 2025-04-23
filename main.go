package main

import (
	"awesomeProject1/scrapers"
	"github.com/go-co-op/gocron/v2"
	"log"
	"time"
)

func main() {
	s, _ := gocron.NewScheduler()
	defer func() { _ = s.Shutdown() }()

	_, _ = s.NewJob(
		gocron.CronJob(
			// standard cron tab parsing
			"* * * * *",
			false,
		),
		gocron.NewTask(func() {
			log.Println("Starting Scraping Scheduler")
			scrapers.ScrapeGreenHouse()
			jobs, _ := scrapers.ScrapePlainHTML("https://www.workatastartup.com/companies?demographic=any&hasEquity=any&hasSalary=any&industry=any&interviewProcess=any&jobType=any&layout=list-compact&sortBy=created_desc&tab=any&usVisaNotRequired=any")
			log.Printf("Headout Jobs: %s\n", jobs)
		}),
	)

	s.Start()

	select {
	case <-time.After(time.Minute * 2):
	}
}
