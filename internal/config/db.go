package config

type closeDB struct {
	Pass   string `yaml:"password"`
	User   string `yaml:"username"`
	DBName string `yaml:"db_name"`
	Addr   string `yaml:"addr"`
}

func (db *closeDB) toExternalStruct() DB {
	return DB{
		pass:   db.Pass,
		user:   db.User,
		dbName: db.DBName,
		addr:   db.Addr,
	}
}

type DB struct {
	pass   string
	user   string
	dbName string
	addr   string
}

func (db *DB) Password() string {
	return db.pass
}

func (db *DB) User() string {
	return db.user
}

func (db *DB) DBName() string {
	return db.dbName
}

func (db *DB) Addr() string {
	return db.addr
}
