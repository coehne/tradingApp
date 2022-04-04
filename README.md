# Trading App
This is my final project for CS50x course @Harvard University.

It is inspired by the C$50 Trading app (https://finance.cs50.net/) which was part of one assignment during the course. You can find the current main branch deployed at https://tradapp.uber.space

I wanted to take this as an opportunity to familiarize myself with Golang, several packages and the clean architecture approach by Robert C. Martin.


## Project Scope

- [x] Build a go REST API with fiber and gorm in MVC architecture
- [x] Build an integration to consume stock market information from a private API
- [x] Use JWT cookie authentication
- [x] Refactoring in to clean architecture
- [x] Build a react frontend
- [x] Use tailwind CSS
- [x] Use Github actions for CI/CD and autodeployment to uberspace
- [ ] Write tests in Go and React

## Dependencies
The project is build on the following dependencies
- Go 1.18
- PostgreSQL 13.6
- Node.js 16

The repo has a .devcontainer which makes it easy to start the dev environment in VS Code if you have docker installed on your system.
It includes Go version 1.18, PostgreSQL 13.6 and Node 16 to start up the backend, db and the frontend.

In order to use it you need the [Remote Container VS Code extension by Microsoft.](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) Once you open the repo in VS Code you will get asked if you want to reopen it in a dev container. If you won't get asked, you can do it manually by clicking in the bottom left corner of vs code and then on "Reopen in Container".

If you don't have VS Code or docker installed, you can install the dependencies manually.

Once you have opened the dev container or installed the required dependencies, you can continue by starting the backend and frontend.

### Environment file
This project uses IEX Cloud for getting current stock market infomration. In order to use it, you need to register a free account at https://iexcloud.io/ and get yourself an API key (starting with pl_*).

Once you have the API key you can rename the .app.env file to app.env in the /api folder and insert it into the "IEXCLOUD_API_KEY" variable.

## Starting The Dev Environment

### Backend
To start the backend, cd into the /api folder and simply run `go run main.go`

### Frontend
To start the frontend dev server, cd into the fe folder and run `npm start`

### API Testing with insomnia
In the .insomnia folder you can also find an environment and request collection to test the api without the react frontend.

## Usage Of The App

After registration you have to deposit some money into your trading account. Once done, you can start buying Stonks by clicking on Buy in your depot / home screen. The app will get the latest price from the IEX Cloud API for you and calculate the total amount which will be deducted from your account when you submit the trade. 

The trade also creates a corresponding transaction in your transaction log where you can revisit the details of the trade again. This is also possible on your trading history page. 

Your depot page / home page calculates the current value of your positions, using the latest prices from the IEX Cloud API. It also shows you how much the total value of your account (cash) + your stonks are.

If you sell stonks for a higher price than you bought it, the corresponding profit will be added to your cash position which you can later withdraw again.
