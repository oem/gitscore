build:
	CGO_ENABLED=0 GOOS=linux go build -o list -a ./cmds/gitscore/
	CGO_ENABLED=0 GOOS=linux go build -o dashboard -a ./cmds/gitscore-dashboard/
