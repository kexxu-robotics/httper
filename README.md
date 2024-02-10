# httper
Http helper functions for Golang

## Respond

Respond Status 'Success'
```
http.HandleFunc(r *http.Request, w http.ResponseWriter){
    httper.RespondSuccess(w)
}
```

Respond Status
```
http.HandleFunc(r *http.Request, w http.ResponseWriter){
    httper.RespondStatus(w, "processing")
}
```

Respond with JSON
```
http.HandleFunc(r *http.Request, w http.ResponseWriter){
    car := Car{ Wheels: 4 }
    httper.RespondJson(w, car)
}
```

## Input

Get values from a GET parameter or POST form
```
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
}
```

Get json from POST body

```
http.HandleFunc(r *http.Request, w http.ResponseWriter){
    
    // basic settings, 1mb max, allow unknown fields    
    car := Car{}
    err := httper.GetJsonBodyDefault(w, r, car)
    if err != nil {
        return
    }

    // custom settings    
    // max 1024 bytes (1kb)
    // don't allow unknown fields (fields not in Train{})
    train := Train{}
    err := httper.GetJsonBody(w, r, train, 1024, false)
    if err != nil {
        return
    }

}
```
