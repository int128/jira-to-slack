const express = require('express');
const bodyParser = require('body-parser');
const morgan = require('morgan');
const webhookHandler = require('./webhook_handler');

const app = express();
app.use(morgan());
app.use(bodyParser.json());

app.get('/', (req, res) => res.send('OK'));
app.get('/healthz', (req, res) => res.send('OK'));

app.post('/webhook', (req, res) =>
  webhookHandler(req)
    .then(value => res.send(value))
    .catch(err => {
      console.error(err);
      res.status(500).send("ERROR");
    }));

app.listen(3000);
