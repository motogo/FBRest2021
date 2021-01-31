package config

import ("database/sql"
        "strconv"
		"log"
		//"null"
        _"github.com/nakagami/firebirdsql"
)

func Conn() (db *sql.DB, err error) {
                   
    var connstr = "SYSDBA:masterkey@localhost:3050/D:/Data/DokuMents/DOKUMENTS30.FDB"
	log.Println("Open database:"+connstr) 
	db, err = sql.Open("firebirdsql",connstr )
	return
}

func ConnLocation(port int,location string, filename string) (db *sql.DB, err error) {
   
    if(port < 1){
	    port = 3050
	}
	if(len(location)) < 1{
	    location = "localhost"
	}
    var connstr = string("SYSDBA:masterkey@"+location+":"+strconv.Itoa(port)+"/"+filename)
  	
    log.Println("Open database:"+connstr) 
	db, err = sql.Open("firebirdsql", connstr) 
	if err != nil {
		log.Println("Error open database") 
	} else {
		log.Println("Success open database") 
	}
	return
}

func TestConnLocation(port int,location string, filename string) (err error) {
   
    if(port < 1){
	    port = 3050
	}
	if(len(location)) < 1{
	    location = "localhost"
	}
    var connstr = string("SYSDBA:masterkey@"+location+":"+strconv.Itoa(port)+"/"+filename)
	var connstr2 = string(location+":"+strconv.Itoa(port)+"/"+filename)
	log.Println("Open database:"+connstr2) 
	var db *sql.DB
	db, err = sql.Open("firebirdsql", connstr) 

	if err != nil {
		return
	}

	err = db.Ping(); 
	
	return
}

func ConnLocation2(connectionstring string) (db *sql.DB, err error) {

   
    log.Println("Open database:"+connectionstring) 
	db, err = sql.Open("firebirdsql", connectionstring)
	return
}


