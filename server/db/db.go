package db

import (
	"database/sql"
	"net/url"
	_ "github.com/lib/pq"
)

type Connection struct {
  Host string
  User, Password string
  DbName string
  DisableSSL bool
}

func (c *Connection) ConnectionURL() string {
  dbUrl := &url.URL{
    Scheme: "postgres",
    User: url.UserPassword(c.User, c.Password),
    Host: c.Host,
    Path: c.DbName,
  }
  if c.DisableSSL {
    dbUrl.RawQuery = url.Values{
      "sslmode": []string{"disable"},
    }.Encode()
  }
  return dbUrl.String()
}

func (c *Connection) Open() (*sql.DB, error) {
  return sql.Open("postgres", c.ConnectionURL())
}
