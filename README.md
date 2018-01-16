
# HelloWorld Golang Lambda

## What it does
    
You post to the lambda some json:
    
```
    {
      "message": "hello",
    }
```

The lambda responds with some json
    
```    
    {
      "message": "you sent the message hello",
      "ok": true
    }
```

## Running tests

```go test```

## Deploying

```./deploy.sh```

    