// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __genPackageInFile_struct_with_unions_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.Type{
		types.MakeStructTypeRef("StructWithUnions",
			[]types.Field{
				types.Field{"a", types.MakeTypeRef(ref.Ref{}, 1), false},
				types.Field{"d", types.MakeTypeRef(ref.Ref{}, 2), false},
			},
			types.Choices{},
		),
		types.MakeStructTypeRef("",
			[]types.Field{},
			types.Choices{
				types.Field{"b", types.MakePrimitiveTypeRef(types.Float64Kind), false},
				types.Field{"c", types.MakePrimitiveTypeRef(types.StringKind), false},
			},
		),
		types.MakeStructTypeRef("",
			[]types.Field{},
			types.Choices{
				types.Field{"e", types.MakePrimitiveTypeRef(types.Float64Kind), false},
				types.Field{"f", types.MakePrimitiveTypeRef(types.StringKind), false},
			},
		),
	}, []ref.Ref{})
	__genPackageInFile_struct_with_unions_CachedRef = types.RegisterPackage(&p)
}

// StructWithUnions

type StructWithUnions struct {
	_a __unionOfBOfFloat64AndCOfString
	_d __unionOfEOfFloat64AndFOfString

	ref *ref.Ref
}

func NewStructWithUnions() StructWithUnions {
	return StructWithUnions{
		_a: New__unionOfBOfFloat64AndCOfString(),
		_d: New__unionOfEOfFloat64AndFOfString(),

		ref: &ref.Ref{},
	}
}

type StructWithUnionsDef struct {
	A __unionOfBOfFloat64AndCOfStringDef
	D __unionOfEOfFloat64AndFOfStringDef
}

func (def StructWithUnionsDef) New() StructWithUnions {
	return StructWithUnions{
		_a:  def.A.New(),
		_d:  def.D.New(),
		ref: &ref.Ref{},
	}
}

func (s StructWithUnions) Def() (d StructWithUnionsDef) {
	d.A = s._a.Def()
	d.D = s._d.Def()
	return
}

var __typeRefForStructWithUnions types.Type

func (m StructWithUnions) Type() types.Type {
	return __typeRefForStructWithUnions
}

func init() {
	__typeRefForStructWithUnions = types.MakeTypeRef(__genPackageInFile_struct_with_unions_CachedRef, 0)
	types.RegisterStruct(__typeRefForStructWithUnions, builderForStructWithUnions, readerForStructWithUnions)
}

func builderForStructWithUnions(values []types.Value) types.Value {
	i := 0
	s := StructWithUnions{ref: &ref.Ref{}}
	s._a = values[i].(__unionOfBOfFloat64AndCOfString)
	i++
	s._d = values[i].(__unionOfEOfFloat64AndFOfString)
	i++
	return s
}

func readerForStructWithUnions(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(StructWithUnions)
	values = append(values, s._a)
	values = append(values, s._d)
	return values
}

func (s StructWithUnions) Equals(other types.Value) bool {
	return other != nil && __typeRefForStructWithUnions.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s StructWithUnions) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s StructWithUnions) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForStructWithUnions.Chunks()...)
	chunks = append(chunks, s._a.Chunks()...)
	chunks = append(chunks, s._d.Chunks()...)
	return
}

func (s StructWithUnions) ChildValues() (ret []types.Value) {
	ret = append(ret, s._a)
	ret = append(ret, s._d)
	return
}

func (s StructWithUnions) A() __unionOfBOfFloat64AndCOfString {
	return s._a
}

func (s StructWithUnions) SetA(val __unionOfBOfFloat64AndCOfString) StructWithUnions {
	s._a = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructWithUnions) D() __unionOfEOfFloat64AndFOfString {
	return s._d
}

func (s StructWithUnions) SetD(val __unionOfEOfFloat64AndFOfString) StructWithUnions {
	s._d = val
	s.ref = &ref.Ref{}
	return s
}

// __unionOfBOfFloat64AndCOfString

type __unionOfBOfFloat64AndCOfString struct {
	__unionIndex uint32
	__unionValue types.Value
	ref          *ref.Ref
}

func New__unionOfBOfFloat64AndCOfString() __unionOfBOfFloat64AndCOfString {
	return __unionOfBOfFloat64AndCOfString{
		__unionIndex: 0,
		__unionValue: types.Float64(0),
		ref:          &ref.Ref{},
	}
}

type __unionOfBOfFloat64AndCOfStringDef struct {
	__unionIndex uint32
	__unionValue types.Value
}

