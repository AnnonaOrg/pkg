package autoload

/*
	You can just read the .env file on import just by doing

		import _ "github.com/AnnonaOrg/pkg/godotenv/autoload"

	And bob's your mother's brother
*/

import "github.com/AnnonaOrg/pkg/godotenv"

func init() {
	godotenv.Load()
}
