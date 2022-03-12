package main

func init() {
	conf := config{
		server: server{
			port: ":1818",
		},
		db: db{
			dbname:          "test",
			user:            "postgres",
			password:        "123",
			host:            "srv.baltpolymer.com",
			port:            "55432",
			sslmode:         "disable",
			connect_timeout: 10,
		},
	}
}

type config struct {
	server
	db
}

type server struct {
	port string
}

type db struct {
	dbname, user, password, host, port, sslmode string
	connect_timeout                             uint8
}
