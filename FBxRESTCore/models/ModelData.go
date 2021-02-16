package models

import (
	
	"database/sql"
	"null"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	_struct "fbrest/FBxRESTCore/struct"
)



type ModelGetData struct {
	DB *sql.DB
}

type Db_init struct {
	DB *sql.DB
}


func (model ModelGetData) GetSQLData(cmd string) (getStruct string, err error) {
	const funcstr = "func models.GetSQLData"
    log.WithFields(log.Fields{"Getting data": cmd,	}).Debug(funcstr)
	
	row, err := model.DB.Query(cmd)
	if err != nil {				
		log.WithFields(log.Fields{"Error": err.Error(),	}).Error(funcstr+"->query command")
		//return  nil,err
		return  "",err
	} else {
		
		colNames, err := row.Columns()
		ct, err := row.ColumnTypes()
		//dtt := ct[0].DatabaseTypeName()
		//log.Debug(ct)
		if err != nil {
			log.WithFields(log.Fields{"Error": err.Error(),	}).Error(funcstr+"->get colnames")
		}
		
		readCols := make([]interface{}, len(colNames))
		writeCols := make([]null.String, len(colNames))
		var _isiStructAll [][] _struct.SqlResponseData
		//var _isiStruct [] string

		for i, _ := range writeCols {
			readCols[i] = &writeCols[i]
		}
		
		for row.Next() {
			_isiStruct2 := make([]_struct.SqlResponseData, len(colNames))
			err := row.Scan(readCols...)			
			if err != nil {
				log.WithFields(log.Fields{"Error": err.Error(),	}).Error(funcstr+"->scan next row")
			} else {	
				
				for i, _ := range ct {
					_isiStruct2[i].Colvalue  = writeCols[i].String
					_isiStruct2[i].Colname   = ct[i].Name()
					_isiStruct2[i].Datatype  = ct[i].DatabaseTypeName()										
					_isiStruct2[i].IsNull    = !writeCols[i].Valid							
				}

				//log.Debug(_isiStruct2)
				//pagesJson, err := json.Marshal(writeCols)
				//if err != nil {
				//	log.WithFields(log.Fields{"Error": err.Error(),	}).Error(funcstr+"->marshal to JSON")
				//}
				//_isiStruct = append(_isiStruct,string(pagesJson))
				_isiStructAll = append(_isiStructAll,_isiStruct2)
			}
		
	  	    
		}
		log.WithFields(log.Fields{"Getting data done": cmd,	}).Debug(funcstr)
		    var erg [] string
			var ergstr string = "["
			for n, data := range _isiStructAll {
				if(n == 0){
					ergstr = ergstr + "{"
				} else {
					ergstr = ergstr + ",{"
				}
				var sep string
		        for i, _ := range data {
					if(data[i].Datatype == "VARYING"){
						sep = "\""
				    } else if(data[i].Datatype == "DATE"){
						sep = "\""
					} else if(data[i].Datatype == "TIMESTAMP"){
						sep = "\""
					} else if(data[i].Datatype == "TIME"){
						sep = "\""
					} else {
						sep = ""
					}
					if(i == 0){
						ergstr = ergstr + "\""+data[i].Colname + "\":"+sep+data[i].Colvalue+sep	
					} else {
						ergstr = ergstr + ",\""+data[i].Colname + "\":"+sep+data[i].Colvalue+sep
					}
			        
		        }
				ergstr = ergstr + "}"
				erg = append(erg,ergstr)
		    }
			ergstr = ergstr + "]"
	    
		return ergstr, nil
		//return _isiStruct, nil
	}
}

func (model ModelGetData) GetSQLDataOld(cmd string) (getStruct [] string, err error) {
	const funcstr = "func models.GetSQLData"
    log.WithFields(log.Fields{"Getting data": cmd,	}).Debug(funcstr)
	
	row, err := model.DB.Query(cmd)
	if err != nil {				
		log.WithFields(log.Fields{"Error": err.Error(),	}).Error(funcstr+"->query command")
		return  nil,err
	} else {
		
		colNames, err := row.Columns()
		
		if err != nil {
			log.WithFields(log.Fields{"Error": err.Error(),	}).Error(funcstr+"->get colnames")
		}
		
		readCols := make([]interface{}, len(colNames))
		writeCols := make([]null.String, len(colNames))
		var _isiStruct [] string

		for i, _ := range writeCols {
			readCols[i] = &writeCols[i]
		}
		
		for row.Next() {
			err := row.Scan(readCols...)			
			if err != nil {
				log.WithFields(log.Fields{"Error": err.Error(),	}).Error(funcstr+"->scan next row")
			} else {						
				pagesJson, err := json.Marshal(writeCols)
				if err != nil {
					log.WithFields(log.Fields{"Error": err.Error(),	}).Error(funcstr+"->marshal to JSON")
				}
				_isiStruct = append(_isiStruct,string(pagesJson))
			}
		}
		log.WithFields(log.Fields{"Getting data done": cmd,	}).Debug(funcstr)
		return _isiStruct, nil
	}
}

func (model ModelGetData) RunSQLData(cmd string) (getStruct [] string, err error) {
	const funcstr = "func models.RunSQLData"
    log.WithFields(log.Fields{"Getting data": cmd,	}).Debug(funcstr)
	
	row, err := model.DB.Query(cmd)
	if err != nil {				
		log.WithFields(log.Fields{"Error": err.Error(),	}).Error(funcstr+"->query command")
		return  nil,err
	} else {
		
		colNames, err := row.Columns()
		
		if err != nil {
				log.WithFields(log.Fields{"Error": err.Error(),	}).Error(funcstr+"->get colnames")
			} else {
		
		}
		
		readCols := make([]interface{}, len(colNames))
		writeCols := make([]null.String, len(colNames))
		var _isiStruct [] string

		for i, _ := range writeCols {
			readCols[i] = &writeCols[i]
		}
		
		for row.Next() {
			err := row.Scan(readCols...)			
			if err != nil {
				log.WithFields(log.Fields{"Error": err.Error(),	}).Error(funcstr+"->scan next row")
			} else {						
				pagesJson, err := json.Marshal(writeCols)
				if err != nil {
					log.WithFields(log.Fields{"Error": err.Error(),	}).Error(funcstr+"->marshal to JSON")
				}
				_isiStruct = append(_isiStruct,string(pagesJson))
			}
		}
		log.WithFields(log.Fields{"Getting data done": cmd,	}).Debug(funcstr)
		return _isiStruct, nil
	}
}
