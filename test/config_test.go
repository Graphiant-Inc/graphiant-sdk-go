package graphiant_sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	graphiant_sdk "github.com/Graphiant-Inc/graphiant-sdk-go"
)

func GetAuthToken() (*graphiant_sdk.APIClient, string) {

	username := os.Getenv("GRAPHIANT_USERNAME")
	password := os.Getenv("GRAPHIANT_PASSWORD")

	configuration := graphiant_sdk.NewConfiguration()
	apiClient := graphiant_sdk.NewAPIClient(configuration)
	authReq := graphiant_sdk.NewV1AuthLoginPostRequestWithDefaults()
	authReq.SetUsername(username)
	authReq.SetPassword(password)

	fmt.Printf("Calling V1AuthLoginPost (url: %s)\n", configuration.Servers[0].URL)
	_, httpRes, _ := apiClient.DefaultAPI.
		V1AuthLoginPost(context.Background()).
		V1AuthLoginPostRequest(*authReq).
		Execute()
	defer httpRes.Body.Close()

	var result struct {
		Auth        bool   `json:"auth"`
		AccountType string `json:"accountType"`
		Token       string `json:"token"`
	}

	if err := json.NewDecoder(httpRes.Body).Decode(&result); err != nil {
		fmt.Printf("Failed to decode auth response: %v\n", err)
	}
	fmt.Printf("Auth token: %v\n", result.Token)
	return apiClient, "Bearer " + result.Token
}

func UpdateDeviceStatus(deviceId int64, deviceStatus string) {
	apiClient, token := GetAuthToken()

	bringupReq := graphiant_sdk.NewV1DevicesBringupPutRequest()
	bringupReq.SetDeviceIds([]int64{deviceId})
	bringupReq.SetStatus(deviceStatus)
	fmt.Printf("Calling V1DevicesBringupPutRequest : %v %v",
		bringupReq.GetDeviceIds(), bringupReq.GetStatus())

	_, err := apiClient.DefaultAPI.
		V1DevicesBringupPut(context.Background()).
		Authorization(token).
		V1DevicesBringupPutRequest(*bringupReq).
		Execute()
	if err != nil {
		fmt.Printf("Failed to update device status: %v\n", err)
	}

}

func ConfigureWan(deviceId int64, interfaceName string, wanCircuit string) {
	apiClient, token := GetAuthToken()

	circuit := map[string]graphiant_sdk.ManaV2CircuitConfig{
		wanCircuit: {
			Name:              graphiant_sdk.PtrString(wanCircuit),
			ConnectionType:    graphiant_sdk.PtrString("internet_dia"),
			Label:             graphiant_sdk.PtrString("internet_dia_4"),
			DiaEnabled:        graphiant_sdk.PtrBool(true),
			QosProfile:        graphiant_sdk.PtrString("gold25"),
			LinkDownSpeedMbps: graphiant_sdk.PtrInt32(100),
			LinkUpSpeedMbps:   graphiant_sdk.PtrInt32(50),
		},
	}
	interfaces := map[string]graphiant_sdk.ManaV2NullableInterfaceConfig{
		interfaceName: {
			Interface: &graphiant_sdk.ManaV2InterfaceConfig{
				AdminStatus: graphiant_sdk.PtrBool(true),
				Circuit:     graphiant_sdk.PtrString(wanCircuit),
				Ipv4: &graphiant_sdk.ManaV2InterfaceIpConfig{
					Dhcp: &graphiant_sdk.ManaV2InterfaceDhcpConfig{
						DhcpClient: graphiant_sdk.PtrBool(true)},
				},
				Ipv6: &graphiant_sdk.ManaV2InterfaceIpConfig{
					Dhcp: &graphiant_sdk.ManaV2InterfaceDhcpConfig{
						DhcpClient: graphiant_sdk.PtrBool(true)},
				},
			},
		},
	}
	configReq := graphiant_sdk.NewV1DevicesDeviceIdConfigPutRequest()
	configReq.Edge = graphiant_sdk.NewManaV2EdgeDeviceConfig()
	configReq.Edge.SetInterfaces(interfaces)
	configReq.Edge.SetCircuits(circuit)
	circuitsJSON, _ := json.MarshalIndent(configReq.Edge.GetCircuits(), "", "  ")
	interfacesJSON, _ := json.MarshalIndent(configReq.Edge.GetInterfaces(), "", "  ")
	fmt.Printf("Calling V1DevicesDeviceIdConfigPut : %v %s %s", deviceId, circuitsJSON, interfacesJSON)
	graphiant_sdk.PollAndPutDeviceConfig(apiClient, token, deviceId, *configReq)
}

