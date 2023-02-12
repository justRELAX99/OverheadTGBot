package entity

type TelegramBotConfig struct {
	AdminPassword   string `json:"adminPassword"`
	HttpToken       string `json:"httpToken"`
	RecipientChatId int64  `json:"recipientChatId"`
}
