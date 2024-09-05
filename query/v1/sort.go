package query

func UseSort[I any](x []*Sort, tx IOrder[I]) I {
	arr := make([]string, 0, len(x))

	for _, v := range x {
		if key := v.sqlOrderKey(); key != "" {
			arr = append(arr, key)
		}
	}

	if len(arr) == 0 {
		return tx.(I)
	}

	return tx.Order(arr...)
}

func (x *Sort) sqlOrderKey() string {
	switch v := x.GetKey().(type) {
	case *Sort_Desc:
		return v.Desc + " DESC"
	case *Sort_Asc:
		return v.Asc
	}
	return ""
}
