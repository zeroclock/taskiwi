package config

type Config struct {
	DSConfig *DatastoreConfig
}

type DatastoreConfig struct {
	// ProjectID string
}

func GetConfig() *Config {
	return &Config{
		DSConfig: &DatastoreConfig{
			// ProjectID: "taskiwi-dev",
		},
	}
}
