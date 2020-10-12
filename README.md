# go-grpc-api

This project is my first step into gRPC and GO. 
The project deals with creating an api wich returns an addition or multiplication of two numbers and it tells also weither a number is prime or not.


```bash
git clone https://github.com/amine7777/go-grpc-api.git
cd go-grpc-api/server
go run main.go
cd ..
cd client 
go run main.go
```
Check your browser on:

```
localhost:8580/add/<FirstNumber>/<SecondNumber>           
localhost:8580/mult/<FirstNumber>/<SecondNumber>           
localhost:8580/prime/<AnyNumber>     
```


