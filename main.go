package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
    "strings"

    "github.com/fatih/color"
)

type Card struct {
    UID    int
    Name   string
    Series string
    Number string
    State  string
}

var collection []Card
var nextUID int

func main() {
    color.Set(color.FgWhite)
    defer color.Unset()

    scanner := bufio.NewScanner(os.Stdin)
    nextUID = 1 // Initialisation de l'UID à 1

    for {
        clearScreen()
		displayTitle()
        displayMenu()
        scanner.Scan()
        option := scanner.Text()

        switch option {
        case "1":
            addCard(scanner)
        case "2":
            removeCard(scanner)
        case "3":
            displayCollection()
        case "4":
            searchCard(scanner)
        case "5":
            saveCollection(scanner)
        case "6":
            loadCollection(scanner)
        case "7":
            fmt.Println("Quitter...")
            return
        default:
            fmt.Println("Option invalide. Essayez encore.")
        }
    }
}

func displayTitle() {
    fmt.Print("\033[38;2;250;115;227m") // Utilisation du rose vif
    
    fmt.Println(`
 __ __              __ __                               
|  \  \ ___  _ _ _ |  \  \ ___ ._ _  ___  ___  ___  _ _ 
|     |/ ._>| | | ||     |<_> || ' |<_> |/ . |/ ._>| '_>
|_|_|_|\___.|__/_/ |_|_|_|<___||_|_|<___|\_. |\___.|_|  
                                         <___'          
`)

    fmt.Print("\033[0m")
}

func displayMenu() {
    color.New(color.FgGreen).Println("1.\t  Ajouter une carte")
    color.New(color.FgRed).Println("2.\t  Supprimer une carte")
    color.New(color.FgCyan).Println("3.\t  Afficher la collection")
    color.New(color.FgMagenta).Println("4.\t  Chercher une carte (WIP)")
    color.New(color.FgBlue).Println("5.\t  Sauvegarder")
    color.New(color.FgYellow).Println("6.\t 󱃭 Charger")
    color.New(color.FgWhite).Println("7.\t 󰌑 Quitter")
    fmt.Print("Choisissez une option: ")
}

func clearScreen() {
    fmt.Print("\033[H\033[2J")
}

func addCard(scanner *bufio.Scanner) {
    fmt.Println("Ajouter une carte:")
    card := Card{}
    card.UID = nextUID // Utilisation du compteur pour l'UID
    nextUID++          // Incrémentation du compteur

    fmt.Print("Nom: ")
    scanner.Scan()
    card.Name = scanner.Text()

    fmt.Print("Série: ")
    scanner.Scan()
    card.Series = scanner.Text()

    fmt.Print("Numéro (xxx/xxx): ")
    scanner.Scan()
    card.Number = scanner.Text()

    fmt.Print("État (MT, NM, EX, GD, LP, PL, PO): ")
    scanner.Scan()
    card.State = scanner.Text()

    collection = append(collection, card)
    fmt.Println("Carte ajoutée. Appuyez sur une touche pour continuer.")
    scanner.Scan()
}

func removeCard(scanner *bufio.Scanner) {
    fmt.Println("Supprimer une carte:")
    fmt.Print("UID de la carte à supprimer: ")
    scanner.Scan()
    uidStr := scanner.Text()
    uid, err := strconv.Atoi(uidStr)
    if err != nil {
        fmt.Println("UID invalide. Appuyez sur une touche pour continuer.")
        scanner.Scan()
        return
    }

    for i, card := range collection {
        if card.UID == uid {
            collection = append(collection[:i], collection[i+1:]...)
            fmt.Println("Carte supprimée. Appuyez sur une touche pour continuer.")
            scanner.Scan()
            return
        }
    }
    fmt.Println("Carte non trouvée. Appuyez sur une touche pour continuer.")
    scanner.Scan()
}

func displayCollection() {
    color.New(color.FgCyan).Println("Collection:")
    for _, card := range collection {
        colorFunc := getColorFuncForState(card.State)
        numParts := strings.Split(card.Number, "/")
        numColor := color.New(color.FgWhite).SprintfFunc()
        if len(numParts) == 2 {
            num1, _ := strconv.Atoi(numParts[0])
            num2, _ := strconv.Atoi(numParts[1])
            if num1 > num2 {
                numColor = color.New(color.FgYellow).SprintfFunc() // Numéro en jaune
            } else {
                numColor = color.New(color.FgWhite).SprintfFunc() // Numéro en blanc
            }
        }

        // Affichage formaté
        fmt.Printf("┌ %s\t%s %s\n", color.New(color.FgWhite).Sprintf(card.Name), colorFunc("󰓹"), colorFunc(getStateDescription(card.State)))
        fmt.Printf("└ UID: %s\tExtension: %s\tNumero: %s\n\n",
            color.New(color.FgHiBlack).Sprintf("%d", card.UID),
            color.New(color.FgHiBlack).Sprintf(card.Series),
            numColor(card.Number))
    }
    fmt.Println("Appuyez sur une touche pour continuer.")
    var input string
    fmt.Scanln(&input)
}

