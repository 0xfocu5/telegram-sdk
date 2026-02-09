package telegram

import (
	"log"
	"testing"
)

const (
	VolumeBotApi     = "8028314379:AAH1SWUMU0FbI5LZ7_AF3i94WEc9KZSoqE0"
	HengPanBotApi    = "7441266721:AAEgg3NgLNj0B0R1oFek-Sl5dsMjbiPm1_U"
	VolumeChannel    = -4887368708
	BscVolumeChannel = -4803318876
	MyselfChannel    = 8065483765
)

func Test_send_msg(t *testing.T) {
	config := &Config{
		Token: VolumeBotApi,
		Debug: true,
	}

	tg, err := NewBot(config)
	if err != nil {
		t.Fatalf("Failed to create bot: %v", err)
	}

	_, err = tg.SendMessage("Hello, world!", BscVolumeChannel, 0, true)
	if err != nil {
		log.Printf("Send Message Warn: %v", err)
		// We log instead of fail because this tries to hit real API
	}
}
