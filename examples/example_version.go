/*
Graphiant Go SDK Example

This example demonstrates how to access the SDK version.

*/

package main

import (
	"fmt"

	graphiant_sdk "github.com/Graphiant-Inc/graphiant-sdk-go"
)

func main() {
	fmt.Printf("Graphiant Go SDK Version: %s\n", graphiant_sdk.Version)
}
