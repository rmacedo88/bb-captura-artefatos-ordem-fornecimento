# bb-captura-artefatos-ordem-fornecimento

Captura o nome dos entregáveis (de baixa plataforma, versionados em repositório GIT) para as ordens de fornecimento
abertas a partir de **01/06/2021** conforme descrito
no [manual](https://fontes.intranet.bb.com.br/stt/publico/atendimento).

## Disclaimer

> Esta ferramenta apenas relaciona os arquivos criados ou alterados pela sua chave, sem discriminação da natureza ou
> validade desses artefatos perante o guia de métricas ou outros itens contratuais entre a fábrica e o BB.

> Artefatos não pertencentes ao guia de métricas ou impróprios para entrega, são de total
> responsabilidade do utilizador. As dúvidas sobre artefatos constantes no relatório produzido por essa ferramenta,
> devem ser encaminhadas ao seu gestor.

## Como utilizar

### Uso da Ferramenta

Baixe o executável correspondente ao seu sistema operacional na
seção **[Releases](https://github.com/rmacedo88/bb-captura-artefatos-ordem-fornecimento/releases)**. Existem compilações
disponíves para Windows, Linux e macOS.
> A Maioria dos erros emitidos pela aplicação são relacionados a repositórios incorretos e são terminais,
> informe todos os parâmetros corretamente para evitá-los.

#### Linux

```shell
$ cd ~/Downloads # supondo que você baixou no diretório Downloads
$ chmod +x captura_artefatos_linux_amd64 # Somente necessário na primeira execução
$ ./captura_artefatos_linux_amd64 -chave=<sua chave> -data=YYYY-MM-DD -arquivo=lista_repositorios.txt
```

#### Windows

```shell
$ cd Downloads # supondo que você baixou no diretório Downloads
$ captura_artefatos_windows_amd64.exe -chave=<sua chave> -data=YYYY-MM-DD -arquivo=lista_repositorios.txt
```

#### macOS

```shell
$ cd ~/Downloads # supondo que você baixou no diretório Downloads
$ chmod +x captura_artefatos_linux_amd64 # Somente necessário na primeira execução
$ ./captura_artefatos_darwin_amd64 -chave=<sua chave> -data=YYYY-MM-DD -arquivo=lista_repositorios.txt
```

###### Parâmetros:

- **-chave** : Sua chave C. Ex.: _**C1234567**_
- **-data** : Data de início do atendimento da OF no formato YYYY-MM-DD. Ex.: _**2021-06-01**_
- **-arquivo** : Arquivo de texto contendo uma lista de caminhos com os seus repositórios.
	Ex.: _**arquivo_com_a_lista_de_repositorios.txt**_

> O conteúdo do arquivo descrito no parâmetro _**-arquivo**_ pode ser informado da seguinte forma:

```shell
# exemplo do conteúdo do arquivo txt no Windows
C:\Users\<seu usuario>\<seu diretório de trabalho>\<repositorio git A>
C:\Users\<seu usuario>\<seu diretório de trabalho>\<repositorio git B>
C:\Users\<seu usuario>\<seu diretório de trabalho>\<repositorio git C>
...

# exemplo do conteúdo do arquivo txt no Linux
/home/<seu usuario>/<seu diretório de trabalho>/<repositorio git A>
/home/<seu usuario>/<seu diretório de trabalho>/<repositorio git B>
/home/<seu usuario>/<seu diretório de trabalho>/<repositorio git C>
...

# exemplo do conteúdo do arquivo txt no Mac
/home/<seu usuario>/<seu diretório de trabalho>/<repositorio git A>
/home/<seu usuario>/<seu diretório de trabalho>/<repositorio git B>
/home/<seu usuario>/<seu diretório de trabalho>/<repositorio git C>
...
```

### Instalação do git localmente

Essa ferramenta depende de uma instalação local do git. Caso você não o tenha
instalado [siga as instruções contidas aqui (inglês)](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

Para verificar a instalação, abra um terminal e digite:

```shell
$ git --version
git version 2.30.2 # Se estiver ok, deverá imprimir a versão instalada
```

> O git deve ser acessível a qualquer terminal, IDE ou outra ferramenta do tipo no seu sistema operacional.
> Se você não puder executá-lo nas condições supracitadas, é recomendado remover a versão existente e seguir o roteiro descrito acima para instalação correta.

## Construir localmente

```shell
$ git clone https://github.com/rmacedo88/bb-captura-artefatos-ordem-fornecimento.git
$ cd bb-captura-artefatos-ordem-fornecimento
$ make -j8
```

> Os binários serão salvos no diretório **bb-captura-artefatos-ordem-fornecimento/Bin** </br>
> Caso deseje compilar somente para a sua plataforma, é necessário editar o arquivo **bb-captura-artefatos-ordem-fornecimento/Makefile**

## Testes unitários

Aos que se interessem em contribuir, é importante se atentar aos testes unitários uma vez que o Makefile possui
instruções para permitir o build somente se os testes passarem.

Após construídos, eles devem ser acionados no script **bb-captura-artefatos-ordem-fornecimento/scripts/test_unit.sh**

## Avisos

Por se tratar de uma ferramenta não oficial e exposta em um domínio externo. Não existe menção ao guia de métricas, suas
tarefas, valores e quaisquer cálculos de produtividade. O intuito é somente produzir material para entrega segundo o
padrão determinado pela equipe de qualidade.

## Notas da versão

- Adição do suporte a multiplos repositórios.
- Depreciação do parâmetro _**-path**_
- Adição do parâmetro _**-arquivo**_
- Toda a saída no terminal refatorada para utilizar logging
- Adição de testes unitários
- Todos os erros foram convertidos em erros terminais

Consulte o [histórico de versões](CHANGELOG.md)

## Contribuições

Esta ferramenta carece de alguns recursos tais como:

- Guardar a chave C do usuário
- Guardar a última data usada para gerar o relatório de entregas
- entre outros ...

Basta clonar, implementar, testar e criar um pull request aqui no github