//go:generate echo $PWD - $GOPACKAGE - $GOFILE
//go:generate mockgen -package server_test -destination mock_contract_test.go github.com/kamilsk/passport/pkg/server Service
package server_test
