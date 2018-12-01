## RxMed
Develop Web Application , Mobile application (ios and android) and necessary APIs, to enable users to order medicine and communicate with doctors. <br />
Project Proposal - [Link](https://docs.google.com/document/d/1gf4INaLJkbKT7FtNATJ1bJ00mcKXdJAwOwt7YhcWVxg/edit?usp=sharing)
<br />
The project will be similar to [carezone.com](https://carezone.com)

## HyperLedger Fabric
## Network Configuration

* 3 Organisations
* 12 Peers
* 1 Certification Authority
* Solo Ordering and will be changed into Kafka on production

##  Prerequisites

This application has been developed on **Ubuntu 16.04** but the Hyperledger Fabric framework is compatible with Mac OS X, Windows and other Linux distributions

Hyperledger Fabric uses **Docker** to easily deploy a blockchain network. In addition, some components (peers) also deploys docker containers to separate data (channel). So make sure that your platform supports this kind of virtualization.

Hyperledger Fabric has been built on **Go** language. In addition, the chaincode (smart contract) is also written in Golang

###  Docker

**Docker version 17.03.0-ce or greater is required.**

#### Linux (Ubuntu)

First of all, in order to install docker correctly we need to install its dependencies:

```bash
sudo apt install apt-transport-https ca-certificates curl software-properties-common
```

Once the dependencies are installed, we can install docker:

```bash
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add - && \
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" && \
sudo apt update && \
sudo apt install -y docker-ce
```
To apply the changes made, you need to logout/login. You can then check your version with:

```bash
docker -v
```
###  Docker Compose

**Docker-compose version 1.8 or greater is required.**

#### Linux

The installation is pretty fast:

```bash
sudo curl -L https://github.com/docker/compose/releases/download/1.21.2/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose && \
sudo chmod +x /usr/local/bin/docker-compose
```

Apply these changes by logout/login and then check its version with:

```bash
docker-compose version
```

###  NodeJS

**Node version 8.10 or greater is required.**

#### Linux

Download NodeJs : https://nodejs.org/en/blog/release/v8.10.0/


### Ensure dependencies

Ensure dependencies by running
```
sudo npm install
```
If error occurs please run

```
sudo npm rebuild
```

## System Requirements 

* Processor : 6 core 12 thread CPU
* Ram : 16 Gb
* HDD : 50 Gb

## Start Network

##### Terminal Window 1

You can start the network by running

```
sudo ./runApp.sh
```

* This lauches the required network on your local machine
* Installs the fabric-client and fabric-ca-client node modules
* And, starts the node app on PORT 4000

## Run Tests

##### Terminal Window 2


In order for the following shell script to properly parse the JSON, you must install ``jq``:

instructions [https://stedolan.github.io/jq/](https://stedolan.github.io/jq/)

With the application started in terminal 1, next, test the APIs by executing the script - **testAPIs.sh**:
```
cd fabric-samples/balance-transfer

## To use golang chaincode execute the following command

./testAPIs.sh -l golang

## OR use node.js chaincode

./testAPIs.sh -l node
```

## Sample REST APIs Requests

### Login Request

* Register and enroll new users in Organization - **Org1**:

`curl -s -X POST http://localhost:4000/users -H "content-type: application/x-www-form-urlencoded" -d 'username=Jim&orgName=Org1'`

**OUTPUT:**

```
{
  "success": true,
  "secret": "RaxhMgevgJcm",
  "message": "Jim enrolled Successfully",
  "token": "<put JSON Web Token here>"
}
```

The response contains the success/failure status, an **enrollment Secret** and a **JSON Web Token (JWT)** that is a required string in the Request Headers for subsequent requests.

### Create Channel request

```
curl -s -X POST \
  http://localhost:4000/channels \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"channelName":"rxmed",
	"channelConfigPath":"../artifacts/channel/channel.tx"
}'
```

Please note that the Header **authorization** must contain the JWT returned from the `POST /users` call

### Join Channel request

```
curl -s -X POST \
  http://localhost:4000/channels/rxmed/peers \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.rxmed.com","peer1.org1.rxmed.com","peer2.org1.rxmed.com","peer3.org1.rxmed.com"]
}'
```

### Update Anchor Peers

```
curl -s -X POST \
  http://localhost:4000/channels/rxmed/anchorpeers \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"configUpdatePath":"../artifacts/channel/Org1MSPanchors.tx"
}'
```


### Install chaincode

```
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d "{
	\"peers\": [\"peer0.org1.rxmed.com\",\"peer1.org1.rxmed.com\",\"peer2.org1.rxmed.com\",\"peer3.org1.rxmed.com\"],
	\"chaincodeName\":\"mycc\",
	\"chaincodePath\":\"$CC_SRC_PATH\",
	\"chaincodeType\": \"$LANGUAGE\",
	\"chaincodeVersion\":\"v0\"
}"
```


### Instantiate chaincode

This is the endorsement policy defined during instantiation.
This policy can be fulfilled when members from both orgs sign the transaction proposal.

```
{
	identities: [{
			role: {
				name: 'member',
				mspId: 'Org1MSP'
			}
		},
		{
			role: {
				name: 'member',
				mspId: 'Org2MSP'
			}
		}
	],
	policy: {
		'2-of': [{
			'signed-by': 0
		}, {
			'signed-by': 1
		}]
	}
}
```

```
curl -s -X POST \
  http://localhost:4000/channels/rxmed/chaincodes \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d "{
	\"chaincodeName\":\"mycc\",
	\"chaincodeVersion\":\"v0\",
	\"chaincodeType\": \"$LANGUAGE\",
	\"fcn\":\"init\"
}"
```
**NOTE:** *chaincodeType* must be set to **node** when node.js chaincode is used

### Invoke request

This invoke request is signed by peers from both orgs, *org1* & *org2*.
```
curl -s -X POST \
  http://localhost:4000/channels/rxmed/chaincodes/mycc \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.rxmed.com","peer0.org2.rxmed.com","peer0.org3.rxmed.com"],
	"fcn":"createDoctor",
	"args":["DOC4","DOCID4","ANAS KOSP", "PMI433", "HOSP4"]
}')
```
**NOTE:** Ensure that you save the Transaction ID from the response in order to pass this string in the subsequent query transactions.

### Chaincode Query on peer0 of Org 1 for Query DOC1

```
curl -s -X GET \
  "http://localhost:4000/channels/rxmed/chaincodes/mycc?peer=peer0.org1.rxmed.com&fcn=query&args=%5B%22DOC1%22%5D" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
```
### Query Chaincode on peer1 of Org1 for Query All Records

```
curl -s -X GET \
  "http://localhost:4000/channels/rxmed/chaincodes/mycc?peer=peer1.org1.rxmed.com&fcn=queryAll" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
```
### Query Chaincode on peer1 of Org1 for Query History of DOC0

```
curl -s -X GET \
  "http://localhost:4000/channels/rxmed/chaincodes/mycc?peer=peer1.org1.rxmed.com&fcn=queryHistory&args=%5B%22DOC0%22%5D" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
```

### Query Chaincode on peer0 of Org1 and 2 to Delete DOC2

```
curl -s -X POST \
  http://localhost:4000/channels/rxmed/chaincodes/mycc \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.rxmed.com","peer0.org2.rxmed.com"],
	"fcn":"delete",
	"args":["DOC2"]
}')
```

### Query Chaincode on peer1 of Org1 for Query History of DOC0

```
curl -s -X GET \
  "http://localhost:4000/channels/rxmed/chaincodes/mycc?peer=peer1.org1.rxmed.com&fcn=queryHistory&args=%5B%22DOC2%22%5D" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
```

### Query Block by BlockNumber

```
curl -s -X GET \
  "http://localhost:4000/channels/rxmed/blocks/1?peer=peer0.org1.rxmed.com" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
```

### Query Transaction by TransactionID

```
curl -s -X GET \
  "http://localhost:4000/channels/rxmed/transactions/$TRX_ID?peer=peer0.org1.rxmed.com" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
```

### Query ChainInfo

```
curl -s -X GET \
  "http://localhost:4000/channels/rxmed?peer=peer0.org1.rxmed.com" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
```

### Query Installed chaincodes

```
curl -s -X GET \
  "http://localhost:4000/chaincodes?peer=peer0.org1.rxmed.com" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
```

### Query Instantiated chaincodes

```
curl -s -X GET \
  "http://localhost:4000/channels/rxmed/chaincodes?peer=peer0.org1.rxmed.com" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
```

### Query Channels

```
curl -s -X GET \
  "http://localhost:4000/channels/rxmed/chaincodes?peer=peer0.org1.rxmed.com" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
```

This work is licensed under NetObjex Inc 2018.






