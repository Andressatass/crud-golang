package configuration

type Configuration struct {
	Database Database
}

type Database struct {
	Port     string
	Host     int
	User     string
	Password string
	Name     string
}
