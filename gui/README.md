# 🖥️ mac-monitor GUI

Esta é a interface gráfica do **mac-monitor**, construída com [Wails v2](https://wails.io/) e [React](https://reactjs.org/).

## 🚀 Como Iniciar

### Pré-requisitos

1. **Go 1.18+** instalado.
2. **Wails CLI** instalado:
   ```bash
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```
3. **Node.js 16+** e **NPM** instalados.

### Desenvolvimento (Hot Reload)

Para rodar o projeto em modo de desenvolvimento, use o comando abaixo na raiz do projeto:

```bash
make gui-dev
```

Isso iniciará o Wails em modo `dev`, fornecendo Hot Reload tanto para o código Go quanto para o frontend React.

### Build de Produção

Para gerar o binário final (macOS `.app` bundle):

```bash
make gui-build
```

O binário será gerado em `gui/build/bin/`.

## 🎨 Design

A interface utiliza o tema **Glassmorphism**, com:
- **Fundo:** Transparência com blur (`backdrop-filter: blur(12px)`).
- **Cores:** Paleta futurista dark baseada no `docs/ui-system-design.md`.
- **Métricas:** Atualização em tempo real (1s) via Wails bindings.

## 🛠️ Estrutura do Projeto

- `gui/app.go`: Lógica Go que interage com a interface (Wails bindings).
- `gui/main.go`: Ponto de entrada da aplicação Wails.
- `gui/frontend/src/App.tsx`: Componente principal do dashboard.
- `gui/frontend/src/style.css`: Estilos globais e tema "Glass".
