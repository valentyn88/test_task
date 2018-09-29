###How to build services?

- To build consumer service run: make build-mac SERVICE=consumer

- To build reader service run: make build-mac SERVICE=reader

###How to run consumer service?

- After build you can run ./consumer

###How to run reader service?

- Reader service is command line service that takes 1 param - it's path to csv file, so ./reader test_data/data.csv

###How to run test

- run: make test

###TODO

- I would add gometalinter to Makefile
- I would add Dockerfile to each services for building application
- I would add docker-compose.yml for running multiple containers
