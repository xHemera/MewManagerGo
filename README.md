# MewManagerGo
The simplest PokemonTCG collection manager because i'm a simple man


## Installation

```bash
  git clone git@github.com:xHemera/MewManagerGo.git
  cd MewManagerGo
  go build -o mew main.go
  echo "POKEMON_API_KEY=yourapikey" >> .env
```
Hint: Grab you api key here : https://pokemontcg.io/. If you don't want to use the Search function, this step is skippable.

## Usage

```bash
  ./mew
```
It's pretty straigh forward, simply input your action.

Save files are saved at the project root in the saves/ folder; in a mew.xxxxxxx.csv format and can be opened by any other software.
