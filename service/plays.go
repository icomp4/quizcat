package service

import (
	"database/sql"
	"log"
)

type PlaysService struct {
	db *sql.DB
}

func NewPlaysService(db *sql.DB) *PlaysService {
	return &PlaysService{
		db: db}
}
func (p *PlaysService) ResetDailyPlays() {
	_, err := p.db.Exec("UPDATE quizzes SET daily_plays = 0")
	if err != nil {
		log.Fatal("error resetting daily plays: %w", err)
	}
}

func (p *PlaysService) ResetWeeklyPlays() {
	_, err := p.db.Exec("UPDATE quizzes SET weekly_plays = 0")
	if err != nil {
		log.Fatal("error resetting weekly plays: %w", err)
	}

}

func (p *PlaysService) ResetMonthlyPlays() {
	_, err := p.db.Exec("UPDATE quizzes SET monthly_plays = 0")
	if err != nil {
		log.Fatal("error resetting monthly plays: %w", err)
	}

}