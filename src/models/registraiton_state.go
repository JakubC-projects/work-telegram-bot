package models

import "cloud.google.com/go/civil"

type RegistrationState struct {
	UserId int64
	Reg    Registration       `json:"reg"`
	Action RegistrationAction `json:"action"`
}

type RegistrationAction struct {
	Type     RegistrationActionType
	SetDate  *civil.Date `json:"setDate,omitempty"`
	SetHours float64     `json:"setHours,omitempty"`
	SetGoal  WorkGoal    `json:"setGoal,omitempty"`
}

type RegistrationActionType int

const (
	ActionNone RegistrationActionType = iota
	ActionSetDate
	ActionSetHour
	ActionSetGoal
	ActionReadyToSend
)
