# Tarefa 1.0: Configuração do Projeto e Engine Core

<critical>Ler os arquivos de prd.md e techspec.md desta pasta, se você não ler esses arquivos sua tarefa será invalidada</critical>

## Visão Geral
Esta tarefa consiste em configurar o ambiente de desenvolvimento Go com suporte a CGO para macOS e implementar o motor principal (Engine) que orquestrará os coletores de métricas.

<skills>
### Conformidade com Skills Padrões
- Go (Golang) com CGO
- Estruturas de Dados de Baixo Nível (Darwin/Mach)
- Concorrência com Goroutines
</skills>

<requirements>
- Ambiente Go 1.21+ configurado.
- Suporte a CGO habilitado para compilação em macOS.
- Implementação da interface `Collector` definida na Tech Spec.
- Motor de loop (Ticker) que respeite o intervalo de 1s de coleta.
</requirements>

## Subtarefas

- [ ] 1.1 Inicializar o módulo Go (`go mod init mac-monitor`).
- [ ] 1.2 Criar a estrutura de diretórios (`pkg/engine`, `pkg/collectors`, `pkg/models`).
- [ ] 1.3 Definir a interface `Collector` e a estrutura `MetricsPayload`.
- [ ] 1.4 Implementar o `Core Engine` (loop de coleta e agregação básica).
- [ ] 1.5 Configurar o Makefile para builds CGO direcionados ao Darwin (Intel/ARM).

## Detalhes de Implementação
Referenciar as seções **"Arquitetura do Sistema"** e **"Design de Implementação"** na `techspec-mac-monitor.md`. O foco deve ser na robustez do loop principal e na capacidade de adicionar novos coletores dinamicamente através da interface definida.

## Critérios de Sucesso
- Binário compila sem erros no macOS.
- O engine inicia e executa o loop de coleta sem deadlocks.
- O uso inicial de CPU do engine (sem coletores pesados) deve ser próximo a 0%.

## Testes da Tarefa

- [ ] Testes de unidade para o agendador do Engine (mocking collectors).
- [ ] Teste de integração verificando se o Engine chama o método `Collect` no intervalo correto.
- [ ] Validação de concorrência (Race Detector: `go test -race`).

<critical>SEMPRE CRIE E EXECUTE OS TESTES DA TAREFA ANTES DE CONSIDERÁ-LA FINALIZADA</critical>

## Arquivos relevantes
- `docs/prd-mac-monitor.md`
- `docs/techspec-mac-monitor.md`
- `main.go`
- `pkg/engine/engine.go`
- `pkg/collectors/collector.go`
