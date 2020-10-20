# App Mesh Demo

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
copilot svc deploy
```

## Cleanup
```bash
copilot app delete
```
