# Mortgage Application

This Mortgage Application is written in Go(for Backend) and React(for Frontend)
<br>
The repo contains some folders as well as files, I will only talk about the important ones:
```
- bank folder contains one file called bank.go which contains all the logic for Bank (creating, updating, edit, deleting etc)
- client folder contains all my frontend codes in react
- db folder is where all database connection happened and the database used is Mongodb
- Makefile contains few codes which runs my Go server by doing `make serve`.
```
Below are the guidelines: <br>
- Clone the the repo by doing `git clone github.com/Ad3bay0c/MortgageApp`
- At the root, type `go mod tidy` to download all the packages used
- Then, run the server by doing  `make serve`, this will go into Makefile file at the root and execute the line containing `serve`
- cd into client by doing `cd client/` and run `npm install` to install all the dependencies used for the client
-Then, run `npm start` to start the react app in your default browser.
- Now you're good to go

Check your browser, now you should see the mortgage app running, if you see any error, make sure you follow the guidelines above properly and check if the go server is running. you can confirm this if after doing `make serve` you see `server sarted at localhost:5000` and a `database connected` log at your terminal.
```
The First Page which is the Home Page contains a form where you can create bank and can edit, update and delete.
Check the Navbar, you will see a mortgage Calculator that calculates monthly payment based on the bank selected. Though,I'm not yet done with this part.
```

`gorilla/mux` is used as the router package, Golang is used to write the API and React for the frontend