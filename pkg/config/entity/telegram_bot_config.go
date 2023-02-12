package entity

type TelegramBotConfig struct {
	AdminPassword   string `json:"adminPassword"`
	HttpToken       string `json:"httpToken"`
	RecipientChatId string `json:"recipientChatId"`
}
