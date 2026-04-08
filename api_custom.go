package graphiant_sdk

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetDeviceStatus(apiClient *APIClient, token string, deviceId int64) bool {

	resp, _, _ := apiClient.DefaultAPI.
		V1EdgesSummaryGet(context.Background()).
		Authorization(token).
		Execute()

	for _, edge := range resp.GetEdgesSummary() {
		if edge.GetDeviceId() == deviceId {
			if (edge.GetStatus() == "active" || edge.GetStatus() == "maintenance" || edge.GetStatus() == "staging") && edge.GetPortalStatus() == "Ready" && edge.GetTtConnCount() == 2 {
				fmt.Printf("Device %s (%d) is ready : Status is %s and Portal Status is Ready\n", edge.GetHostname(), deviceId, edge.GetStatus())
				return true
			} else {
				fmt.Printf("Device %s (%d) is not ready : Status is %s and Portal Status is %s\n", edge.GetHostname(), deviceId, edge.GetStatus(), edge.GetPortalStatus())
			}
		}
	}
	return false
}

func PutDeviceConfig(apiClient *APIClient, token string, deviceId int64, deviceConfig V1DevicesDeviceIdConfigPutRequest) *http.Response {

	fmt.Printf("Executing config put request for device %d...\n", deviceId)

	_, httpRes, err := apiClient.DefaultAPI.
		V1DevicesDeviceIdConfigPut(context.Background(), deviceId).
		Authorization(token).
		V1DevicesDeviceIdConfigPutRequest(deviceConfig).Execute()

	if err != nil {
		if httpRes != nil {
			// Drain and discard response body to avoid logging potentially sensitive data.
			_, _ = io.ReadAll(httpRes.Body)
			fmt.Printf("Error executing config put request for device %d (HTTP Status: %d). Error details omitted from log.\n", deviceId, httpRes.StatusCode)
		} else {
			fmt.Printf("Error executing config put request for device %d. Error details omitted from log.\n", deviceId)
		}
		return nil
	}
	// Read and discard response body without logging its potentially sensitive contents.
	_, _ = io.ReadAll(httpRes.Body)
	fmt.Printf("Config put request for device %d completed successfully with HTTP status: %v\n", deviceId, httpRes.StatusCode)
	return httpRes

}

func PollAndPutDeviceConfig(apiClient *APIClient, token string, deviceId int64, deviceConfig V1DevicesDeviceIdConfigPutRequest) *http.Response {

	// Poll device status every 30 seconds for 10 attempts
	for i := 0; i < 10; i++ {
		fmt.Printf("Polling attempt %d/5 for device %d...\n", i+1, deviceId)

		if GetDeviceStatus(apiClient, token, deviceId) {

			return PutDeviceConfig(apiClient, token, deviceId, deviceConfig)
		}

		if i < 9 { // Don't sleep after the last attempt
			fmt.Printf("Device not ready, waiting 30 seconds...\n")
			time.Sleep(30 * time.Second)
		}
	}

	fmt.Printf("Device %d did not become ready after 10 attempts\n", deviceId)
	return nil
}