package services

import (
	"errors"

	"github.com/ripu2/blahblah/internal/models"
)

func CreateChanelService(channel *models.Channel, creatorId int) error {
	channel.CreatedBy = creatorId
	err := channel.CreateChanel()
	if err != nil {
		return errors.New("failed to create chanel") // return early if there is an error saving the event. No need to continue with the rest of the function.
	}
	return nil
}
