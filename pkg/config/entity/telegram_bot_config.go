package entity

type TelegramBotConfig struct {
	Name            string `json:"name"`
	UserName        string `json:"userName"`
	HttpToken       string `json:"httpToken"`
	OutputPaths     string `json:"outputPaths"`
	RecipientChatId string `json:"recipientChatId"`
}