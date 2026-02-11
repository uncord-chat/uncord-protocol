package events

type Opcode int

const (
	OpcodeDispatch       Opcode = 0
	OpcodeHeartbeat      Opcode = 1
	OpcodeIdentify       Opcode = 2
	OpcodeResume         Opcode = 6
	OpcodeReconnect      Opcode = 7
	OpcodeInvalidSession Opcode = 9
	OpcodeHello          Opcode = 10
	OpcodeHeartbeatACK   Opcode = 11
)
