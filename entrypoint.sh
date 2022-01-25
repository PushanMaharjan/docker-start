wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

# this line watches for go file changes
CompileDaemon --build="go build -o main main.go" --command=./main