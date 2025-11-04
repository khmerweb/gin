package db

func CreateSchema() {
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

	sql_post := `CREATE TABLE IF NOT EXISTS Post (
        id TEXT PRIMARY KEY,
        title TEXT NOT NULL,
        content TEXT,
		categories TEXT NOT NULL,
		thumb TEXT NOT NULL,
		date TEXT NOT NULL,
		videos TEXT,
		author TEXT NOT NULL,
		created_at TEXT
    );`

	_, err_post := mydb.Exec(sql_post)
	if err_post != nil {
		panic(err_post)
	}

}
