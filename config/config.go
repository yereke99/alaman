// config/config.go
package config

import (
	"os"
)

// Config contains application configuration parameters
type Config struct {
	Port              string `json:"port"`
	Token             string `json:"token"`
	BaseURL           string `json:"base_url"`
	DBName            string `json:"db_name"`
	SavePaymentsDir   string `json:"save_payments_dir"`
	AdminID           int64  `json:"admin_id"`
	AdminID2          int64  `json:"admin_id2"`
	AdminID3          int64  `json:"admin_id3"`
	StartPhotoId      string `json:"start_photo_id"`
	StartVideoId      string `json:"start_video_id"`
	InstructorVideoId string `json:"instructor_video"`
	Cost              int    `json:"cost"`
	BotUsername       string `json:"bot_username"`
	Bin               string `json:"bin"`
	Bin2              string `json:"bin2"`
	Bin3              string `json:"bin3"`
	Bin4              string `json:"bin4"`
}

// NewConfig creates and returns a new configuration instance
func NewConfig() (*Config, error) {
	cfg := &Config{
		Port:              ":8081",
		Token:             "8050408674:AAGrHR3LHyIWEfF2ZNVt2x8wr29MxUzvM9g",
		BaseURL:           "https://ccc8-89-219-13-135.ngrok-free.app", // Update this with your actual domain
		DBName:            "alaman.db",
		SavePaymentsDir:   "./payment",
		AdminID:           800703982,
		AdminID2:          7854239462,
		AdminID3:          685953723,
		StartPhotoId:      "AgACAgIAAxkBAANSaFP5emhGuJ5qTUamzTYon-yyPv4AAszxMRuxzqBKW2jULQVc0e4BAAMCAAN5AAM2BA",
		StartVideoId:      "BAACAgIAAxkBAAIGQ2hs996Wo5tLH-aZu32XGWhcBjMxAALFeQACM7hoSwWQNDUxWvt-NgQ",
		InstructorVideoId: "BAACAgIAAxkBAAIExWhf1MIAAZ0mGONHcGxOWRPHa4SRLAACXnUAAj8UAUt-qpkmBZGhqjYE",
		Cost:              18900,
		BotUsername:       "Alaman",
		Bin:               "870304301209",
		Bin2:              "060301551728",
		Bin3:              "011225600097",
		Bin4:              "010514551360",
	}

	// Override with environment variables if set
	if port := os.Getenv("PORT"); port != "" {
		cfg.Port = ":" + port
	}

	if token := os.Getenv("BOT_TOKEN"); token != "" {
		cfg.Token = token
	}

	if baseURL := os.Getenv("BASE_URL"); baseURL != "" {
		cfg.BaseURL = baseURL
	}

	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		cfg.DBName = dbName
	}

	if savePaymentsDir := os.Getenv("SAVE_PAYMENTS_DIR"); savePaymentsDir != "" {
		cfg.DBName = savePaymentsDir
	}

	return cfg, nil
}
