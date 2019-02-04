// even_test project main.go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"httprouter-master"
	"io"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Record struct {
	Table_name string `json:"table_name"`
	Key        string `json:"key"`
	Value      string `json:"value"`
}

func main() {
	router := httprouter.New()
	router.PUT("/api/table", addTable)
	router.DELETE("/api/table/:name", deleteTable)
	router.PUT("/api/record", addRecord)
	router.DELETE("/api/record/:key", deleteRecord)
	router.GET("/api/record/:key", getRecord)
	http.ListenAndServe(":8080", router)

}

func Create_table(table_name string) error {

	db, err := sql.Open("sqlite3", "productdb.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, error := db.Exec("CREATE TABLE $1(key_str text NOT NULL UNIQUE, value text)", table_name)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	return (error)
}

func delete_table(table_name string) error {

	db, err := sql.Open("sqlite3", "productdb.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, error := db.Exec("DROP TABLE $1", table_name)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	return (error)
}

func add_date(table_name string, key string, value string) error {

	db, err := sql.Open("sqlite3", "productdb.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	crKey := []byte("example key 1234")

	value = encrypt(crKey, value)
	result, error := db.Exec("insert into $1 (key_str, value) values ($2, $3)", table_name, key, value)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	return (error)
}

func delete_data(table_name string, key string) error {

	db, err := sql.Open("sqlite3", "productdb.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, error := db.Exec("delete from $1 where key_str = $2", table_name, key)
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	return (error)
}

func get_date(table_name string, key string) (value string) {
	db, err := sql.Open("sqlite3", "productdb.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	row := db.QueryRow("SELECT first value FROM $1 WHERE key_str = $2", table_name, key)

	error := row.Scan(&value)
	if error != nil {
		panic(error)
	}

	return
}

func addTable(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var rec Record
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil {
		panic(err)
	}
	error := Create_table(rec.Table_name)
	if error != nil {
		panic(error)
	}
}

func deleteTable(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var rec Record
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil {
		panic(err)
	}
	error := delete_table(rec.Table_name)
	if error != nil {
		panic(error)
	}
}

func addRecord(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var rec Record
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil {
		panic(err)
	}
	error := add_date(rec.Table_name, rec.Key, rec.Value)
	if error != nil {
		panic(error)
	}
}

func deleteRecord(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var rec Record
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil {
		panic(err)
	}
	error := delete_data(rec.Table_name, rec.Key)
	if error != nil {
		panic(error)
	}
}

func getRecord(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var rec Record
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil {
		panic(err)
	}
	value := get_date(rec.Table_name, rec.Key)
	crKey := []byte("example key 1234")
	rec.Value = decrypt(crKey, value)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	error := json.NewEncoder(w).Encode(rec)
	if error != nil {
		panic(error)
	}
}

func encrypt(key []byte, text string) string {
	plaintext := []byte(text)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext)
}

func decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}
