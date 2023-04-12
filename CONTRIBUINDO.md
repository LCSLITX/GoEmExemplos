## Contribuindo

Obrigado pelo interesse em contribuir com o Go Em Exemplos!

* Ao enviar uma PR que afeta o conteúdo exibido no conteúdo do site, atualizar os arquivos HTML no diretório `public` é insuficiente, uma vez que a fonte do site está no diretório `examples`.
  
  Para corrigir de maneira apropriada, atualize os arquivos respectivos no
  diretório `examples` e execute o comando `tools/build` localmente para gerar
  o HTML; inclua ambos os arquivos (de `examples` e `public`) no seu PR.

  O conteúdo exibido no site só será atualizado ao enviar outra PR para a branch `gh-pages` com o diretório `docs` atualizado. Este diretório é igual ao diretório `public` gerado a partir do comando `tools/build`. O `docs` é utilizado somente porque o deploy do github pages necessita de um diretório com este nome.

  Passo-a-passo:
  1. Alterar o arquivo respectivo no diretório `examples`;
  2. Executar o comando `tools/build` para um novo diretório `public` ser gerado;
  3. Abrir uma nova PR com a branch base sendo a master (main);
  4. Renomear o novo diretório `public` para `docs`
  5. Abrir uma nova PR, apenas contendo o diretório `docs` com a branch base sendo a gh-pages.

  Se você não pretende criar as PRs, sinta-se a vontade para apenas abrir um issue com os seus apontamentos.

* Estamos abertos a adicionar mais exemplos no site. No entanto, eles devem ser sobre coisas realmente utilizadas pelos programadores e que apenas requerem a biblioteca padrão. Se você está interessado em adicionar um exemplo, por favor, antes de fazê-lo, abra um [issue](https://github.com/LCSLITX/GoEmExemplos/issues) para discutir.

* Não iremos alterar a navegação do site, tampouco adicionar um link de "seção anterior" ou um link de índice além do título do texto.
