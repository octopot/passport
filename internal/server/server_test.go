//go:generate echo $PWD - $GOPACKAGE - $GOFILE
//go:generate mockgen -package server_test -destination mock_contract_test.go go.octolab.org/ecosystem/passport/internal/server Service
package server_test
