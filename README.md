# Go言語でQtを使ってみる
Go言語でQt(GUIツールキット)を扱えるパッケージのコードを読みながら試してみたQtの自作コードサンプル集<br>
[https://github.com/therecipe/qt](https://github.com/therecipe/qt)を利用する。<br>
Qtについて：[Qt | Cross-platform software development for embedded &amp; desktop](https://www.qt.io/)<br><br>

## サンプルコードの試し方<br>
**Go言語でQtを試すための環境構築**<br>
[Go言語でQtを扱ってみる on Windows - saitodev.co](https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7Qt%E3%82%92%E6%89%B1%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B_on_Windows)<br>
[Go言語でQtを扱ってみる on Ubuntu - saitodev.co](https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7Qt%E3%82%92%E6%89%B1%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B_on_Ubuntu)
<br><br>
**サンプルコードを試すためのコマンド**
```
// QFormLayoutを試したい場合(ディレクトリ名がレイアウトディレクトリ内のform)
cd サンプルコードのクローンを配置したいディレクトリ
git clone https://github.com/ryoko-saito/qt_on_golang.git
cd qt_on_golang/layout
qtdeploy test desktop ./form
```
