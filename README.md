# Site de Artigos com Go + GitHub Pages

Este repositório contém o código para gerar automaticamente uma página com seus artigos acadêmicos, buscando dados do seu ORCID via API do OpenAlex.

## Estrutura
- `main.go`: Script em Go que busca os dados e gera o arquivo `index.html`.
- `.github/workflows/update.yml`: Configuração do GitHub Actions que roda o script automaticamente toda semana e faz o commit no repositório.

## Como usar
1. Extraia o conteúdo desta pasta no seu computador.
2. Suba o conteúdo (incluindo a pasta oculta `.github`) para um novo repositório no seu GitHub.
3. Vá na aba "Actions" do seu repositório no GitHub, selecione "Atualizar Site de Artigos" e clique em "Run workflow" para rodar pela primeira vez e gerar o arquivo `index.html`.
4. Vá em "Settings" > "Pages", selecione a branch `main` e a pasta `/(root)` e salve.
5. Seu site estará disponível em alguns minutos!
