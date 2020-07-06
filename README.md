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

import "github.com/securenative/securenative-go/securenative"

func main() {
    sn, err :=  securenative.InitSDK()
    if err != nil {
         // Do some error handling
    }

    defer sn.ReleaseSDK()
}
```
### Option 2: Initialize via API Key

```go
package main

import "github.com/securenative/securenative-go/securenative"

func main() {
    sn, err :=  securenative.InitSDKWithApiKey("YOUR_API_KEY")
    if err != nil {
         // Do some error handling
    }

    defer sn.ReleaseSDK()
}
```

### Option 3: Initialize via ConfigurationBuilder
```go
package main

import (
    "github.com/securenative/securenative-go/securenative"
    . "github.com/securenative/securenative-go/securenative/config"
)

func main() {
    configBuilder := securenative.GetConfigBuilder()
    sn, err := securenative.InitSDKWithOptions(configBuilder.WithApiKey("API_KEY").WithMaxEvents(10).WithLogLevel("ERROR").Build())
    if err != nil {
         // Do some error handling
    }

    defer sn.ReleaseSDK()
}
```

## Getting SecureNative instance
Once initialized, sdk will create a singleton instance which you can get: 
```go
package main

import "github.com/securenative/securenative-go/securenative"

func main() {
    sn := securenative.GetInstance()
    
    defer sn.ReleaseSDK()
}
```

## Tracking events

Once the SDK has been initialized, tracking requests sent through the SDK
instance. Make sure you build event with the EventBuilder:

 ```go
package main

import (
    "github.com/securenative/securenative-go/securenative"
    . "github.com/securenative/securenative-go/securenative/config"
    . "github.com/securenative/securenative-go/securenative/context"
    . "github.com/securenative/securenative-go/securenative/enums"
    . "github.com/securenative/securenative-go/securenative/models"
)

func main() {
    sn := securenative.GetInstance()
    contextBuilder := securenative.GetContextBuilder()
    eventOptionsBuilder := sn.GetEventOptionsBuilder(EventTypes.LogIn)
    
    defer sn.ReleaseSDK()
    
    context := contextBuilder.WithIp("127.0.0.1").WithClientToken("SECURED_CLIENT_TOKEN").WithHeaders(map[string]string{"user-agent": "Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405"}).Build()
    eventOptions, err := eventOptionsBuilder.WithUserId("1234").WithUserTraits(UserTraits{Name:"Your Name", Email:"name@gmail.com"}).WithContext(context).WithProperties(map[string]string{"prop1": "CUSTOM_PARAM_VALUE", "prop2": "true", "prop3": "3"}).Build()
    if err != nil {
        // Do some error handling
    }
    
    sn.Track(eventOptions)
}
 ```

You can also create request context from requests:

```go
package main

import (
    "github.com/securenative/securenative-go/securenative"
    . "github.com/securenative/securenative-go/securenative/config"
    . "github.com/securenative/securenative-go/securenative/context"
    . "github.com/securenative/securenative-go/securenative/enums"
    . "github.com/securenative/securenative-go/securenative/models"
)

func main() {
    sn := securenative.GetInstance()
    contextBuilder := securenative.GetContextBuilder()
    eventOptionsBuilder := sn.GetEventOptionsBuilder(EventTypes.LogIn)
    
    context := contextBuilder.FromHttpRequest(request)
    
    defer sn.ReleaseSDK()
    
    eventOptions, err := eventOptionsBuilder.WithUserId("1234").WithUserTraits(UserTraits{Name:"Your Name", Email:"name@gmail.com"}).WithContext(context).WithProperties(map[string]string{"prop1": "CUSTOM_PARAM_VALUE", "prop2": "true", "prop3": "3"}).Build()
    if err != nil {
        // Do some error handling
    }
    
    sn.Track(eventOptions)
}
```

## Verify events

**Example**

```go
package main

import (
	"github.com/securenative/securenative-go/securenative"
    . "github.com/securenative/securenative-go/securenative/enums"
    . "github.com/securenative/securenative-go/securenative/models"
)

func main() {
    sn := securenative.GetInstance()
    contextBuilder := securenative.GetContextBuilder()
    eventOptionsBuilder := sn.GetEventOptionsBuilder(EventTypes.LogIn)
    
    defer sn.ReleaseSDK()
    
    context := contextBuilder.WithIp("127.0.0.1").WithClientToken("SECURED_CLIENT_TOKEN").WithHeaders(map[string]string{"user-agent": "Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405"}).Build()
    eventOptions, err := eventOptionsBuilder.WithUserId("1234").WithUserTraits(UserTraits{Name:"Your Name", Email:"name@gmail.com"}).WithContext(context).WithProperties(map[string]string{"prop1": "CUSTOM_PARAM_VALUE", "prop2": "true", "prop3": "3"}).Build()
    
    if err != nil {
        // Do some error handling
    }
        
    verifyResult := sn.Verify(eventOptions)
    verifyResult.RiskLevel  // Low, Medium, High
    verifyResult.Score  // Risk score: 0 -1 (0 - Very Low, 1 - Very High)
    verifyResult.Triggers  // ["TOR", "New IP", "New City"]
}
```

## WebHook signature verification

Apply our filter to verify the request is from us, for example:

```go
package demo

import (
    "github.com/securenative/securenative-go/securenative"
    "net/http"
)

func WebhookEndpoint(request *http.Request) bool {
    sn := securenative.GetInstance()
    defer sn.ReleaseSDK()
    
    return sn.VerifyRequestPayload(request)
}
 ```
    