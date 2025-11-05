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
		created_at TEXT DEFAULT CURRENT_TIMESTAMP,
        updated_at TEXT DEFAULT CURRENT_TIMESTAMP
    );

	CREATE TRIGGER IF NOT EXISTS delete_old_posts 
        AFTER INSERT ON Post
        BEGIN
            DELETE FROM Post 
            WHERE created_at < DATE('now', '-90 days') AND categories LIKE '%news%';
        END;

    CREATE TRIGGER IF NOT EXISTS update_timestamp_trigger
        AFTER UPDATE ON Post
        FOR EACH ROW
        BEGIN
            UPDATE Post
            SET updated_at = CURRENT_TIMESTAMP
    		WHERE id = NEW.id;
        END;
	`

	_, err_post := mydb.Exec(sql_post)
	if err_post != nil {
		panic(err_post)
	}

}
