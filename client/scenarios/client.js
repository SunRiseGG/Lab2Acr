'use strict';
const http = require('../common/http');

const clientApi = baseUrl => {
  const client = http.client(baseUrl);
  return {
    accountsList: () => client.get('/banking'),
    transaction: ({ giverId, receiverId, sum }) => client.post('/banking', { giverId, receiverId, sum })
  }
};

module.exports = { clientApi };
