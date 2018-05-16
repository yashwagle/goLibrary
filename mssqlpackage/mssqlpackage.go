package mssqlpackage

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
	//Import to get the MSSQL driver
	_ "github.com/denisenkom/go-mssqldb"
)

//Create Context for timeout of query
func createContext(timeout int) context.Context {
	timestr := strconv.Itoa(timeout)
	timestr = timestr + `s`
	d, _ := time.ParseDuration(timestr)
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, d)
	return ctx
}

func dbconnect(conn string) (*sql.DB, error) { // connect to the database
	db, err := sql.Open("mssql", conn)
	if err != nil {
		//Cannot connect to DB
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		//Cannot Connect to DB
		return nil, err
	}
	return db, nil
}

//UpdateQuery is used to execute DML operations like insert, update, delete
func UpdateQuery(username string, password string, host string, port string, dbname string, query string, timeout int) (string, error) {
	dsn := "server=" + host + ";user id=" + username + ";password=" + password + ";port=" + port + ";database=" + dbname //constructing the URL
	db, err := dbconnect(dsn)
	if err != nil {
		return "Cannot Connect", err
	}
	defer closeconnection(db)

	ctx := createContext(timeout)
	rows, err := db.ExecContext(ctx, query)
	if err != nil {
		return "Error executing query", err
	}

	num, err := rows.RowsAffected() //Executing the query
	if err != nil {
		return "", err
	}

	op := `{"numberOfRowsAffected":"` + fmt.Sprintf("%v", num) + `"}` //Getting number of Rows Affected
	return op, nil
}

func closeconnection(dbconn *sql.DB) {
	dbconn.Close()

}

//FireQuery is used to execute Select Queries
func FireQuery(username string, password string, host string, port string, dbname string, query string, timeout int) (string, error) {
	dsn := "server=" + host + ";user id=" + username + ";password=" + password + ";port=" + port + ";database=" + dbname //constructing the URL
	db, err := dbconnect(dsn)
	if err != nil {
		return "", err
	}
	defer closeconnection(db)
	var result string
	result, err = exec(db, query, timeout) //Calling the execute function
	if err != nil {

		return "", err
	}
	return result, nil
}

//CreateQuery is a function to execute the DDL Queries
func CreateQuery(username string, password string, host string, port string, dbname string, query string, timeout int) (string, error) {
	dsn := "server=" + host + ";user id=" + username + ";password=" + password + ";port=" + port + ";database=" + dbname //constructing the URL
	db, err := dbconnect(dsn)
	if err != nil {
		return "", err
	}
	defer closeconnection(db)
	ctx := createContext(timeout)
	_, err = db.ExecContext(ctx, query) //Executing the DDL Query
	if err != nil {
		return "", err
	}

	op := `{"Query Status":"Operation Successful"}`
	return op, nil

}

func exec(db *sql.DB, cmd string, timeout int) (result string, err error) {
	var op string
	ctx := createContext(timeout)
	rows, err := db.QueryContext(ctx, cmd) //Executing the query
	if err != nil {
		return "", err
	}
	defer rows.Close()
	cols, err := rows.Columns() //Getting the column names
	if err != nil {
		return "", err
	}
	if cols == nil {
		return "", nil
	}

	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ { //making an interface to store the row values
		vals[i] = new(interface{})

	}

	var currRow string
	var str string
	op = `{ "rows": [`
	var currEle string
	for rows.Next() { //iterating through the rows

		currRow = ""
		err = rows.Scan(vals...) //getting one row
		if err != nil {
			continue
		}
		for i := 0; i < len(vals); i++ { //iterating through columns

			v := (vals[i].(*interface{}))
			str = fmt.Sprintf("%v", (*v))
			currEle = `{"column":{"name":"` + cols[i] + `",` + `"value":"` + str + `"}}`
			currRow = currRow + currEle + ","
		}
		currRow = strings.TrimSuffix(currRow, `,`)
		op = op + `{"row":[` + currRow + `]},`

	}
	if rows.Err() != nil {
		return "", rows.Err()
	}
	op = strings.TrimSuffix(op, `,`)
	op = op + `]}`

	return op, nil
}
