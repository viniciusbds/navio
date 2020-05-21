package container

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

func readContainersDB() {
	db := openDBConn()
	defer db.Close()

	selDB, err := db.Query("SELECT * FROM containers ORDER BY name DESC")
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, name, imageID, status, root, command string
		err := selDB.Scan(&id, &name, &imageID, &status, &root, &command)
		if err != nil {
			panic(err.Error())
		}
		containers[name] = NewContainer(id, name, imageID, status, root, command)
	}
}

func insertContainersDB(container *Container) {
	db := openDBConn()
	defer db.Close()

	insForm, err := db.Prepare("INSERT INTO containers(id, name, imageID, status, root, command) VALUES(?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(container.ID, container.Name, container.ImageID, container.Status, container.Root, container.Command)
}

func removeContainerDB(name string) error {
	db := openDBConn()
	defer db.Close()

	delForm, err := db.Prepare("DELETE FROM containers WHERE name=?")
	if err != nil {
		return err
	}
	delForm.Exec(name)
	return nil
}
