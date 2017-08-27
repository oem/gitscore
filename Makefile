build: list dashboard

list:
	CGO_ENABLED=0 GOOS=linux go build -o list -a ./cmds/gitscore/

dashboard:
	CGO_ENABLED=0 GOOS=linux go build -o dashboard -a ./cmds/gitscore-dashboard/