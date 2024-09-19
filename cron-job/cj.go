package cronjob

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {

	c := cron.New()

	c.AddFunc("@every 10s", func() {
		fmt.Println("Cron job executed at", time.Now())
	})

	c.Start()

	select {}
}
