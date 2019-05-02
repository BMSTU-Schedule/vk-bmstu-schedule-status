package vk_bmstu_schedule_status

import (
	"errors"
	"net/url"
	"strconv"

	vkapi "github.com/dimonchik0036/vk-api"
)

const (
	MaxChatId uint = 100000000
)

type VkApiClient interface {
	Connect(token string) error
	EditChat(chatID uint, title string) error
}

type DefaultVkApiClient struct {
	instance *vkapi.Client
}

func (client *DefaultVkApiClient) Connect(login string, password string) error {
	instance, err := vkapi.NewClientFromLogin(login, password, vkapi.ScopeMessages)
	if err != nil {
		return err
	}
	client.instance = instance
	return nil
}

func (client *DefaultVkApiClient) EditChat(chatID uint, title string) error {
	if chatID > MaxChatId {
		return errors.New("invalid chat ID")
	}

	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(int(chatID)))
	values.Add("title", title)

	if _, err := client.instance.Do(vkapi.NewRequest("messages.editChat", "", values)); err != nil {
		return errors.New("request failed")
	}
	return nil
}
