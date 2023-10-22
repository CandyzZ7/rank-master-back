api:
	goctl api go -api ./api/app.api -dir . -style go_zero -home=./tpl
swagger:
	goctl api plugin -plugin goctl-swagger="swagger -filename doc/swagger/app.json" -api api/app.api -dir .
run:
	go run app.go
exe:
	go build app.go