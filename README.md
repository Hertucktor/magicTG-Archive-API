# magicTGArchive

##Overview - What is this project about?

##How does it work (an abbreviation)
Type in a Magic: The Gathering card name and set name, the Golang backend makes an api request to the fan made mtg api (https://magicthegathering.io/).
The information given by the API REST will be stored in a MongoDB.

##How to run the current project
The project is currently in its cli stage which means that all input needs to be done by cli:
```
go run cmd/main.go --lang <en/de>
```
Then you will be asked for the card name and set name -> enter both after each other e voila your card is stored in the database

##Things to come
- An API to request information from the mongoDB making the card information available per http request in a tidy JSON format
- A UI with options making full use of the CRUD operation of the database; a table like structure in wich the contents of the card archive can be displayed user-friendly heavily oriented on the official Magic: The Gathering hompage : https://magic.wizards.com/en/articles/archive/card-image-gallery/kaldheim


###MTG Fan API Rate Limits
```
Rate limits are enforced for all third-party applications and services. 
This document is an overview of the rate limits themselves, as well as how they are enforced and 
best practices for handling the errors returned when a rate limit is reached.

Third-party applications are currently throttled to 5000 requests per hour. 
As this API continues to age, the rate limits may be updated to provide better performance to users
```
###Configure Credentials for MongoDB in Docker
```
nano init-mongo.js : update credentials and database name to your liking
if you change the db name or user credentials, don't forget to change it in client.go at line 13 and in the docker-compose.yml
```
###Makefile
```
program execution automation with Makefile:
make db -B; to execute the docker-compose.yml to get the mongoDB up and running
make app -B; To test compile and execute the app
```