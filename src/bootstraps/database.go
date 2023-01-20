package bootstraps

import "core/src/database"

// InitDatabase estabilishes database connection.
func InitDatabase() {
	_, err := database.Connect()

	if err != nil {
		panic(err)
	}
}
