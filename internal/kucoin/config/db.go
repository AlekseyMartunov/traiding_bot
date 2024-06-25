package config

type DB struct {
	Pass   string `yaml:"password"`
	User   string `yaml:"username"`
	DBName string `yaml:"db_name"`
	Addr   string `yaml:"addr"`
}

func (db *DB) GetPassword() string {
	return db.Pass
}

func (db *DB) GetUser() string {
	return db.User
}

func (db *DB) GetDBName() string {
	return db.DBName
}

func (db *DB) GetDBAddr() string {
	return db.Addr
}
