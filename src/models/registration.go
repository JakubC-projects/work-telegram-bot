package models

import "cloud.google.com/go/civil"

type Registration struct {
	Name            string
	WorkDescription string

	Date     *civil.Date
	Hours    float64
	Color    Color
	WorkGoal WorkGoal
}

type Color string

const (
	ColorGreen  Color = "🟢 GREEN"
	ColorRed    Color = "🔴 RED"
	ColorOrange Color = "🟠 ORANGE"
	ColorBlue   Color = "🔵 BLUE"
)

type WorkGoal string

const (
	GoalSamvirk WorkGoal = "Samvirk"
	GoalBUK     WorkGoal = "BUK"
	GoalOther   WorkGoal = "Inne"
)
