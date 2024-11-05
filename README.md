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

My favorite search engine is [Duck Duck Go](https://duckduckgo.com "The best search engine for privacy").
```
ngrok config add-authtoken <token>
```
