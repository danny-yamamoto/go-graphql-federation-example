# go-graphql-federation-example
Implement GraphQL Federation.

```bash
go install github.com/99designs/gqlgen@latest
go get github.com/99designs/gqlgen/graphql/handler
```

## Running Subquery.
```bash
vscode ➜ /workspaces/go-graphql-federation-example/users (main) $ go run server.go 
2023/07/25 04:31:12 connect to http://localhost:4000/ for GraphQL playground
```

```bash
vscode ➜ /workspaces/go-graphql-federation-example (main) $ curl -X POST -H "Content-Type: Application/json" -d '{"query":"{ todos { id } }"}' http://localhost:4000/query
{"data":{"todos":[{"id":"TODO-1"},{"id":"TODO-2"}]}}
vscode ➜ /workspaces/go-graphql-federation-example (main) $ 
```

## Running Bramble.
```bash
go run main.go -conf ./config.json 
```

## Querying Bramble.
### `id` only
```bash
vscode ➜ /workspaces/go-graphql-federation-example (main) $ curl -X POST -H "Content-Type: Application/json" -d '{"query":"{ todos { id } }"}' http://localhost:8082/query
{"data":{"todos":[{"id":"TODO-1"},{"id":"TODO-2"}]}}
vscode ➜ /workspaces/go-graphql-federation-example (main) $ 
```

### `id` and `text`
### 
```bash
vscode ➜ /workspaces/go-graphql-federation-example (main) $ curl -X POST -H "Content-Type: Application/json" -d '{"query":"{ todos { id text } }"}' http://localhost:8082/query
{"data":{"todos":[{"id":"TODO-1","text":"My Todo 1"},{"id":"TODO-2","text":"My Todo 2"}]}}vscode ➜ /workspaces/go-graphql-federation-example (main) $
```
