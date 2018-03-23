package main

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	//"net/http"
	"os"
	"strings"
	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

type dockerHubStatsType struct {
	Count int `json:"count"`
}

func sanitizeInput(input string) string {
	parts := strings.Split(input, "\n")
	return strings.Trim(parts[0], " ")
}

func translateTextToEnglish(text string) (string, error) {
	ctx := context.Background()

        targetLanguage := "En"

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", err
	}
        authOpt := option.WithAPIKey(os.Getenv("APIKEY"))
	client, err := translate.NewClient(ctx, authOpt)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", err
	}
	return resp[0].Text, nil
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("Unable to read standard input:", err)
	}
	arg := string(input)
	if len(input) == 0 {
		log.Fatalln("A word or a sentance required to detect language")
	}

	arg = sanitizeInput(arg)
	translated, terr := translateTextToEnglish(arg)
        if terr != nil {
		fmt.Printf("Error: %v\n", terr)
	}        

	fmt.Printf("%s\n", translated)
}
