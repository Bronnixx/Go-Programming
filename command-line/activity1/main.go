package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	choices = []string{"File input", "Type in input"}
)
type model struct{
	cursor int
	choice string
}

func main(){
	p := tea.NewProgram(model{})
	m, err := p.Run()
	if err != nil{
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	if m, ok := m.(model); ok && m.choice != "" {
		fmt.Printf("\n---\nYou chose %s!\n", m.chose)
	}
	if m, ok := m.(model); ok && m.choice != "" && m.choice == "Type in input"{
		processStdin()
	}
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type){
		case "ctrl+c", "q", "esc":
		return m, tea.Quit
		case "enter":
		m.choice = choices[m.cursor]
		return m, tea.Quit
		case "down", "j":
		m.cursor++
		if m.cursor >= len(choices){
			m.cursor = 0
		}
		case "up", "k":
		m.cursor--
		if m.cursor == len(choices){
			m.cursor = len(choices) - 1
		}
	}
	return m, nil
}
//TUI view
func (m model) View() string{
	s:= strings.Builder{}
	s.WriteString("Select if you would like to work with file input or type in input:\n\n")
	for i := 0; i < len(choices); i++{
		if m.cursor == i {
			s.WriteString("(•) ")
		}else{
			s.WriteString("()")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")
	return s.String()
}
//encoding func
func rot13(s string)string{
	result := make([]byte, len(s))
	for i:= 0; i < len(s); i++{
		char := s[i]
		switch {
			case char >= 'a' && char <= 'z':
			result[i] = 'a' + (char-'a' + 13)%26
			default:
			result[i] = char
		}
	}
	return string(result)
}
//read data from stdin
func processStdin(){
	reader := bufio.NewReadWriter(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}else if err != nil{
			fmt.Println("Error reading stdin:", err)
			return
		}
		encoded := rot13(input)
		fmt.Print(encoded)
	}
}
//file input
func processFile(filename string){
var inputReader io.Reader
file, err := os.Open(filename)
if err != nil{
	fmt.Println("Error opening file:", err)
	return
}
defer file.Close()
inputReader = file
//process input and apply rot13
scanner := bufio.NewScanner(inputReader)
for scanner.Scan(){
	encoded := rot13(scanner.Text())
		fmt.Println(encoded)
}
if err := scanner.Err(); err != nil{
	fmt.Println("Error reading input:", err)
}
}
