httper
===
Http helper functions for Golang

![tests](https://github.com/kexxu-robotics/httper/actions/workflows/go.yml/badge.svg)

## Respond

Respond Status 'Success'
```go
http.HandleFunc(r *http.Request, w http.ResponseWriter){
    httper.RespondSuccess(w)
}
```

Respond Status
```go
http.HandleFunc(r *http.Request, w http.ResponseWriter){
    httper.RespondStatus(w, "processing")
}
```

Respond with JSON
```go
http.HandleFunc(r *http.Request, w http.ResponseWriter){
    car := Car{ Wheels: 4 }
    httper.RespondJson(w, car)
}
```

Respond with JSON gzipped, reducing network load for larger responses
```go
http.HandleFunc(r *http.Request, w http.ResponseWriter){
    car := Car{ Wheels: 4 }
    httper.RespondJsonGzipped(w, car)
}
```


## Input

Get values from a GET parameter or POST form
```go
http.HandleFunc(r *http.Request, w http.ResponseWriter){

    // get an int
    val, err := httper.CheckFromInt(r, w, "value")
    if err != nil {
        return
    }

    // get an int64
    id, err := httper.CheckFromInt64(r, w, "id")
    if err != nil {
        return
    }

    // get string
    hash, err := httper.CheckFormString(r, w, "hash")
    if err != nil {
        return
    }

    // get string with default fallback
    remarks := httper.CheckFormStringDefault(r, w, "remarks", "No remarks")

}
```

Get json from POST body

```go
http.HandleFunc(r *http.Request, w http.ResponseWriter){
    
    // basic settings, 1mb max, allow unknown fields    
    car := Car{}
    err := httper.GetJsonBodyDefault(r, w, &car)
    if err != nil {
        return
    }

    // custom settings    
    // max 1024 bytes (1kb)
    // don't allow unknown fields (fields not in Train{})
    train := Train{}
    err := httper.GetJsonBody(r, w, &train, 1024, false)
    if err != nil {
        return
    }

}
```
