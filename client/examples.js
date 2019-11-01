'use strict';
const scenarios = require('./scenarios/client');

const client = scenarios.clientApi('http://localhost:8000');

// Scenario 1: Display available accounts.
client.accountsList()
    .then(list => {
        console.log('-----Scenario 1-----');
        console.log('Available accounts: ');
        console.dir(list);
    })
    .catch(e => {
        console.log(`Problem listing available accounts: ${e.message}`);
    });

client.transaction({ giverId: 2, receiverId: 1, sum: 2000.14 })
  .then(res => {
    console.log('-----Scenario 2-----');
    console.log('Create transaction response:', res);
    return client.accountsList()
    .then(list => list.map(c => c.name).join(', '));
    console.log('Transaction completed');
  })
  .catch(e => {
      console.log(`Problem creating a new transaction: ${e.message}`);
  });
