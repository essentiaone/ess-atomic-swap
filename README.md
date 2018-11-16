# Ess-atomic-swap
Service to cross-chain transactions between two or more people

[What is Essentia atomic swap?](#what)

[Why Essentia atomic swap?](#why)

[Console commands](#console)

[Run the server](#run)


<a name="what"><h2>What is Essentia atomic swap?</h2></a>

Essentia atomic swap is presented as an independent service that is able to serve from two users. It specifies the path to the trusted nodes, which are specified manually, through which the cryptocurrency exchange will take place. The ethereum smart contract is also used as an immutable database in which pools of confirmed and non-confirmed swaps are stored. For greater security, a smart contract can be redeployed to the ethereum network by the user.

<a name="why"><h2>Why Essentia atomic swap?</h2></a>

When you want to exchange your cryptocurrency with someone without third party services, or want to use this blockchain technology in your solution but you cant trust anyone - Essentia atomic swap is the right choice!

Essentia atomic swap does the following for you:

1. Create orders for atomic swap
2. View inits of other atomic swaps
3. Audit inits of other atomic swaps
4. Cancel your init to atomic swap
5. Create many orders to exchange
6. In case you don't want to launch your nodes, you can use default nodes
7. Can be used as ready enterprise solution

<a name="console"><h2>Console commands</h2></a>

Examples of the commands:

```go
init(_addressFrom, _addressTo, 1, BTC, "S0me $trong p@ssw0rd!")
confirmInit(_addressFrom, "S0me $trong p@ssw0rd!", _transactionHash)
auditInit(_addressOfInitiator)
redeem(_addressOfInitiator)
```

<a name="run"><h2>Run the server with the folowing commands:</h2></a>

1. Set `ENV` variables from `.env` file
	```
	$ export PORT=1604
	$ export DEV_PORT=1605
	$ export BTC_ADDR=https://address_to_btc_node
	$ export ETH_ADDR=https://address_to_eth_node
	$ export SCDB_ADDR=https://address_to_scdb (smart contract)
	$ export PRIVATE_KEY=your_private_key_for_ETH_ADDR_node
	$ export PRODUCTION=""
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
