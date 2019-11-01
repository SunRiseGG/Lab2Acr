package scenarios

import (
	"encoding/json"
	"github.com/SunRiseGG/lab2/server/tools"
	"log"
	"net/http"
)

type HttpHandlerFunc http.HandlerFunc

func HttpHandler(db *DataBase) HttpHandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			handleAccountsList(db, res)
		} else if req.Method == "POST" {
			handleTransaction(req, res, db)
		} else {
			res.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleTransaction(req *http.Request, res http.ResponseWriter, db *DataBase) {
	var tr TransactionHelper
	if err := json.NewDecoder(req.Body).Decode(&tr); err != nil {
		log.Printf("Error decoding scenario input: %s", err)
		tools.WriteJsonBadRequest(res, "bad JSON payload")
		return
	}
	_, err := db.Transaction(&tr)
	if err == nil {
		tools.WriteJsonOk(res, &tr)
	} else {
		log.Printf("Error updating values: %s", err)
		tools.WriteJsonInternalError(res)
	}
}

func handleAccountsList(db *DataBase, res http.ResponseWriter) {
	result, err := db.AccountsList()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(res)
		return
	}
	tools.WriteJsonOk(res, result)
}
