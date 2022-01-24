wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

# this line watches for go file changes
CompileDaemon --build="go build -0 main.go" --command=./main