# GEMINI.md - Contexto Instrucional para mac-monitor

Este arquivo fornece contexto e diretrizes para o projeto **mac-monitor**. Como o projeto está em sua fase inicial de planejamento e estruturação, siga estas orientações para manter a consistência com o fluxo de trabalho estabelecido.

## Visão Geral do Projeto

O **mac-monitor** é um projeto destinado ao monitoramento do sistema macOS. O repositório está estruturado para suportar um ciclo de desenvolvimento rigoroso, utilizando templates padronizados para documentação e gestão de tarefas.

### Tecnologias Inferidas (Baseadas em Templates)
Embora o código fonte ainda não esteja presente no diretório raiz, os templates sugerem o seguinte stack tecnológico planejado:
- **Linguagem Principal:** Go (Golang)
- **Testes E2E:** Playwright
- **Monitoramento:** Prometheus e Grafana
- **Fluxo de Trabalho:** Baseado em PRD (Product Requirements Document) e Especificações Técnicas.

## Estrutura do Diretório

- `/docs`: Destinado a armazenar a documentação final do projeto (PRDs, TechSpecs, etc.). Atualmente vazio.
- `/templates`: Contém os modelos fundamentais para o desenvolvimento:
  - `prd-template.md`: Template para Documento de Requisitos de Produto.
  - `techspec-template.md`: Template para Especificação Técnica, incluindo seções para arquitetura, modelos de dados e monitoramento.
  - `tasks-template.md`: Resumo de tarefas de implementação.
  - `task-template.md`: Detalhamento de tarefas individuais, com foco em testes e critérios de sucesso.

## Fluxo de Desenvolvimento Sugerido

Sempre que iniciar uma nova funcionalidade ou correção:
1. **Definição:** Utilize o `templates/prd-template.md` para definir os requisitos no diretório `/docs`.
2. **Design Técnico:** Utilize o `templates/techspec-template.md` para desenhar a solução técnica.
3. **Quebra de Tarefas:** Utilize `templates/tasks-template.md` e `templates/task-template.md` para organizar o trabalho.
4. **Implementação e Testes:** Siga as especificações técnicas, garantindo a criação de testes de unidade, integração e E2E (Playwright) conforme sugerido nos templates.

## Convenções de Idioma

A documentação e os templates estão em **Português**. Mantenha a consistência idiomática ao criar novos documentos de requisitos ou especificações técnicas, a menos que instruído de outra forma.

## Comandos Úteis (Placeholder)

Como o projeto ainda não possui um sistema de build definido:
- **TODO:** Definir comandos para `go build`, `go test` e execução do Playwright quando a estrutura de código for estabelecida.
