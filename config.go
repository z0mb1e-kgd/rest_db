package main

import "time"

const (
	serverPort = ":1818"

	dbName            = "test"
	dbUser            = "postgres"
	dbPassword        = "123"
	dbHost            = "srv.baltpolymer.com"
	dbPort            = "55432"
	dbSslmode         = "disable"
	dbConnectTimeout  = 10
	dbConnectAttempts = 5
	dbRetryPeriod     = time.Second * 5
)
