package pointers

func ChangeInt(PointerToInt *int) {
	// temp := 50
	*PointerToInt = 50
	// whenever accepting any argument as a pointer, datatype should be *datatype(ex. PointerToInt as *int)
	//* is used to convert a pointer to actual variable

}
