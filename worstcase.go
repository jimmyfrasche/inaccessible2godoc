package inaccessible2godoc

type evenIfAccessibleUseless int

func f() evenIfAccessibleUseless {
	return 0
}

//X is created by f which returns an evenIfAccessibleUseless.
//This type is technically part of the package API,
//but, if it's otherwise not referenced, it would not be clear
//why this type has been made accessible as f is not.
//
//Perhaps it should be made accessible and golint could test for this case.
var X = f()
