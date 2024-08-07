# SNS-Test By Thepsirin

- This service struct in 3 layers(handlers, services, repository)
- In This API service you can
   - Create, Read, Update, Delete user and authenticate
   - Create Friend, View all friend Request, List all friend, Update friend status and Delete friend (In this feature we used redis to store the friend request and list of friend)
   - Create Post, View all user post, Read specific post, Update and Delete post
Add like to post and delete like
- This service has provided **gracefully shutdown and use Echo V4, Gorm, Maria DB, Redis** and etc.
- For The future we will created API for managing Comment and cached high frequency usage function and will separate the function to new service for scalability and try to use websocket for Post comment and like.
- For API Document is in the Postman Collection document

## Prerequisites Software
1. Golang **1.22.3**
1. Docker and Docker compose
1. Postman

## Usage step
0. Clone and open this repo
1. Compose the docker-compose.yaml file (docker-compose up -d) for creating Redis and Maria DB
```bash
make up
```
2. Run the service (go run .)
```bash
make run
```

3. Migrate user data (Create 10 users for use in the service (default password is 12345678))
```bash
make migrate-user
```

## Unit test
- Unit test for service and handler
```bash
make test-cover
```

- Unit test for service and handler (Full test function)
```bash
make test-full-cover
```


## License

[MIT](https://choosealicense.com/licenses/mit/)