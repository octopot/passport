//go:generate echo $PWD - $GOPACKAGE - $GOFILE
//go:generate mockgen -package service_test -destination mock_contract_test.go github.com/kamilsk/passport/pkg/service Storage
package service_test
