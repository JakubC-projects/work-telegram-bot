package workapi

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/JakubC-projects/work-telegram-bot/src/config"
	"github.com/JakubC-projects/work-telegram-bot/src/models"
)

func SaveRegistration(ctx context.Context, reg models.Registration) error {

	data := url.Values{}
	data.Add("entry.1654063466", reg.Name)
	data.Add("entry.1172831946", fmt.Sprintf("%.2f", reg.Hours))
	data.Add("entry.36799388", reg.WorkDescription)
	data.Add("entry.805261923_year", fmt.Sprint(reg.Date.Year))
	data.Add("entry.805261923_month", fmt.Sprint(int(reg.Date.Month)))
	data.Add("entry.805261923_day", fmt.Sprint(reg.Date.Day))
	data.Add("entry.886802696", string(reg.Color))
	data.Add("entry.1705999664", string(reg.WorkGoal))
	data.Add("usp", "pp_url")
	data.Add("submit", "Submit")

	fullUrl := config.C.WorkApi.FormUrl + "?" + data.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullUrl, nil)
	if err != nil {
		return fmt.Errorf("cannot create request: %w", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("cannot send request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("cannot read response: %w", err)
	}

	if !strings.Contains(string(body), "Your response has been recorded") {
		return fmt.Errorf("response not successful")
	}

	return nil
}
