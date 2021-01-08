package test

import (
	"database/sql"
	"fmt"
	_ "github.com/zhaobingss/go-sqlcipher/v4"
	"log"
	"testing"
)

func Test01(t *testing.T)  {

	dbname := "./foo.db?" +
		"_pragma_key=123456" +
		"&_pragma_cipher_page_size=1024" +
		"&_pragma_kdf_iter=64000"+
		"&_pragma_cipher_kdf_algorithm=PBKDF2_HMAC_SHA1" +
		"&_pragma_cipher_hmac_algorithm=HMAC_SHA1" +
		"&_pragma_cipher_use_hmac=OFF"

	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(id INTEGER, name TEXT);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO users(id, name) VALUES (1, '张三');`)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT * FROM users;`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rows)

}