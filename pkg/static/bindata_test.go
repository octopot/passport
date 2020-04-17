//go:generate echo $PWD - $GOPACKAGE - $GOFILE
package static_test

import (
	"io/ioutil"
	"path/filepath"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kamilsk/passport/pkg/static"
)

func TestAsset(t *testing.T) {
	tests := []struct {
		name   string
		asset  string
		golden string
	}{
		{"common migration", "static/migrations/1_common.sql", "./migrations/1_common.sql"},
		{"account migration", "static/migrations/2_account.sql", "./migrations/2_account.sql"},
		{"domain migration", "static/migrations/3_domain.sql", "./migrations/3_domain.sql"},
		{"client script", "static/scripts/passport.min.js", "./scripts/passport.min.js"},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			expected, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			obtained, err := static.Asset(filepath.Join(tc.asset))
			assert.NoError(t, err)
			assert.Equal(t, expected, obtained)
		})
	}
}

func TestMustAsset(t *testing.T) {
}

func TestAssetInfo(t *testing.T) {
}

func TestAssetNames(t *testing.T) {
}

func TestAssetDir(t *testing.T) {
	tests := []struct {
		name     string
		assetDir string
		expected []string
	}{
		{"root", "static", []string{"migrations", "scripts"}},
		{"migrations", "static/migrations", []string{"1_common.sql", "2_account.sql", "3_domain.sql"}},
		{"scripts", "static/scripts", []string{"passport.min.js"}},
		{"not found", "static/migrations/unknown", nil},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			files, err := static.AssetDir(tc.assetDir)
			sort.Strings(tc.expected)
			sort.Strings(files)
			assert.Equal(t, tc.expected, files)
			if len(files) == 0 {
				assert.Error(t, err)
			}
		})
	}
}

func TestRestoreAsset(t *testing.T) {
}

func TestRestoreAssets(t *testing.T) {
}
