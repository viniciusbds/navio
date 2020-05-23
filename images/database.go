package images

import (
	"database/sql"

	// Mysql Driver
	_ "github.com/go-sql-driver/mysql"
)

// Função openDBConn, abre a conexão com o banco de dados
func openDBConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "navio"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func readImagesDB() {
	db := openDBConn()
	defer db.Close()

	selDB, err := db.Query("SELECT * FROM images ORDER BY name DESC")
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var name, base, version, size, url string
		var id int
		err := selDB.Scan(&id, &name, &base, &version, &size, &url)
		if err != nil {
			panic(err.Error())
		}
		images[name] = NewImage(name, base, version, size, url)
	}
}

func insertImageDB(image *Image) {
	db := openDBConn()
	defer db.Close()

	insForm, err := db.Prepare("INSERT INTO images(name, base, version, size, url) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(image.Name, image.Base, image.Version, image.Size, image.URL)
}

func removeImageDB(imgName string) {
	db := openDBConn()
	defer db.Close()

	delForm, err := db.Prepare("DELETE FROM images WHERE name=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(imgName)
}
