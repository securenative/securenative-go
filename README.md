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
SecureNative can automatically load your config from *securenative.ini* file or from the file that is specified in your *SECURENATIVE_CONFIG_FILE* env variable:

```go
package demo

import . "github.com/securenative/securenative-go/securenative"

secureative :=  InitSDK()
```
### Option 2: Initialize via API Key

```go
package demo

import . "github.com/securenative/securenative-go/securenative"

secureNative :=  InitSDKWithApiKey("YOUR_API_KEY")
```

### Option 3: Initialize via ConfigurationBuilder
```go
package demo

import (
    . "github.com/securenative/securenative-go/securenative"
    . "github.com/securenative/securenative-go/securenative/config"
)


configBuilder := NewConfigurationBuilder()
secureNative := InitSDKWithOptions(configBuilder
                                    .WithApiKey("API_KEY")
                                    .WithMaxEvents(10)
                                    .WithLogLevel("ERROR")
                                    .Build())
```

## Getting SecureNative instance
Once initialized, sdk will create a singleton instance which you can get: 
```go
package demo

import . "github.com/securenative/securenative-go/securenative"

secureNative := GetSDKInstance()
```

## Tracking events

Once the SDK has been initialized, tracking requests sent through the SDK
instance. Make sure you build event with the EventBuilder:

 ```go
package demo

import (
    . "github.com/securenative/securenative-go/securenative"
    . "github.com/securenative/securenative-go/securenative/config"
    . "github.com/securenative/securenative-go/securenative/context"
    . "github.com/securenative/securenative-go/securenative/enums"
    . "github.com/securenative/securenative-go/securenative/models"
)

secureNative := GetSDKInstance()
configBuilder := NewConfigurationBuilder()
contextBuilder := NewContextBuilder()
eventOptionsBuilder := EventOptionsBuilder(EventTypes.LOG_IN)

context := contextBuilder.
        WithIp("127.0.0.1").
        WithClientToken("SECURED_CLIENT_TOKEN").
        WithHeaders({"user-agent", "Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405"}).
        Build()

eventOptions := eventOptionsBuilder.
	    WithUserId("1234").
        WithUserTraits(UserTraits("Your Name", "name@gmail.com")).
        WithContext(context).
        WithProperties({"prop1": "CUSTOM_PARAM_VALUE", "prop2": True, "prop3": 3}).
        Build()

secureNative.Track(eventOptions)
 ```

You can also create request context from requests:

```go
package demo

import (
    . "github.com/securenative/securenative-go/securenative"
    . "github.com/securenative/securenative-go/securenative/config"
    . "github.com/securenative/securenative-go/securenative/context"
    . "github.com/securenative/securenative-go/securenative/enums"
    . "github.com/securenative/securenative-go/securenative/models"
)

secureNative := GetSDKInstance()
configBuilder := NewConfigurationBuilder()
contextBuilder := NewContextBuilder()
eventOptionsBuilder := EventOptionsBuilder(EventTypes.LOG_IN)

context := contextBuilder.FromHttpRequest(request)

eventOptions := eventOptionsBuilder.
	    WithUserId("1234").
        WithUserTraits(UserTraits("Your Name", "name@gmail.com")).
        WithContext(context).
        WithProperties({"prop1": "CUSTOM_PARAM_VALUE", "prop2": True, "prop3": 3}).
        Build()

secureNative.Track(eventOptions)
```

## Verify events

**Example**

```go
package demo

import (
    . "github.com/securenative/securenative-go/securenative"
    . "github.com/securenative/securenative-go/securenative/config"
    . "github.com/securenative/securenative-go/securenative/context"
    . "github.com/securenative/securenative-go/securenative/enums"
    . "github.com/securenative/securenative-go/securenative/models"
)

secureNative := GetSDKInstance()
configBuilder := NewConfigurationBuilder()
contextBuilder := NewContextBuilder()
eventOptionsBuilder := EventOptionsBuilder(EventTypes.LOG_IN)

context := contextBuilder.
        WithIp("127.0.0.1").
        WithClientToken("SECURED_CLIENT_TOKEN").
        WithHeaders({"user-agent", "Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405"}).
        Build()

eventOptions := eventOptionsBuilder.
	    WithUserId("1234").
        WithUserTraits(UserTraits("Your Name", "name@gmail.com")).
        WithContext(context).
        WithProperties({"prop1": "CUSTOM_PARAM_VALUE", "prop2": True, "prop3": 3}).
        Build()
    
verifyResult := secureNative.Verify(eventOptions)
verifyResult.RiskLevel  // Low, Medium, High
verifyResult.Score  // Risk score: 0 -1 (0 - Very Low, 1 - Very High)
verifyResult.Triggers  // ["TOR", "New IP", "New City"]
```

## Webhook signature verification

Apply our filter to verify the request is from us, for example:

```go
package demo

import . "github.com/securenative/securenative-go/securenative"

func WebhookEndpoint(request) bool {
    secureNative := GetSDKInstance()
    
    return secureNative.VerifyRequestPayload(request)
}
 ```
    