package models

// HelloData is the payload for opcode 10 (Hello), sent by the server immediately after a WebSocket connection is
// established. The client must begin sending heartbeats at the specified interval.
type HelloData struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

// IdentifyData is the payload for opcode 2 (Identify), sent by the client to authenticate a new connection. The token
// is a valid JWT access token.
type IdentifyData struct {
	Token string `json:"token"`
}

// ResumeData is the payload for opcode 6 (Resume), sent by the client to restore a previously disconnected session.
// The server replays any events with sequence numbers greater than seq.
type ResumeData struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Seq       int64  `json:"seq"`
}

// ReadyData is the payload for the READY dispatch event (opcode 0, type "READY"), sent after a successful Identify.
// It contains the initial state the client needs to populate its local cache.
type ReadyData struct {
	SessionID string       `json:"session_id"`
	User      User         `json:"user"`
	Server    ServerConfig `json:"server"`
	Channels  []Channel    `json:"channels"`
	Roles     []Role       `json:"roles"`
	Members   []Member     `json:"members"`
}

// MessageDeleteData is the payload for the MESSAGE_DELETE dispatch event, identifying the deleted message and the
// channel it belonged to.
type MessageDeleteData struct {
	ID        string `json:"id"`
	ChannelID string `json:"channel_id"`
}
