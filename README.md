<p align="center">
  <a href="https://www.securenative.com"><img src="https://user-images.githubusercontent.com/45174009/77826512-f023ed80-7120-11ea-80e0-58aacde0a84e.png" alt="SecureNative Logo"/></a>
</p>

<p align="center">
  <b>A Cloud-Native Security Monitoring and Protection for Modern Applications</b>
</p>
<p align="center">
  <a href="https://github.com/securenative/securenative-go">
    <img alt="Github Actions" src="https://github.com/securenative/securenative-go/workflows/CI/badge.svg">
  </a>
  <a href="https://codecov.io/gh/securenative/securenative-go">
    <img src="https://codecov.io/gh/securenative/securenative-go/branch/master/graph/badge.svg" />
  </a>
  <a href="https://badge.fury.io/go/github.com%2Fsecurenative%2Fsecurenative-go"><img src="https://badge.fury.io/go/github.com%2Fsecurenative%2Fsecurenative-go.svg" alt="Go project version" height="18"></a>
</p>
<p align="center">
  <a href="https://docs.securenative.com">Documentation</a> |
  <a href="https://docs.securenative.com/quick-start">Quick Start</a> |
  <a href="https://blog.securenative.com">Blog</a> |
  <a href="">Chat with us on Slack!</a>
</p>
<hr/>


