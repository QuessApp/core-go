package bootstraps

// InitApp boostraps the app, like init database, env, router, etc.
// The term ``bootstrap`` here means something like `initializer`.
// A function that initializes something is a function that will run at the moment the app runs.
func InitApp() {
	InitEnv()
	InitDatabase()
	InitRouter()
}
