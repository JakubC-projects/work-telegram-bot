package workapi

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"github.com/JakubC-projects/work-telegram-bot/src/models"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestSaveRegistration(t *testing.T) {
	reg := models.Registration{
		Name:            "Jakub Czyz",
		WorkDescription: "Test",
		Date:            lo.ToPtr(civil.DateOf(time.Now())),
		Hours:           0.01,
		Color:           models.ColorGreen,
		WorkGoal:        models.GoalSamvirk,
	}

	err := SaveRegistration(context.Background(), reg)
	assert.NoError(t, err)
}
