# post message

```http request
POST /api/v1/messages/post
```

keys:

- none base:
    - user_id: string `wxid_xxxxxx` <font color=red>__!!!force need!!!__</font>
    - type: new_order, finish_order, del_order <font color=red>__!!!force need!!!__</font>
    - full_order_id: number


- base on: new_order
    - selected_date: date `1919-08-10` <font color=red>__!!!force need!!!__</font>
    - selected_time: time `11:45:14 base on UTC+8 Asia/beijing time` <font color=red>__!!!force need!!!__</font>
    - press_doctor_wxid: string <font color=red>__!!!force need!!!__</font>


- base on: finish_order
    - press_doctor_wxid: string <font color=red>__!!!force need!!!__</font>


- base on: del_order
    - press_doctor_wxid: string

## Example:

### add order

request body:

```json
{
  "user_id": "wxid_haoersenpai",
  "type": "new_order",
  "full_order_id": 114514,
  "selected_date": "1919-08-10",
  "selected_time": "11:45:14",
  "press_doctor_wxid": "wxid_xxxxxx"
}
```

response body:

- if success:

```json
{
  "success": true,
  "full_order_id": 114514,
  "short_order_id": 64
}
```

- if fail:

```json
{
  "success": false,
  "error": "error message",
  "error_code": "error code"
}
```

### finish order

request body:

```json
{
  "user_id": "wxid_haoersenpai",
  "type": "finish_order",
  "selected_date": "1919-08-10",
  "selected_time": "11:45:14",
  "full_order_id": 114514,
  "press_doctor_wxid": "wxid_xxxxxx"
}
```

response body:

- if success:

```json
{
  "success": true
}
```

- if fail:

```json
{
  "success": false,
  "error": "error message",
  "error_code": "error code"
}
```

### del order

request body:

```json
{
  "user_id": "wxid_haoersenpai",
  "type": "del_order",
  "full_order_id": 114514,
  "press_doctor_wxid": "wxid_xxxxxx"
}
```

response body:

- if success:

```json
{
  "success": true
}
```

- if fail:

```json
{
  "success": false,
  "error": "error message",
  "error_code": "error code"
}
```