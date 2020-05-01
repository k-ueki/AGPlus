![](./src/assets/AGPlus6.png)
![Test](https://github.com/k-ueki/AGPlus/workflows/server_ci/badge.svg)

<div style="text-align: center;">
このサービスは青山学院大学の授業や教員に関するレビューを投稿および検索できるwebサービスです。
</div>
## Project setup
```
git clone https://github.com/k-ueki/AGPluss
```

### 環境構築手順
1. 以下のコマンドを叩き、docker imageを作成する
```
docker-compose build
```
2. コンテナを作成する
```
docker-compose up
```

※ server側のcodeを更新した場合は、1,2をやり直す必要あり(現時点)

