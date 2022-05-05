// Read .env for Apollo metrics and schema
require('dotenv').config()

const { ApolloServer } = require('apollo-server');
const { ApolloGateway } = require("@apollo/gateway");


const gateway = new ApolloGateway();

const server = new ApolloServer({
    gateway,
    introspection: process.env.NODE_ENV !== 'prod'
});

server.listen().then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
});
