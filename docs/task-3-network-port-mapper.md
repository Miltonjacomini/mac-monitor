# Tarefa 3.0: Mapeamento de Rede e Portas (Network Collector)

<critical>Ler os arquivos de prd.md e techspec.md desta pasta, se você não ler esses arquivos sua tarefa será invalidada</critical>

## Visão Geral
Implementação do coletor de rede focado na identificação de portas TCP/UDP abertas e seu mapeamento para processos específicos (PID e Nome), além de capturar taxas de tráfego instantâneas.

<skills>
### Conformidade com Skills Padrões
- Go (Golang) com CGO
- Biblioteca `libproc` do macOS
- Network Socket Analysis
- Mapeamento de PID para Nome de Processo
</skills>

<requirements>
- Listagem em tempo real de portas TCP/UDP em uso.
- Mapeamento preciso de cada porta para seu PID e Nome de processo correspondente.
- Cálculo de taxas de upload e download instantâneas por interface.
- Respeito aos limites de permissão do macOS para monitoramento de rede.
</requirements>

## Subtarefas

- [ ] 3.1 Implementar o coletor de rede em `pkg/collectors/network.go`.
- [ ] 3.2 Utilizar `libproc` para listar conexões e mapear PIDs a sockets.
- [ ] 3.3 Implementar o cálculo de throughput de rede (bytes/s).
- [ ] 3.4 Filtrar conexões para exibir apenas as portas ativas relevantes (Listening/Established).
- [ ] 3.5 Lidar com permissões de segurança necessárias para acessar dados de rede de outros processos.

## Detalhes de Implementação
Referenciar a seção **"Network Collector"** na `techspec-mac-monitor.md`. O desafio principal é a eficiência no mapeamento PID -> Port, que deve ser feito sem varrer excessivamente o sistema de arquivos ou o kernel.

## Critérios de Sucesso
- Listagem de portas coincide com o comando `lsof -nP -i`.
- Identificação correta dos nomes dos processos vinculados às portas.
- Taxas de rede estáveis e precisas.

## Testes da Tarefa

- [ ] Teste de unidade para o mapeador PID/Porta.
- [ ] Teste de integração verificando a detecção de uma porta de teste aberta propositalmente.
- [ ] Validação de performance para garantir que a varredura de rede não impacte o sistema.

<critical>SEMPRE CRIE E EXECUTE OS TESTES DA TAREFA ANTES DE CONSIDERÁ-LA FINALIZADA</critical>

## Arquivos relevantes
- `docs/prd-mac-monitor.md`
- `docs/techspec-mac-monitor.md`
- `pkg/collectors/network.go`
- `pkg/models/network.go`
