# [Create handler by proxy reverse](#lintest)
<a name="lintest"></a>

### Instalation
```
 go get -u github.com/ecsavigne/proxy-reverse@latest
```

### Example
```
    import (
        "github.com/ecsavigne/proxy-reverse/proxy"
        "github.com/gin-gonic/gin"
        )

    func main() {
        p := NewProxyReverse(proxy.ProxyReverse{
            Host    : "localhost"
            // Port    : ""
            Protocol: "http"
        }) 

        g := gin.Default()
        g.GET("/test1", p.RequestProxy())
        g.Run(":7895")
    }

```
