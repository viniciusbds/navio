package main

import (
	"database/sql"

	// Mysql Driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/viniciusbds/navio/utilities"
)

func main() {

	db, err := sql.Open("mysql", utilities.DBuser+":"+utilities.DBpass+"@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS " + utilities.DBname)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE DATABASE " + utilities.DBname)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + utilities.DBname)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS containers (" +
		"id varchar(30) NOT NULL," +
		"name varchar(30) NOT NULL," +
		"image varchar(30) NOT NULL," +
		"status varchar(10) NOT NULL," +
		"root varchar(300) NOT NULL," +
		"command varchar(300) NOT NULL," +
		"params varchar(300)," +
		"PRIMARY KEY (id) )  DEFAULT CHARSET=latin1")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS images (" +
		"id int(6) unsigned NOT NULL AUTO_INCREMENT," +
		"name varchar(30) NOT NULL," +
		"base varchar(30) NOT NULL," +
		"version varchar(10) NOT NULL," +
		"size varchar(10) NOT NULL," +
		"url varchar(130) NOT NULL," +
		"PRIMARY KEY (`id`) ) AUTO_INCREMENT=0 DEFAULT CHARSET=latin1")
	if err != nil {
		panic(err)
	}

	insForm, err := db.Prepare("INSERT INTO images(name, base, version, size, url) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec("alpine", "alpine", "v3.11", "2.7M", "http://dl-cdn.alpinelinux.org/alpine/v3.11/releases/x86_64/alpine-minirootfs-3.11.6-x86_64.tar.gz")
	insForm.Exec("busybox", "busybox", "v4.0", "1.5M", "https://raw.githubusercontent.com/teddyking/ns-process/4.0/assets/busybox.tar")
	insForm.Exec("ubuntu", "ubuntu", "v20.04", "90.0M", "http://cloud-images.ubuntu.com/minimal/releases/focal/release/ubuntu-20.04-minimal-cloudimg-amd64-root.tar.xz")
}
