import type { DispatchEvent } from "./dispatch";
import type { Opcode } from "./opcodes";

/** Wire-format structure for all WebSocket messages. */
export interface Frame {
  op: Opcode;
  s?: number | null;
  t?: DispatchEvent | null;
  d?: unknown;
}
