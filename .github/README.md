<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/uncord-chat/.github/main/profile/logo-banner-dark.png">
    <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/uncord-chat/.github/main/profile/logo-banner-light.png">
    <img alt="Uncord" src="https://raw.githubusercontent.com/uncord-chat/.github/main/profile/logo-banner-light.png" width="1280">
  </picture>
</p>

Shared protocol definitions for the [Uncord](https://github.com/uncord-chat) project. Contains type definitions, constants, and schemas used by both the server and client with no business logic.

### Contents

| Package | Description |
|---------|-------------|
| `permissions/` | Permission bitfield constants and helpers (Go + TypeScript) |
| `events/` | Gateway opcodes and dispatch event type names (Go + TypeScript) |
| `errors/` | API error code string constants (Go + TypeScript) |

### Usage

**Go (server)**

```go
import "github.com/uncord-chat/uncord-protocol/permissions"

perms := permissions.SendMessages | permissions.AddReactions
if perms.Has(permissions.SendMessages) {
    // allowed
}
```

**TypeScript (client)**

```typescript
import { SendMessages, AddReactions, has } from "@uncord-chat/protocol/permissions/bitfield";

const perms = SendMessages | AddReactions;
if (has(perms, SendMessages)) {
  // allowed
}
```

### Related repositories

| Repository | Description |
|-----------|-------------|
| [uncord-server](https://github.com/uncord-chat/uncord-server) | Go server. REST API, WebSocket gateway, permission engine, plugin system |
| [uncord-client](https://github.com/uncord-chat/uncord-client) | React Native client for Windows, macOS, Linux, iOS, and Android |
| [uncord-docs](https://github.com/uncord-chat/uncord-docs) | User and admin documentation |

### License

MIT
