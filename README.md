# proof-verification-relayer

## Description

This service listens to `RootUpdated` events and saves it to database. 
See [endpoint doc](https://rarimo.github.io/proof-verification-relayer/#tag/State/operation/GetStateV2) to get signed state data.
Once you get it, send signed state to [on-chain replicator](https://github.com/rarimo/passport-contracts/blob/master/contracts/sdk/RegistrationSMTReplicator.sol#L102).

## Install

  ```
  git clone github.com/rarimo/proof-verification-relayer
  cd proof-verification-relayer
  go build main.go
  export KV_VIPER_FILE=./config_transit_state.yaml
  ./main migrate up
  ./main run service
  ```

## Configuration

### Transit state mode

Use [config_transit_state.yaml](config_transit_state.yaml) and set:

- `private_key` of your signer on smart contract
- `address` of your state replica
- `db_url` Postgres database to store state roots

## Documentation

We do use openapi:json standard for API. We use swagger for documenting our API.

To open online documentation, go to [swagger editor](http://localhost:8080/swagger-editor/) here is how you can start it
```
  cd docs
  npm install
  npm start
```
To build documentation use `npm run build` command,
that will create open-api documentation in `web_deploy` folder.

To generate resources for Go models run `./generate.sh` script in root folder.
use `./generate.sh --help` to see all available options.

Note: if you are using Gitlab for building project `docs/spec/paths` folder must not be
empty, otherwise only `Build and Publish` job will be passed.  

## Running from docker 
  
Make sure that docker installed.

use `docker run ` with `-p 8080:80` to expose port 80 to 8080

  ```
  docker build -t github.com/rarimo/proof-verification-relayer .
  docker run -e KV_VIPER_FILE=/config.yaml github.com/rarimo/proof-verification-relayer
  ```

## Running from Source

* Set up environment value with config file path `KV_VIPER_FILE=./config.yaml`
* Provide valid config file
* Launch the service with `migrate up` command to create database schema
* Launch the service with `run service` command


### Database
For services, we do use ***PostgresSQL*** database. 
You can [install it locally](https://www.postgresql.org/download/) or use [docker image](https://hub.docker.com/_/postgres/).


### Third-party services


## Contact

Responsible 
The primary contact for this project is  [//]: # (TODO: place link to your telegram and email)
