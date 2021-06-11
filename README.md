# bb-captura-artefatos-ordem-fornecimento

Captura dos nomes dos entregáveis (de baixa plataforma, versionados em repositório GIT) para as ordens de fornecimento
abertas a partir de **01/06/2021** conforme descrito no [manual](https://fontes.intranet.bb.com.br/stt/publico/atendimento)

## Uso

Baixar o executável correspondente ao seu sistema operacional em **bb-captura-artefatos-ordem-fornecimento/Bin**.
Existem compilações disponíves para [Windows](https://github.com/rmacedo88/bb-captura-artefatos-ordem-fornecimento/blob/main/Bin/captura_artefatos_windows_amd64.exe), [Linux](https://github.com/rmacedo88/bb-captura-artefatos-ordem-fornecimento/blob/main/Bin/captura_artefatos_linux_amd64) e [macOS](https://github.com/rmacedo88/bb-captura-artefatos-ordem-fornecimento/blob/main/Bin/captura_artefatos_darwin_amd64.dmg).

Antes de qualquer coisa, esta ferramenta depende que o git esteja instalado localmente.
[Siga as instruções contidas aqui (inglês)](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

Após a instalação do git, abra um terminal e confirme a instalação com:

```shell
$ git --version
```

Em seguida,

## Construir localmente

```shell
$ git clone https://github.com/rmacedo88/bb-captura-artefatos-ordem-fornecimento.git
$ cd bb-captura-artefatos-ordem-fornecimento
$ make -j8
```

> Os binários serão salvos no diretório **bb-captura-artefatos-ordem-fornecimento/Bin**

## Testes unitários

**NÂO** estão disponíveis por enquanto. Assim que possível estes serão criados e acionados pelo script
**bb-captura-artefatos-ordem-fornecimento/scripts/test_unit.sh**

## Avisos

Por se tratar de uma ferramenta não oficial e exposta em um domínio externo. Não existe menção ao guia de métricas, suas
tarefas, valores e quaisquer cálculos de produtividade. O intuito é somente produzir material para entrega segundo o
padrão estipulado para

## Contribuições

Essa ferramenta carece de alguns recursos tais como:

- Testes unitários
- Guardar a chave C do usuário
- Guardar a última data usada para gerar o relatório de entregas
- Suporte a múltiplos repositórios
