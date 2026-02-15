package events

import "encoding/json"

// Frame is the wire-format structure for all WebSocket messages. Dispatch events (op 0) carry a sequence number and
// event type; control frames use only op and optionally d.
type Frame struct {
	Op   Opcode          `json:"op"`
	Seq  *int64          `json:"s,omitempty"`
	Type *DispatchEvent  `json:"t,omitempty"`
	Data json.RawMessage `json:"d,omitempty"`
}
