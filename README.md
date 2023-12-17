# Build-Pokedex

This is a Read Evaluate Print Loop (REPL) program that uses PokeAPI to fetch information based on user input. I have been learning the basics of Golang and wanted to create a project that would challenge me and also use http requests. I decided to go with this project because I have fond memories of watching and playing Pokemon while I was younger and came upon the PokeAPI recently. After looking through the documentation, I came up with the idea to create a Pokedex that will help users explore areas in Pokemon, catch Pokemon, and let the user learn the stats of the pokemon in the user's Pokedex.

## Setup

### Installing
* Clone this repository: https://github.com/ejosborn/Build-Pokedex
* This project uses Go version 1.21.4, other versions can be used as long as it is after 1.20

### Running the Program
* First navigate to the project folder and run
```bash
go build
```
* This will create an executable so you can run the program

* Next in the terminal, type
```bash
./Build-Pokedex
```

## Usage

### Commands
* `help` - This will output all of the commands that you can use in this 
* `map` - This command will list the next page of locations a user can go to
* `mapb` - This command will list the previous page of locations a user can go to
* `explore [insert-map-area]` - This command will list all of the pokemon found at the inserted map area
* `catch [insert pokemon-name]` - This command will try to catch a pokemon. The more baseExp a pokemon has, the harder it is to catch
* `inspect [pokemon-name]` - This command will check the pokedex for the pokemon name and will output its base stats
* `pokedex` - This command will list all of the pokemon that you have caught and are listed inside your pokedex

## Contributing
### Submit a pull request
If you'd like to contribute, please fork the repository and open a pull request to the `main` branch

## Author
* **Ellis Osborn** - [Ejosborn](https://github.com/ejosborn)
