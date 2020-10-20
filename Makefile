.PHONY: deploy-frontend deploy-api deploy-all delete

deploy-frontend:
	copilot init -a demo -s frontend  -t 'Load Balanced Web Service' -d './frontend/Dockerfile'

deploy-api:
	copilot init -a demo -s api  -t 'Backend Service' -d './api/Dockerfile'

deploy-all:
	copilot svc deploy

delete:
	copilot app delete
