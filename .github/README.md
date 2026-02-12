<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/uncord-chat/.github/main/profile/logo-banner-dark.png">
    <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/uncord-chat/.github/main/profile/logo-banner-light.png">
    <img alt="Uncord" src="https://raw.githubusercontent.com/uncord-chat/.github/main/profile/logo-banner-light.png" width="1280">
  </picture>
</p>

Shared protocol definitions for the [Uncord](https://github.com/uncord-chat) project. Contains type definitions, constants, and API schemas used by both the server and client. Every definition exists in both Go and TypeScript, kept in sync manually.

### Contents

| Package | Description |
|---------|-------------|
| `models/` | Shared entity types (User) and API request/response types (Go + TypeScript) |
| `permissions/` | Permission bitfield constants and helpers (Go + TypeScript) |
| `events/` | Gateway opcodes and dispatch event type names (Go + TypeScript) |
| `errors/` | API error code string constants (Go + TypeScript) |

### Usage

**Go (server)**

```go
import (
    "github.com/uncord-chat/uncord-protocol/models"
    "github.com/uncord-chat/uncord-protocol/permissions"
)

perms := permissions.SendMessages | permissions.AddReactions
if perms.Has(permissions.SendMessages) {
    // allowed
}

resp := models.AuthResponse{
    User:        models.User{ID: "...", Email: "...", Username: "..."},
    AccessToken: "...",
}
```

**TypeScript (client)**

```typescript
import * as permissions from "@uncord-chat/protocol/permissions/bitfield";
import type { AuthResponse } from "@uncord-chat/protocol/models/auth";

const perms = permissions.SendMessages | permissions.AddReactions;
if (permissions.has(perms, permissions.SendMessages)) {
  // allowed
}

const resp: AuthResponse = await fetch("/api/v1/auth/login", { ... }).then(r => r.json());
```

### Development

```bash
make fmt            # Format Go and TypeScript
make lint           # Lint Go (go vet) and TypeScript (biome)
make test           # Run Go and TypeScript tests
make fmt-check      # CI check — fails if formatting is off
```

### Related repositories

| Repository | Description |
|-----------|-------------|
| [uncord-server](https://github.com/uncord-chat/uncord-server) | Go server. REST API, WebSocket gateway, permission engine, plugin system |
| [uncord-client](https://github.com/uncord-chat/uncord-client) | React Native client for Windows, macOS, Linux, iOS, and Android |
| [uncord-docs](https://github.com/uncord-chat/uncord-docs) | User and admin documentation |

### License

MIT
