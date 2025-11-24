package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

// Define the LLM model name globally (can be configured via env var in real app)
const LLMModelName = "llama3"

// Data structure to pass to the HTML template
type PageData struct {
	LLMName     string
	PromptTitle string
	Response    string
	Error       string
}

// Function to truncate a string to the first N words. Used in HTML template.
func truncateWords(text string, n int) string {
	words := strings.Fields(text)
	if len(words) <= n {
		return text
	}
	return strings.Join(words[:n], " ") + "..."
}

// Parse the template files and register the custom function
var tmpl = template.Must(template.New("index.html").Funcs(template.FuncMap{
	"truncate": truncateWords,
}).ParseFiles("templates/index.html"))

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ask", askHandler)
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Pass the model name even on the initial load
	data := PageData{LLMName: LLMModelName}
	err := tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func askHandler(w http.ResponseWriter, r *http.Request) {
	// ... (method check and prompt check) ...
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusMethodNotAllowed)
		return
	}

	prompt := r.FormValue("prompt")
	if prompt == "" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}

	// Call Ollama via Command Line
	// ... (cmd execution logic) ...
	cmd := exec.Command("ollama", "run", LLMModelName, prompt)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		log.Printf("Command failed with %s\nStderr: %s\n", err, stderr.String())
		data := PageData{
			LLMName:     LLMModelName,
			PromptTitle: prompt,
			Error:       "Failed to communicate with Ollama server. Error: " + stderr.String(),
		}
		tmpl.ExecuteTemplate(w, "index.html", data)
		return
	}

	// Prepare data to render back to the user
	data := PageData{
		LLMName:     LLMModelName,
		PromptTitle: prompt, // Pass the original prompt back
		Response:    out.String(),
	}

	// Render the page again
	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
