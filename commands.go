package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore <area>",
			description: "Shows pokemon in the area",
			callback:    commandExplore,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Shows stats for a pokemon (if caught!)",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "List next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 locations",
			callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List names of all pokemon caught!",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
