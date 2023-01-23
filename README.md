<p align="center">
<br />
<h1 align="center">NFTir</h1>
<h5 align="center">SUNY OSwego </h3>
<h6 align="center">CSC 482 <h4>
</p>

###### Currently work in progress

## Overview
NFTir agent is a golang server is engineered to periodically query the NFTGo server, with a 6 hours frequency, to acquire a dataset of the top 25 NFTs ranked by trading volume. The agent then breaks down the dataset into single items and push them into an Amazon DynamoDB table which will be used by the [NFTir RESTful server](https://github.com/NFTir/server) to process the data.

## Getting Started

### Requirement

- [git](https://git-scm.com/)
- [golang](https://go.dev/)
- [docker](https://www.docker.com/)
- [aws-cli](https://aws.amazon.com/cli/)

### Clone the repo

```bash
git clone https://github.com/NFTir/agent.git
cd agent
```

### Set up environment variables

At the root of the directory, create a .env file using .env.example as the template and fill out the variables.

### Running the project

1. Build and run `agent` locally using `Make` scripts
```bash
make go-build-local
```

2. Build and run `agent` on Docker using `Make` scripts

```bash
make build-app
```

### Resources

- [NFTGo.io](https://nftgo.io/)
- [Amazon DynamoDB](https://aws.amazon.com/dynamodb/)
- [AWS-CLI](https://aws.amazon.com/cli/)
- [Golang](https://go.dev)
