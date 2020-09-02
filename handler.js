'use strict';

const template = fs.readFileSync('static/index.html', 'utf-8');

module.exports.hello = async event => {
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
