package config_test

import (
	"testing"

	"github.com/SarunasBucius/tables-management/config"
)

func TestGetDbURI(t *testing.T) {
	testValues := []struct {
		Name          string
		FilePath      string
		ExpectedValue string
		HasError      bool
	}{
		{
			"URI with options",
			"testdata/conf_w_options_test.yaml",
			"mongodb+srv://user:password@host/?options",
			false,
		},
		{
			"URI without options",
			"testdata/conf_wo_options_test.yaml",
			"mongodb+srv://user:password@host",
			false,
		},
		{
			"File not found error",
			"random_file.yaml",
			"",
			true,
		},
		{
			"Missing values error",
			"testdata/empty.yaml",
			"",
			true,
		},
	}

	for _, tt := range testValues {
		got, err := config.GetDbURI(tt.FilePath)
		if err == nil && tt.HasError {
			t.Errorf("[%v] expected to get error but got none", tt.Name)
		}
		if err != nil && !tt.HasError {
			t.Errorf("[%v] did not expect error but got: %v", tt.Name, err)
			continue
		}

		if got != tt.ExpectedValue {
			t.Errorf("[%v] value expected: %v, got: %v", tt.Name, tt.ExpectedValue, got)
		}
	}
}
