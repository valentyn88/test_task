###How to build services?

- Each service (consumer, reader) contains Makefile, so you can run make build-mac or make build-linux

###How to run consumer service?

- Consumer service is just http server, so after build you can run ./consumer

###How to run reader service?

- Reader service is command line service that takes 1 param - it's path to csv file, so ./reader test_data/data.csv

###TODO

- If I had enough time, I would use grpc for communication between 2 services
- I would add gometalinter to Makefile
- I would take out user package, because this package is common for both services
- I would add Dockerfile to each services for building application
- I would add docker-compose.yml for running multiple containers
- I would add channel for reading rows from file and send data to consumer service in a few goroutines