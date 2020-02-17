package remodel

const (
	// MySQL table column types
	BigInt    ColumnType = "bigint"
	MediumInt ColumnType = "mediumint"
	SmallInt  ColumnType = "smallint"
	TinyInt   ColumnType = "tinyint"
	Int       ColumnType = "int"
	Float     ColumnType = "float"
	Double    ColumnType = "double"
	Bit       ColumnType = "bit"
	Date      ColumnType = "date"
	Datetime  ColumnType = "datetime"
	Timestamp ColumnType = "timestamp"
	Varchar   ColumnType = "varchar"
	Blob      ColumnType = "blob"
	Text      ColumnType = "text"
	Enum      ColumnType = "enum"
	Set       ColumnType = "set"

	// Golang types
	Uint64      EntityType = "uint64"
	Uint32      EntityType = "uint32"
	Uint16      EntityType = "uint16"
	Uint8       EntityType = "uint8"
	Int64       EntityType = "int64"
	Int32       EntityType = "int32"
	Int16       EntityType = "int16"
	Int8        EntityType = "int8"
	Bool        EntityType = "bool"
	Float64     EntityType = "float64"
	Float32     EntityType = "float32"
	TimePtr     EntityType = "*time.Time"
	String      EntityType = "string"
	ByteSlice   EntityType = "[]byte"
	StringSlice EntityType = "[]string"

	// outside library for generate code
	ErrorsLib   = "github.com/juju/errors"
	LogLib      = "github.com/labstack/gommon/log"
	RapidashLib = "go.knocknote.io/rapidash"

	EntityPackageName = "entity"
	DaoPackageName    = "dao"
)
