package db

func CreateSchemas() {
	mydb := Connect()
	//defer mydb.Close()
	sql_user := `CREATE TABLE IF NOT EXISTS User (
        id TEXT PRIMARY KEY,
        title TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
		role TEXT NOT NULL,
		thumb TEXT,
		content TEXT,
		date TEXT NOT NULL
    );`

	_, err_user := mydb.Exec(sql_user)
	if err_user != nil {
		panic(err_user)
	}

}
