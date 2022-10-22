# GORM DuckDB Driver
A DuckDB driver for [gorm.io](https://gorm.io/).
## Quickstart
```go
import (
    "github.com/c0deltin/duckdb-driver"
    "gorm.io/gorm"
)

db, err := gorm.Open(duckdb.Open("path/to/database.db"))
```

## Datatypes
### Lists
````go
type Entity struct {
	// use types.StringArray for a list of strings
	StringList types.StringArray `gorm:"type:varchar[]"`
	// use types.Int32Array for a list of integers
	IntList    types.Int32Array  `gorm:"type:integer[]"`
}
````


## Notice
> :warning: **This repository is in a non-stable status.**   Some features of DuckDB are not implemented yet or untested.

This package imports [github.com/marcboeker/go-duckdb](https://github.com/marcboeker/go-duckdb).    
Please take care of the instructions there, before compiling.   