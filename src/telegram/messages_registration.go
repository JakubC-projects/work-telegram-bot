package telegram

import (
	"fmt"

	"cloud.google.com/go/civil"
	"github.com/JakubC-projects/work-telegram-bot/src/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/samber/lo"
)

func newRegistrationMessage(chatId int64, state models.RegistrationState) tgbotapi.Chattable {
	text := getRegistrationBase(state.Reg)

	actionText, buttons := getRegistrationAction(state.Action)

	text += "\n" + actionText

	return tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:      chatId,
			ReplyMarkup: buttons,
		},
		Text: text,
	}
}

func updateRegMessage(chatId int64, messageID int, state models.RegistrationState) tgbotapi.Chattable {
	text := getRegistrationBase(state.Reg)

	actionText, buttons := getRegistrationAction(state.Action)
	text += "\n" + actionText
	return tgbotapi.NewEditMessageTextAndMarkup(chatId, messageID, text, buttons)
}

func getRegistrationAction(action models.RegistrationAction) (string, tgbotapi.InlineKeyboardMarkup) {
	switch action.Type {
	case models.ActionSetDate:
		return getSetDateAction(*action.SetDate)

	case models.ActionSetHour:
		return getSetHoursAction(action.SetHours)

	case models.ActionSetGoal:
		return getSetGoalAction(action.SetGoal)
	}
	return "", tgbotapi.InlineKeyboardMarkup{}
}

func getSetDateAction(d civil.Date) (string, tgbotapi.InlineKeyboardMarkup) {
	text := fmt.Sprintf("Set Date: %s", d)
	buttons := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{{Text: "-1 Day", CallbackData: setDateActionCallback(d, -1)}, {Text: "+1 Day", CallbackData: setDateActionCallback(d, 1)}},
			{{Text: "Save", CallbackData: saveDateActionCallback(d)}},
		},
	}
	return text, buttons
}

func setDateActionCallback(d civil.Date, dayDiff int) *string {
	return lo.ToPtr(fmt.Sprintf(`updateReg:{"action":{"setDate":"%s"}}`, d.AddDays(dayDiff)))
}

func saveDateActionCallback(d civil.Date) *string {
	return lo.ToPtr(fmt.Sprintf(`updateReg:{"reg":{"date": "%s"},"action":{"type": 2}}`, d))
}

func getSetHoursAction(h float64) (string, tgbotapi.InlineKeyboardMarkup) {
	text := fmt.Sprintf("Set Hours: %.2f", h)
	buttons := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{{Text: "-1 Hour", CallbackData: setHourActionCallback(h, -1)}, {Text: "+1 Hour", CallbackData: setHourActionCallback(h, 1)}},
			{{Text: "-15 Minutes", CallbackData: setHourActionCallback(h, -0.25)}, {Text: "+15 Minutes", CallbackData: setHourActionCallback(h, 0.25)}},
			{{Text: "Save", CallbackData: saveHourActionCallback(h)}},
		},
	}
	return text, buttons
}

func setHourActionCallback(h float64, diff float64) *string {
	newHours := h + diff
	if newHours < 0 {
		newHours = 0
	}
	return lo.ToPtr(fmt.Sprintf(`updateReg:{"action":{"setHours":%f}}`, newHours))
}

func saveHourActionCallback(h float64) *string {
	return lo.ToPtr(fmt.Sprintf(`updateReg:{"reg":{"hours": %f},"action":{"type": 3}}`, h))
}

func getSetGoalAction(g models.WorkGoal) (string, tgbotapi.InlineKeyboardMarkup) {
	text := fmt.Sprintf("Set Goal: %s", g)
	buttons := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{{Text: string(models.GoalBUK), CallbackData: setGoalActionCallback(models.GoalBUK)}},
			{{Text: string(models.GoalSamvirk), CallbackData: setGoalActionCallback(models.GoalSamvirk)}},
			{{Text: string(models.GoalOther), CallbackData: setGoalActionCallback(models.GoalOther)}},
			{{Text: "Save", CallbackData: saveGoalActionCallback(g)}},
		},
	}
	return text, buttons
}

func setGoalActionCallback(g models.WorkGoal) *string {
	return lo.ToPtr(fmt.Sprintf(`updateReg:{"action":{"setGoal":"%s"}}`, g))
}

func saveGoalActionCallback(g models.WorkGoal) *string {
	if g == "" {
		return lo.ToPtr("noop")
	}
	return lo.ToPtr(fmt.Sprintf(`updateReg:{"reg":{"goal": "%s"},"action":{"type": 0}}`, g))
}
