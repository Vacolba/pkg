package model

// Request Auth Headers
const (
	HeaderToken  = "token"
	HeaderBearer = "BEARER"
	HeaderAuth   = "Authorization"
)

// Config is the config format for the main application.
type Config struct {
	Storage Storage `json:"storage"`
	Web     Web     `json:"web"`
	GRPC    GRPC    `json:"grpc"`
	Logger  Logger  `json:"logger"`
	OpenID  OpenID  `json:"auth"`
}

// Storage holds app's storage configuration.
type Storage struct {
	DB      DBStorage    `json:"db"`
	Session RedisStorage `json:"session"`
	Cache   RedisStorage `json:"cache"`
}

// DBStorage settings for SQL Databse connection
type DBStorage struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Key      string `json:"key"`
}

// Web is the config format for the HTTP server.
type Web struct {
	HTTP           string   `json:"http"`
	HTTPS          string   `json:"https"`
	TLSCert        string   `json:"tlsCert"`
	WS             string   `json:"ws"`
	TLSKey         string   `json:"tlsKey"`
	AllowedOrigins []string `json:"allowedOrigins"`
}

// GRPC is the config for the gRPC API.
type GRPC struct {
	Addr        string `json:"addr"`
	TLSCert     string `json:"tlsCert"`
	TLSKey      string `json:"tlsKey"`
	TLSClientCA string `json:"tlsClientCA"`
}

// OpenID Connector settings
type OpenID struct {
	ID       string `json:"id"`
	Secret   string `json:"secret"`
	Redirect string `json:"redirect"`
	Discover string `json:"discover"`
}

// Logger holds configuration required to customize logging for vreports.
type Logger struct {
	Level  string `json:"level"`  // Level sets logging level severity.
	Format string `json:"format"` // Format specifies the format to be used for logging.
}
