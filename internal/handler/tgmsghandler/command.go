package tgmsghandler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// This function should only be called when the message is a command.
func HandleCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) error {
	cmd := msg.Command()

	var response string
	var err error
	switch cmd {
	case HelpCmdWord:
		response, err = HandleHelpCommand(msg)
	case AuthCmdWord:
		response, err = HandleAuthCommand(msg)
	case StartCmdWord:
		response, err = HandleStartCommand(msg)
	case AskCmdWord:
		response, err = HandleAskCommand(msg)
	case QueryCmdWord:
		response, err = HandleQueryCommand(msg)
	case RequestCmdWord:
		response, err = HandleRequestCommand(msg)
	}

	// Handle error
	SendTextMessage(bot, msg, response)

	if err == nil {
		aiResponse := GetAIResponse(msg.Chat.ID)
		return SendTextMessage(bot, msg, aiResponse)
	}

	return err
}