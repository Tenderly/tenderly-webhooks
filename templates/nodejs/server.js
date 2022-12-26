// server.js
//
// Use this sample code to handle webhook events
// 
// Run the server on http://localhost:8008
// npm start

const express = require('express');
const crypto = require('crypto');
const app = express();

// replace this with key taken from dashboard for specific webhook
// if you are testing webhook before creation set it on empty string
const signingKey = '';

//set req method get to verify that a webhook endpoint exists.
app.get('/webhook', (request, response) => {
  response.send();
})

app.post('/webhook', express.raw({type: 'application/json'}), (request, response) => {
  const signature = request.headers['x-tenderly-signature'];
  const timestamp = request.headers['date'];
  
  if (!isValidSignature(signature, request.body, timestamp)) {
		console.log('Error signature not valid');
    response.status(401).send('Error signature not valid');
    return;
  }

  let body;
  try {
    body = JSON.parse(request.body.toString())
  } catch(e) {
    response.status(400).send(`Webhook error: ${e}`);
    return;
  }
  const eventType = body.event_type
  switch (eventType) {
    case 'TEST':
      // Then define and call a function to handle the test event
      break;
    case 'ALERT':
    	// Then define and call a function to handle the alert event
      break;
    // ... handle other event types
    default:
      console.log(`Unhandled event type ${eventType}`);
  }

  // Return a 200 response
  response.send();
});

const port = 8008;
app.listen(port, () => console.log(`Running on port ${port}`));

function isValidSignature(signature, body, timestamp) {
  const hmac = crypto.createHmac("sha256", signingKey); // Create a HMAC SHA256 hash using the signing key
  hmac.update(body.toString(), 'utf8'); // Update the hash with the request body using utf8
  hmac.update(timestamp) // Update the hash with the request timestamp
  const digest = hmac.digest("hex");
  return crypto.timingSafeEqual(Buffer.from(signature), Buffer.from(digest))
}