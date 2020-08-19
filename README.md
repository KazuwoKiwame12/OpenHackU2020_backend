# OpenHackU2020_backend
## APIの設計
### 感情が地図に載る(感情の共有)
**各県の色16進数を返す**  
- APIのURL: `/emotion/prefectures`  
- Method: Get  
- フロントエンドから受け取るデータ  
```
None
```
- フロントエンドへと渡すデータ
```
[
  {
    prefecture: string,
    color: string
  },
  ....
]
```

**コメントの一覧表示 or（感情の場所を指すピン）を各県ごとに分けて渡す(選択された県のみ渡す方が良い)**
- APIのURL: `/emotion/{prefecture}/comments`
- Method: Get  
- フロントエンドから受け取るデータ
```
prefecture: string  
```
- フロントエンドへと渡すデータ
```
[
    {
      id: int,
      emotion: int,
      latitude: double,
      longtitude: double,
      user_name: string,
      dateTime: dateTime,
    },
    .....
]
```
**ピンを押して、コメント内容確認(紐ついた返信も)**
- APIのURL: `/emotion/{prefecture}/comments/{id}` 
- Method: Get   
- フロントエンドから受け取るデータ  
```
id: int
```
- フロントエンドへと渡すデータ
```
[
    {
        comment: string,
        user: {id: int, name: string},
        emotion: int,
        dateTime: dateTime,
        responses: [
            {
              id: int,
              user: {id: int, name: string},
              comment: string
            },
            ......
        ]
    ]
]
```
### ユーザ情報登録
**ユーザ情報登録(名前が重複しないようにする)**  
- APIのURL: `/user/resister`
- Method: Post
- フロントエンドから受け取るデータ    
```
name: string  
```
- フロントエンドへと渡すデータ
```
[
  {
    id: int,
    name: string
  },
]
```
**ユーザ情報編集** 
- APIのURL: /user/edit
- Method: Update
- フロントエンドから受け取るデータ  
```
newName: string
```
- フロントエンドへと渡すデータ
```
[
  {
    id: int,
    name: string
  },
]
```

## DBの設計  
### User
- id: int、bigIncrement、primary key
- name: string、unique(認証に使う)
- point: int __(追加機能)

### Emotion
- id: int、bigIncrement、primary key
- emotion: int(frontendで番号に対応した数字を)
- point: int __(追加機能)

### Comment
- id: int、bigIncrement、primary key
- user_id: foreign key
- emotion_id: foreign key
- prefecture: string
- latitude: double(緯度)
- longtitude: double(経度)
- comment: text
- dateTime: dateTime

### Response
- id: int、bigIncrement、primary key
- user_id: foreign key
- comment_id: foreign key
- comment: text
- dateTime: dateTime

### 関係性
- User Comment: 一対多
- User Response: 一対多
- Comment Response: 一対多
- Comment Emotion: 一対一
