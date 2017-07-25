package controllers

import (
  "app/app/models"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
  "github.com/revel/revel"
)

type GormController struct {
  *revel.Controller
  Tx *gorm.DB
}

var Db *gorm.DB

func InitDB() {
  var err error
  Db, err = gorm.Open("mysql", "homestead:secret@tcp(192.168.10.10:3306)/myapp?collation=utf8mb4_unicode_ci")

  if err != nil {
    revel.ERROR.Println("FATAL", err)
    panic(err)
  }

  // defer Db.Close()

  tab := &models.User{}
  Db.AutoMigrate(tab)
  Db.Model(tab).AddUniqueIndex("idx_user__gmail", "gmail")
  Db.Model(tab).AddUniqueIndex("idx_user__pu_mail", "pu_mail")

}
func (c *GormController) Begin() revel.Result {
  txn := Db.Begin()
  if txn.Error != nil {
    panic(txn.Error)
  }
  c.Tx = txn
  revel.INFO.Println("c.Tx init", c.Tx)
  return nil
}
func (c *GormController) Commit() revel.Result {
  if c.Tx == nil {
    return nil
  }
  c.Tx.Commit()
  if err := c.Tx.Error; err != nil && err != sql.ErrTxDone {
    panic(err)
  }
  c.Tx = nil
  revel.INFO.Println("c.Tx commited (nil)")
  return nil
}

func (c *GormController) Rollback() revel.Result {
  if c.Tx == nil {
    return nil
  }
  c.Tx.Rollback()
  if err := c.Tx.Error; err != nil && err != sql.ErrTxDone {
    panic(err)
  }
  c.Tx = nil
  return nil
}
