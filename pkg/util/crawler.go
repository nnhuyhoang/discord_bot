package util

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/nnhuyhoang/discord_bot/pkg/consts"
	"github.com/nnhuyhoang/discord_bot/pkg/model"
)

func GetCoronaIndex() []model.Country {
	c := colly.NewCollector()
	var nations []model.Country

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML(`table[id=main_table_countries_yesterday]`, func(e *colly.HTMLElement) {
		index := 0
		e.ForEachWithBreak("tbody tr", func(_ int, h *colly.HTMLElement) bool {
			if index >= 10 {
				return false
			}
			if h.Attr("class") == "" {
				index = index + 1
				nation := model.Country{}

				h.ForEach(`td`, func(ei int, el *colly.HTMLElement) {
					switch el.Index {
					case 1:

						nation.Name = strings.TrimSpace(el.ChildText(".mt_a"))
					case 2:
						totalCase := strings.TrimSpace(el.Text)
						if totalCase == "N/A" || totalCase == "" {
							nation.TotalCase = "N/A"
						}
						nation.TotalCase = totalCase
					case 3:
						newCase := strings.TrimSpace(el.Text)
						if newCase == "N/A" || newCase == "" {
							nation.NewCase = "N/A"
						} else {
							nation.NewCase = newCase
						}
					case 4:
						totalDeath := strings.TrimSpace(el.Text)
						if totalDeath == "N/A" || totalDeath == "" {
							nation.TotalDeath = "N/A"
						} else {
							nation.TotalDeath = totalDeath
						}
					case 5:
						newDeath := strings.TrimSpace(el.Text)
						if newDeath == "N/A" || newDeath == "" {
							nation.NewDeath = "N/A"
						} else {
							nation.NewDeath = newDeath
						}
					}

				})
				nations = append(nations, nation)
			}
			return true
		})
	})
	c.Visit(consts.CoronaIndexPage)
	return nations
}

func GetCoronaIndexByCountryName(name string) model.Country {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	nation := model.Country{
		Name:       "N/A",
		TotalCase:  "N/A",
		NewCase:    "N/A",
		TotalDeath: "N/A",
		NewDeath:   "N/A",
	}

	c.OnHTML(`table[id=main_table_countries_yesterday]`, func(e *colly.HTMLElement) {
		found := false
		index := 0
		e.ForEachWithBreak("tbody tr", func(_ int, h *colly.HTMLElement) bool {
			if h.Attr("class") == "" {
				index = index + 1
				h.ForEachWithBreak(`td`, func(ei int, el *colly.HTMLElement) bool {
					switch el.Index {
					case 1:

						nationValue := strings.TrimSpace(el.ChildText(".mt_a"))
						if !strings.EqualFold(nationValue, name) {
							return false
						}
						nation.Name = nationValue
						found = true
					case 2:
						totalCase := strings.TrimSpace(el.Text)
						if totalCase == "N/A" || totalCase == "" {
							nation.TotalCase = "N/A"
						}
						nation.TotalCase = totalCase
					case 3:
						newCase := strings.TrimSpace(el.Text)
						if newCase == "N/A" || newCase == "" {
							nation.NewCase = "N/A"
						} else {
							nation.NewCase = newCase
						}
					case 4:
						totalDeath := strings.TrimSpace(el.Text)
						if totalDeath == "N/A" || totalDeath == "" {
							nation.TotalDeath = "N/A"
						} else {
							nation.TotalDeath = totalDeath
						}
					case 5:
						newDeath := strings.TrimSpace(el.Text)
						if newDeath == "N/A" || newDeath == "" {
							nation.NewDeath = "N/A"
						} else {
							nation.NewDeath = newDeath
						}
					}
					return true
				})
			}
			return !found
		})
	})
	c.Visit(consts.CoronaIndexPage)
	return nation
}
