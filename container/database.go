package container

import (
	"database/sql"
	"strings"

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

func readContainersDB() {
	db := openDB()
	defer db.Close()

	selDB, err := db.Query("SELECT * FROM containers ORDER BY name DESC")
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, name, image, status, root, command, params string
		err := selDB.Scan(&id, &name, &image, &status, &root, &command, &params)
		if err != nil {
			panic(err.Error())
		}
		containers[id] = NewContainer(id, name, image, status, root, command, strings.Split(params, ","))
	}
}

func insertContainersDB(container *Container) {
	db := openDB()
	defer db.Close()

	params := ""
	if len(container.Params) > 0 {
		params = container.Params[0]
		for i, param := range container.Params {
			if i != 0 {
				params += "," + param
			}
		}
	}

	insForm, err := db.Prepare("INSERT INTO containers(id, name, image, status, root, command, params) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(container.ID, container.Name, container.Image, container.Status, container.Root, container.Command, params)
}

func updateContainerNameDB(ID, name string) error {
	db := openDB()
	defer db.Close()

	sqlStatement := `
	UPDATE containers
	SET name = ?
	WHERE id = ?;`

	insForm, err := db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	_, err = insForm.Exec(name, ID)
	return err
}

func updateContainerStatusDB(ID, status string) error {
	db := openDB()
	defer db.Close()

	sqlStatement := `
	UPDATE containers
	SET status = ?
	WHERE id = ?;`

	insForm, err := db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	_, err = insForm.Exec(status, ID)
	return err
}

func removeContainerDB(ID string) error {
	db := openDB()
	defer db.Close()

	delForm, err := db.Prepare("DELETE FROM containers WHERE id=?")
	if err != nil {
		return err
	}
	_, err = delForm.Exec(ID)
	return err
}
