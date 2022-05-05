// Read .env for Apollo metrics and schema
require('dotenv').config()

const { ApolloServer } = require('apollo-server');
const { ApolloGateway } = require('@apollo/gateway');
const { ApolloServerPluginLandingPageLocalDefault,
    ApolloServerPluginLandingPageProductionDefault
} = require('apollo-server-core');


const gateway = new ApolloGateway();

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
