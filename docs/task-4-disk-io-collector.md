# Tarefa 4.0: Coletor de Disco e I/O

<critical>Ler os arquivos de prd.md e techspec.md desta pasta, se você não ler esses arquivos sua tarefa será invalidada</critical>

## Visão Geral
Implementação da coleta de métricas de armazenamento, incluindo espaço disponível em volumes e taxas de leitura/escrita (I/O) em tempo real.

<skills>
### Conformidade com Skills Padrões
- Go (Golang) com CGO
- Syscalls de Sistema de Arquivos (`getfsstat`, `statfs`)
- Monitoramento de I/O de Baixo Nível
</skills>

<requirements>
- Captura de taxas de leitura e escrita (I/O) globais e por volume principal.
- Reporte de espaço total, usado e disponível em todos os volumes montados.
- Identificação de pontos de montagem e tipos de sistema de arquivos (APFS, etc).
</requirements>

## Subtarefas

- [ ] 4.1 Implementar o coletor de disco em `pkg/collectors/disk.go`.
- [ ] 4.2 Utilizar `statfs` para obter informações de uso de volume.
- [ ] 4.3 Utilizar CGO para interfacear com métricas de I/O do kernel (IOKit ou counters de sistema).
- [ ] 4.4 Normalizar os dados em estruturas legíveis para a interface.
- [ ] 4.5 Garantir que volumes externos e de rede não causem lentidão na coleta.

## Detalhes de Implementação
Referenciar a seção **"Disk Collector"** na `techspec-mac-monitor.md`. Focar na obtenção de métricas de I/O sem polling excessivo, aproveitando os contadores do próprio Darwin.

## Critérios de Sucesso
- Espaço em disco reportado coincide com o `df -h`.
- Gráficos de I/O mostram atividade durante operações intensas de disco (ex: cópia de arquivos).
- Impacto nulo em discos em estado de repouso (não deve acordar discos desnecessariamente).

## Testes da Tarefa

- [ ] Teste de unidade para conversão de bytes/blocos para tamanhos legíveis (GB/TB).
- [ ] Teste de integração validando a leitura de volumes APFS.
- [ ] Validação de timeouts para volumes lentos ou inacessíveis.

<critical>SEMPRE CRIE E EXECUTE OS TESTES DA TAREFA ANTES DE CONSIDERÁ-LA FINALIZADA</critical>

## Arquivos relevantes
- `docs/prd-mac-monitor.md`
- `docs/techspec-mac-monitor.md`
- `pkg/collectors/disk.go`
- `pkg/models/disk.go`
