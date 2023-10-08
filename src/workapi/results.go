package workapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JakubC-projects/work-telegram-bot/src/config"
	"github.com/JakubC-projects/work-telegram-bot/src/models"
)

type resultType struct {
	Colors models.Results `json:"colors"`
}

func GetResults(ctx context.Context) (models.Results, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, config.C.WorkApi.ResultsUrl, nil)
	if err != nil {
		return models.Results{}, fmt.Errorf("cannot create request: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.Results{}, fmt.Errorf("cannot send request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return models.Results{}, fmt.Errorf("invalid response code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Results{}, fmt.Errorf("cannot read response: %w", err)
	}

	var result resultType

	if err := json.Unmarshal(body, &result); err != nil {
		return result.Colors, fmt.Errorf("cannot parse response: %w", err)
	}

	return result.Colors, nil
}
