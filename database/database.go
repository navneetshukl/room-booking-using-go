package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)


func DB_Connect() (*sql.DB,error){

	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=Bookings_demo user=postgres password=password")
	if err != nil {
		log.Fatalf(fmt.Sprintf("Unable to connect: %v\n", err))
		return nil,err
	}

	//defer conn.Close()

	log.Println("Connected to database")

	err=conn.Ping()
	if err!=nil{
		log.Fatal("Cannot Ping the database")
		return nil,err
	}
	log.Println("pinged database")

	return conn,nil

}

func InsertRow(first,last,email,phone,start,end string){
	conn,err:=DB_Connect()
	defer conn.Close()
	if err!=nil{
		panic(err)
	}
	query:=`insert into users(first_name,last_name,phone_number,email,arrival,departure) values($1,$2,$3,$4,$5,$6)`

	_,err =conn.Exec(query,first,last,phone,email,start,end)

	if err!=nil{
		log.Fatal(err)
	}

	log.Println("Inserted a row")
}
