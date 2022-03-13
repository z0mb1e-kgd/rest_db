package main

import "time"

const (
	serverDefaultPort    = "1818"
	serverMaxHeaderBytes = 1 << 20
	serverReadTimeout    = time.Second * 10
	serverWriteTimeout   = time.Second * 10

	dbDbname         = "test"
	dbUser           = "postgres"
	dbPassword       = "123"
	dbHost           = "srv.baltpolymer.com"
	dbPort           = "55432"
	dbSslmode        = "disable"
	dbConnectTimeout = 10
)
