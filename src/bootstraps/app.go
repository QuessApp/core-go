package bootstraps

// InitApp boostraps the app, like init database, env, router, etc.
func InitApp() {
	InitEnv()
	InitDatabase()
	InitRouter()
}
