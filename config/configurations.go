package config

type Database struct {
	Host    string
	Port    int
	User    string
	DbName  string
	Pass    string
	SslMode string
}
type Jwt struct {
	SecretKey []byte
}

type Config struct {
	Database Database
	Jwt      Jwt
}

func LoadConfigs() Config {
	return Config{Database{"localhost", 5432, "postgres", "dqu_tdkt", "123456", "disable"},
		Jwt{[]byte("tdkt-dhqn")}}
}
