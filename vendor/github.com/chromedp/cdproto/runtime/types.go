package runtime

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// ScriptID unique script identifier.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-ScriptId
type ScriptID string

// String returns the ScriptID as string value.
func (t ScriptID) String() string {
	return string(t)
}

// RemoteObjectID unique object identifier.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-RemoteObjectId
type RemoteObjectID string

// String returns the RemoteObjectID as string value.
func (t RemoteObjectID) String() string {
	return string(t)
}

// UnserializableValue primitive value which cannot be JSON-stringified.
// Includes values -0, NaN, Infinity, -Infinity, and bigint literals.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-UnserializableValue
type UnserializableValue string

// String returns the UnserializableValue as string value.
func (t UnserializableValue) String() string {
	return string(t)
}

// RemoteObject mirror object referencing original JavaScript object.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-RemoteObject
type RemoteObject struct {
	Type                Type                `json:"type"`                          // Object type.
	Subtype             Subtype             `json:"subtype,omitempty"`             // Object subtype hint. Specified for object type values only. NOTE: If you change anything here, make sure to also update subtype in ObjectPreview and PropertyPreview below.
	ClassName           string              `json:"className,omitempty"`           // Object class (constructor) name. Specified for object type values only.
	Value               easyjson.RawMessage `json:"value,omitempty"`               // Remote object value in case of primitive values or JSON values (if it was requested).
	UnserializableValue UnserializableValue `json:"unserializableValue,omitempty"` // Primitive value which can not be JSON-stringified does not have value, but gets this property.
	Description         string              `json:"description,omitempty"`         // String representation of the object.
	ObjectID            RemoteObjectID      `json:"objectId,omitempty"`            // Unique object identifier (for non-primitive values).
	Preview             *ObjectPreview      `json:"preview,omitempty"`             // Preview containing abbreviated property values. Specified for object type values only.
	CustomPreview       *CustomPreview      `json:"customPreview,omitempty"`
}

// CustomPreview [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-CustomPreview
type CustomPreview struct {
	Header       string         `json:"header"`                 // The JSON-stringified result of formatter.header(object, config) call. It contains json ML array that represents RemoteObject.
	BodyGetterID RemoteObjectID `json:"bodyGetterId,omitempty"` // If formatter returns true as a result of formatter.hasBody call then bodyGetterId will contain RemoteObjectId for the function that returns result of formatter.body(object, config) call. The result value is json ML array.
}

// ObjectPreview object containing abbreviated remote object value.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-ObjectPreview
type ObjectPreview struct {
	Type        Type               `json:"type"`                  // Object type.
	Subtype     Subtype            `json:"subtype,omitempty"`     // Object subtype hint. Specified for object type values only.
	Description string             `json:"description,omitempty"` // String representation of the object.
	Overflow    bool               `json:"overflow"`              // True iff some of the properties or entries of the original object did not fit.
	Properties  []*PropertyPreview `json:"properties"`            // List of the properties.
	Entries     []*EntryPreview    `json:"entries,omitempty"`     // List of the entries. Specified for map and set subtype values only.
}

// PropertyPreview [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-PropertyPreview
type PropertyPreview struct {
	Name         string         `json:"name"`                   // Property name.
	Type         Type           `json:"type"`                   // Object type. Accessor means that the property itself is an accessor property.
	Value        string         `json:"value,omitempty"`        // User-friendly property value string.
	ValuePreview *ObjectPreview `json:"valuePreview,omitempty"` // Nested value preview.
	Subtype      Subtype        `json:"subtype,omitempty"`      // Object subtype hint. Specified for object type values only.
}

// EntryPreview [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-EntryPreview
type EntryPreview struct {
	Key   *ObjectPreview `json:"key,omitempty"` // Preview of the key. Specified for map-like collection entries.
	Value *ObjectPreview `json:"value"`         // Preview of the value.
}

// PropertyDescriptor object property descriptor.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-PropertyDescriptor
type PropertyDescriptor struct {
	Name         string        `json:"name"`                // Property name or symbol description.
	Value        *RemoteObject `json:"value,omitempty"`     // The value associated with the property.
	Writable     bool          `json:"writable,omitempty"`  // True if the value associated with the property may be changed (data descriptors only).
	Get          *RemoteObject `json:"get,omitempty"`       // A function which serves as a getter for the property, or undefined if there is no getter (accessor descriptors only).
	Set          *RemoteObject `json:"set,omitempty"`       // A function which serves as a setter for the property, or undefined if there is no setter (accessor descriptors only).
	Configurable bool          `json:"configurable"`        // True if the type of this property descriptor may be changed and if the property may be deleted from the corresponding object.
	Enumerable   bool          `json:"enumerable"`          // True if this property shows up during enumeration of the properties on the corresponding object.
	WasThrown    bool          `json:"wasThrown,omitempty"` // True if the result was thrown during the evaluation.
	IsOwn        bool          `json:"isOwn,omitempty"`     // True if the property is owned for the object.
	Symbol       *RemoteObject `json:"symbol,omitempty"`    // Property symbol object, if the property is of the symbol type.
}

