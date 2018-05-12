package mssqlpackage

import (
	"database/sql"
	"fmt"
	"strings"
	//Import to get the MSSQL driver
	_ "github.com/denisenkom/go-mssqldb"
)

//UpdateQuery is used to execute DML operations like insert, update, delete
func UpdateQuery(username string, password string, host string, port string, dbname string, query string) (string, error) {
	dsn := "server=" + host + ";user id=" + username + ";password=" + password + ";port=" + port + ";database=" + dbname //constructing the URL
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error()) //Cannot connect to DB
		return "Cannot connect: ", err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error()) //Cannot Connect to DB
		return "Cannot connect: ", err
	}
	defer closeconnection(db)

	rows, err := db.Exec(query)
	if err != nil {
		return "Error executing query", err
	}

	//  fmt.Println(rows.RowsAffected())
	num, err := rows.RowsAffected()
	if err != nil {
		return "Error", err
	}

	op := `{"numberOfRowsAffected":"` + fmt.Sprintf("%v", num) + `"}` //Getting number of Rows Affected
	return op, nil
}

func closeconnection(dbconn *sql.DB) {
	dbconn.Close()

}

//FireQuery is used to execute Select Queries
func FireQuery(username string, password string, host string, port string, dbname string, query string) (string, error) {
	dsn := "server=" + host + ";user id=" + username + ";password=" + password + ";port=" + port + ";database=" + dbname //constructing the URL
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		//fmt.Println("Cannot connect: ", err.Error())														//Cannot connect to DB
		return "Cannot connect: ", err
	}
	err = db.Ping()
	if err != nil {
		//fmt.Println("Cannot connect: ", err.Error())												//Cannot Connect to DB
		return "Cannot connect: ", err
	}
	defer closeconnection(db)
	var result string
	result, err = exec(db, query) //Calling the execute function
	if err != nil {
		//fmt.Println(err)
		return "Some error ", err
	}
	return result, nil
}

//CreateQuery is a function to execute the DDL Queries
func CreateQuery(username string, password string, host string, port string, dbname string, query string) (string, error) {
	dsn := "server=" + host + ";user id=" + username + ";password=" + password + ";port=" + port + ";database=" + dbname //constructing the URL
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error()) //Cannot connect to DB
		return "Cannot connect: ", err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error()) //Cannot Connect to DB
		return "Cannot connect: ", err
	}
	defer closeconnection(db)

	_, err = db.Exec(query)
	if err != nil {
		return "Error executing query", err
	}

	op := `"Query Status":"Operation Successful"` //Getting The row
	return op, nil

}

func exec(db *sql.DB, cmd string) (result string, err error) {
	var op string
	rows, err := db.Query(cmd) //Executing the query
	if err != nil {
		return "Erro executing query", err
	}
	defer rows.Close()
	cols, err := rows.Columns() //Getting the columns
	if err != nil {
		return "", err
	}
	if cols == nil {
		return "No Columns", nil
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
			if i != 0 {
				fmt.Print("\t")
			}
			v := (vals[i].(*interface{}))
			str = fmt.Sprintf("%v", (*v))
			currEle = `{"column":{"name":"` + cols[i] + `",` + `"value":"` + str + `"}}`
			currRow = currRow + currEle + ","
		}
		currRow = strings.TrimSuffix(currRow, `,`)
		op = op + `{"row":[` + currRow + `]},`

	}
	if rows.Err() != nil {
		return "Some error", rows.Err()
	}
	op = strings.TrimSuffix(op, `,`)
	op = op + `]}`

	return op, nil
}
