package main


import (
  "bufio"
  "fmt"
  "os"
  "strings"
)


func startRepl() {

  // Retrieving a pointer of an insatnce of Scanner struct
  reader := bufio.NewScanner(os.Stdin)

  // This outer for loop is for REPL mechanism 
  for {
    fmt.Print("pokedex> ")

    // Scan fuction read given command from console
    reader.Scan()

    // This Err method return the first error in the event of reading by Scan mathod
    if err := reader.Err(); err != nil {
      fmt.Fprintln(os.Stderr, "This error occured:", err)
      continue
    }

    words := cleanInput(reader.Text()) 
    if len(words) == 0 {
      continue
    }

    commandName := words[0]

    // Checking if given command exist as key in commandMap
    if command, commandExists := getCommands()[commandName]; !commandExists {
      fmt.Println("invalid command")
    } else {
      // If given command is valid command(exist in commandMap) then invoke it's callback fuction
      err := command.callback()
      if err != nil {
        fmt.Printf("failed to execute command %s: %v", commandName, err)
      }
    }
  }

}


func cleanInput(input string) []string {
  output := strings.ToLower(input)
  words := strings.Fields(output)
  return words
}
