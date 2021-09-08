package etc

type Configuration struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
	Db      db     `toml:"db"`
	Web     web    `toml:"web"`
	Mq      mq     `toml:"mq"`
}

type db struct {
	Host     string `toml:"host"`
	Database string `toml:"database"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Ssl		 string	`toml:"ssl"`
}

type web struct {
	Listen string `toml:"listen"`
}

type mq struct {
	MQConn              string
	MQTopicLoraToInflux string
}
