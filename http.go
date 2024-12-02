package boticordgo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) SendStats(payload StatsPayload) error {
	resp, err := c.DoRequest("POST", "/stats", payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %v", err)
	}
	return nil
}

func (c *Client) GetBotInfo(BotID string) (*BotInfo, error) {
	resp, err := c.DoRequest("GET", fmt.Sprintf("/bots/%s", BotID), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var BotInfo BotInfo
	if err := json.NewDecoder(resp.Body).Decode(&BotInfo); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return &BotInfo, nil
}

func (c *Client) PostBotStats(payload StatsPayload) (*[]BotInfo, error) {
	resp, err := c.DoRequest("POST", "/stats", payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var BotsInfo []BotInfo

	if err := json.NewDecoder(resp.Body).Decode(&BotsInfo); err != nil {
		return nil, err
	}
	return &BotsInfo, nil
}
