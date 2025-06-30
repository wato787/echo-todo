# Go Echo プロジェクト構成

このドキュメントでは、Go EchoベストプラクティスのためのCRUD のプロジェクト構成について説明します。

## フォルダ構成

```
echo-todo/
├── cmd/                    # アプリケーションのエントリーポイント
│   └── server/            # サーバーアプリケーション
│       └── main.go        # メインアプリケーション
├── internal/              # プライベートなアプリケーションコード
│   ├── config/           # 設定管理
│   │   └── config.go     # アプリケーション設定
│   ├── handlers/         # HTTPハンドラー（コントローラー）
│   │   └── todo_handler.go
│   ├── middleware/       # カスタムミドルウェア
│   │   └── auth.go      # 認証・認可ミドルウェア
│   ├── repository/       # データアクセス層
│   │   └── todo_repository.go
│   └── services/         # ビジネスロジック層
│       └── todo_service.go
├── pkg/                  # 他のプロジェクトでも使用可能なライブラリコード
│   ├── models/          # データモデル
│   │   └── todo.go
│   └── utils/           # ユーティリティ関数
│       ├── response.go  # レスポンス形式
│       └── validator.go # バリデーション
├── docs/                # ドキュメント
├── bin/                 # ビルド済みバイナリ（gitignore推奨）
├── go.mod              # Go モジュール定義
├── go.sum              # 依存関係のハッシュ
└── README.md           # プロジェクト説明
```

## 各ディレクトリの役割

### `/cmd`
- アプリケーションのエントリーポイント
- 複数のアプリケーション（server, cli toolsなど）がある場合はそれぞれサブディレクトリを作成
- main関数を含む

### `/internal`
- このアプリケーション専用のプライベートコード
- 他のプロジェクトからimportされない
- アプリケーションのコアロジック

#### `/internal/handlers`
- HTTPリクエストを処理するハンドラー
- Echo Contextを受け取り、レスポンスを返す
- ルーティングとHTTP層の責務

#### `/internal/services`
- ビジネスロジック層
- ドメインルールと業務要件を実装
- repositoryを呼び出してデータを操作

#### `/internal/repository`
- データアクセス層
- データベースへのCRUD操作
- インターフェースとして定義し、実装を交換可能にする

#### `/internal/config`
- アプリケーション設定の管理
- 環境変数の読み込みと設定構造体

#### `/internal/middleware`
- カスタムミドルウェア
- 認証、ログ、リクエスト処理など

### `/pkg`
- 他のプロジェクトでも再利用可能なライブラリコード
- 外部に公開してもよいコード

#### `/pkg/models`
- データモデルの定義
- JSONタグ、バリデーションタグなど

#### `/pkg/utils`
- ユーティリティ関数
- 汎用的な処理

## アーキテクチャパターン

### Layered Architecture
```
HTTP Request
     ↓
[Handlers] ← ルーティング、HTTPレスポンス
     ↓
[Services] ← ビジネスロジック
     ↓
[Repository] ← データアクセス
     ↓
Database/External APIs
```

### Dependency Injection
- インターフェースベースの設計
- テストしやすい構造
- 実装の交換が容易

### 責務分離
- **Handlers**: HTTPリクエスト/レスポンスの処理
- **Services**: ビジネスロジック
- **Repository**: データアクセス
- **Models**: データ構造定義

## ビルドと実行

### 開発環境
```bash
# サーバー起動
go run cmd/server/main.go

# ビルド
go build -o bin/server cmd/server/main.go

# 実行
./bin/server
```

### 依存関係管理
```bash
# モジュール整理
go mod tidy

# 依存関係追加
go get package-name
```

## 推奨事項

### ネーミング規則
- **Handlers**: `{Entity}Handler` (例: `TodoHandler`)
- **Services**: `{Entity}Service` (例: `TodoService`)
- **Repository**: `{Entity}Repository` (例: `TodoRepository`)
- **Models**: エンティティ名 (例: `Todo`)

### インターフェース設計
- 小さく、焦点を絞ったインターフェース
- mockしやすい設計
- 依存関係の逆転

### エラーハンドリング
- カスタムエラー型の活用
- 適切なHTTPステータスコード
- 統一されたエラーレスポンス形式

### テスト
- 各層ごとにテストを作成
- モックを活用した単体テスト
- 統合テストでエンドツーエンドの動作確認

この構成により、保守性が高く、テストしやすく、拡張性のあるGo Echoアプリケーションを構築できます。