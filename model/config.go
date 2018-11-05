package model

// Request Auth Headers
const (
	HeaderToken  = "token"
	HeaderBearer = "BEARER"
	HeaderAuth   = "Authorization"
)

// Config is the config format for the main application.
type Config struct {
	Storage Storage `json:"storage" yaml:"storage"`
	Web     Web     `json:"web" yaml:"web"`
	GRPC    GRPC    `json:"grpc" yaml:"grpc"`
	Logger  Logger  `json:"logger" yaml:"logger"`
	OpenID  OpenID  `json:"auth" yaml:"auth"`
}

// Storage holds app's storage configuration.
type Storage struct {
	DB      DBStorage    `json:"db" yaml:"db"`
	Session RedisStorage `json:"session" yaml:"session"`
	Cache   RedisStorage `json:"cache" yaml:"cache"`
}

// DBStorage settings for SQL Databse connection
type DBStorage struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	User     string `json:"user yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
	Key      string `json:"key" yaml:"key"`
}

// RedisStorage settings to connect to Redis
type RedisStorage struct {
	Type     string `json:"type" yaml:"type"`
	Master   string `json:"master" yaml:"master"`
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Size     int    `json:"size" yaml:"size"`
	Database string `json:"database" yaml:"database"`
	HashKey  string `json:"key" yaml:"key"`
	BlockKey string `json:"blockkey" yaml:"blockkey"`
}

// Web is the config format for the HTTP server.
type Web struct {
	HTTP           string   `json:"http" yaml:"http"`
	HTTPS          string   `json:"https" yaml:"https"`
	TLSCert        string   `json:"tlsCert" yaml:"tlsCert"`
	WS             string   `json:"ws" yaml:"ws"`
	TLSKey         string   `json:"tlsKey" yaml:"tlsKey"`
	AllowedOrigins []string `json:"allowedOrigins" yaml:"allowedOrigins"`
}

// GRPC is the config for the gRPC API.
type GRPC struct {
	Addr        string `json:"addr" yaml:"addr"`
	TLSCert     string `json:"tlsCert" yaml:"tlsCert"`
	TLSKey      string `json:"tlsKey" yaml:"tlsKey"`
	TLSClientCA string `json:"tlsClientCA" yaml:"tlsClientCA"`
}

// OpenID Connector settings
type OpenID struct {
	ID       string `json:"id" yaml:"id"`
	Secret   string `json:"secret" yaml:"secret"`
	Redirect string `json:"redirect" yaml:"redirect"`
	Discover string `json:"discover" yaml:"discover"`
}

// Logger holds configuration required to customize logging for vreports.
type Logger struct {
	Level  string `json:"level" yaml:"level"`   // Level sets logging level severity.
	Format string `json:"format" yaml:"format"` // Format specifies the format to be used for logging.
}