// InternalPropertyDescriptor object internal property descriptor. This
// property isn't normally visible in JavaScript code.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-InternalPropertyDescriptor
type InternalPropertyDescriptor struct {
	Name  string        `json:"name"`            // Conventional property name.
	Value *RemoteObject `json:"value,omitempty"` // The value associated with the property.
}

// PrivatePropertyDescriptor object private field descriptor.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-PrivatePropertyDescriptor
type PrivatePropertyDescriptor struct {
	Name  string        `json:"name"`            // Private property name.
	Value *RemoteObject `json:"value,omitempty"` // The value associated with the private property.
	Get   *RemoteObject `json:"get,omitempty"`   // A function which serves as a getter for the private property, or undefined if there is no getter (accessor descriptors only).
	Set   *RemoteObject `json:"set,omitempty"`   // A function which serves as a setter for the private property, or undefined if there is no setter (accessor descriptors only).
}

// CallArgument represents function call argument. Either remote object id
// objectId, primitive value, unserializable primitive value or neither of (for
// undefined) them should be specified.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-CallArgument
type CallArgument struct {
	Value               easyjson.RawMessage `json:"value,omitempty"`               // Primitive value or serializable javascript object.
	UnserializableValue UnserializableValue `json:"unserializableValue,omitempty"` // Primitive value which can not be JSON-stringified.
	ObjectID            RemoteObjectID      `json:"objectId,omitempty"`            // Remote object handle.
}

// ExecutionContextID ID of an execution context.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-ExecutionContextId
type ExecutionContextID int64

// Int64 returns the ExecutionContextID as int64 value.
func (t ExecutionContextID) Int64() int64 {
	return int64(t)
}

// ExecutionContextDescription description of an isolated world.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-ExecutionContextDescription
type ExecutionContextDescription struct {
	ID       ExecutionContextID  `json:"id"`       // Unique id of the execution context. It can be used to specify in which execution context script evaluation should be performed.
	Origin   string              `json:"origin"`   // Execution context origin.
	Name     string              `json:"name"`     // Human readable name describing given context.
	UniqueID string              `json:"uniqueId"` // A system-unique execution context identifier. Unlike the id, this is unique across multiple processes, so can be reliably used to identify specific context while backend performs a cross-process navigation.
	AuxData  easyjson.RawMessage `json:"auxData,omitempty"`
}

// ExceptionDetails detailed information about exception (or error) that was
// thrown during script compilation or execution.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-ExceptionDetails
type ExceptionDetails struct {
	ExceptionID        int64               `json:"exceptionId"`                  // Exception id.
	Text               string              `json:"text"`                         // Exception text, which should be used together with exception object when available.
	LineNumber         int64               `json:"lineNumber"`                   // Line number of the exception location (0-based).
	ColumnNumber       int64               `json:"columnNumber"`                 // Column number of the exception location (0-based).
	ScriptID           ScriptID            `json:"scriptId,omitempty"`           // Script ID of the exception location.
	URL                string              `json:"url,omitempty"`                // URL of the exception location, to be used when the script was not reported.
	StackTrace         *StackTrace         `json:"stackTrace,omitempty"`         // JavaScript stack trace if available.
	Exception          *RemoteObject       `json:"exception,omitempty"`          // Exception object if available.
	ExecutionContextID ExecutionContextID  `json:"executionContextId,omitempty"` // Identifier of the context where exception happened.
	ExceptionMetaData  easyjson.RawMessage `json:"exceptionMetaData,omitempty"`
}

// Error satisfies the error interface.
func (e *ExceptionDetails) Error() string {
	var b strings.Builder
	// TODO: watch script parsed events and match the ExceptionDetails.ScriptID
	// to the name/location of the actual code and display here
	fmt.Fprintf(&b, "exception %q (%d:%d)", e.Text, e.LineNumber, e.ColumnNumber)
	if obj := e.Exception; obj != nil {
		fmt.Fprintf(&b, ": %s", obj.Description)
	}
	return b.String()
}

