package db

func CreateSchema() {
	mydb := Connect()

	sql := `CREATE TABLE IF NOT EXISTS User (
        id TEXT PRIMARY KEY,
        title TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
		role TEXT NOT NULL,
		thumb TEXT,
		content TEXT,
		date TEXT NOT NULL
    );`

	mydb.Exec(sql)
	defer mydb.Close()
}
