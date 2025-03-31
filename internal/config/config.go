package config

import ()

type Config struct {
	Server Server
}

type Server struct {
	Addres string
	Port string
	ConnNumber int
}