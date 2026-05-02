package controllers

import (
	"errors"
)

type LoginAdapter interface {
	send() error
}

type TelegramLoginAdapter struct {
}

func (adapter *TelegramLoginAdapter) send() error {
	var err = errors.New("Not implemented")
	return err
}

func SendLoginCode(adapter LoginAdapter) {
	adapter.send()
}
