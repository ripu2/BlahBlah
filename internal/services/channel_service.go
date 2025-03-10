package services

import (
	"errors"

	"github.com/ripu2/blahblah/internal/models"
)

func CreateChanelService(channel *models.Channel) error {
	err := channel.CreateChanel()
	if err != nil {
		return errors.New(err.Error()) // return early if there is an error saving the event. No need to continue with the rest of the function.
	}
	return nil
}

func GetAllChannelsService(id int64) ([]models.Channel, error) {
	channels, err := models.GetAllChannels(id)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return channels, nil
}

func GetChannelByIdService(id int64) ([]models.Channel, error) {
	channels, err := models.GetChannelByOwnerId(id)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return channels, nil
}

func InsertUserInChannelService(user *models.ChannelUser) error {
	return user.AddToChanel()
}
