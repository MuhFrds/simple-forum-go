package configs

type (
	Config struct {
		Service Service `mapstructure:"service"`
	}

	Service struct {
		Port string `mapstructure:"port" `
	}
)