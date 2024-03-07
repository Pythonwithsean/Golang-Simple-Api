package main

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		// DBAdress:   fmt.SSprintf()
	}
}

func getEnv(key, fallback string) string {
	if value, ok := LookupEnv(key); ok {
		return value
	}
}
