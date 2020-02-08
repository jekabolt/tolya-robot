package server

type Server struct {
	Port string `env:"SERVER_PORT" envDefault:"8080"`
}
