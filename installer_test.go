package gowin32

import (
	"testing"
)

func TestGetInstalledProducts(t *testing.T) {
	products, err := GetInstalledProducts()
	if err != nil {
		t.Errorf("didn't expect an error, but got %v", err)
	}
	if len(products) == 0 {
		t.Errorf("expected to get some products, but didn't get any!")
	}

	for _, pc := range products {
		name, err := GetInstalledProductProperty(pc, InstallPropertyProductName)
		if err != nil {
			t.Errorf("didn't expect an error getting name, but got %v", err)
		}
		if name == "" {
			continue
		}
		version, err := GetInstalledProductProperty(pc, InstallPropertyVersionString)
		if err != nil {
			t.Errorf("didn't expect an error getting version for %s, but got %v", name, err)
		}
		if version == "" {
			t.Errorf("expected to get name and version, but got %s %s", name, version)
		}
	}
}
