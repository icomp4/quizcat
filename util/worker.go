package util

import (
	"database/sql"
	"fmt"
	"quizcat/service"

	"github.com/robfig/cron/v3"
)

func StartPlayCountResetScheduler(db *sql.DB) (*cron.Cron, error) {
	c := cron.New()
	playsService := service.NewPlaysService(db)

	if _, err := c.AddFunc("@midnight", playsService.ResetDailyPlays); err != nil {
		return nil, fmt.Errorf("error scheduling daily play reset: %w", err)
	}

	if _, err := c.AddFunc("@weekly", playsService.ResetWeeklyPlays); err != nil {
		return nil, fmt.Errorf("error scheduling weekly play reset: %w", err)
	}

	if _, err := c.AddFunc("@monthly", playsService.ResetMonthlyPlays); err != nil {
		return nil, fmt.Errorf("error scheduling monthly play reset: %w", err)
	}

	c.Start()

	return c, nil
}
