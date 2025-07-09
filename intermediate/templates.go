package intermediate

import (
	"bufio"
	"fmt"
	"html/template"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdio)
	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	templates := map[string]string{
		"welcome":      "Welcome, {{.name}} we are glad you joined",
		"notification": "{{.name}}, yuo have received a new notification: {{.notification}}",
		"error":        "Oops, there is an {{.error}}",
	}

	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

}
