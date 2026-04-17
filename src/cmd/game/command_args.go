package main

import "flag"

type cliArgs struct {
	gameDataFolder string
}

func bindFlags(dst *cliArgs) {
	flag.StringVar(&dst.gameDataFolder, "data", "game_data",
		"game data folder location")

	flag.Parse()
}