func ConfigureLan(deviceId int64, interfaceName string, lan string, ipv4Address string, ipv6Address string) {
	apiClient, token := GetAuthToken()

	interfaces := map[string]graphiant_sdk.ManaV2NullableInterfaceConfig{
		interfaceName: {
			Interface: &graphiant_sdk.ManaV2InterfaceConfig{
				Lan:         graphiant_sdk.PtrString(lan),
				AdminStatus: graphiant_sdk.PtrBool(true),
				Ipv4: &graphiant_sdk.ManaV2InterfaceIpConfig{
					Address: &graphiant_sdk.ManaV2NullableAddress{
						Address: graphiant_sdk.PtrString(ipv4Address),
					},
				},
				Ipv6: &graphiant_sdk.ManaV2InterfaceIpConfig{
					Address: &graphiant_sdk.ManaV2NullableAddress{
						Address: graphiant_sdk.PtrString(ipv6Address),
					},
				},
			},
		},
	}
	configReq := graphiant_sdk.NewV1DevicesDeviceIdConfigPutRequest()
	configReq.Edge = graphiant_sdk.NewManaV2EdgeDeviceConfig()
	configReq.Edge.SetInterfaces(interfaces)
	interfacesJSON, _ := json.MarshalIndent(configReq.Edge.GetInterfaces(), "", "  ")

	fmt.Printf("Calling V1DevicesDeviceIdConfigPut : %v %s\n", deviceId, interfacesJSON)
	graphiant_sdk.PollAndPutDeviceConfig(apiClient, token, deviceId, *configReq)
}

func ConfigureDefaultLan(deviceId int64, interfaceName string, defaultLan string) {
	apiClient, token := GetAuthToken()

	interfaces := map[string]graphiant_sdk.ManaV2NullableInterfaceConfig{
		interfaceName: {
			Interface: &graphiant_sdk.ManaV2InterfaceConfig{
				Lan: graphiant_sdk.PtrString(defaultLan),
			},
		},
	}
	configReq := graphiant_sdk.NewV1DevicesDeviceIdConfigPutRequest()
	configReq.Edge = graphiant_sdk.NewManaV2EdgeDeviceConfig()
	configReq.Edge.SetInterfaces(interfaces)
	interfacesJSON, _ := json.MarshalIndent(configReq.Edge.GetInterfaces(), "", "  ")

	fmt.Printf("Calling V1DevicesDeviceIdConfigPut : %v %s\n", deviceId, interfacesJSON)
	graphiant_sdk.PollAndPutDeviceConfig(apiClient, token, deviceId, *configReq)
}

func ConfigureSubInterface(deviceId int64, interfaceName string, lan string, vlan int32, ipv4Address string, ipv6Address string) {
	apiClient, token := GetAuthToken()

	interfaces := map[string]graphiant_sdk.ManaV2NullableInterfaceConfig{
		interfaceName: {
			Interface: &graphiant_sdk.ManaV2InterfaceConfig{
				Subinterfaces: &map[string]graphiant_sdk.ManaV2NullableInterfaceVlanConfig{
					fmt.Sprintf("%d", vlan): {
						Interface: &graphiant_sdk.ManaV2InterfaceVlanConfig{
							AdminStatus: graphiant_sdk.PtrBool(true),
							Ipv4: &graphiant_sdk.ManaV2InterfaceIpConfig{
								Address: &graphiant_sdk.ManaV2NullableAddress{
									Address: graphiant_sdk.PtrString(ipv4Address),
								},
							},
							Ipv6: &graphiant_sdk.ManaV2InterfaceIpConfig{
								Address: &graphiant_sdk.ManaV2NullableAddress{
									Address: graphiant_sdk.PtrString(ipv6Address),
								},
							},
							Lan:  graphiant_sdk.PtrString(lan),
							Vlan: graphiant_sdk.PtrInt32(vlan),
						},
					},
				},
			},
		},
	}
	configReq := graphiant_sdk.NewV1DevicesDeviceIdConfigPutRequest()
	configReq.Edge = graphiant_sdk.NewManaV2EdgeDeviceConfig()
	configReq.Edge.SetInterfaces(interfaces)
	interfacesJSON, _ := json.MarshalIndent(configReq.Edge.GetInterfaces(), "", "  ")

	fmt.Printf("Calling V1DevicesDeviceIdConfigPut : %v %s\n", deviceId, interfacesJSON)
	graphiant_sdk.PollAndPutDeviceConfig(apiClient, token, deviceId, *configReq)
}

