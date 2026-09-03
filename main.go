package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type OpenAlexResponse struct {
	Results []struct {
		Title           string `json:"title"`
		PublicationYear int    `json:"publication_year"`
		PrimaryLocation struct {
			Source struct {
				DisplayName string `json:"display_name"`
			} `json:"source"`
		} `json:"primary_location"`
		Doi string `json:"doi"`
	} `json:"results"`
}

func main() {
	url := "https://api.openalex.org/works?filter=author.orcid:https://orcid.org/0000-0002-9996-9983,type:article&sort=publication_year:desc"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erro ao buscar dados do OpenAlex: %v", err)
	}
	defer resp.Body.Close()

	var data OpenAlexResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Erro ao decodificar JSON: %v", err)
	}

	const tmpl = `
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Publicações - Gerlan Silva</title>
    <style>
        body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif; max-width: 800px; margin: 40px auto; padding: 0 20px; line-height: 1.6; color: #333; }
        h1 { border-bottom: 2px solid #eee; padding-bottom: 10px; }
        .article { margin-bottom: 25px; padding: 20px; background: #f9f9f9; border-radius: 8px; border-left: 4px solid #0366d6; }
        .title { font-size: 1.2em; font-weight: bold; margin-bottom: 5px; color: #24292e; }
        .journal { color: #586069; font-style: italic; margin-bottom: 5px; }
        .year { display: inline-block; background: #0366d6; color: white; padding: 2px 8px; border-radius: 12px; font-size: 0.85em; font-weight: bold; }
        a { color: #0366d6; text-decoration: none; margin-top: 10px; display: inline-block; }
        a:hover { text-decoration: underline; }
    </style>
</head>
<body>
    <h1>Artigos Publicados</h1>
    <p>Lista atualizada automaticamente via ORCID e OpenAlex.</p>
    
    {{range .Results}}
    <div class="article">
        <div class="title">{{.Title}}</div>
        <div class="journal">{{if .PrimaryLocation.Source.DisplayName}}{{.PrimaryLocation.Source.DisplayName}}{{else}}Revista não especificada{{end}}</div>
        <div class="year">{{.PublicationYear}}</div>
        {{if .Doi}}
        <br><a href="{{.Doi}}" target="_blank" rel="noopener noreferrer">Acessar Artigo (DOI) &rarr;</a>
        {{end}}
    </div>
    {{end}}
</body>
</html>
`
	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		log.Fatalf("Erro ao criar template: %v", err)
	}

	f, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("Erro ao criar o arquivo index.html: %v", err)
	}
	defer f.Close()

	if err := t.Execute(f, data); err != nil {
		log.Fatalf("Erro ao executar template: %v", err)
	}
	
	log.Println("Site gerado com sucesso em index.html")
}
