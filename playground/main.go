package main

func main() {
	/*	host, empty := os.LookupEnv("DB_HOST")
		if empty == false {
			fmt.Println("Error retrieving DB_HOST:is empty")
			panic("Failed to retrieve DB_HOST environment variable")
		}
		port, empty := os.LookupEnv("DB_PORT")
		if empty == false {
			fmt.Println("Error retrieving DB_PORT:is empty")
			panic("Failed to retrieve DB_PORT environment variable")
		}
		user, empty := os.LookupEnv("DB_USER")
		if empty == false {
			fmt.Println("Error retrieving DB_USER:is empty")
			panic("Failed to retrieve DB_USER environment variable")
		}
		password, empty := os.LookupEnv("DB_PASSWORD")
		if empty == false {
			fmt.Println("Error retrieving DB_PASSWORD:is empty")
			panic("Failed to retrieve DB_PASSWORD environment variable")
		}
		dbName, empty := os.LookupEnv("DB_DATABASE")
		if empty == false {
			fmt.Println("Error retrieving DB_DATABASE:is empty")
			panic("Failed to retrieve DB_DATABASE environment variable")
		}

		fmt.Printf("Host: %s\nUser: %s\nPassword: %s\nDatabase: %s\n", host, user, password, dbName)

		pgconn := pgdriver.NewConnector(
			pgdriver.WithNetwork("tcp"),
			pgdriver.WithAddr(host+":"+port),
			pgdriver.WithUser(user),
			pgdriver.WithPassword(password),
			pgdriver.WithDatabase(dbName),
			//pgdriver.WithTLSConfig(),
			pgdriver.WithTimeout(5*time.Second),
		)
		//var conn, err = bun.Conn{os.Getenv("DB_DATABASE"), pgconn}
	*/
}