// Timestamp number of milliseconds since epoch.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-Timestamp
type Timestamp time.Time

// Time returns the Timestamp as time.Time value.
func (t Timestamp) Time() time.Time {
	return time.Time(t)
}

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t Timestamp) MarshalEasyJSON(out *jwriter.Writer) {
	v := float64(time.Time(t).UnixNano() / int64(time.Millisecond))

	out.Buffer.EnsureSpace(20)
	out.Buffer.Buf = strconv.AppendFloat(out.Buffer.Buf, v, 'f', -1, 64)
}

// MarshalJSON satisfies json.Marshaler.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *Timestamp) UnmarshalEasyJSON(in *jlexer.Lexer) {
	*t = Timestamp(time.Unix(0, int64(in.Float64()*float64(time.Millisecond))))
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *Timestamp) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// TimeDelta number of milliseconds.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-TimeDelta
type TimeDelta float64

// Float64 returns the TimeDelta as float64 value.
func (t TimeDelta) Float64() float64 {
	return float64(t)
}

// CallFrame stack entry for runtime errors and assertions.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-CallFrame
type CallFrame struct {
	FunctionName string   `json:"functionName"` // JavaScript function name.
	ScriptID     ScriptID `json:"scriptId"`     // JavaScript script id.
	URL          string   `json:"url"`          // JavaScript script name or url.
	LineNumber   int64    `json:"lineNumber"`   // JavaScript script line number (0-based).
	ColumnNumber int64    `json:"columnNumber"` // JavaScript script column number (0-based).
}

// StackTrace call frames for assertions or error messages.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-StackTrace
type StackTrace struct {
	Description string        `json:"description,omitempty"` // String label of this stack trace. For async traces this may be a name of the function that initiated the async call.
	CallFrames  []*CallFrame  `json:"callFrames"`            // JavaScript function name.
	Parent      *StackTrace   `json:"parent,omitempty"`      // Asynchronous JavaScript stack trace that preceded this stack, if available.
	ParentID    *StackTraceID `json:"parentId,omitempty"`    // Asynchronous JavaScript stack trace that preceded this stack, if available.
}

// UniqueDebuggerID unique identifier of current debugger.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-UniqueDebuggerId
type UniqueDebuggerID string

// String returns the UniqueDebuggerID as string value.
func (t UniqueDebuggerID) String() string {
	return string(t)
}

// StackTraceID if debuggerId is set stack trace comes from another debugger
// and can be resolved there. This allows to track cross-debugger calls. See
// Runtime.StackTrace and Debugger.paused for usages.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-StackTraceId
type StackTraceID struct {
	ID         string           `json:"id"`
	DebuggerID UniqueDebuggerID `json:"debuggerId,omitempty"`
}

// Type object type.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-RemoteObject
type Type string

// String returns the Type as string value.
func (t Type) String() string {
	return string(t)
}

