# tenderly-webhooks
Webhook setups in golang/nodejs/python

Steps to receive event notifications via webhook

You can start receiving event notifications in your app using the steps in this section:

1. Identify the events (currently only alert events) you want to monitor.
2. Create a webhook endpoint as an HTTP endpoint (URL) on your local server.
3. Handle webhook requests from Tenderly and return **`2xx`** response status codes.
4. Test that your webhook endpoint is working properly using the [ngrok](https://ngrok.com/).
5. Deploy your webhook endpoint so it’s a publicly accessible HTTPS URL.
6. Register your publicly accessible HTTPS URL in the Tenderly dashboard (currently only you can only register webhooks as delivery channel for alerts). Note here, before create of a webhook we will send GET request to same url to check existence. Please be sure to set get method with immediately return success. Also timeout for webhook POST request is 5 seconds, so please be sure to return success within 5 seconds window.