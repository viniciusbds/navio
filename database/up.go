package main

import (
	"database/sql"

	// Mysql Driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/viniciusbds/navio/constants"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS " + constants.DBname)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE DATABASE " + constants.DBname)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE USER IF NOT EXISTS '" + constants.DBuser + "'@'localhost' IDENTIFIED BY '" + constants.DBpass + "'")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("GRANT ALL PRIVILEGES ON " + constants.DBname + " . * TO '" + constants.DBuser + "'@'localhost';")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("FLUSH PRIVILEGES")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + constants.DBname)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS containers (" +
		"id varchar(30) NOT NULL," +
		"name varchar(30) NOT NULL," +
		"image varchar(30) NOT NULL," +
		"status varchar(10) NOT NULL," +
		"rootfs varchar(300) NOT NULL," +
		"command varchar(300) NOT NULL," +
		"params varchar(300)," +
		"cgroups varchar(300)," +
		"PRIMARY KEY (id) )  DEFAULT CHARSET=latin1")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS images (" +
		"id int(6) unsigned NOT NULL AUTO_INCREMENT," +
		"name varchar(30) NOT NULL," +
		"base varchar(30) NOT NULL," +
		"version varchar(10) NOT NULL," +
		"size float(10) NOT NULL," +
		"url varchar(130) NOT NULL," +
		"PRIMARY KEY (`id`) ) AUTO_INCREMENT=0 DEFAULT CHARSET=latin1")
	if err != nil {
		panic(err)
	}

	insForm, err := db.Prepare("INSERT INTO images(name, base, version, size, url) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = insForm.Exec("alpine", "alpine", "v3.11", 2.7, "http://dl-cdn.alpinelinux.org/alpine/v3.11/releases/x86_64/alpine-minirootfs-3.11.6-x86_64.tar.gz")
	if err != nil {
		panic(err.Error())
	}
	_, err = insForm.Exec("busybox", "busybox", "v4.0", 1.5, "https://raw.githubusercontent.com/teddyking/ns-process/4.0/assets/busybox.tar")
	if err != nil {
		panic(err.Error())
	}
	_, err = insForm.Exec("ubuntu", "ubuntu", "v20.04", 90.0, "http://cloud-images.ubuntu.com/minimal/releases/focal/release/ubuntu-20.04-minimal-cloudimg-amd64-root.tar.xz")
	if err != nil {
		panic(err.Error())
	}
}
