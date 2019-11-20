package config

import (
	"reflect"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	expectedConfig := Config{
		ServiceName:    serviceName,
		ServiceAddress: defaultAddress,
	}

	gotConfig := New()

	if !reflect.DeepEqual(expectedConfig, gotConfig) {
		t.Fatalf("\nexpected:\n %+v\nbut got:\n %+v", expectedConfig, gotConfig)
	}
}
