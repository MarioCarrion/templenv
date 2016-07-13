# Docker Dynamic Composer (WIP)

Project for creating dynamic composer files after a sequence of commands.

## Workflow

```
ddc-mc init --name example                                # Initiates the "example" workspace
ddc-mc add --template template1 --container container-123 # Adds "container-123" using the template "template1"
ddc-mc add --template templateA --container container-abc # Adds "container-abc" using the template "templateA"
ddc-mc write                                              # Writes a docker compose file called "example.xml"
```

With that `docker-compose up --file example.yml` can be used
