package telegram

import (
	"fmt"
	"sort"

	"github.com/JakubC-projects/work-telegram-bot/src/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type teamResult struct {
	team  models.Color
	value float64
}

func newResultsMessage(chatId int64, res models.Results) tgbotapi.Chattable {
	text := "Results:\n"

	teamResults := resultsToTeamResults(res)

	for i, res := range teamResults {
		text += fmt.Sprintf("%d. %s: %.2f\n", i+1, res.team, res.value)
	}

	return tgbotapi.NewMessage(chatId, text)
}

func resultsToTeamResults(res models.Results) []teamResult {
	teamResults := []teamResult{
		{models.ColorGreen, res.Green},
		{models.ColorRed, res.Red},
		{models.ColorOrange, res.Orange},
		{models.ColorBlue, res.Blue},
	}

	sort.Slice(teamResults, func(i, j int) bool {
		return teamResults[i].value > teamResults[j].value
	})

	return teamResults

}