func (def __unionOfBOfFloat64AndCOfStringDef) New() __unionOfBOfFloat64AndCOfString {
	return __unionOfBOfFloat64AndCOfString{
		__unionIndex: def.__unionIndex,
		__unionValue: def.__unionValue,
		ref:          &ref.Ref{},
	}
}

func (s __unionOfBOfFloat64AndCOfString) Def() (d __unionOfBOfFloat64AndCOfStringDef) {
	d.__unionIndex = s.__unionIndex
	d.__unionValue = s.__unionValue
	return
}

var __typeRefFor__unionOfBOfFloat64AndCOfString types.Type

func (m __unionOfBOfFloat64AndCOfString) Type() types.Type {
	return __typeRefFor__unionOfBOfFloat64AndCOfString
}

func init() {
	__typeRefFor__unionOfBOfFloat64AndCOfString = types.MakeTypeRef(__genPackageInFile_struct_with_unions_CachedRef, 1)
	types.RegisterStruct(__typeRefFor__unionOfBOfFloat64AndCOfString, builderFor__unionOfBOfFloat64AndCOfString, readerFor__unionOfBOfFloat64AndCOfString)
}

func builderFor__unionOfBOfFloat64AndCOfString(values []types.Value) types.Value {
	i := 0
	s := __unionOfBOfFloat64AndCOfString{ref: &ref.Ref{}}
	s.__unionIndex = uint32(values[i].(types.UInt32))
	i++
	s.__unionValue = values[i]
	i++
	return s
}

func readerFor__unionOfBOfFloat64AndCOfString(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(__unionOfBOfFloat64AndCOfString)
	values = append(values, types.UInt32(s.__unionIndex))
	values = append(values, s.__unionValue)
	return values
}

func (s __unionOfBOfFloat64AndCOfString) Equals(other types.Value) bool {
	return other != nil && __typeRefFor__unionOfBOfFloat64AndCOfString.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s __unionOfBOfFloat64AndCOfString) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s __unionOfBOfFloat64AndCOfString) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefFor__unionOfBOfFloat64AndCOfString.Chunks()...)
	chunks = append(chunks, s.__unionValue.Chunks()...)
	return
}

func (s __unionOfBOfFloat64AndCOfString) ChildValues() (ret []types.Value) {
	ret = append(ret, s.__unionValue)
	return
}

func (s __unionOfBOfFloat64AndCOfString) B() (val float64, ok bool) {
	if s.__unionIndex != 0 {
		return
	}
	return float64(s.__unionValue.(types.Float64)), true
}

func (s __unionOfBOfFloat64AndCOfString) SetB(val float64) __unionOfBOfFloat64AndCOfString {
	s.__unionIndex = 0
	s.__unionValue = types.Float64(val)
	s.ref = &ref.Ref{}
	return s
}

func (def __unionOfBOfFloat64AndCOfStringDef) B() (val float64, ok bool) {
	if def.__unionIndex != 0 {
		return
	}
	return float64(def.__unionValue.(types.Float64)), true
}

func (def __unionOfBOfFloat64AndCOfStringDef) SetB(val float64) __unionOfBOfFloat64AndCOfStringDef {
	def.__unionIndex = 0
	def.__unionValue = types.Float64(val)
	return def
}

func (s __unionOfBOfFloat64AndCOfString) C() (val string, ok bool) {
	if s.__unionIndex != 1 {
		return
	}
	return s.__unionValue.(types.String).String(), true
}

func (s __unionOfBOfFloat64AndCOfString) SetC(val string) __unionOfBOfFloat64AndCOfString {
	s.__unionIndex = 1
	s.__unionValue = types.NewString(val)
	s.ref = &ref.Ref{}
	return s
}

func (def __unionOfBOfFloat64AndCOfStringDef) C() (val string, ok bool) {
	if def.__unionIndex != 1 {
		return
	}
	return def.__unionValue.(types.String).String(), true
}

func (def __unionOfBOfFloat64AndCOfStringDef) SetC(val string) __unionOfBOfFloat64AndCOfStringDef {
	def.__unionIndex = 1
	def.__unionValue = types.NewString(val)
	return def
}

// __unionOfEOfFloat64AndFOfString

type __unionOfEOfFloat64AndFOfString struct {
	__unionIndex uint32
	__unionValue types.Value
	ref          *ref.Ref
}

