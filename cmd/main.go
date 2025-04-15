package main

import (
	"github.com/urfave/cli/v2"
	"go-cheater/internal/google_doc"
	"log"
	"os"
)

func main() {

	app := &cli.App{

		Name:                 "go-cheater",
		Version:              "0.0.1",
		Usage:                "A CLI tool to help you avoid AI detection",
		EnableBashCompletion: true,
		Suggest:              true,

		Commands: []*cli.Command{
			{
				Name:    "parse",
				Usage:   "Parse a document",
				Aliases: []string{"p"},
				Action: func(cCtx *cli.Context) error {
					docId := cCtx.Args().Get(0)
					doc, err := google_doc.RequestDocData(docId)
					if err != nil {
						panic(err)
					} else {
						log.Println("Document Title: ", doc.Title)
						output := google_doc.GetContentString(doc.Body)
						log.Println("Document Content: ", output)
						log.Println("Document ID: ", doc.DocumentId)
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
