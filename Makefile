build: list dashboard

docker: clean build
	docker build -t gitscore .

list:
	CGO_ENABLED=0 GOOS=linux go build -o list -a ./cmds/gitscore/

dashboard:
	CGO_ENABLED=0 GOOS=linux go build -o dashboard -a ./cmds/gitscore-dashboard/

clean:
	rm list
	rm dashboard