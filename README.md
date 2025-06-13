# Graphiant-SDK-Go
Go SDK for [Graphiant NaaS](https://www.graphiant.com).

Refer [Graphiant Documentation](https://docs.graphiant.com/) to get started with our services.


## Install

1. **Clone the repository:**
   ```sh
   git clone https://github.com/Graphiant-Inc/graphiant-sdk-go
   cd graphiant-sdk-go/graphiant-sdk
   ```

2. **Install dependencies:**
   ```sh
   go mod tidy
   ```

## Run Tests

1. **Set environment variables for authentication:**
   ```sh
   export username=your_username
   export password=your_password
   ```

2. **Run a quick test to validate the SDK:**
   ```sh
   cd test
   go test -v -run Test_edge_summary

   Note: This test should return the Edge devices in the enterprise
   ```

## Additional Information

- Refer to [Graphiant Documentation](https://docs.graphiant.com/) for API usage and examples.
- For any issues, please open an issue on this repository.

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