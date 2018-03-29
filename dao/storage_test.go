//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package dao_test -destination $PWD/dao/mock_db_test.go database/sql/driver Conn,Driver,Stmt,Rows
package dao_test
