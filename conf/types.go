package conf

type Configuration interface {
	Logging() logging
	Db() db
}

type Db interface {
	ConnString() string
}

type config struct {
	Name    string
	Version string
	logging logging
	db      db
}

type configData struct {
	Name       string  `json:"name"`
	Version    string  `json:"version"`
	LoggingDef logging `json:"logging"`
	DbDef      db      `json:"db"`
}

type logging struct {
	Level          string `json:"level"`
	File           string `json:"file"`
	ConsoleEnabled bool   `json:"consoleEnabled"`
}

type db struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}
