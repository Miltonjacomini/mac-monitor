# Tarefa 6.0: Integração Final e Otimização de Performance

<critical>Ler os arquivos de prd.md e techspec.md desta pasta, se você não ler esses arquivos sua tarefa será invalidada</critical>

## Visão Geral
Fase final de integração de todos os coletores, polimento da interface e otimização rigorosa de performance para garantir que o mac-monitor cumpra sua promessa de ser a ferramenta mais leve do sistema.

<skills>
### Conformidade com Skills Padrões
- Profiling de Performance em Go (`pprof`)
- Benchmarking de Sistema
- Segurança e Sandboxing de macOS
- Empacotamento de Aplicações para macOS
</skills>

<requirements>
- Integração estável de todos os sub-coletores (CPU, Mem, Net, Disk).
- Garantia de consumo de CPU < 1% em idle.
- Implementação de auto-monitoramento (self-monitoring).
- Configuração de permissões de segurança e conformidade com Sandbox.
- Build final otimizado para Intel e Apple Silicon.
</requirements>

## Subtarefas

- [ ] 6.1 Realizar a integração final de todos os módulos no `main.go`.
- [ ] 6.2 Executar sessões de profiling com `pprof` para identificar gargalos.
- [ ] 6.3 Otimizar o uso de memória e CPU do próprio mac-monitor.
- [ ] 6.4 Implementar diálogos de solicitação de permissão (Acessibilidade/Rede).
- [ ] 6.5 Configurar o pipeline de build para gerar binários universais.
- [ ] 6.6 Validar o comportamento térmico e de energia da aplicação.

## Detalhes de Implementação
Referenciar as seções **"Monitoramento e Observabilidade"** e **"Considerações Técnicas"** na `techspec-mac-monitor.md`. O foco é o refinamento final e a prontidão para distribuição.

## Critérios de Sucesso
- Consumo de CPU < 1% confirmado em idle.
- Aplicação estável sem memory leaks em execuções de longa duração (24h+).
- Binário único funcionando nativamente em Apple Silicon e Intel.

## Testes da Tarefa

- [ ] Testes de stress (long-run stability tests).
- [ ] Benchmarks comparativos de uso de recursos (CPU/RAM).
- [ ] Validação de segurança e integridade do binário.

<critical>SEMPRE CRIE E EXECUTE OS TESTES DA TAREFA ANTES DE CONSIDERÁ-LA FINALIZADA</critical>

## Arquivos relevantes
- `main.go`
- `Makefile`
- `scripts/build.sh`
- `docs/techspec-mac-monitor.md`
