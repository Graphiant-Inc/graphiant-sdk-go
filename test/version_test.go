/*
Graphiant Go SDK

Testing Version constant

*/

package graphiant_sdk

import (
	"testing"

	openapiclient "github.com/Graphiant-Inc/graphiant-sdk-go"
)

func TestVersion(t *testing.T) {
	expectedVersion := "v25.12.1"
	if openapiclient.Version != expectedVersion {
		t.Errorf("Expected version %s, got %s", expectedVersion, openapiclient.Version)
	}
}
