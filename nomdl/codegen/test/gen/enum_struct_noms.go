// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __genPackageInFile_enum_struct_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeEnumTypeRef("Handedness", "right", "left", "switch"),
		types.MakeStructTypeRef("EnumStruct",
			[]types.Field{
				types.Field{"hand", types.MakeTypeRef(ref.Ref{}, 0), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	__genPackageInFile_enum_struct_CachedRef = types.RegisterPackage(&p)
}

// Handedness

type Handedness uint32

const (
	Right Handedness = iota
	Left
	Switch
)

func NewHandedness() Handedness {
	return Handedness(0)
}

var __typeRefForHandedness types.TypeRef

func (e Handedness) TypeRef() types.TypeRef {
	return __typeRefForHandedness
}

func init() {
	__typeRefForHandedness = types.MakeTypeRef(__genPackageInFile_enum_struct_CachedRef, 0)
	types.RegisterFromValFunction(__typeRefForHandedness, func(v types.Value) types.Value {
		return Handedness(uint32(v.(types.UInt32)))
	})
}

func (e Handedness) InternalImplementation() uint32 {
	return uint32(e)
}

func (e Handedness) Equals(other types.Value) bool {
	if other, ok := other.(Handedness); ok {
		return e == other
	}
	return false
}

func (e Handedness) Ref() ref.Ref {
	throwaway := ref.Ref{}
	return types.EnsureRef(&throwaway, e)
}

func (e Handedness) Chunks() []types.Future {
	return nil
}

// EnumStruct

type EnumStruct struct {
	m   types.Map
	ref *ref.Ref
}

func NewEnumStruct() EnumStruct {
	return EnumStruct{types.NewMap(
		types.NewString("hand"), NewHandedness(),
	), &ref.Ref{}}
}

type EnumStructDef struct {
	Hand Handedness
}

func (def EnumStructDef) New() EnumStruct {
	return EnumStruct{
		types.NewMap(
			types.NewString("hand"), def.Hand,
		), &ref.Ref{}}
}

func (s EnumStruct) Def() (d EnumStructDef) {
	d.Hand = s.m.Get(types.NewString("hand")).(Handedness)
	return
}

var __typeRefForEnumStruct types.TypeRef

func (m EnumStruct) TypeRef() types.TypeRef {
	return __typeRefForEnumStruct
}

func init() {
	__typeRefForEnumStruct = types.MakeTypeRef(__genPackageInFile_enum_struct_CachedRef, 1)
	types.RegisterFromValFunction(__typeRefForEnumStruct, func(v types.Value) types.Value {
		return EnumStructFromVal(v)
	})
}

func EnumStructFromVal(val types.Value) EnumStruct {
	// TODO: Do we still need FromVal?
	if val, ok := val.(EnumStruct); ok {
		return val
	}
	// TODO: Validate here
	return EnumStruct{val.(types.Map), &ref.Ref{}}
}

func (s EnumStruct) InternalImplementation() types.Map {
	return s.m
}

func (s EnumStruct) Equals(other types.Value) bool {
	if other, ok := other.(EnumStruct); ok {
		return s.Ref() == other.Ref()
	}
	return false
}

func (s EnumStruct) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s EnumStruct) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s EnumStruct) Hand() Handedness {
	return s.m.Get(types.NewString("hand")).(Handedness)
}

func (s EnumStruct) SetHand(val Handedness) EnumStruct {
	return EnumStruct{s.m.Set(types.NewString("hand"), val), &ref.Ref{}}
}