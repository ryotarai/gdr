# gdr

CLI tool for Google Drive

## Usage

### `sheets values get`

```
$ gdr sheets values get --service-account-file '...' --range 'A1:C3' --sheet-id '...'
{"majorDimension":"ROWS","range":"'Sheet1!A1:C3","values":[["1"]]}
```

### `sheets values update`

```
$ echo '{"values": [["1"]]}' | gdr sheets values update --service-account-file '...' --range 'A1:C3' --sheet-id '...'
```
