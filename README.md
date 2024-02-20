# dogo - Docker Desktop alternative

`dogo` is a simple cli tool that allows you to manage your docker containers and images from a web interface. It runs in a browser window, and makes use of the Docker SDK for Golang for interacting with the Docker API. The web interface is built using HTMX and tailwindcss.

## Running dogo

If you are a WSL2 Docker user like me, run the SetUnixEnvGolang.ps1 script to set the golang compiler to target linux. Ater this running the Run.ps1 script will compile and run the application server.

## TODOS

- [ ] List all containers and images
- [ ] Start, stop, and remove containers
- [ ] Remove images
- [ ] View container logs
- [ ] View container stats
- [ ] View and remove volumes
- [ ] View and remove networks
- [ ] View and remove secrets
- [ ] View and remove configs