[SecureNative](https://www.securenative.com/) performs user monitoring by analyzing user interactions with your application and various factors such as network, devices, locations and access patterns to stop and prevent account takeover attacks.


## Install the SDK
```bash
go get github.com/securenative/securenative-go
```

## Initialize the SDK

To get your *API KEY*, login to your SecureNative account and go to project settings page:

### Option 1: Initialize via Config file
SecureNative can automatically load your config from *securenative.yml* file or from the file that is specified in your *SECURENATIVE_CONFIG_FILE* env variable:

```go
package main

import (
    "github.com/securenative/securenative-go/sdk"
    "log"
)

func main() {
    sn, err :=  sdk.InitSDK(sdk.WithConfigFile("path/to/securenative.yml"))
    if err != nil || sn == nil {
         log.Fatal("Do some error handling")
    }

    defer sn.Stop()
}
```
### Option 2: Initialize via API Key

```go
package main

import (
    "github.com/securenative/securenative-go/sdk"
    "log"
)

func main() {
    sn, err :=  sdk.InitSDK(sdk.WithApiKey("YOUR_API_KEY"))
    if err != nil {
         log.Fatal("Do some error handling")
    }

    defer sn.Stop()
}
```

### Option 3: Initialize via ConfigurationBuilder
```go
package main

import (
    "github.com/securenative/securenative-go/config"
    "github.com/securenative/securenative-go/sdk"
    "log"
)

func main() {
    options := config.DefaultSecureNativeOptions()
    options.ApiKey = "YOUR_API_KEY"
    options.MaxEvents = 10
    options.LogLevel = "ERROR"
    sn, err := sdk.InitSDK(options)
    if err != nil {
         log.Fatal("Do some error handling")
    }

    defer sn.Stop()
}
```

## Getting SecureNative instance
Once initialized, sdk will create a singleton instance which you can get: 
```go
package main

import (
    "github.com/securenative/securenative-go/sdk"
    "log"
)

func main() {
    sn, err := sdk.GetInstance()
    if err != nil {
        log.Fatal("Do some error handling")
    }
    
    defer sn.Stop()
}
```

## Tracking events

Once the SDK has been initialized, tracking requests sent through the SDK
instance. Make sure you build event with the EventBuilder:

 ```go
package main

import (
    "github.com/securenative/securenative-go/context"
    "github.com/securenative/securenative-go/enums"
    "github.com/securenative/securenative-go/events"
    "github.com/securenative/securenative-go/models"
    "github.com/securenative/securenative-go/sdk"
    "log"
)

func main() {
    sn, err := sdk.GetInstance()
    if err != nil {
            log.Fatal("Do some error handling")
    }

    c := &context.SecureNativeContext{
        ClientToken:    "SECURED_CLIENT_TOKEN",
        Ip:             "127.0.0.1",
        Headers:        map[string]string{"user-agent": "Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405"},
    }
    eventOptions := models.EventOptions{
        Event: enums.EventTypes.LogIn,
        UserId: "1234",
        UserTraits: models.UserTraits{Name:"Your Name", Email:"name@gmail.com", Phone: "+1234567890"},
        Context: c,
        Properties: map[string]interface{}{"prop1": "CUSTOM_PARAM_VALUE", "prop2": "true", "prop3": "3"},    
    }
    
    defer sn.Stop()
    
    sn.Track(eventOptions)
}
 ```

You can also create request context from requests:

```go
package demo

import (
    "github.com/securenative/securenative-go/context"
    "github.com/securenative/securenative-go/enums"
    "github.com/securenative/securenative-go/events"
    "github.com/securenative/securenative-go/models"
    "github.com/securenative/securenative-go/sdk"
    "log"
    "net/http"
)

func Track(request *http.Request) {
    sn, err := sdk.GetInstance()
    if err != nil {
        log.Fatal("Do some error handling")
    }

    c := context.FromHttpRequest(request)
    eventOptions := models.EventOptions{
        Event: enums.EventTypes.LogIn,
        UserId: "1234",
        Context: c,
    }

    defer sn.Stop()
      
    sn.Track(eventOptions)
}
```

## Verify events

**Example**

```go
package main

import (
    "github.com/securenative/securenative-go/context"
    "github.com/securenative/securenative-go/enums"
    "github.com/securenative/securenative-go/events"
    "github.com/securenative/securenative-go/models"
    "github.com/securenative/securenative-go/sdk"
    "log"
)

func main() {
    sn, err := sdk.GetInstance()
    if err != nil {
        log.Fatal("Do some error handling")
    }
    
    eventOptions := models.EventOptions{Event:enums.EventTypes.LogIn}
    defer sn.Stop()
    
    c := &context.SecureNativeContext{
            ClientToken:    "SECURED_CLIENT_TOKEN",
            Ip:             "127.0.0.1",
            Headers:        map[string]string{"user-agent": "Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405"},
    }
        
    verifyResult := sn.Verify(eventOptions)
    verifyResult.RiskLevel  // Low, Medium, High
    verifyResult.Score  // Risk score: 0 - 1 (0 - Very Low, 1 - Very High)
    verifyResult.Triggers  // ["TOR", "New IP", "New City"]
}
```

## WebHook signature verification

Apply our filter to verify the request is from us, for example:

```go
package demo

import (
    "github.com/securenative/securenative-go/sdk"
    "log"
    "net/http"
)

func VerifyWebHook(request *http.Request) bool {
    sn, err := sdk.GetInstance()
    if err != nil {
        log.Fatal("Do some error handling")
    }
    defer sn.Stop()
    
    return sn.VerifyRequestPayload(request)
}
 ```

## Extract proxy headers from cloud providers

You can specify custom header keys to allow extraction of client ip from different providers.
This example demonstrates the usage of proxy headers for ip extraction from Cloudflare.

### Option 1: Using config file
```yaml
SECURENATIVE_API_KEY: "YOUR_API_KEY"
SECURENATIVE_PROXY_HEADERS: ["CF-Connecting-IP"]
```

Initialize sdk as shown above.

### Options 2: Using ConfigurationBuilder

```go
options := config.DefaultSecureNativeOptions()
options.ApiKey = "YOUR_API_KEY"
options.ProxyHeaders = []string{"CF-Connecting-IP"}
    
sn, err := sdk.InitSDK(options)
if err != nil {
     log.Fatal("Do some error handling")
}
```
    
## Remove PII Data From Headers

By default, SecureNative SDK remove any known pii headers from the received request.
We also support using custom pii headers and regex matching via configuration, for example:

### Option 1: Using config file
```yaml
SECURENATIVE_API_KEY: "YOUR_API_KEY"
SECURENATIVE_PII_HEADERS: ["apiKey"]
```

Initialize sdk as shown above.

### Options 2: Using ConfigurationBuilder

```go
options := config.DefaultSecureNativeOptions()
options.ApiKey = "YOUR_API_KEY"
options.PiiHeaders = []string{"authentication"}
    
sn, err := sdk.InitSDK(options)
if err != nil {
     log.Fatal("Do some error handling")
}
```