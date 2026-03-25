# Especificação Técnica - mac-monitor

## Resumo Executivo
O **mac-monitor** será implementado como uma aplicação de alta performance em Go (Golang), utilizando `cgo` para interfacear diretamente com as APIs de baixo nível do macOS (Darwin). A arquitetura será baseada em um modelo de "Coletores" modulares que consultam o kernel via `sysctl`, `libproc` e `IOKit`. Esta abordagem garante a precisão nativa dos dados com o mínimo de overhead de CPU, cumprindo o requisito de < 1% de uso em idle.

## Arquitetura do Sistema

### Visão Geral dos Componentes

- **Core Collector (Engine):** Orquestra o ciclo de coleta de dados de todos os sub-coletores.
- **CPU Collector:** Utiliza `host_processor_info` para métricas por núcleo e `sysctl` para frequências.
- **Memory Collector:** Interfaceia com `host_statistics64` para obter métricas de pressão de memória e tipos de alocação (Wired, Compressed, etc).
- **Network Collector:** Utiliza `libproc` e filtros de rede para mapear PIDs a portas TCP/UDP abertas.
- **Disk Collector:** Utiliza `getfsstat` e `statfs` para I/O e uso de disco.
- **UI/Display Layer:** Inicialmente uma interface de terminal (TUI) de alta fidelidade ou uma integração com a Menu Bar via `systray` para manter a natureza nativa e leve.

### Fluxo de Dados
1. O **Engine** dispara tiques (padrão: 1s).
2. Cada **Collector** executa chamadas de sistema (syscalls) ou chamadas CGO.
3. Os dados brutos são normalizados em estruturas Go.
4. O **Aggregator** prepara os dados para a camada de exibição.

## Design de Implementação

### Interfaces Principais

```go
// Interface base para todos os coletores de métricas
type Collector interface {
    // Collect captura as métricas atuais do sistema
    Collect(ctx context.Context) (MetricsPayload, error)
    // Name retorna o identificador do coletor (cpu, mem, net, etc)
    Name() string
}

// Payload de métricas normalizado
type MetricsPayload struct {
    Timestamp time.Time
    Data      map[string]interface{}
}
```

### Modelos de Dados

```go
type CPUMetrics struct {
    TotalUsage float64
    PerCore    []float64
    Frequency  int64 // MHz
    TopProcs   []ProcessInfo
}

type MemoryMetrics struct {
    Used       uint64
    Wired      uint64
    Compressed uint64
    Pressure   float64 // 0-100%
    SwapUsed   uint64
}

type NetworkPort struct {
    Port     int
    Protocol string // "TCP" | "UDP"
    PID      int
    Process  string
}
```

## Pontos de Integração

- **macOS Kernel (Darwin):** Integração via CGO com as bibliotecas do sistema.
- **Prometheus (Opcional):** Endpoint local para exportação de métricas se configurado.

## Abordagem de Testes

### Testes Unidade
- Mock de payloads do sistema para validar a lógica de normalização.
- Validação de cálculos de porcentagem de CPU e conversão de bytes.

### Testes de Integração
- Verificação de consistência entre os dados coletados e as ferramentas nativas (`top`, `netstat`, `Activity Monitor`).

### Testes de E2E
- Testes automatizados de interface usando Playwright (se houver interface web/GUI) ou scripts de validação de CLI.

## Sequenciamento de Desenvolvimento

### Ordem de Construção
1. **Squeleto do Engine e CGO:** Configurar o ambiente de build para Darwin.
2. **CPU & Memory Collectors:** Funcionalidades base e mais críticas para performance.
3. **Network & Port Mapper:** Implementação do mapeamento PID -> Port (complexidade média).
4. **Display Layer:** Implementação da Menu Bar/TUI.
5. **Integração e Polimento:** Ajustes de performance e suporte a Modo Escuro.

### Dependências Técnicas
- Ferramentas de linha de comando do Xcode (Clang/LLVM) para compilação CGO.
- Biblioteca `libproc` e headers do kernel macOS.

## Monitoramento e Observabilidade
- **Logs:** Nível `INFO` para eventos de ciclo de vida e `DEBUG` para falhas de leitura de syscalls.
- **Self-Monitoring:** O mac-monitor reportará seu próprio uso de CPU/RAM como parte da telemetria.

## Considerações Técnicas

### Decisões Principais
- **Go vs Swift:** Go foi escolhido pela eficiência em concorrência (Goroutines para coletores paralelos) e facilidade de criar binários estáticos, mantendo a performance próxima ao C.
- **CGO:** Necessário para acessar APIs que não possuem wrappers puros em Go.

### Riscos Conhecidos
- **Permissões (Sandbox):** O acesso a certas informações de rede (como portas de outros processos) pode exigir privilégios de `root` ou permissões específicas de "Acessibilidade" no macOS.
- **Mudanças de API da Apple:** APIs de baixo nível podem mudar entre versões do macOS (ex: Monterey para Ventura/Sonoma).

### Arquivos relevantes e dependentes
- `docs/prd-mac-monitor.md`
- `main.go` (A ser criado)
- `pkg/collectors/` (A ser criado)
