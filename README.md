# Crypto Price Alert

Application diagram:
![](diagram.png "Diagram")

I opted out to overengineering this project in order to demonstrate some concepts and technologies. In order to do that, I followed some good practices, such as: Clean Architecture and SOLID principles 

# Assumptions

- One currency will be supported: USD
- Three coins will be supported: BTC, ETH, BNB
- Each job execution will notify the user about the price change

## Commands

- for run application asap you can use docker-compose, as the images are already builded and uploaded to docker hub:
```shell
    make docker_compose
```

- for build the application:
```shell
    make docker_build
```

## TODO's

- Implement DLQ (Dead Letter Queue)
- Finish graceful shutdown implementation
- Implement better test coverage
- To be production ready, must follow 12-factor app, such as:
  - separate codebase (the three projects)
  - create script for database migration
  - change the price alert job to be as a one-time execution, using kubernetes job scheduler as an example