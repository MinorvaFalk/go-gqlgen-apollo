// Read .env for Apollo metrics and schema
require('dotenv').config()

const { ApolloServer } = require('apollo-server');
const { ApolloGateway, IntrospectAndCompose  } = require('@apollo/gateway');
const { ApolloServerPluginLandingPageLocalDefault,
    ApolloServerPluginLandingPageProductionDefault
} = require('apollo-server-core');

const goAPI = process.env["GO_HOST"] || "localhost:4001";

const gateway = new ApolloGateway({
    supegraphSdl: new IntrospectAndCompose({
        subgraphs: [
            {name: 'go-api', url: `http://${goAPI}/query`}
        ]
    })
});

const server = new ApolloServer({
    gateway,
    introspection: process.env.NODE_ENV !== 'prod',
    plugins: [
        // Install a landing page plugin based on NODE_ENV
        process.env.NODE_ENV === 'prod'
          ? ApolloServerPluginLandingPageProductionDefault({
              footer: false,
            })
          : ApolloServerPluginLandingPageLocalDefault({ footer: false }),
      ],
});

server.listen().then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
});
