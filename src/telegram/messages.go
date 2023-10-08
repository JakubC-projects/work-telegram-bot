package telegram

import (
	"encoding/json"
	"fmt"

	"github.com/JakubC-projects/work-telegram-bot/src/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func newWelcomeMessage(chatId int64) tgbotapi.Chattable {
	return tgbotapi.NewMessage(chatId,
		`Welcome to WORK the work bot`)
}

func newResultsMessage(chatId int64, results models.Results) tgbotapi.Chattable {
	text, _ := json.Marshal(results)
	return tgbotapi.NewMessage(chatId,
		`Results: `+string(text))
}

func getRegistrationBase(reg models.Registration) string {
	var result = "Your registration\n"

	if reg.Name != "" {
		result += "Name: " + reg.Name + "\n"
	}
	if reg.WorkDescription != "" {
		result += "Description: " + reg.WorkDescription + "\n"
	}
	if reg.Color != "" {
		result += "Color: " + reg.Color + "\n"
	}
	if reg.Date != nil {
		result += "Date: " + reg.Date.String() + "\n"
	}

	if reg.Hours != 0 {
		result += fmt.Sprintf("Hours: %.2f\n", reg.Hours)
	}

	if reg.WorkGoal != "" {
		result += fmt.Sprintf("Work goal: %s\n", reg.WorkGoal)
	}

	return result
}
