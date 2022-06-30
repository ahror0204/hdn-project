CURRENT_DIR=$(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	./scripts/gen-proto.sh	${CURRENT_DIR}
	ls genproto/*.pb.go | xargs -n1 -IX bash -c "sed -e '/bool/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"

create-migrations:
	migrate create -ext sql -dir migrations -seq addtokencolumn

migrate-up:
	migrate -path migrations/ -database postgres://postgres:1@localhost:5432/businessdb?sslmode=disable up

migrate-down:
	migrate -path migrations/ -database postgres://postgres:1@localhost:5432/businessdb?sslmode=disable down

migrate-forse:
	migrate -path migrations/ -database postgres://postgres:1@localhost:5432/businessdb?sslmode=disable force 1

#  migrate create -ext sql -dir migrations -seq init_schema