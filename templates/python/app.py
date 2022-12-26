# app.py
#
# Use this sample code to handle webhook events
#
#  Install dependencies
#   pip3 install flask
#
#  Run the server on http://localhost:8008
#   python3 -m flask run --port=8008

import hmac
import hashlib
import string

from flask import Flask, jsonify, request

#  replace this with key taken from dashboard for specific webhook
#  if you are testing webhook before creation set it on empty string
signingKey = '';
app = Flask(__name__)

@app.route('/webhook', methods=['GET', 'POST'])
def webhook():
    if request.method == 'GET':
       return jsonify(success=True)
    
    body = request.data
    signature = request.headers['X-Tenderly-Signature']
    timestamp = request.headers['Date']

    if (isValidSignature(signature, body, timestamp) == False):
      print('Error signature not valid')
      return jsonify(success=False)
    
    payload = request.json
    eventType = payload.get('event_type')
    # Handle the event
    if eventType == 'TEST':
      # Then define and call a function to handle the test event
      print('Event type {}'.format(eventType))
    elif eventType == 'ALERT':
      # Then define and call a function to handle the alert event
      print('Event type {}'.format(eventType))
      # ... handle other event types
    else:
      print('Unhandled event type {}'.format(eventType))

    return jsonify(success=True)

def isValidSignature(signature: string, body: bytes, timestamp: string):
    h = hmac.new(str.encode(signingKey), body, hashlib.sha256)
    h.update(str.encode(timestamp))
    digest = h.hexdigest()
    return hmac.compare_digest(signature, digest)