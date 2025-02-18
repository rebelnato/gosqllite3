# gosqlite3

[![Go Reference](https://pkg.go.dev/badge/github.com/rebelnato/gosqlite3.svg)](https://pkg.go.dev/github.com/rebelnato/gosqlite3)
[![License: GPL-3.0](https://img.shields.io/badge/License-GPL-3)](https://github.com/rebelnato/gosqlite3/blob/main/LICENSE)


`gosqlite3` provides basic CRUD functionalities and a predefined users tables that contains below columns

| Column name | Type |
|-------------|------|
| Id | INTEGER PRIMARY KEY AUTOINCREMENT |
| username | TEXT NOT NULL UNIQUE |
| password | TEXT |

## Getting started with gosqlite3 package <mark>[Critical step]</mark>

1. Create "db/config" folder under parent program folder .
2. Create "config.yml" file inside config dir .
3. Add below config in the "config.yml" file . Please note that the file can be updated based on users requirement .

```go
dbConfig:
  path: "./db" # Path where the db will be created and stored
  name: "mydb.db" # Name of the db
```


[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Frebelnato%2Fgosqlite3.svg?type=large&issueType=license)](https://app.fossa.com/projects/git%2Bgithub.com%2Frebelnato%2Fgosqlite3?ref=badge_large&issueType=license)