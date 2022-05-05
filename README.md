# Quickstart
Steps :
1. Make sure you have an `Apollo Studio` account.
   > Create account apollo studio account [here](https://studio.apollographql.com/)
2. Create a new `Federation Graph` on Apollo Studio
3. Create .env file inside `./apollo-gateway` then copy APOLLO_KEY and APOLLO_GRAPH_REF
  > Check .env-example file inside ./apollo-gateway for .env format

  > For apollo graph ref `rover subgraph publish <your graph-ref>`
4. Publish your schema using this [instruction](#Publishing)
5. Run your `Apollo Gateway` and `Go-API`

> Using **normal start** :
>```bash
># First terminal
># ...
>
># Make sure to add .env with API-Key and Graph-Ref
># Check ./apollo-gateway/.env-example for >example
>$ cd apollo-gateway
>$ node index.js
>
># Second terminal
># ...
>
>$ cd go-api
>$ go run .
>```

>Using **docker compose**:
>```bash
>$ docker compose up --build
>```

### Publishing Schema to Apollo Studio : <a id="Publishing"></a>
1. Make sure to add API-KEY using `rover config auth`
2. Run this command:
```bash
rover subgraph publish <GRAPH-REF> \
  --name monolith-go-api --schema ./go-api/graph/schema.graphqls \
  --routing-url http://localhost:4001
```