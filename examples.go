//Package inaccessible2godoc demonstrates all the cases
//where an unexported type is part of the public API of a package
//but is not shown by godoc.
//
//It does not actually do anything.
package inaccessible2godoc

//unexportedProperly should only be shown when ?m=all.
//Everything else should be in some way accessible by default.
var unexportedProperly = 0

//unexportedButUsedInAVar is not accessible.
type unexportedButUsedInAVar struct{}

//Foo is inaccessible.
func (unexportedButUsedInAVar) Foo() {}

//ExportedVarOfUnexportedButUsedInAVar is exported but the docs and methods of its type are inaccessible.
var ExportedVarOfUnexportedButUsedInAVar = unexportedButUsedInAVar{}

//unexportedButUsedInAVar is not accessible.
type unexportedButUsedInAVar2 struct{}

//unexportedButUsedInAVar3 is not accessible.
type unexportedButUsedInAVar3 struct{}

//unexportedAndNotUsedInAVar should not be accessible.
type unexportedAndNotUsedInAVar struct{}

//Ex2 and Ex3 are exported so their types should be as well,
//but unexportedAndNotUsedInAVar should not be since it doesn't escape
//to the public API.
var Ex2, Ex3, _ = func() (ex2 unexportedButUsedInAVar2, ex3 unexportedButUsedInAVar3, ex4 unexportedAndNotUsedInAVar) {
	return
}()

//unexportedButUsedInAConst is not accessible.
type unexportedButUsedInAConst int

//Foo is inaccessible.
func (unexportedButUsedInAConst) Foo() {}

//ExportedConstOfUnexportedButUsedInAConst is exported but the docs and methods of its type are inaccessible.
const ExportedConstOfUnexportedButUsedInAConst = unexportedButUsedInAConst(0)

//unexportedConstCorner is not accessible to godoc.
type unexportedConstCorner int

const (
	_ unexportedConstCorner = iota
	//ExportedConstCornerCase was created with an unexported type and iota
	ExportedConstCornerCase
)

//unexportedButUsedInAField is not accessible.
type unexportedButUsedInAField struct{}

//Foo is inaccessible.
func (unexportedButUsedInAField) Foo() {}

//UsesUnexportedInAField uses an unexported type in a field.
type UsesUnexportedInAField struct {
	//Nodocs is exported but the docs and methods of its type are inaccessible and unlinked.
	Nodocs unexportedButUsedInAField
}

//unexportedButEmbedded is not accessible.
type unexportedButEmbedded struct {
	//Exported is not accessible.
	Exported int
}

//Foo is now accessible on EmbedsExportedType as godoc handles this as a special case,
//but it is still inaccessible on EmbedsUnexportedTypeRecursive.
func (unexportedButEmbedded) Foo() {}

//EmbedsUnexportedType embeds an unexported type that contains exported fields which are not listed.
type EmbedsUnexportedType struct {
	unexportedButEmbedded
}

//Bar and Foo is accessible.
func (EmbedsUnexportedType) Bar() {}

//unexportedRecursive is not accessible.
//
//It should only be accessible because it's embedded in an exported struct and embeds a struct which has exported fields and methods.
type unexportedRecursive struct {
	unexportedButEmbedded
}

//EmbedsUnexportedTypeRecursive embeds an unexported type that embeds a type that contains exported methods and fields.
type EmbedsUnexportedTypeRecursive struct {
	//unexportedRecursively embeds a type that should be accessible because it embeds a type that should be accessible.
	unexportedRecursive
}

//unexportedButEmbeddedInterface is not accessible.
type unexportedButEmbeddedInterface interface {
	Foo()
}

//EmbedsUnexportedInterface embeds an unexported interface.
//
//Every inaccessible type in this package satisfies this interface, but you wouldn't know by looking.
type EmbedsUnexportedInterface interface {
	unexportedButEmbeddedInterface
}

//unexportedButReturned is not accessible.
type unexportedButReturned struct{}

//Foo is inaccessible.
func (unexportedButReturned) Foo() {}

//ReturnUnexportedButReturned returns an unexported type which is inaccessible.
//
//Note: golint warns about this case.
func ReturnUnexportedButReturned() *unexportedButReturned {
	return nil
}

//unexportedButUsedInParameter is not accessible.
type unexportedButUsedInParameter int

//Foo is inaccessible.
func (unexportedButUsedInParameter) Foo() {}

//UsesUnexportedAsParameter takes an unexported type as a parameter.
func UsesUnexportedAsParameter(a unexportedButUsedInParameter) {}

var (
	_ EmbedsUnexportedInterface = unexportedButUsedInAVar{}
	_ EmbedsUnexportedInterface = unexportedButUsedInAConst(0)
	_ EmbedsUnexportedInterface = unexportedButUsedInAField{}
	_ EmbedsUnexportedInterface = unexportedButEmbedded{}
	_ EmbedsUnexportedInterface = unexportedRecursive{}
	_ EmbedsUnexportedInterface = unexportedButReturned{}
	_ EmbedsUnexportedInterface = unexportedButUsedInParameter(0)
)

//unexportedButAliased is not accessible.
type unexportedButAliased struct{}

//AliasesUnexportedType is accessible but unexportedButAliased is not.
type AliasesUnexportedType = unexportedButAliased
