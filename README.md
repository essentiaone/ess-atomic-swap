# Ess-atomic-swap
Service to cross-chain transactions between two or more people

## Run the server with the folowing commands:

Export variables from `.env` file:
```
export ESS_ATOMIC_SWAP_PORT=your_port
```

Use commant from `Makefile`:

1. Default command for `build` and `start` project:
	```
	$ make
	```

2. Command for only `build` project:
	```
	$ make build
	```

3. Command for only `start` project:
	```
	$ make start
	```

## Run `docker` container for project:

Export variables from `.env` file:
```
export ESS_ATOMIC_SWAP_PORT=your_port
```

Use commant for start `docker`:

1. Build `docker` container:
	```
	$ docker-compose -f docker-development-compose.yml build
	```

2. Run `docker` container:
	```
	$ docker-compose -f docker-development-compose.yml up
	```
