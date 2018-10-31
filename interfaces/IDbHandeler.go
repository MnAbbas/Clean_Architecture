package interfaces

type IDbHandeler interface {
	Execute(statement string) error
	Query(statement string) (IRow, error)
}

type IRow interface {
	Scan(dst ...interface{}) error
	Next() bool
}
