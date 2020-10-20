# Copilot Demo

This repo contains a project of two microservices which can be extended to multiple microservices.

The frontend is a simple nodejs application which calls api for dynamic giphy images. The backend is a golang application which calls giphy API to retrieve images and responds back with previous history of giphy responses.

## Requirements

- Install copilot

```bash
curl -Lo $HOME/.local/bin/copilot https://github.com/aws/copilot-cli/releases/download/v0.4.0/copilot-$(echo `uname -s` | tr '[A-Z]' '[a-z]')-v0.4.0 \
&& chmod +x $HOME/.local/bin/copilot \
&& copilot --help
```

## Deployment

Deploy the frontend

```bash
make deploy-frontend
```

Deploy the api

```bash
make deploy-api
```

## Cleanup
```bash
copilot app delete
OR
make delete
```
