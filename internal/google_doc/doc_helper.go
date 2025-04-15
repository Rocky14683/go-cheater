package google_doc

import "google.golang.org/api/docs/v1"

func GetContentString(body *docs.Body) string {
	var output string
	for _, element := range body.Content {
		if element.Paragraph != nil {
			for _, elem := range element.Paragraph.Elements {
				if elem.TextRun != nil {
					output += elem.TextRun.Content
				}
			}
		}
	}
	return output
}
