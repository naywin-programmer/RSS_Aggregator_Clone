package configs

import (
	"time"

	"github.com/naywin-programmer/RSS_Aggregator/services"
)

func RunServices() {
	go services.RssScraperRun(DBConnection.DB, 10, 30*time.Second, time.Now().UTC().Add(-24*time.Hour))
}
