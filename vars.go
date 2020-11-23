package cielo

import "os"

var (
	// ProductionEnvironment sets the environment to production
	ProductionEnvironment = Environment{
		APIUrl:      os.Getenv("PRODUCTION_API_URL"),
		APIQueryURL: os.Getenv("PRODUCTION_API_QUERY_URL"),
	}
	// SandboxEnvironment sets the environment to sandbox
	SandboxEnvironment = Environment{
		APIUrl:      os.Getenv("SANDBOX_API_URL"),
		APIQueryURL: os.Getenv("SANDBOX_API_QUERY_URL"),
	}
)
