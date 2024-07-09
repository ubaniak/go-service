package db

func Build() (*ConnectionRegistry, error) {
	r := NewConnectionRegistry()

	sqlite, err := NewConnectionBuilder().SetDriver(SqLite).SetDatabase("test.db").Build()
	if err != nil {
		return nil, err
	}

	r.Register("Test", sqlite)

	return r, nil
}
