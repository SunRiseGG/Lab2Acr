package scenarios

import (
	"database/sql"
	"fmt"
)

type Account struct {
  Id int64 `json:"id"`
  Balance float64 `json:"balance"`
  LastOperationTime string `json:"lastoperationtime"`
}

type TransactionHelper struct {
  GiverId int64
  ReceiverId int64
  Sum float64
}

type DataBase struct {
  Db *sql.DB
}

func NewDataBase(db *sql.DB) *DataBase {
  return &DataBase{Db: db}
}

func (db *DataBase) AccountsList() ([]*Account, error) {
  rows, err := db.Db.Query("SELECT id, balance, lastoperationtime FROM banking")
  if err != nil {
    return nil, err
  }
  defer rows.Close()
  var result []*Account
  for rows.Next() {
    var acc Account
    err := rows.Scan(&acc.Id, &acc.Balance, &acc.LastOperationTime)
    if err != nil {
      return nil, err
    }
    result = append(result, &acc)
  }
  if result == nil {
    result = make([]*Account, 0)
  }
  return result, nil
}

func (db *DataBase) Transaction(tr *TransactionHelper) (bool, error) {
  giverBalance, err1 := db.Db.Query("SELECT balance FROM banking WHERE id = $1", &tr.GiverId)
  receiverBalance, err2 := db.Db.Query("SELECT balance FROM banking WHERE id = $1", &tr.ReceiverId)
  if err1 != nil {
    return false, err1
  }
	defer giverBalance.Close()
	if err2 != nil {
    return false, err2
  }
	defer receiverBalance.Close()
  var balance1, balance2 float64
  giverBalance.Next()
  giverBalance.Scan(&balance1)
  receiverBalance.Next()
  receiverBalance.Scan(&balance2)
  if balance1 < tr.Sum {
    return false, fmt.Errorf("Specified amount is unavailable!")
  }
  balance1 -= tr.Sum
  balance2 += tr.Sum
  _, execErr1 := db.Db.Exec("UPDATE banking SET balance = $1 WHERE id = $2", balance1, &tr.GiverId)
  _, execErr2 := db.Db.Exec("UPDATE banking SET balance = $1 WHERE id = $2", balance2, &tr.ReceiverId)
  if execErr1 != nil {
    return false, execErr1
  } else if execErr2 != nil {
    return false, execErr2
  }
  return true, nil
}
