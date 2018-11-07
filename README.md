# Ess-atomic-swap
Service to cross-chain transactions between two or more people

## Run the server with the folowing commands:

1. Set `ENV` variables from `.env` file"
	```
	$ export PORT=1604
	$ export DEV_PORT=1605
	$ export BTC_ADDR=https://address_to_btc_node
	$ export ETH_ADDR=https://address_to_eth_node
	$ export SCDB_ADDR=https://address_to_scdb (smart contract)
	$ export PRIVATE_KEY=your_private_key_for_ETH_ADDR_node
	...
	```

2. For `development` you should use:
	```
	$ docker-compose -f docker-development-compose.yml build
	$ docker-compose -f docker-development-compose.yml up
	```

3. For `production` you should use"
	```
	$ docker-compose -f docker-production-compose.yml build
	$ docker-compose -f docker-production-compose.yml up
	```
