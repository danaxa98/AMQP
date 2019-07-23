package pkg

import (
	"database/sql"
	"log"
	"time"
)

type repository struct {
	db			*sql.DB
}

func (rep *repository) insert(msg [] byte){
	tx, err := rep.db.Begin()
	checkError(err)

	_, err = tx.Exec("INSERT INTO amqp(message,time) VALUES(?,?)", string(msg), time.Now())
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	checkError(tx.Commit())
}


