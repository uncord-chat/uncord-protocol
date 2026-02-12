// Package events defines the WebSocket opcodes and dispatch event types
// used by the Uncord gateway protocol.
package events

// Opcode identifies the type of a WebSocket frame. The gateway uses
// opcodes to distinguish control frames (heartbeat, identify) from
// dispatch events that carry application data.
type Opcode int

const (
	// OpcodeDispatch delivers a dispatch event (server → client).
	OpcodeDispatch Opcode = 0
	// OpcodeHeartbeat is a keep-alive ping (bidirectional).
	OpcodeHeartbeat Opcode = 1
	// OpcodeIdentify authenticates a new connection (client → server).
	OpcodeIdentify Opcode = 2
	// OpcodeResume resumes a dropped session (client → server).
	OpcodeResume Opcode = 6
	// OpcodeReconnect asks the client to reconnect (server → client).
	OpcodeReconnect Opcode = 7
	// OpcodeInvalidSession tells the client its session is invalid (server → client).
	OpcodeInvalidSession Opcode = 9
	// OpcodeHello is sent on connect with the heartbeat interval (server → client).
	OpcodeHello Opcode = 10
	// OpcodeHeartbeatACK acknowledges a heartbeat (server → client).
	OpcodeHeartbeatACK Opcode = 11
)
