# Generate an API file
apis:
	goctl api go -api ./api/app.api -dir . -style go_zero -home=./tpl
# Generate a swagger file
swagger:
	goctl api plugin -plugin goctl-swagger="swagger -filename doc/swagger/app.json" -api api/app.api -dir .
# Generate a swagger file
gen:
	goctl api go -api ./api/app.api -dir . -style go_zero -home=./tpl
	goctl api plugin -plugin goctl-swagger="swagger -filename doc/swagger/app.json" -api api/app.api -dir .
	go generate ./...
	go mod tidy
# Generate an executable file
exe:
	make gen
	go build app.go
# Run the executable file
run:
	make gen
	go mod tidy
	go run app.go