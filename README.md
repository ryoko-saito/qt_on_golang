# Go言語でQtを使ってみる
Go言語でQt(GUIツールキット)を扱えるパッケージのコードを読みながら試してみたQtの自作コードサンプル集<br>
<a href="https://github.com/therecipe/qt" target="_blank">https://github.com/therecipe/qt</a>を利用する。<br>
Qtについて：<a href="https://www.qt.io/" target="_blank">Qt | Cross-platform software development for embedded &amp; desktop</a><br><br>

## サンプルコードの試し方<br>
**Go言語でQtを試すための環境構築**<br>
[Go言語でQtを扱ってみる on Windows - saitodev.co](https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7Qt%E3%82%92%E6%89%B1%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B_on_Windows){:target="_blank"}<br>
[Go言語でQtを扱ってみる on Ubuntu - saitodev.co](https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7Qt%E3%82%92%E6%89%B1%E3%81%A3%E3%81%A6%E3%81%BF%E3%82%8B_on_Ubuntu){:target="_blank"}
<br><br>
**サンプルコードを試すためのコマンド**
```
// QFormLayoutを試したい場合(ディレクトリ名がform)
cd サンプルコードのクローンを配置したいディレクトリ
git clone https://github.com/ryoko-saito/qt_on_golang.git
cd qt_on_golang
qtdeploy test desktop ./form
```
