'use strict';
const fs = require('fs');
const template = fs.readFileSync('static/index.html', 'utf-8');

module.exports.index = async event => {
  const html = Mustache.render(template, view);

  const response = {
    statusCode: 200,
    headers: {
      'Content-Type': 'text/html; charset=UTF-8'
    },
    body: html
  }

  return response;
};
