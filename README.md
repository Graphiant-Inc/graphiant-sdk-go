# Graphiant-SDK-Go

Go SDK for [Graphiant NaaS](https://www.graphiant.com).

Refer [Graphiant-SDK-Go User Guide](https://docs.graphiant.com/docs/graphiant-sdk-go) under [Automation Section](https://docs.graphiant.com/docs/automation) in [Graphiant Documentation](https://docs.graphiant.com/) for getting started instructions.

## Import

Graphiant-SDK-Go package can be imported using Go tools. Refer [graphiant-sdk-go](https://pkg.go.dev/github.com/Graphiant-Inc/graphiant-sdk-go) package on public Go package registry.

```sh
go get github.com/Graphiant-Inc/graphiant-sdk-go
```

## Install from source code

Steps to clone and install Graphiant-SDK-Go from source code.

### Clone the graphiant-sdk-go repository

```sh
git clone https://github.com/Graphiant-Inc/graphiant-sdk-go
cd graphiant-sdk-go
```

### Install dependencies
```sh
go mod tidy
```

## Test

Steps to sanity test Graphiant-SDK-Go.

### Set environment variables for Graphiant API authentication
```sh
export username=<your_graphiant_portal_login_email>
export password=<your_graphiant_portal_login_password>
```

### Run sanity test
```sh
cd test
go test -v -run Test_edge_summary
```

This test should return your enterprise edge device summary on Graphiant portal.

```sh
=== RUN   Test_edge_summary
=== RUN   Test_edge_summary/Test_DefaultAPIService_V1AuthLoginPost
    sanity_test.go:37: Calling V1AuthLoginPost
    sanity_test.go:54: Auth token: {true graphiant gr-auth-848fc9fc-fcbc-417c-8292-c5629de06b0c-55b8a92b-7786-4c9c-9b96-f64aa8e7788f}
    sanity_test.go:58: Calling V1EdgesSummaryGet
    sanity_test.go:69: Edge summary response: [
          {
            "assignedOn": {
              "nanos": 852650000,
              "seconds": 1724282838
            },
            "deviceId": 30000056161,
            "discoveredLocation": "Amsterdam, Saskatchewan, Canada",
            "enterpriseId": 10000000000,
            "firstAppearedOn": {
              "nanos": 953831000,
              "seconds": 1724282838
            },

<cropped>
--- PASS: Test_edge_summary (1.80s)
    --- PASS: Test_edge_summary/Test_DefaultAPIService_V1AuthLoginPost (1.80s)
PASS
ok      github.com/Graphiant-Inc/graphiant-sdk-go/test  2.155s

```

## Generate

Steps to generate Graphiant-SDK-Go locally using Graphiant API docs. 

```sh
openapi-generator generate -i graphiant_api_docs_v25.6.2.json -g go --git-user-id Graphiant-Inc --git-repo-id graphiant-sdk-go --package-name graphiant_sdk
```

Note: Latest version of Graphiant API docs can be downloaded from Graphiant portal under "Support Hub" > "Developer Tools".

## License

Copyright (c) 2025 Graphiant-Inc

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.