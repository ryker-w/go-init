package etc

// 全局环境变量配置。（按需添加）
type Configuration struct {
	Web   web   `toml:"web"`
	Db    db    `toml:"db"`
	Amqp  amqp  `toml:"amqp"`
	Mqtt  mqtt  `toml:"mqtt"`
	Redis redis `toml:"redis"`
	Token token `toml:"token"`
}

type web struct {
	Listen string `toml:"listen"`
}

type db struct {
	Host     string `toml:"host"`
	Database string `toml:"database"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Ssl      string `toml:"ssl"`
}

type amqp struct {
	Conn string `toml:"conn"`
}

type mqtt struct {
	Broker   string `toml:"broker"`
	UserName string `toml:"username"`
	Password string `toml:"password"`
}

type redis struct {
	Enable   bool   `toml:"enable"`
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
	// other
}

type token struct {
	Enable bool   `toml:"enable"`
	Issuer string `toml:"issuer"`
	Key    string `toml:"key"`
	Ttl    int    `toml:"ttl"`
}
