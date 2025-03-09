package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/ripu2/blahblah/internal/config/db"
)

type Channel struct {
	ID        int64           `json:"id" db:"id"`
	Name      string          `json:"name" db:"name"`
	CreatedBy int64           `json:"created_by" db:"created_by"`
	IsPrivate bool            `json:"is_private" db:"is_private"`
	Metadata  json.RawMessage `json:"metadata" db:"metadata"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
}

type ChannelUser struct {
	ChannelID int64     `json:"channel_id" db:"channel_id"`
	UserID    int64     `json:"user_id" db:"user_id"`
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
	metadataBytes, _ := json.Marshal(chanel.Metadata)
	err = db.DB.QueryRow(query, chanel.Name, chanel.CreatedBy, chanel.IsPrivate, metadataBytes).Scan(&chanel.ID)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil

}

func GetAllChannels() ([]Channel, error) {
	query := `SELECT * FROM communication_channel`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	channels := make([]Channel, 0, 10)

	for rows.Next() {
		var chanel Channel
		var metadata json.RawMessage

		err := rows.Scan(
			&chanel.ID,
			&chanel.Name,
			&chanel.CreatedBy,
			&chanel.IsPrivate,
			&metadata, // ✅ JSONB type ab Scan properly karega
			&chanel.CreatedAt,
		)
		chanel.Metadata = metadata // ✅ Directly assign `json.RawMessage`
		if err != nil {
			return nil, err
		}

		channels = append(channels, chanel)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return channels, nil
}

func GetChannelByOwnerId(ownerId int64) ([]Channel, error) {
	query := `SELECT * FROM communication_channel WHERE created_by = $1`
	rows, err := db.DB.Query(query, ownerId)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	channels := make([]Channel, 0, 10)

	for rows.Next() {
		var chanel Channel
		var metadata json.RawMessage

		err := rows.Scan(
			&chanel.ID,
			&chanel.Name,
			&chanel.CreatedBy,
			&chanel.IsPrivate,
			&metadata, // ✅ JSONB type ab Scan properly karega
			&chanel.CreatedAt,
		)
		chanel.Metadata = metadata // ✅ Directly assign `json.RawMessage`
		if err != nil {
			return nil, err
		}

		channels = append(channels, chanel)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return channels, nil
}
