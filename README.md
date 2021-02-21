# magicTGArchive
##Creators:
API | Database | Backend | Frontend :   https://github.com/Hertucktor & Android App: https://github.com/HimeUrikohime

##Overview - What is this project about?
This Project will make storing digital copies of your favorite Magic: The Gathering cards as easy as opening a new booster pack.
You will be able to store your own collection of mtg cards and build decks with these cards and share your collection information with your friends.

##How does it work (an abbreviation)
Using http request to the REST API you will be able to manipulate your own mtg collection (for further API Usage Info read docs/api)
The Frontend will allow you to manage your collection in a user friendly way
The Android App helps you take pictures of your cards and identify your cards and therefore you will be faster with 
inserting cards into your collection.

##How to run the current project
```
run docker localy (you will need docker desktop : https://www.docker.com/products/docker-desktop )
make db

after your db is up and running execute the card importer (this will take up to an hour)
make importer

finally your database is running and filled with all needed information to start the main event
make api

now use the api as described in the api doc and have fun
```

###Configure Credentials for MongoDB in Docker
If you don't want to run your app with the default credentials change them like so:
```
nano init-mongo.js : update credentials and database name to your liking
if you change the db name or user credentials, don't forget to change it in client.go at line 13 and in the docker-compose.yml
```