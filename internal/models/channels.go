package models

import (
	"errors"
	"time"

	"github.com/ripu2/blahblah/internal/config/db"
)

type Channel struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedBy int       `json:"created_by" db:"created_by"`
	IsPrivate bool      `json:"is_private" db:"is_private"`
	Metadata  JSONB     `json:"metadata" db:"metadata"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type ChannelUser struct {
	ChannelID int       `json:"channel_id" db:"channel_id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Role      string    `json:"role" db:"role"` // Admin, member, etc.
	JoinedAt  time.Time `json:"joined_at" db:"joined_at"`
}

// JSONB type to handle JSON data
type JSONB map[string]interface{}

func (chanel *Channel) CreateChanel() error {
	var exists *bool
	if db.DB == nil {
		return errors.New("database connection is nil or uninitialized")
	}
	err := db.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM communication_channel WHERE id = $1)", chanel.ID).Scan(&exists)
	if err != nil {
		return err
	}
	if *exists {
		return errors.New("channel already exists")
	}
	query := `
	INSERT INTO communication_channel (name, created_by, is_private, metadata, created_at)
	VALUES ($1, $2, $3, $4, NOW()) RETURNING id;
	`
	err = db.DB.QueryRow(query, chanel.Name, chanel.CreatedBy, chanel.IsPrivate, chanel.Metadata).Scan(&chanel.ID)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil

}
