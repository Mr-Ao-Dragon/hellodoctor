# add user

```http request
PUT /api/v1/user/add
```

keys:

- none base:
    - wxid: `wxid_1234567890` <font color=red>__!!!force need!!!__</font>
    - "access_token": `${access_token}`
    - name: `${wechat_name}`
    - avatar: `${wechat_avatar_url}`
    - phone: `${phone}`

## example

request body:

```json
{
  "wxid": "wxid_1234567890",
  "access_token": "access_token",
  "name": "name",
  "avatar": "avatar"
}
```

response body:

- if success:

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "id": 1234567890,
    "wxid": "wxid_1234567890",
    "name": "name",
    "avatar": "avatar"
  }
}
```

- if fail:

```json
{
  "code": "${error_code}",
  "msg": "${error_msg}"
}
```

# query user

```http request
GET /api/v1/user/query
```

keys:

- none base:
    - name: `${wechat_name}`
    - phone: `${phone}`

tps: just one of them is required.

## example

request body:

```json
{
  "name": "name"
}
```

or:

```json
{
  "phone": "phone"
}
```

if we have both:

```json
{
  "name": "name",
  "phone": "phone"
}
```

response body:

- if success:

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "id": 1234567890,
    "wxid": "wxid_1234567890",
    "name": "name",
    "avatar": "avatar"
  }
}
```

- if fail:

```json
{
  "code": "${error_code}",
  "msg": "${error_msg}"
}
```

# alter user

```http request
PATCH /api/v1/user/alter
```

keys:

- none base:
    - wxid: `wxid_1234567890` <font color=red>__!!!force need!!!__</font>
    - name: `${wechat_name}`
    - phone: `${phone}`
    - avatar: `${wechat_avatar_url}`

## example

request body:

```json
{
  "wxid": "wxid_1234567890",
  "name": "name",
  "phone": "phone",
  "avatar": "wechat_avatar_url"
}
```

tips: With the exception to wxid, key values will directly overwrite existing data.
response body:

- if success:

```json
{
  "code": 0,
  "msg": "success"
}
```

- if fail:

```json
{
  "code": "${error_code}",
  "msg": "${error_msg}"
}
```

# login

```http request
POST /api/v1/user/login
```

keys:

- none base:
    - wxid: `wxid_1234567890` <font color=red>__!!!force need!!!__</font>
    - access_token: `${access_token}` <font color=red>__!!!force need!!!__</font>

## example

request body:

```json
{
  "wxid": "wxid_123456789",
  "access_token": "access_token"
}
```