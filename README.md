# go-server

## Run App

* Clone the repo

* Go to root dir
```bash
$ cd go-server
```

* Run server
```bash
$ go run ./
```

* Go to UI dir
```bash
$ cd ui
```

* Run Angular dev server
```bash
$ ng serve --open
```

UI host: http://localhost:4200
Server host: http://localhost:9091

**NOTE**: Cross domen requests are blocked by CORS policy.
Run chrome `disable-web-security` mode. For mac: 
```bash
$ open -a Google\ Chrome --args --disable-web-security --user-data-dir
```
Or add Chrome Exension:
https://chrome.google.com/webstore/detail/allow-control-allow-origi/



## API:
- `GET /api/getItem?id={id}`


- `POST /api/addItem`

```json
    { "message": "First Item" }
```

- `DELETE /api/deleteItem`
