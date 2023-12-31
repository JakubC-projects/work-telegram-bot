package telegram

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/civil"
	"github.com/JakubC-projects/work-telegram-bot/src/db"
	"github.com/JakubC-projects/work-telegram-bot/src/log"
	"github.com/JakubC-projects/work-telegram-bot/src/models"
	"github.com/JakubC-projects/work-telegram-bot/src/workapi"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/samber/lo"
)

func startRegistration(ctx context.Context, m *tgbotapi.Message) error {
	_, description, _ := strings.Cut(m.Text, " ")

	state := models.RegistrationState{
		UserId: m.From.ID,
		Reg: models.Registration{
			Name:            m.From.FirstName + " " + m.From.LastName,
			Color:           models.ColorGreen,
			WorkDescription: description,
		},
		Action: models.RegistrationAction{
			Type:    models.ActionSetDate,
			SetDate: lo.ToPtr(civil.DateOf(time.Now())),
		},
	}

	msg, err := client.Send(newRegistrationMessage(m.Chat.ID, state))
	if err != nil {
		log.L.Error("cannot send message",
			"err", err)
		return err
	}

	if err := db.SaveRegistration(ctx, msg.Chat.ID, msg.MessageID, state); err != nil {
		log.L.Error("cannot create new registration",
			"err", err)
		return nil
	}

	return nil
}

func updateReg(ctx context.Context, c *tgbotapi.CallbackQuery) error {
	_, updateData, _ := strings.Cut(c.Data, ":")

	reg, err := db.GetRegistration(ctx, c.Message.Chat.ID, c.Message.MessageID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil
		}
		return err
	}

	if err := json.Unmarshal([]byte(updateData), &reg); err != nil {
		return fmt.Errorf("cannot unmarshal json: %w", err)
	}

	if err := db.SaveRegistration(ctx, c.Message.Chat.ID, c.Message.MessageID, reg); err != nil {
		return fmt.Errorf("cannot save updated state: %w", err)
	}

	_, err = client.Send(updateRegMessage(c.Message.Chat.ID, c.Message.MessageID, reg))

	return err
}

func sendRegistration(ctx context.Context, c *tgbotapi.CallbackQuery) error {
	reg, err := db.GetRegistration(ctx, c.Message.Chat.ID, c.Message.MessageID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil
		}
		return err
	}

	err = workapi.SaveRegistration(ctx, reg.Reg)
	if err != nil {
		err = fmt.Errorf("cannot save registration: %w", err)
		log.L.Error("cannot save registration", "err", err)
		client.Send(tgbotapi.NewMessage(c.Message.Chat.ID, err.Error()))
		return err
	}

	_, err = client.Send(sendRegMessage(c.Message.Chat.ID, c.Message.MessageID, reg))
	if err != nil {
		return err
	}

	results, err := workapi.GetResults(ctx)
	if err != nil {
		err = fmt.Errorf("cannot get results: %w", err)
		log.L.Error("cannot get results", "err", err)
		client.Send(tgbotapi.NewMessage(c.Message.Chat.ID, err.Error()))
		return nil
	}

	_, err = client.Send(newResultsMessage(c.Message.Chat.ID, results))

	return err
}
