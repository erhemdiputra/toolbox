## Toolbox

### Helper package to help your programming activity

1. [HTTP Request](#http-request)

---

### <a name="#http-request">HTTP Request</a>
`GET`
- Parameter:
    - URL: string
    - header: http.Header
    - param: map[string]string (query string)
    - out: interface{} (response from api)

- Example Request:
```
getResp := make(map[string]interface{})
httpRequestInst := httprequest.NewHTTPRequest(5 * time.Second)
err := httpRequestInst.Get("https://jsonplaceholder.typicode.com/todos/1", http.Header{}, map[string]string{
    "hello":    "world",
    "kenpachi": "zaraki",
}, &getResp)
```

`POST`
- Parameter: 
    - URL: string
    - header: http.Header
    - param: map[string]string (query string, application/x-www-form-urlencoded)
    - payload: interface{} (json payload, application/json)
    - out: interface{} (response from api)

- If param `payload` not nil, header `Content-Type` will set to `application/json`, otherwise `application/x-www-form-urlencoded`

- Example Request
```
postResp := make(map[string]interface{})
err = httpRequestInst.Post("https://jsonplaceholder.typicode.com/posts",
    http.Header{}, map[string]string{
        "shinigami": "ichigo kurosaki, byakuya kuchiki, renji abarai!",
    }, nil, &postResp,
)
```
