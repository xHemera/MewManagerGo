# MewManagerGo
The simplest PokemonTCG collection manager because i'm a simple man

## Prerequisites
Install [GoLang](https://go.dev/) and any [NerdFont](https://www.nerdfonts.com/).

## Installation from pre-built binaries
Create an UTF-8 .env and fill it with "POKEMON_API_KEY=yourapikey".
Execute the binary.

## Installation from sources (recommended)

```bash
  git clone https://github.com/xHemera/MewManagerGo.git
  cd MewManagerGo
  go build -o mew main.go
  echo "POKEMON_API_KEY=yourapikey" >> .env
```
Hint: Grab you api key here : https://pokemontcg.io/. Or use mine by asking gently on discord :3

## Usage

```bash
  ./mew
```
It's pretty straigh forward, simply input your action.

Save files are saved at the project root in the saves/ folder; in a mew.xxxxxxx.csv format and can be opened by any other software.
