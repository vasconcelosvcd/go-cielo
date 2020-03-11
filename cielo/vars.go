package cielo

var (
	// ProductionEnvironment sets the environment to production
	ProductionEnvironment = Environment{
		APIUrl:      "https://api.cieloecommerce.cielo.com.br",
		APIQueryURL: "https://apiquery.cieloecommerce.cielo.com.br",
	}
	// SandboxEnvironment sets the environment to sandbox
	SandboxEnvironment = Environment{
		APIUrl:      "https://apisandbox.cieloecommerce.cielo.com.br",
		APIQueryURL: "https://apiquerysandbox.cieloecommerce.cielo.com.br",
	}
)
