# go-generate-code-smaple

template/template.goを雛形とし、コードを生成する。
生成したコードはout/main.goとして出力される。

今回は簡易的にfmt.Println関数の引数を変えたものを生成し、出力。
ast.Walk関数やprinter.Fprint関数を利用し、ASTの変換、ASTからのコード生成を実施した。
