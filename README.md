# On local

## Installation ngrok
```
curl -sSL https://ngrok-agent.s3.amazonaws.com/ngrok.asc \
	| sudo tee /etc/apt/trusted.gpg.d/ngrok.asc >/dev/null \
	&& echo "deb https://ngrok-agent.s3.amazonaws.com buster main" \
	| sudo tee /etc/apt/sources.list.d/ngrok.list \
	&& sudo apt update \
	&& sudo apt install ngrok
```
## Configure and run
### Add your authtoken:
Don’t have an authtoken? [Sign up](https://dashboard.ngrok.com/get-started/setup/linux) for a free account.
```
ngrok config add-authtoken <token>
```
## Deploy your app online
```
ngrok http http://localhost:8080
```

## Go to your repo in Github
Settings -> Look at the left sidebar -> Webhooks -> Add Webhook -> Fill information

### Example:
Payload URL: Enter the URL of your server where you want to receive the webhook events. This URL should point to an endpoint in your application that can handle POST requests (e.g., https://example.com/webhooks/github).

Content type: Choose application/json to get the payload data in JSON format.

Secret: (Optional) Add a secret token. This is useful for verifying that requests come from GitHub. You’ll use this token to validate the signature of incoming webhook requests.

Events: Select the event type(s) you want to receive. 
For pull requests, select:
Let me select individual events and then check the Pull requests option.

Active: Ensure this box is checked so that the webhook is active.
