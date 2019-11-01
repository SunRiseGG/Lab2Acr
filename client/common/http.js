'use strict';

const request = require('request');

const client = baseUrl => {
  const respHandler = response => {
    if (response.ok) {
      return response.json();
    }
    throw new Error(`Unexpected response from the server ${response.status} ${response.statusText}`)
  };
  return {
    get: path => {
      return new Promise((resolve, reject) => {
        request(`${baseUrl}${path}`, { json: true }, (err, res, body) => {
          if (err) {
            reject(err);
            return;
          }
          resolve(body);
        });
      });
    },
    post: async (path, data) => {
      return new Promise((resolve, reject) => {
        request(`${baseUrl}${path}`, { json: true, method: 'POST', body: data }, (err, res, body) => {
          if (err) {
            reject(err);
            return;
          }
          resolve(body);
        });
      });
    }
  };
};

module.exports = { client };
