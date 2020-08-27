## APIの設計
### 感情が地図に載る(感情の共有)
**各県の色16進数を返す**  
- APIのURL: `/emotion/prefectures`  
- Method: GET
- フロントエンドから受け取るデータ  
```
None
```
- フロントエンドへと渡すデータ
```
[
  {
    Prefecture: string,
    Color: string
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
      ID: int,
      Emotion: int,
      Latitude: double,
      Longtitude: double,
      UserName: string,
      DateTime: dateTime,
    },
    .....
]
```
**ピンを押して、コメント内容確認(紐ついた返信も)**
- APIのURL: `/emotion/{prefecture}/comments/{comment_id}` 
- Method: Get   
- フロントエンドから受け取るデータ  
```
comment_id: int
```
- フロントエンドへと渡すデータ
```
{
  CommentContent: {
    Comment: string,
    Emotion: int,
    UserName: string
    DateTime: dateTime,
  }
  Response: [
      {
        ID: int,
        UserName: string,
        Comment: string,
        DateTime: dateTime,
      },
      ......
  ]
}
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
{
  ID: int,
  Name: string,
  Point: int
},

```
**ユーザ情報編集** 
- APIのURL: `/user/edit`
- Method: PATCH
- フロントエンドから受け取るデータ  
```
id: int
newName: string
```
- フロントエンドへと渡すデータ
```
{
  ID: int,
  Name: string,
  Point: int
}
```

### 感情の登録など
**コメントの登録**
- APIのURL: `/comment/register`
- Method: POST
- フロントエンドから受け取るデータ
```
[    
  {      
    user_id: int,
    emotion: int,   
    comment: string,   
    latitude: double,      
    longtitude: double,      
    prefectutre: string,      
    dateTime: dateTime, 
  }
]
```
- フロントエンドへと渡すデータ
```
result: boolean
```

**コメントに対する返信の保存**
- APIのURL: `/comment/response`
- Method: POST
- フロントエンドから受け取るデータ
```
[    
  {      
    user_id: int,
    comment_id: int,   
    comment: string,        
    dateTime: dateTime, 
  }
]
```
- フロントエンドへと渡すデータ
```
result: boolean
```
**コメント削除**
- APIのURL: `/comment/delete`
- Method: DELETE
- フロントエンドから受け取るデータ
```
[    
  {      
    comment_id: int
  }
]
```
- フロントエンドへと渡すデータ
```
result: boolean
```
**返信削除**
- APIのURL: `/comment/response/delete`
- Method: DELETE
- フロントエンドから受け取るデータ
```
[    
  {      
    response_id: int
  }
]
```
- フロントエンドへと渡すデータ
```
result: boolean
```