func getColorFuncForState(state string) func(format string, args ...interface{}) string {
    switch state {
    case "MT":
        return color.New(color.FgCyan).SprintfFunc() // Parfait
    case "NM":
        return color.New(color.FgGreen).SprintfFunc() // Près du parfait
    case "EX":
        return color.New(color.FgGreen).SprintfFunc() // Excellent
    case "GD":
        return color.New(color.FgYellow).SprintfFunc() // Bon
    case "LP":
        return color.New(color.FgYellow).SprintfFunc() // Moyen
    case "PL":
        return color.New(color.FgRed).SprintfFunc() // Peu mieux faire
    case "PO":
        return color.New(color.FgHiRed).SprintfFunc() // Horrible
    default:
        return color.New(color.FgWhite).SprintfFunc() // Non défini
    }
}

func getStateDescription(state string) string {
    switch state {
    case "MT":
        return "MT"
    case "NM":
        return "NM"
    case "EX":
        return "EX"
    case "GD":
        return "GD"
    case "LP":
        return "LP"
    case "PL":
        return "PL"
    case "PO":
        return "PO"
    default:
        return "Inconnu"
    }
}

func saveCollection(scanner *bufio.Scanner) {
    fmt.Print("Nom du fichier de sauvegarde: ")
    scanner.Scan()
    filename := scanner.Text()

    if err := os.MkdirAll("saves", os.ModePerm); err != nil {
        fmt.Println("Erreur lors de la création du dossier de sauvegarde.")
        return
    }

    file, err := os.Create(fmt.Sprintf("saves/mew.%s.csv", filename))
    if err != nil {
        fmt.Println("Erreur lors de la création du fichier de sauvegarde.")
        return
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, card := range collection {
        err := writer.Write([]string{
            strconv.Itoa(card.UID),
            card.Name,
            card.Series,
            card.Number,
            card.State,
        })
        if err != nil {
            fmt.Println("Erreur lors de l'écriture dans le fichier de sauvegarde.")
            return
        }
    }

    fmt.Println("Sauvegarde terminée. Appuyez sur une touche pour continuer.")
    scanner.Scan()
}

func loadCollection(scanner *bufio.Scanner) {
    files, err := os.ReadDir("saves")
    if err != nil {
        fmt.Println("Erreur lors de la lecture du dossier de sauvegarde.")
        return
    }

    fmt.Println("Fichiers de sauvegarde:")
    for i, file := range files {
        if file.IsDir() {
            continue
        }
        fmt.Printf("%d. %s\n", i+1, file.Name())
    }

    fmt.Print("Choisissez un fichier: ")
    scanner.Scan()
    choice := scanner.Text()
    index, err := strconv.Atoi(choice)
    if err != nil || index < 1 || index > len(files) {
        fmt.Println("Choix invalide.")
        return
    }

    file := files[index-1]
    f, err := os.Open(fmt.Sprintf("saves/%s", file.Name()))
    if err != nil {
        fmt.Println("Erreur lors de l'ouverture du fichier.")
        return
    }
    defer f.Close()

    reader := csv.NewReader(f)
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Erreur lors de la lecture du fichier.")
        return
    }

    collection = nil
    for _, record := range records {
        uid, _ := strconv.Atoi(record[0])
        collection = append(collection, Card{
            UID:    uid,
            Name:   record[1],
            Series: record[2],
            Number: record[3],
            State:  record[4],
        })
    }

    fmt.Println("Chargement terminé. Appuyez sur une touche pour continuer.")
    scanner.Scan()
}

func searchCard(scanner *bufio.Scanner) {
    fmt.Print("Nom de la carte à chercher: ")
    scanner.Scan()
    searchName := scanner.Text()

    fmt.Printf("Résultats pour '%s':\n", searchName)
    for _, card := range collection {
        if strings.Contains(strings.ToLower(card.Name), strings.ToLower(searchName)) {
            fmt.Printf("UID: %d, Nom: %s, Série: %s, Numéro: %s, État: %s\n",
                card.UID, card.Name, card.Series, card.Number, card.State)
        }
    }
    fmt.Println("Recherche terminée. Appuyez sur une touche pour continuer.")
    scanner.Scan()
}

