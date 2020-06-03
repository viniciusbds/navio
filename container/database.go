package container

import (
	"database/sql"
	"strings"

	// Mysql Driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/viniciusbds/navio/utilities"
)

// Função openDBConn, abre a conexão com o banco de dados
func openDBConn() (db *sql.DB) {
	db, err := sql.Open("mysql", utilities.DBuser+":"+utilities.DBpass+"@/"+utilities.DBname)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func readContainersDB() {
	db := openDBConn()
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
		containers[name] = NewContainer(id, name, image, status, root, command, strings.Split(params, ","))
	}
}

func insertContainersDB(container *Container) {
	db := openDBConn()
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

func updateContainerDB(container *Container) {
	db := openDBConn()
	defer db.Close()

	// Is possiblie update only the name of container and the status
	sqlStatement := `
	UPDATE containers
	SET name = ?, status = ?
	WHERE id = ?;`

	insForm, err := db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(container.Name, container.Status, container.ID)
}

func removeContainerDB(ID string) error {
	db := openDBConn()
	defer db.Close()

	delForm, err := db.Prepare("DELETE FROM containers WHERE id=?")
	if err != nil {
		return err
	}
	delForm.Exec(ID)
	return nil
}
