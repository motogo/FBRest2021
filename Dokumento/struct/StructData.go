package _struct

import (
	"null"
)

type StructData struct {
 ID null.String `json:”ID”`
 BEZ null.String `json:”BEZ”`
 SCHLUESSEL null.String `json:”SCHLUESSEL”`
}

type StructVANFORDERUNGALLEData struct {
 ID string `json:”ID”`
 BEZ string `json:”BEZ”`
 BESCHREIBUNG string `json:”BESCHREIBUNG”`
}


type Location struct {
 ID string `json:"ID"`
 BEZ string `json:"BEZ"`
 GUELTIG int64 `json:"GUELTIG"`
}

type Response_weight struct {
 ID string `json:"ID"`
 BEZ string `json:"BEZ"`
 GUELTIG int64 `json:"GUELTIG"`
}

type SQLAttributes struct {
 TABLE string `json:"TABLE"`
 FILTER string `json:"FILTER"`
 CMD string  `json:"CMD"`
 LOCATION string  `json:"LOCATION"`
 DATABASE string  `json:"DATABASE"`
 PORT int  `json:"PORT"`
 INFO string `json:"INFO"`
}

type StructKey struct {
 ID string `json:”ID”`
 BEZ string `json:”BEZ”`	
}

