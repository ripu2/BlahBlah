package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type Message struct {
	ID          int       `json:"id" db:"id"`
	ChannelID   uuid.UUID `json:"channel_id" db:"channel_id"`
	SenderID    uuid.UUID `json:"sender_id" db:"sender_id"`
	Content     string    `json:"content" db:"content"`
	MessageType string    `json:"message_type" db:"message_type"` // text, image, video, etc.
	Metadata    JSONB     `json:"metadata" db:"metadata"`
	SentAt      time.Time `json:"sent_at" db:"sent_at"`
}

type ArchivedMessage struct {
	ID          int        `json:"id" db:"id"`
	ChannelID   *uuid.UUID `json:"channel_id,omitempty" db:"channel_id"`
	SenderID    *uuid.UUID `json:"sender_id,omitempty" db:"sender_id"`
	Content     string     `json:"content" db:"content"`
	MessageType string     `json:"message_type" db:"message_type"`
	Metadata    JSONB      `json:"metadata" db:"metadata"`
	SentAt      time.Time  `json:"sent_at" db:"sent_at"`
}
