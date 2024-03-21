apis:
	goctl api go -api ./api/app.api -dir . -style go_zero -home=./tpl
swagger:
	goctl api plugin -plugin goctl-swagger="swagger -filename doc/swagger/app.json" -api api/app.api -dir .
gen:
	goctl api go -api ./api/app.api -dir . -style go_zero -home=./tpl
	goctl api plugin -plugin goctl-swagger="swagger -filename doc/swagger/app.json" -api api/app.api -dir .
	go generate ./...
	go mod tidy
exe:
	make gen
	go build app.go
run:
	make gen
	go mod tidy
	go run app.go