func New__unionOfEOfFloat64AndFOfString() __unionOfEOfFloat64AndFOfString {
	return __unionOfEOfFloat64AndFOfString{
		__unionIndex: 0,
		__unionValue: types.Float64(0),
		ref:          &ref.Ref{},
	}
}

type __unionOfEOfFloat64AndFOfStringDef struct {
	__unionIndex uint32
	__unionValue types.Value
}

func (def __unionOfEOfFloat64AndFOfStringDef) New() __unionOfEOfFloat64AndFOfString {
	return __unionOfEOfFloat64AndFOfString{
		__unionIndex: def.__unionIndex,
		__unionValue: def.__unionValue,
		ref:          &ref.Ref{},
	}
}

func (s __unionOfEOfFloat64AndFOfString) Def() (d __unionOfEOfFloat64AndFOfStringDef) {
	d.__unionIndex = s.__unionIndex
	d.__unionValue = s.__unionValue
	return
}

var __typeRefFor__unionOfEOfFloat64AndFOfString types.Type

func (m __unionOfEOfFloat64AndFOfString) Type() types.Type {
	return __typeRefFor__unionOfEOfFloat64AndFOfString
}

func init() {
	__typeRefFor__unionOfEOfFloat64AndFOfString = types.MakeTypeRef(__genPackageInFile_struct_with_unions_CachedRef, 2)
	types.RegisterStruct(__typeRefFor__unionOfEOfFloat64AndFOfString, builderFor__unionOfEOfFloat64AndFOfString, readerFor__unionOfEOfFloat64AndFOfString)
}

func builderFor__unionOfEOfFloat64AndFOfString(values []types.Value) types.Value {
	i := 0
	s := __unionOfEOfFloat64AndFOfString{ref: &ref.Ref{}}
	s.__unionIndex = uint32(values[i].(types.UInt32))
	i++
	s.__unionValue = values[i]
	i++
	return s
}

func readerFor__unionOfEOfFloat64AndFOfString(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(__unionOfEOfFloat64AndFOfString)
	values = append(values, types.UInt32(s.__unionIndex))
	values = append(values, s.__unionValue)
	return values
}

func (s __unionOfEOfFloat64AndFOfString) Equals(other types.Value) bool {
	return other != nil && __typeRefFor__unionOfEOfFloat64AndFOfString.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s __unionOfEOfFloat64AndFOfString) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s __unionOfEOfFloat64AndFOfString) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefFor__unionOfEOfFloat64AndFOfString.Chunks()...)
	chunks = append(chunks, s.__unionValue.Chunks()...)
	return
}

func (s __unionOfEOfFloat64AndFOfString) ChildValues() (ret []types.Value) {
	ret = append(ret, s.__unionValue)
	return
}

func (s __unionOfEOfFloat64AndFOfString) E() (val float64, ok bool) {
	if s.__unionIndex != 0 {
		return
	}
	return float64(s.__unionValue.(types.Float64)), true
}

func (s __unionOfEOfFloat64AndFOfString) SetE(val float64) __unionOfEOfFloat64AndFOfString {
	s.__unionIndex = 0
	s.__unionValue = types.Float64(val)
	s.ref = &ref.Ref{}
	return s
}

func (def __unionOfEOfFloat64AndFOfStringDef) E() (val float64, ok bool) {
	if def.__unionIndex != 0 {
		return
	}
	return float64(def.__unionValue.(types.Float64)), true
}

func (def __unionOfEOfFloat64AndFOfStringDef) SetE(val float64) __unionOfEOfFloat64AndFOfStringDef {
	def.__unionIndex = 0
	def.__unionValue = types.Float64(val)
	return def
}

func (s __unionOfEOfFloat64AndFOfString) F() (val string, ok bool) {
	if s.__unionIndex != 1 {
		return
	}
	return s.__unionValue.(types.String).String(), true
}

func (s __unionOfEOfFloat64AndFOfString) SetF(val string) __unionOfEOfFloat64AndFOfString {
	s.__unionIndex = 1
	s.__unionValue = types.NewString(val)
	s.ref = &ref.Ref{}
	return s
}

func (def __unionOfEOfFloat64AndFOfStringDef) F() (val string, ok bool) {
	if def.__unionIndex != 1 {
		return
	}
	return def.__unionValue.(types.String).String(), true
}

func (def __unionOfEOfFloat64AndFOfStringDef) SetF(val string) __unionOfEOfFloat64AndFOfStringDef {
	def.__unionIndex = 1
	def.__unionValue = types.NewString(val)
	return def
}