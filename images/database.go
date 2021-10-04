package images

import (
	"database/sql"

	// Mysql Driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/viniciusbds/navio/constants"
)

func openDB() (db *sql.DB) {
	db, err := sql.Open("mysql", constants.DBuser+":"+constants.DBpass+"@/"+constants.DBname)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func readImagesDB() {
	db := openDB()
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

func insertImageDB(image *Image) error {
	db := openDB()
	defer db.Close()

	insForm, err := db.Prepare("INSERT INTO images(name, base, version, size, url) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = insForm.Exec(image.Name, image.Base, image.Version, image.Size, image.URL)
	return err
}

func removeImageDB(imgName string) {
	db := openDB()
	defer db.Close()

	delForm, err := db.Prepare("DELETE FROM images WHERE name=?")
	if err != nil {
		panic(err.Error())
	}
	_, err = delForm.Exec(imgName)
	if err != nil {
		panic(err.Error())
	}

}
