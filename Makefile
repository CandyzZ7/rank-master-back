.PHONY: api
# generate api proto
api:
	goctl api go -api ./api/app.api -dir . -style go_zero -home=./tpl
.PHONY: swagger
# Generate a swagger file
swagger:
	goctl api plugin -plugin goctl-swagger="swagger -filename doc/swagger/app.json" -api api/app.api -dir .
.PHONY: gen
# Generate a swagger file
gen:
	goctl api go -api ./api/app.api -dir . -style go_zero -home=./tpl
	goctl api plugin -plugin goctl-swagger="swagger -filename doc/swagger/app.json" -api api/app.api -dir .
	go generate ./...
	wire
	go mod tidy
.PHONY: exe
# Generate an executable file
exe:
	make gen
	go build app.go
.PHONY: run
# Run the executable file
run:
	make gen
	go mod tidy
	go run app.go wire_gen.go