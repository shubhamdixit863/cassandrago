package main

import (
	"github.com/shubhamdixit863/cassandrago/db"
)

func main() {

	args := db.Args{
		Host:     "cassandra-40b13fa-shubhamdixit863-a24d.aivencloud.com",
		Port:     14287,
		Username: "avnadmin",
		Password: "AVNS_PCCW8mCga8uqr2G",
		CaPath:   "./ca.pem",
	}
	db.CassandraConnect(args)

}