func DeleteSubInterface(deviceId int64, interfaceName string, vlan int32) {
	apiClient, token := GetAuthToken()

	interfaces := map[string]graphiant_sdk.ManaV2NullableInterfaceConfig{
		interfaceName: {
			Interface: &graphiant_sdk.ManaV2InterfaceConfig{
				Subinterfaces: &map[string]graphiant_sdk.ManaV2NullableInterfaceVlanConfig{
					fmt.Sprintf("%d", vlan): {},
				},
			},
		},
	}
	configReq := graphiant_sdk.NewV1DevicesDeviceIdConfigPutRequest()
	configReq.Edge = graphiant_sdk.NewManaV2EdgeDeviceConfig()
	configReq.Edge.SetInterfaces(interfaces)
	interfacesJSON, _ := json.MarshalIndent(configReq.Edge.GetInterfaces(), "", "  ")

	fmt.Printf("Calling V1DevicesDeviceIdConfigPut : %v %s\n", deviceId, interfacesJSON)
	graphiant_sdk.PollAndPutDeviceConfig(apiClient, token, deviceId, *configReq)
}

func Test_configure_wan_interface(t *testing.T) {
	// Note: This test requires GRAPHIANT_USERNAME and GRAPHIANT_PASSWORD environment variables to be set
	// export GRAPHIANT_USERNAME="your_username"
	// export GRAPHIANT_PASSWORD="your_password"

	t.Skip("skip test") // remove to run test

	ConfigureWan(30000048276, "GigabitEthernet7/0/0", "c-gigabitethernet7-0-0")
}

func Test_configure_lan_interface(t *testing.T) {
	// Note: This test requires GRAPHIANT_USERNAME and GRAPHIANT_PASSWORD environment variables to be set
	// export GRAPHIANT_USERNAME="your_username"
	// export GRAPHIANT_PASSWORD="your_password"

	t.Skip("skip test") // remove to run test

	ConfigureLan(30000048276, "GigabitEthernet7/0/0", "lan-7-test", "10.2.7.2/24", "2001:10:2:7::2/64")
}

func Test_deconfigure_lan_interface(t *testing.T) {
	// Note: This test requires GRAPHIANT_USERNAME and GRAPHIANT_PASSWORD environment variables to be set
	// export GRAPHIANT_USERNAME="your_username"
	// export GRAPHIANT_PASSWORD="your_password"

	t.Skip("skip test") // remove to run test

	ConfigureDefaultLan(30000048276, "GigabitEthernet7/0/0", "default-10000000335")
}

func Test_configure_sub_interface(t *testing.T) {
	// Note: This test requires GRAPHIANT_USERNAME and GRAPHIANT_PASSWORD environment variables to be set
	// export GRAPHIANT_USERNAME="your_username"
	// export GRAPHIANT_PASSWORD="your_password"

	t.Skip("skip test") // remove to run test

	ConfigureSubInterface(30000048276, "GigabitEthernet7/0/0", "lan-8-test", 28, "10.2.8.2/24", "2001:10:2:8::2/64")
}

func Test_delete_sub_interface(t *testing.T) {
	// Note: This test requires GRAPHIANT_USERNAME and GRAPHIANT_PASSWORD environment variables to be set
	// export GRAPHIANT_USERNAME="your_username"
	// export GRAPHIANT_PASSWORD="your_password"

	t.Skip("skip test") // remove to run test

	DeleteSubInterface(30000048276, "GigabitEthernet7/0/0", 28)
}
