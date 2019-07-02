# gdr

CLI tool for Google Drive

## Usage

### `sheets values get`

```
$ gdr sheets values get --service-account-file '...' --range 'A1:C3' --sheet-id '...'
```

### `sheets values update`

```
$ echo '{"values": [["1"]]}' | gdr sheets values update --service-account-file '...' --range 'A1:C3' --sheet-id '...'
```
