.PHONY: api
# generate api proto
api:
	goctl api go -api ./api/v1/app.api -dir . -style go_zero -home=./tpl
	make swagger
.PHONY: swagger
# Generate a swagger file
swagger:
	goctl api plugin -plugin goctl-swagger="swagger -filename doc/swagger/app.json" -api api/v1/app.api -dir .
.PHONY: gen
# Generate a swagger file
gen:
	make api
	make swagger
	go generate ./...
	wire
	go mod tidy
.PHONY: exe
# Generate an executable file
exe:
	make gen
	go build app.go wire_gen.go
.PHONY: run
# Run the executable file
run:
	make gen
	go run app.go wire_gen.go