# 🖥️ mac-monitor

O **mac-monitor** é uma ferramenta de monitoramento de sistema de alta performance para macOS, desenvolvida em Go. Ele oferece uma interface gráfica moderna com estética futurista ("Glassmorphism") via Wails e uma interface de linha de comando (CLI) leve.

---

## ✨ Funcionalidades

- 📊 **Dashboard em Tempo Real:** Monitoramento de CPU, Memória, Disco e Rede com atualização a cada 1 segundo.
- 🧊 **Glassmorphism UI:** Interface translúcida com efeitos de desfoque (blur) e design dark futurista.
- ⚙️ **Cgo Native Integration:** Coleta direta de métricas através de APIs nativas do macOS (`mach`, `sysctl`, `libproc`).
- 🌐 **Network Insights:** Lista de portas abertas, processos associados e estatísticas de interfaces de rede.
- 💾 **Storage Analytics:** Visualização de volumes montados, taxas de leitura/escrita e espaço disponível.
- 📟 **CLI Mode:** Versão ultraleve para terminal.

---

## 🚀 Como Iniciar

### Pré-requisitos

1.  **macOS:** Necessário para as APIs nativas (`darwin`).
2.  **Go 1.21+**
3.  **Wails CLI:** Para a interface gráfica.
    ```bash
    go install github.com/wailsapp/wails/v2/cmd/wails@latest
    ```
4.  **Node.js & NPM:** Para o frontend React.

### Configuração do Ambiente (Fish Shell)

Se você usa **fish** e os comandos `wails` ou `go` não forem encontrados, adicione-os ao seu PATH:

```fish
fish_add_path (go env GOPATH)/bin
```

---

## 🛠️ Desenvolvimento e Build

O projeto utiliza um `Makefile` para facilitar as operações comuns:

### Interface Gráfica (GUI)
```bash
# Rodar em modo desenvolvimento com Hot Reload
make gui-dev

# Gerar o binário de produção (.app bundle)
make gui-build
```
*O binário da GUI será gerado em: `gui/build/bin/mac-monitor-gui`*

### Linha de Comando (CLI)
```bash
# Build da versão CLI
make build

# Executar a versão CLI
make run
```
*O binário da CLI será gerado em: `bin/mac-monitor`*

---

## 🎨 Design System ("The Glass Theme")

Baseado no documento `docs/ui-system-design.md`, a interface utiliza:
- **Background:** `#0B0E11` (Dark Slate)
- **Superfície:** `rgba(30, 34, 39, 0.7)` com `backdrop-filter: blur(12px)`
- **Acentos:** Azul Primário (`#3B82F6`) e Verde Saudável (`#10B981`)
- **Tipografia:** San Francisco (System Default)

---

## 📂 Estrutura do Repositório

- `/pkg/collectors`: Implementações nativas em Go/Cgo para coleta de métricas.
- `/pkg/models`: Definições de estruturas de dados compartilhadas.
- `/gui`: Código fonte da aplicação Wails (Go + React).
- `/gui/frontend`: Interface React com Recharts e CSS customizado.
- `/docs`: Documentação técnica, PRDs e especificações de design.

---

## ⚖️ Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
