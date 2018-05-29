# bow-api
Some test stuff

To build:
```
make local
```

To run:
```
go run main.go
```
or 
```
./bow-api
```

1. REGISTER:
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"player 1"}' \
  http://localhost:8080/api/v1/register
```

2. RECORD:
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"userName":"player 1","frameID":11,"score":10}' \
  http://localhost:8080/api/v1/record
```

3. DISPLAY:
```
curl --header "Content-Type: application/json" \
  --request GET \
  http://localhost:8080/api/v1/display
```
