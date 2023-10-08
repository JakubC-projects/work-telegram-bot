package models

import "cloud.google.com/go/civil"

type Registration struct {
	Name            string
	WorkDescription string

	Date     *civil.Date
	Hours    float64
	Color    string
	WorkGoal string
}

type Color string

const (
	ColorGreen Color = "🟢 GREEN"
)

type WorkGoal string

const (
	GoalSamvirk WorkGoal = "Samvirk"
	GoalBUK     WorkGoal = "BUK"
	GoalOther   WorkGoal = "Inne"
)

// entry.1654063466=Test&entry.1172831946=1&entry.36799388=Test&entry.805261923_year=2023&entry.805261923_month=6&entry.805261923_day=10&entry.886802696=%F0%9F%9F%A2+GREEN&entry.1705999664=BUK&dlut=1696782652951&entry.886802696_sentinel=&entry.1705999664_sentinel=&fvv=1&partialResponse=%5Bnull%2Cnull%2C%22-6928041915760223976%22%5D&pageHistory=0&fbzx=-6928041915760223976

// https://docs.google.com/forms/u/0/d/e/1FAIpQLSeTCI2_De-uWM_XAk-SaqfsddGAtMvqetIsoDYWN_FQzJuFlg/formResponse
