# Tarefa 5.0: Camada de Visualização (Menu Bar / TUI)

<critical>Ler os arquivos de prd.md e techspec.md desta pasta, se você não ler esses arquivos sua tarefa será invalidada</critical>

## Visão Geral
Desenvolvimento da interface de usuário do mac-monitor. O foco é uma experiência nativa, leve e informativa, suportando visualizações compactas e detalhadas.

<skills>
### Conformidade com Skills Padrões
- UI em Terminal (TUI) de alta fidelidade (ex: bubbletea, tcell)
- Integração com Menu Bar macOS (ex: systray)
- Apple Human Interface Guidelines
- Suporte a Modo Claro/Escuro
</skills>

<requirements>
- Opção de visualização compacta na Menu Bar.
- Janela detalhada ou TUI completa para análise aprofundada.
- Gráficos interativos para CPU e Memória.
- Suporte total a Modo Claro e Escuro.
- Navegação intuitiva via teclado e suporte a VoiceOver.
</requirements>

## Subtarefas

- [ ] 5.1 Implementar a integração com a Menu Bar usando `systray`.
- [ ] 5.2 Desenvolver a visualização detalhada (TUI ou Window) em `pkg/ui/`.
- [ ] 5.3 Criar componentes gráficos (charts) que consomem os dados do Engine.
- [ ] 5.4 Implementar a lógica de troca entre Modo Claro e Escuro.
- [ ] 5.5 Garantir a acessibilidade da interface seguindo as diretrizes da Apple.

## Detalhes de Implementação
Referenciar a seção **"UI/Display Layer"** na `techspec-mac-monitor.md`. A interface deve ser "viva", com atualizações suaves de 1 segundo, mantendo o consumo de recursos da UI mínimo.

## Critérios de Sucesso
- Interface se sente integrada ao macOS.
- Mudanças de tema (Light/Dark) são refletidas instantaneamente.
- Navegação fluida sem stuttering durante as atualizações de dados.

## Testes da Tarefa

- [ ] Testes de UI para verificar renderização correta em diferentes tamanhos de janela.
- [ ] Validação visual do suporte a Modo Escuro.
- [ ] Teste de usabilidade com teclado e VoiceOver.

<critical>SEMPRE CRIE E EXECUTE OS TESTES DA TAREFA ANTES DE CONSIDERÁ-LA FINALIZADA</critical>

## Arquivos relevantes
- `docs/prd-mac-monitor.md`
- `docs/techspec-mac-monitor.md`
- `pkg/ui/main_ui.go`
- `pkg/ui/components/`
