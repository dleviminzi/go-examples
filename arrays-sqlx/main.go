package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	. "github.com/snowflakedb/gosnowflake"
)

func main() {
	var db *sqlx.DB
	db, err := sqlx.Open("snowflake", "secret")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	recordIds := []string{"recordId1", "recordId2", "opm"}
	recordIds = append(recordIds, "recordId3")
	convIds := []string{"convId1", "convId2", "convId3", "yufg"}

	// NOTE: it is very important to terminate the statement. if you don't, the number of rows
	// returned from this kind of query are unpredictable.
	r, err := db.Queryx("select * from (values (?, ?));", Array(&recordIds), Array(&convIds))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	for r.Next() {
		var rId *string
		var cId *string
		err = r.Scan(&rId, &cId)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(*rId, *cId)
		fmt.Println("/////////--/////////")
		// fmt.Println(r.Columns())
	}
}
