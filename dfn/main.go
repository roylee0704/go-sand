package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	//packagename-to-mysql-impl, instruction to connect
	db, err := sql.Open("mssql", "server=126.32.3.39;database=Metafile_Test;user id=SSTAuto;password=1qQA2wWS3eED;encrypt=disable")
	if err != nil {
		log.Println("open sql: ", err)
	}
	defer db.Close()

	var p []byte
	var dfn string
	dict := make(map[string][]byte)

	q := "select ts, DFN from RptDFN"
	rows, err := db.Query(q)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&p, &dfn)

		_, ok := dict[fmt.Sprintf("%s", p)]

		if !ok {
			dict[fmt.Sprintf("%s", p)] = p
		} else {
			fmt.Println("repeated")
		}

		fmt.Printf("%+v : %s   %s \n", p, p, dfn)
	}

	var keys []string

	for k := range dict {
		keys = append(keys, k)
	}
	for _, k := range keys {
		fmt.Printf("dict: %+v : %s \n", dict[k], k)
	}
}
