# DynamoDB セットアップガイド

このドキュメントでは、echo-todoアプリケーションでDynamoDBを使用するための設定手順を説明します。

## 前提条件

- AWSアカウントが必要
- AWS CLI がインストールされている
- 適切なIAM権限が設定されている

## 1. AWS CLI のセットアップ

### AWS CLI のインストール
```bash
# macOS (Homebrew)
brew install awscli

# Linux
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install

# Windows
# AWS公式サイトからインストーラーをダウンロード
```

### AWS認証情報の設定
```bash
aws configure
```

以下の情報を入力：
- AWS Access Key ID
- AWS Secret Access Key  
- Default region name (例: `us-east-1`)
- Default output format (例: `json`)

## 2. DynamoDBテーブルの作成

### AWS Console を使用する場合

1. [AWS DynamoDB Console](https://console.aws.amazon.com/dynamodb/) にアクセス
2. 「テーブルの作成」をクリック
3. 以下の設定でテーブルを作成：
   - **テーブル名**: `todos`
   - **パーティションキー**: `id` (文字列)
   - **テーブル設定**: デフォルト設定またはオンデマンド

### AWS CLI を使用する場合

```bash
# テーブル作成
aws dynamodb create-table \
    --table-name todos \
    --attribute-definitions \
        AttributeName=id,AttributeType=S \
    --key-schema \
        AttributeName=id,KeyType=HASH \
    --billing-mode PAY_PER_REQUEST \
    --region us-east-1

# テーブル作成完了の確認
aws dynamodb describe-table --table-name todos --region us-east-1
```

### Terraform を使用する場合

```hcl
resource "aws_dynamodb_table" "todos" {
  name           = "todos"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "id"

  attribute {
    name = "id"
    type = "S"
  }

  tags = {
    Name        = "TodosTable"
    Environment = "development"
  }
}
```

## 3. IAM権限の設定

アプリケーションがDynamoDBにアクセスするために、適切なIAM権限が必要です。

### IAMポリシーの例

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "dynamodb:GetItem",
                "dynamodb:PutItem",
                "dynamodb:UpdateItem",
                "dynamodb:DeleteItem",
                "dynamodb:Scan",
                "dynamodb:Query"
            ],
            "Resource": "arn:aws:dynamodb:us-east-1:ACCOUNT-ID:table/todos"
        }
    ]
}
```

### ローカル開発用の設定

1. IAMユーザーを作成し、上記ポリシーをアタッチ
2. アクセスキーを生成
3. `aws configure` で認証情報を設定

### 本番環境での設定

- EC2: IAMロールを使用
- Lambda: 実行ロールに権限を付与
- ECS: タスクロールを使用

## 4. 環境変数の設定

アプリケーションの設定に使用する環境変数：

```bash
# .env ファイルまたは環境変数として設定
export DYNAMODB_TABLE_NAME=todos
export AWS_REGION=us-east-1
export PORT=1323

# AWS認証情報（推奨：IAMロール使用）
export AWS_ACCESS_KEY_ID=your-access-key
export AWS_SECRET_ACCESS_KEY=your-secret-key
```

## 5. ローカルでのテスト

### DynamoDB Local の使用（オプション）

開発環境でDynamoDB Localを使用することも可能：

```bash
# Docker でDynamoDB Local を起動
docker run -p 8000:8000 amazon/dynamodb-local

# ローカルテーブル作成
aws dynamodb create-table \
    --table-name todos \
    --attribute-definitions AttributeName=id,AttributeType=S \
    --key-schema AttributeName=id,KeyType=HASH \
    --billing-mode PAY_PER_REQUEST \
    --endpoint-url http://localhost:8000 \
    --region us-east-1
```

ローカル環境用の環境変数：
```bash
export AWS_ENDPOINT_URL=http://localhost:8000
export AWS_ACCESS_KEY_ID=dummy
export AWS_SECRET_ACCESS_KEY=dummy
```

## 6. アプリケーション起動

```bash
# 依存関係のインストール
go mod tidy

# アプリケーション起動
go run main.go

# または
go build -o echo-todo
./echo-todo
```

## 7. 動作確認

```bash
# ヘルスチェック
curl http://localhost:1323/health

# レスポンス例
{"status":"healthy"}
```

## トラブルシューティング

### よくあるエラーと対処法

1. **認証エラー**
   ```
   UnableToLocateCredentialsError: Unable to locate credentials
   ```
   → AWS認証情報が正しく設定されているか確認

2. **テーブルが見つからない**
   ```
   ResourceNotFoundException: Table not found
   ```
   → テーブル名と地域が正しいか確認

3. **権限エラー**
   ```
   AccessDeniedException: User is not authorized
   ```
   → IAM権限が適切に設定されているか確認

4. **地域エラー**
   ```
   ValidationException: Invalid region
   ```
   → AWS_REGIONが正しく設定されているか確認

## 次のステップ

DynamoDB接続が確立できたら、以下の機能を実装できます：

1. [登録API](https://github.com/wato787/echo-todo/issues/1)
2. [取得API](https://github.com/wato787/echo-todo/issues/4)  
3. [更新API](https://github.com/wato787/echo-todo/issues/2)
4. [削除API](https://github.com/wato787/echo-todo/issues/3)