package main

import (
	"database/sql"
	"europe/handlers"
	"fmt"
	"net"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type country struct {
	name       string
	capital    string
	square     int
	population int
}

func main() {
	// build server
	listener, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	//build DB
	connStr := "user=mydb1 password=123456 dbname=mydb1 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("DB is connected!")
	//launch a loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn, db)
	}

}

func handleConnection(conn net.Conn, db *sql.DB) {
	defer conn.Close()
	for {
		//1 Server takes random country and sends to Client
		target := randomCountry(db)
		conn.Write([]byte(target.name))

		//4 Server recieves the answer, compares it and sends the message if it is right or not

		answer, err := handlers.Recieve(conn)
		if err != nil {
			return
		}
		if strings.EqualFold(answer, target.capital) {
			conn.Write([]byte("right "))
		} else {
			conn.Write([]byte(fmt.Sprint("wrong. ", target.capital, " is right ")))
		}
		time.Sleep(time.Millisecond)

	}
}

func randomCountry(db *sql.DB) country {
	rows, err := db.Query("select * from europe.general order by random() limit 1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	state := country{}

	for rows.Next() {

		err := rows.Scan(&state.name, &state.capital, &state.square, &state.population)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return state
}
