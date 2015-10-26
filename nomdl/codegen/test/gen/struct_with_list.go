// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __genPackageInFile_struct_with_list_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("StructWithList",
			[]types.Field{
				types.Field{"l", types.MakeCompoundTypeRef(types.ListKind, types.MakePrimitiveTypeRef(types.UInt8Kind)), false},
				types.Field{"b", types.MakePrimitiveTypeRef(types.BoolKind), false},
				types.Field{"s", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"i", types.MakePrimitiveTypeRef(types.Int64Kind), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	__genPackageInFile_struct_with_list_CachedRef = types.RegisterPackage(&p)
}

// StructWithList

type StructWithList struct {
	m   types.Map
	ref *ref.Ref
}

func NewStructWithList() StructWithList {
	return StructWithList{types.NewMap(
		types.NewString("l"), NewListOfUInt8(),
		types.NewString("b"), types.Bool(false),
		types.NewString("s"), types.NewString(""),
		types.NewString("i"), types.Int64(0),
	), &ref.Ref{}}
}

type StructWithListDef struct {
	L ListOfUInt8Def
	B bool
	S string
	I int64
}

func (def StructWithListDef) New() StructWithList {
	return StructWithList{
		types.NewMap(
			types.NewString("l"), def.L.New(),
			types.NewString("b"), types.Bool(def.B),
			types.NewString("s"), types.NewString(def.S),
			types.NewString("i"), types.Int64(def.I),
		), &ref.Ref{}}
}

func (s StructWithList) Def() (d StructWithListDef) {
	d.L = s.m.Get(types.NewString("l")).(ListOfUInt8).Def()
	d.B = bool(s.m.Get(types.NewString("b")).(types.Bool))
	d.S = s.m.Get(types.NewString("s")).(types.String).String()
	d.I = int64(s.m.Get(types.NewString("i")).(types.Int64))
	return
}

var __typeRefForStructWithList types.TypeRef

func (m StructWithList) TypeRef() types.TypeRef {
	return __typeRefForStructWithList
}

func init() {
	__typeRefForStructWithList = types.MakeTypeRef(__genPackageInFile_struct_with_list_CachedRef, 0)
	types.RegisterFromValFunction(__typeRefForStructWithList, func(v types.Value) types.Value {
		return StructWithListFromVal(v)
	})
}

func StructWithListFromVal(val types.Value) StructWithList {
	// TODO: Do we still need FromVal?
	if val, ok := val.(StructWithList); ok {
		return val
	}
	// TODO: Validate here
	return StructWithList{val.(types.Map), &ref.Ref{}}
}

func (s StructWithList) InternalImplementation() types.Map {
	return s.m
}

func (s StructWithList) Equals(other types.Value) bool {
	if other, ok := other.(StructWithList); ok {
		return s.Ref() == other.Ref()
	}
	return false
}

func (s StructWithList) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s StructWithList) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s StructWithList) L() ListOfUInt8 {
	return s.m.Get(types.NewString("l")).(ListOfUInt8)
}

func (s StructWithList) SetL(val ListOfUInt8) StructWithList {
	return StructWithList{s.m.Set(types.NewString("l"), val), &ref.Ref{}}
}

func (s StructWithList) B() bool {
	return bool(s.m.Get(types.NewString("b")).(types.Bool))
}

func (s StructWithList) SetB(val bool) StructWithList {
	return StructWithList{s.m.Set(types.NewString("b"), types.Bool(val)), &ref.Ref{}}
}

func (s StructWithList) S() string {
	return s.m.Get(types.NewString("s")).(types.String).String()
}

func (s StructWithList) SetS(val string) StructWithList {
	return StructWithList{s.m.Set(types.NewString("s"), types.NewString(val)), &ref.Ref{}}
}

func (s StructWithList) I() int64 {
	return int64(s.m.Get(types.NewString("i")).(types.Int64))
}

func (s StructWithList) SetI(val int64) StructWithList {
	return StructWithList{s.m.Set(types.NewString("i"), types.Int64(val)), &ref.Ref{}}
}