const express = require('express');
const bodyParser = require('body-parser');
const morgan = require('morgan');
const webhookHandler = require('./webhook_handler');

const app = express();
app.use(morgan('combined', {skip: req => req.path === '/healthz'}));
app.use(bodyParser.json());

app.get('/healthz', (req, res) => res.send('OK'));

app.get('/', (req, res) => res.send('OK'));

app.post('/', (req, res) =>
  webhookHandler(req)
    .then(() => res.send('OK'))
    .catch(err => {
      console.error(err);
      res.status(500).send('ERROR');
    }));

app.listen(3000);
