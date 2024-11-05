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
Don’t have an authtoken? Sign up for a free account.

Don’t have an authtoken? [Sign up](https://dashboard.ngrok.com/get-started/setup/linux) for a free account.
```
ngrok config add-authtoken <token>
```
## Deploy your app online
```
ngrok http http://localhost:8080
```
