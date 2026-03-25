# Tarefa 2.0: Implementação dos Coletores de CPU e Memória

<critical>Ler os arquivos de prd.md e techspec.md desta pasta, se você não ler esses arquivos sua tarefa será invalidada</critical>

## Visão Geral
Esta tarefa foca na implementação dos coletores core do sistema: CPU e Memória. Utilizando CGO e APIs nativas do Darwin, o objetivo é capturar métricas precisas com o mínimo de overhead.

<skills>
### Conformidade com Skills Padrões
- Go (Golang) com CGO
- APIs de Kernel macOS (`host_processor_info`, `host_statistics64`)
- Cálculo de Métricas de Sistema (Load Average, Memory Pressure)
</skills>

<requirements>
- Uso de `host_processor_info` para métricas por núcleo individual.
- Uso de `sysctl` para obter a frequência atual do processador.
- Uso de `host_statistics64` para tipos de memória: Usada, Comprimida, Wired e Cache.
- Cálculo do indicador de Pressão de Memória (Memory Pressure).
- Listagem dos processos com maior consumo de CPU (Top 5).
</requirements>

## Subtarefas

- [ ] 2.1 Implementar o coletor de CPU em `pkg/collectors/cpu.go` usando CGO.
- [ ] 2.2 Implementar o coletor de Memória em `pkg/collectors/memory.go` usando CGO.
- [ ] 2.3 Implementar a lógica de normalização para as estruturas `CPUMetrics` e `MemoryMetrics`.
- [ ] 2.4 Integrar os coletores ao Engine Core.
- [ ] 2.5 Validar a precisão dos dados comparando com o `top` e o `Activity Monitor`.

## Detalhes de Implementação
Referenciar as seções **"CPU Collector"** e **"Memory Collector"** na `techspec-mac-monitor.md`. Utilize syscalls para garantir a performance e a fidelidade dos dados nativos.

## Critérios de Sucesso
- Métricas de CPU por núcleo refletem a realidade do sistema.
- Pressão de memória calculada corretamente seguindo a lógica do macOS.
- O consumo de CPU desses coletores somados deve ser < 0.5% em idle.

## Testes da Tarefa

- [ ] Testes de unidade para normalização de bytes para GB/MB.
- [ ] Testes de integração para verificar se as chamadas de sistema retornam dados válidos.
- [ ] Benchmarking do tempo de coleta para garantir latência mínima.

<critical>SEMPRE CRIE E EXECUTE OS TESTES DA TAREFA ANTES DE CONSIDERÁ-LA FINALIZADA</critical>

## Arquivos relevantes
- `docs/prd-mac-monitor.md`
- `docs/techspec-mac-monitor.md`
- `pkg/collectors/cpu.go`
- `pkg/collectors/memory.go`
- `pkg/models/metrics.go`