// Type values.
const (
	TypeObject    Type = "object"
	TypeFunction  Type = "function"
	TypeUndefined Type = "undefined"
	TypeString    Type = "string"
	TypeNumber    Type = "number"
	TypeBoolean   Type = "boolean"
	TypeSymbol    Type = "symbol"
	TypeBigint    Type = "bigint"
	TypeAccessor  Type = "accessor"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t Type) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t Type) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *Type) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch Type(in.String()) {
	case TypeObject:
		*t = TypeObject
	case TypeFunction:
		*t = TypeFunction
	case TypeUndefined:
		*t = TypeUndefined
	case TypeString:
		*t = TypeString
	case TypeNumber:
		*t = TypeNumber
	case TypeBoolean:
		*t = TypeBoolean
	case TypeSymbol:
		*t = TypeSymbol
	case TypeBigint:
		*t = TypeBigint
	case TypeAccessor:
		*t = TypeAccessor

	default:
		in.AddError(errors.New("unknown Type value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *Type) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Subtype object subtype hint. Specified for object type values only. NOTE:
// If you change anything here, make sure to also update subtype in
// ObjectPreview and PropertyPreview below.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#type-RemoteObject
type Subtype string

// String returns the Subtype as string value.
func (t Subtype) String() string {
	return string(t)
}

// Subtype values.
const (
	SubtypeArray             Subtype = "array"
	SubtypeNull              Subtype = "null"
	SubtypeNode              Subtype = "node"
	SubtypeRegexp            Subtype = "regexp"
	SubtypeDate              Subtype = "date"
	SubtypeMap               Subtype = "map"
	SubtypeSet               Subtype = "set"
	SubtypeWeakmap           Subtype = "weakmap"
	SubtypeWeakset           Subtype = "weakset"
	SubtypeIterator          Subtype = "iterator"
	SubtypeGenerator         Subtype = "generator"
	SubtypeError             Subtype = "error"
	SubtypeProxy             Subtype = "proxy"
	SubtypePromise           Subtype = "promise"
	SubtypeTypedarray        Subtype = "typedarray"
	SubtypeArraybuffer       Subtype = "arraybuffer"
	SubtypeDataview          Subtype = "dataview"
	SubtypeWebassemblymemory Subtype = "webassemblymemory"
	SubtypeWasmvalue         Subtype = "wasmvalue"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t Subtype) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t Subtype) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *Subtype) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch Subtype(in.String()) {
	case SubtypeArray:
		*t = SubtypeArray
	case SubtypeNull:
		*t = SubtypeNull
	case SubtypeNode:
		*t = SubtypeNode
	case SubtypeRegexp:
		*t = SubtypeRegexp
	case SubtypeDate:
		*t = SubtypeDate
	case SubtypeMap:
		*t = SubtypeMap
	case SubtypeSet:
		*t = SubtypeSet
	case SubtypeWeakmap:
		*t = SubtypeWeakmap
	case SubtypeWeakset:
		*t = SubtypeWeakset
	case SubtypeIterator:
		*t = SubtypeIterator
	case SubtypeGenerator:
		*t = SubtypeGenerator
	case SubtypeError:
		*t = SubtypeError
	case SubtypeProxy:
		*t = SubtypeProxy
	case SubtypePromise:
		*t = SubtypePromise
	case SubtypeTypedarray:
		*t = SubtypeTypedarray
	case SubtypeArraybuffer:
		*t = SubtypeArraybuffer
	case SubtypeDataview:
		*t = SubtypeDataview
	case SubtypeWebassemblymemory:
		*t = SubtypeWebassemblymemory
	case SubtypeWasmvalue:
		*t = SubtypeWasmvalue

	default:
		in.AddError(errors.New("unknown Subtype value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *Subtype) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// APIType type of the call.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Runtime#event-consoleAPICalled
type APIType string

// String returns the APIType as string value.
func (t APIType) String() string {
	return string(t)
}

// APIType values.
const (
	APITypeLog                 APIType = "log"
	APITypeDebug               APIType = "debug"
	APITypeInfo                APIType = "info"
	APITypeError               APIType = "error"
	APITypeWarning             APIType = "warning"
	APITypeDir                 APIType = "dir"
	APITypeDirxml              APIType = "dirxml"
	APITypeTable               APIType = "table"
	APITypeTrace               APIType = "trace"
	APITypeClear               APIType = "clear"
	APITypeStartGroup          APIType = "startGroup"
	APITypeStartGroupCollapsed APIType = "startGroupCollapsed"
	APITypeEndGroup            APIType = "endGroup"
	APITypeAssert              APIType = "assert"
	APITypeProfile             APIType = "profile"
	APITypeProfileEnd          APIType = "profileEnd"
	APITypeCount               APIType = "count"
	APITypeTimeEnd             APIType = "timeEnd"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t APIType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t APIType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *APIType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch APIType(in.String()) {
	case APITypeLog:
		*t = APITypeLog
	case APITypeDebug:
		*t = APITypeDebug
	case APITypeInfo:
		*t = APITypeInfo
	case APITypeError:
		*t = APITypeError
	case APITypeWarning:
		*t = APITypeWarning
	case APITypeDir:
		*t = APITypeDir
	case APITypeDirxml:
		*t = APITypeDirxml
	case APITypeTable:
		*t = APITypeTable
	case APITypeTrace:
		*t = APITypeTrace
	case APITypeClear:
		*t = APITypeClear
	case APITypeStartGroup:
		*t = APITypeStartGroup
	case APITypeStartGroupCollapsed:
		*t = APITypeStartGroupCollapsed
	case APITypeEndGroup:
		*t = APITypeEndGroup
	case APITypeAssert:
		*t = APITypeAssert
	case APITypeProfile:
		*t = APITypeProfile
	case APITypeProfileEnd:
		*t = APITypeProfileEnd
	case APITypeCount:
		*t = APITypeCount
	case APITypeTimeEnd:
		*t = APITypeTimeEnd

	default:
		in.AddError(errors.New("unknown APIType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *APIType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
