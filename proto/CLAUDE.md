# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Protocol buffer definitions for the Aruka chat application. Generated code is output to sibling directories: `../arukabe/` (Go backend) and `../arukaweb/` (React frontend).

## Commands

```bash
buf lint          # lint proto files (STANDARD rules)
buf breaking      # check for breaking changes (FILE rules)
buf generate      # generate Go + TypeScript code from proto definitions
```

## Architecture

Protos are split into two packages:

- **`models/v1/`** — shared message types: `Provider`, `Model`, `Chat`, `Message`, `MessageContent` (oneof for text/thinking/tool blocks), `Pagination`, `TimeData`
- **`chat/v1/`** — `ChatService` RPC definitions and their request/response types
- **`provider/v1/`** — `ProviderService` RPC definitions and their request/response types

### Code generation targets (`buf.gen.yaml`)

| Plugin | Output | Purpose |
|---|---|---|
| `protoc-gen-go` | `../arukabe/gen/connect/aruka` | Go message types |
| `protoc-gen-connect-go` | `../arukabe/gen/connect/aruka` | Connect-RPC handler interfaces (simple mode) |
| `protoc-gen-es` | `../arukaweb/src/connect/arukabe` | TypeScript types (with `include_imports`) |

After editing any `.proto` file, run `buf generate` to update generated code in both downstream directories.

### Services

**ChatService** (`chat/v1/services.proto`): `ListChats`, `NewChat`, `EditChat`, `ChatMessages`, `ChatMessage` (streaming), `AutoChatTitleUpdate`

**ProviderService** (`provider/v1/services.proto`): `ListProviders`, `UpdateProviderModels`
