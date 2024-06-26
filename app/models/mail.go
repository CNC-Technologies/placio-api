package models

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"placio-pkg/logger"
	"strings"
	"time"
)

type EmailContent struct {
	Title   string `json:"title"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Button  struct {
		Label string `json:"label"`
		URL   string `json:"url"`
	} `json:"button"`
}

var emails map[string]EmailContent

func init() {
	// Load email content from JSON file
	// content, err := ioutil.ReadFile("emails/content.json")
	// if err != nil {
	// 	panic(err)
	// }
	// err = json.Unmarshal(content, &emails)
	// if err != nil {
	// 	panic(err)
	// }
}

func ParseTemplate(data interface{}) (content string, err error) {

	// ParseFiles creates a new 	Template and parses the template definitions from
	// the named files. The returned template's name will have the (base) name and
	// (parsed) contents of the first file. There must be at least one file.
	// If an error occurs, parsing stops and the returned *Template is nil.
	tmpl, err := template.ParseFiles("cmd/app/emails/template.html")
	if err != nil {
		return "", err
	}

	// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
	// The zero value for Buffer is an empty buffer ready to use.
	buf := new(bytes.Buffer)

	// Execute applies a parsed template to the specified data object,
	// writing the output to wr.
	// If an error occurs executing the template or writing its output,
	// execution stops, but partial results may already have been written to
	// the output writer.
	// A template may be executed safely in parallel.
	if err := tmpl.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (e *EmailContent) SendEmailToTerminal(to string) error {
	// Validate email address
	if !isValidEmail(to) {
		return errors.New("Invalid email address")
	}

	logger.Info(context.Background(), fmt.Sprintf("Email sent to %s", to))
	return nil
}

func (e *EmailContent) Send(to string, template string, data map[string]string) error {
	// Validate email address
	if !isValidEmail(to) {
		return errors.New("Invalid email address")
	}

	// Get email content
	content, ok := emails[template]
	if !ok {
		return errors.New("Email template not found")
	}

	// Create email HTML
	html, err := createEmail("template", content, data)
	if err != nil {
		return err
	}

	// Send email via Mailgun API
	values := url.Values{}
	values.Set("from", os.Getenv("MAILGUN_FROM"))
	values.Set("to", to)
	values.Set("subject", content.Subject)
	values.Set("html", html)
	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/messages", os.Getenv("Base_URL"), os.Getenv("Domain")), bytes.NewBufferString(values.Encode()))
	if err != nil {
		return err
	}
	req.SetBasicAuth("api", os.Getenv("MAILGUN_API_KEY"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Email not sent, status code: %d", resp.StatusCode)
	}

	fmt.Printf("Email sent to: %s\n", to)
	return nil
}

func isValidEmail(to string) bool {
	return strings.Contains(to, "@") && strings.Contains(to, ".")
}

func createEmail(template string, content EmailContent, values map[string]string) (string, error) {
	// Load email template
	templateFile := path.Join("emails", template+".html")
	templateContent, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return "", err
	}

	// Inject values into template
	email := string(templateContent)
	email = strings.ReplaceAll(email, "{{domain}}", values["domain"])

	if content.Title == "" {
		content.Title = content.Subject
	}
	if strings.Contains(content.Button.URL, "{{domain}}") {
		content.Button.URL = strings.ReplaceAll(content.Button.URL, "{{domain}}", values["domain"])
	}
	bodyLines := strings.Split(content.Body, "\n")
	if len(bodyLines) > 0 && bodyLines[0] != "" {
		bodyLines[0] = fmt.Sprintf("Hi %s,", content.Title)
	}
	body := ""
	for _, line := range bodyLines {
		body += fmt.Sprintf(`<p style="color: #7e8890; font-family: 'Source Sans Pro', helvetica, sans-serif; font-size: 15px; font-weight: normal; Margin: 0; Margin-bottom: 15px; line-height: 1.6;">%s</p>`, line)
	}
	email = strings.ReplaceAll(email, "{{title}}", content.Title)
	email = strings.ReplaceAll(email, "{{body}}", body)
	email = strings.ReplaceAll(email, "{{buttonURL}}", content.Button.URL)
	email = strings.ReplaceAll(email, "{{buttonLabel}}", content.Button.Label)

	email = strings.ReplaceAll(email, "{{year}}", fmt.Sprintf("%d", time.Now().Year()))

	if strings.Contains(email, "{{") {
		return "", errors.New("Email template contains invalid values")
	}

	return email, nil

}
