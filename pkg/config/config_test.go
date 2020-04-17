package config_test

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.octolab.org/sequence"
	"gopkg.in/yaml.v2"

	. "github.com/kamilsk/passport/pkg/config"
)

var update = flag.Bool("update", false, "update .golden files")

func TestApplicationConfig_Dump(t *testing.T) {
	testCases := []struct {
		name    string
		in      string
		out     string
		marshal func(interface{}) ([]byte, error)
	}{
		{"JSON dump", "fixtures/config.yml", "fixtures/dump.json.golden", json.Marshal},
		{"YAML dump", "fixtures/config.yml", "fixtures/dump.yml.golden", yaml.Marshal},
	}

	for _, test := range testCases {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			raw, err := ioutil.ReadFile(tc.in)
			assert.NoError(t, err)

			var cnf ApplicationConfig
			err = yaml.UnmarshalStrict(raw, &cnf)
			assert.NoError(t, err)

			actual, err := tc.marshal(cnf)
			assert.NoError(t, err)

			if *update {
				err = ioutil.WriteFile(tc.out, actual, os.ModePerm)
				assert.NoError(t, err)
			}
			expected, err := ioutil.ReadFile(tc.out)
			assert.NoError(t, err)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestDatabaseConfig_DriverName(t *testing.T) {
	config, wg := DatabaseConfig{DSN: "postgres://postgres:postgres@127.0.0.1:5432/postgres"}, sync.WaitGroup{}
	for range sequence.Simple(runtime.GOMAXPROCS(0) + 1) {
		wg.Add(1)
		go func() {
			assert.Equal(t, "postgres", config.DriverName())
			wg.Done()
		}()
	}
	wg.Wait()
}
