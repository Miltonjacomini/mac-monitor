# PRD - mac-monitor: Ferramenta Nativa de Monitoramento para macOS

## Visão Geral
O **mac-monitor** é uma aplicação nativa para macOS projetada para fornecer visibilidade em tempo real sobre a saúde e a utilização dos recursos do sistema. O objetivo é oferecer uma ferramenta leve, precisa e integrada ao ecossistema Apple, permitindo que desenvolvedores e usuários avançados monitorem CPU, memória, portas de rede e outros indicadores vitais sem a sobrecarga de ferramentas multiplataforma pesadas.

## Objetivos
- **Monitoramento Abrangente:** Centralizar métricas de CPU, RAM, Disco e Rede em uma única interface.
- **Alta Performance:** Garantir que o próprio monitor consuma o mínimo possível de recursos (CPU < 1% em idle).
- **Experiência Nativa:** Utilizar APIs nativas do macOS para garantir precisão nos dados e uma interface que se sinta parte do sistema.
- **Métricas de Sucesso:** Baixa latência na atualização de dados e interface intuitiva que dispense manuais.

## Histórias de Usuário
- **Como desenvolvedor**, eu quero ver o uso de CPU por núcleo para identificar processos que não estão paralelizando bem.
- **Como administrador de sistemas**, eu quero listar rapidamente quais portas estão abertas e por quais processos, para diagnosticar conflitos de rede.
- **Como usuário avançado**, eu quero monitorar a "Pressão de Memória" (Memory Pressure) para entender se preciso de mais hardware ou fechar abas do navegador.
- **Como usuário de laptop**, eu quero ver o impacto térmico e de energia para gerenciar a duração da bateria.

## Funcionalidades Principais

### 1. Monitoramento de CPU
- Visualização de uso total e por núcleo individual.
- Frequência atual do processador.
- Listagem dos processos com maior consumo de CPU (Top 5).

### 2. Monitoramento de Memória
- Exibição de Memória Usada, Comprimida, Wired e Cache.
- Gráfico de Pressão de Memória (indicador nativo do macOS).
- Monitoramento de uso de Swap.

### 3. Monitoramento de Rede e Portas
- Listagem em tempo real de portas TCP/UDP em uso.
- Mapeamento de portas para processos específicos (PID e Nome).
- Taxas de upload e download instantâneas.

### 4. Monitoramento de Disco
- Taxas de leitura e escrita (I/O) em tempo real.
- Espaço disponível em volumes montados.

### 5. Interface de Usuário
- Suporte a Modo Claro e Escuro.
- Opção de visualização compacta na Menu Bar ou janela detalhada.

## Experiência do Usuário
- **Interface:** Limpa e minimalista, seguindo as Apple Human Interface Guidelines.
- **Acessibilidade:** Suporte total a VoiceOver e navegação por teclado.
- **Interatividade:** Gráficos interativos que mostram detalhes ao passar o mouse.

## Restrições Técnicas de Alto Nível
- **Plataforma:** Exclusivo para macOS (arquiteturas Intel e Apple Silicon).
- **Tecnologia:** Implementação utilizando APIs de baixo nível do sistema (ex: `libproc`, `sysctl`, `IOKit`).
- **Segurança:** A aplicação deve solicitar permissões explicitamente para monitoramento de rede e processos, respeitando o Sandbox do macOS onde aplicável.
- **Performance:** A coleta de dados não deve exceder intervalos de 1 segundo para evitar impacto no sistema monitorado.

## Fora de Escopo
- Monitoramento remoto de outros computadores via rede nesta fase inicial.
- Persistência de dados históricos em banco de dados por longos períodos (foco em tempo real).
- Sincronização em nuvem.
- Controle de processos (matar processos, mudar prioridade) - foco apenas em monitoramento no MVP.
