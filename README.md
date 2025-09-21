# Pokedex CLI

A command-line Pokédex application written in Go. This tool allows you to explore Pokémon locations, catch Pokémon, inspect their stats, and manage your own Pokédex, all from your terminal.

## Features
- Explore Pokémon location areas
- List Pokémon in a location area
- Catch Pokémon with a random chance
- Inspect caught Pokémon for stats and types
- View all Pokémon in your Pokédex
- Caching of API responses for performance

## Commands
- `help` — Prints help menu
- `map` — List some location areas
- `mapb` — List the previous page of location areas
- `explore {location_area}` — List Pokémon in a location area
- `catch {pokemon_name}` — Attempt to catch a Pokémon and add it to your Pokédex
- `inspect {pokemon_name}` — View information about a caught Pokémon
- `pokedex` — View all Pokémon in your Pokédex
- `exit` — Exit the Pokédex CLI

## Usage
1. **Build the project:**
   ```sh
   go build -o pokedex-cli
   ```
2. **Run the CLI:**
   ```sh
   ./pokedex-cli
   ```
3. **Type commands at the prompt (`>`):**
   - Example: `explore route-1`
   - Example: `catch pikachu`

## Development
- Written in Go
- Uses [PokéAPI](https://pokeapi.co/) for Pokémon data
- Caching implemented for API responses
- Thread-safe cache with automatic reaping of expired entries

## Testing
Run tests with:
```sh
go test ./...
```

## Project Structure
- `main.go` — Entry point
- `repl.go` — REPL loop and command parsing
- `command_*.go` — Command implementations
- `internal/pokeapi/` — PokéAPI client code
- `internal/pokecache/` — Cache implementation

## Requirements
- Go 1.18 or newer
- Internet connection (for PokéAPI)

