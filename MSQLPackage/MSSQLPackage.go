package MSQLPackage

import (
"database/sql"
"fmt"
"strings"
 _"github.com/denisenkom/go-mssqldb"
)

func FireQuery(username string, password string, host string, port string, dbname string, query string) (string,error){
	dsn := "server=" + host + ";user id=" + username + ";password=" + password + ";port="+port+";database=" + dbname 									//constructing the URL
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())														//Cannot connect to DB
		return "Cannot connect: ", err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())												//Cannot Connect to DB
		return "Cannot connect: ", err
	}
	defer db.Close()
	var result string
 err,result = exec(db, query)																									//Calling the execute function
		if err != nil {
			fmt.Println(err)
			return "Some error ",err
		}
		return result,nil
}

func exec(db *sql.DB, cmd string) (err error,result string) {
	var op string
rows, err := db.Query(cmd)																										//Executing the query
	if err != nil {
		return err,"Erro executing query"
	}
	defer rows.Close()
	cols, err := rows.Columns()																										//Getting the columns
	if err != nil {
		return err,""
	}
	if cols == nil {
		return nil,"No Columns"
	}

	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {																			//making an interface to store the row values
		vals[i] = new(interface{})
		if i != 0 {
			fmt.Print("\t")
		}

	}

		var currRow string;
		var str string;
		op=`{ "rows": [`
		var currEle string
	for rows.Next() {																									//iterating through the rows

		currRow="";
		err = rows.Scan(vals...)																					//getting one row
		if err != nil {
			fmt.Println(err)
			fmt.Println("Hey there")
			continue
		}
		for i := 0; i < len(vals); i++ 		{																//iterating through columns
			if i != 0 {
				fmt.Print("\t")
			}
				v:=(vals[i].(*interface{}))
				str=fmt.Sprintf("%v",(*v))
				currEle=`{"column":{"name":"`+cols[i]+`",`+`"value":"`+ str +`"}}`
			currRow=currRow+currEle+","
		}
		currRow=strings.TrimSuffix(currRow,`,`)
		op=op+`{"row":[`+currRow+`]},`

}
	if rows.Err() != nil {
		return rows.Err(),"Some error"
	}
	op=strings.TrimSuffix(op,`,`)
	op=op+`]}`

	return nil,op
}
