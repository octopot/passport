//go:generate echo $PWD - $GOPACKAGE - $GOFILE
//go:generate mockgen -package service_test -destination mock_contract_test.go go.octolab.org/ecosystem/passport/internal/service Storage
package service_test
