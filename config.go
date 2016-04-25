package intacct

// Config hold the required and optional credentials for a query
type Config struct {
	Sender         string
	SenderPassword string
	User           string
	UserPassword   string
	Company        string
	Location       string // Optional
}

// IsValid returns true if sender, sender password, user, user password, and
// company are all set
func (config Config) IsValid() bool {
	return config.Sender != "" && config.SenderPassword != "" && config.User != "" && config.UserPassword != "" && config.Company != ""
}

// TODO Parse a JSON file
func Parse() (config Config) {
	return
}

// TODO Create a config from ENV variables
func Env() Config {
	return Config{}
}
