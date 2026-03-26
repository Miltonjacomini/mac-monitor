# 🖥️ System Monitor v1.2.3 - Technical Documentation

Este documento serve como guia de implementação para o **System Monitor**, uma aplicação desktop de alta performance desenvolvida com Electron e React. O foco é fornecer métricas de hardware em tempo real com uma interface moderna e baixo consumo de recursos.

---

## 🏗️ 1. Arquitetura de Software

A aplicação segue o padrão de **Processos Separados** do Electron para garantir segurança e fluidez na UI.

### 🛰️ Camada de Dados (Main Process)
O processo principal é responsável pela coleta de telemetria bruta.
- **Library:** `systeminformation` (Node.js).
- **Polling:** Ciclo de 1 segundo para métricas dinâmicas (CPU, RAM, Rede) e 5 segundos para estáticas (Disco, Info de Hardware).
- **Eficiência:** O polling é pausado automaticamente quando a aplicação está em background ou ocultada.

### 🌉 Ponte (Preload Script)
Utiliza `contextBridge` para expor APIs seguras ao React sem expor o módulo `ipcRenderer` completo.

### 🎨 Interface (Renderer Process)
Single Page Application (SPA) construída com React e Vite.
- **State Management:** Hooks nativos (`useState`, `useEffect`) para estados locais de métricas.
- **Rendering:** Componentes memoizados para evitar re-renders causados pelo fluxo constante de dados do IPC.

---

## 🎨 2. Design System (The "Glass" Theme)

O design baseia-se em uma estética **Dark Futuristic** com transparências e bordas suaves.

### 🌑 Paleta de Cores
| Elemento | Hex/RGBA | Uso |
| :--- | :--- | :--- |
| **Background** | `#0B0E11` | Fundo principal da janela. |
| **Card Surface** | `rgba(30, 34, 39, 0.7)` | Fundo dos widgets com `backdrop-filter: blur(12px)`. |
| **Primary Accent** | `#3B82F6` | Gráficos de CPU e botões ativos. |
| **Healthy State** | `#10B981` | Status de Memória e Rede estáveis. |
| **Critical State** | `#EF4444` | Alertas de temperatura ou uso de disco cheio. |
| **Text Primary** | `#F3F4F6` | Títulos e métricas principais. |
| **Text Muted** | `#9CA3AF` | Labels secundárias e unidades (GB, MHz). |

### 📐 Tipografia e Espaçamento
- **Font Stack:** Inter ou San Francisco (System Default).
- **Border Radius:** `12px` para cards, `6px` para botões internos.
- **Grid:** Layout baseado em CSS Grid (3 colunas no Dashboard, 1 coluna no Popover).

---

## 💻 3. Implementação dos Componentes Core

### 📊 Gráficos de Performance (Charts)
Para os gráficos de linha (CPU/Rede), recomenda-se o uso de **SVG nativo** ou **Recharts** com as seguintes propriedades:
- `isAnimationActive={false}` (Melhora a performance em updates de 1s).
- Gradientes lineares no preenchimento inferior da linha.

### 🗔 Gerenciamento de Janelas (Tray & Popover)
O comportamento da versão minimizada na barra de menu exige lógica específica no Electron:

```javascript
// Exemplo de lógica para o Tray Popover
tray.on('click', (event, bounds) => {
  const { x, y } = bounds;
  const { width, height } = popoverWindow.getBounds();
  
  // Posicionamento centralizado abaixo do ícone
  const posX = Math.round(x - width / 2 + bounds.width / 2);
  const posY = Math.round(y);

  popoverWindow.setPosition(posX, posY, false);
  popoverWindow.isVisible() ? popoverWindow.hide() : popoverWindow.show();
});