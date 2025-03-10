package db

const (
	CreateUserTableQuery = `CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(25) NOT NULL UNIQUE,
    first_name VARCHAR(25) NOT NULL,
    last_name VARCHAR(25) NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    dob DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

	CreateChanelTableQuery = `CREATE TABLE IF NOT EXISTS communication_channel (
    id SERIAL PRIMARY KEY,
    name TEXT,
    created_by INT NOT NULL,
    is_private BOOLEAN DEFAULT TRUE, 
    metadata JSONB, -- Extra details (photo, rules, pinned message, etc.)
    created_at TIMESTAMP DEFAULT now(),
    
    CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL
);`

	CreateChannelUserTableQuery = `CREATE TABLE IF NOT EXISTS channel_users (
    channel_id INT NOT NULL,
    user_id INT NOT NULL UNIQUE,
    role TEXT CHECK (role IN ('admin', 'member', 'moderator')), -- Role validation
    joined_at TIMESTAMP DEFAULT now(),

    PRIMARY KEY (channel_id, user_id),
    CONSTRAINT fk_channel FOREIGN KEY (channel_id) REFERENCES communication_channel(id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);`

	CreateMessageTableQuery = `CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    channel_id INT NOT NULL,
    sender_id INT NOT NULL,
    content TEXT NOT NULL,
    message_type TEXT CHECK (message_type IN ('text', 'image', 'video', 'file')), -- Type validation
    metadata JSONB, -- Extra details like attachments, reactions, etc.
    sent_at TIMESTAMP DEFAULT now(),

    CONSTRAINT fk_channel FOREIGN KEY (channel_id) REFERENCES communication_channel(id) ON DELETE CASCADE,
    CONSTRAINT fk_sender FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE SET NULL
);`

	CreateArchivedMessageTableQuery = `CREATE TABLE IF NOT EXISTS archived_messages (
    id SERIAL PRIMARY KEY,
    channel_id INT,
    sender_id INT,
    content TEXT NOT NULL,
    message_type TEXT CHECK (message_type IN ('text', 'image', 'video', 'file')),
    metadata JSONB,
    sent_at TIMESTAMP NOT NULL,

    CONSTRAINT fk_channel FOREIGN KEY (channel_id) REFERENCES communication_channel(id) ON DELETE SET NULL,
    CONSTRAINT fk_sender FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE SET NULL
);`
)
