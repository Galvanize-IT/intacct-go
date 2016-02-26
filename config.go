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

// TODO Parse a JSON file
func Parse() (config Config) {
	return
}

// TODO Create a config from ENV variables
func Env() Config {
	return Config{}
}
