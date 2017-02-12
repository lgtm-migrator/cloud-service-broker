// Code generated by protoc-gen-go.
// source: google/devtools/clouddebugger/v2/data.proto
// DO NOT EDIT!

package clouddebugger

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_devtools_source_v1 "google.golang.org/genproto/googleapis/devtools/source/v1"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf2 "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Enumerates references to which the message applies.
type StatusMessage_Reference int32

const (
	// Status doesn't refer to any particular input.
	StatusMessage_UNSPECIFIED StatusMessage_Reference = 0
	// Status applies to the breakpoint and is related to its location.
	StatusMessage_BREAKPOINT_SOURCE_LOCATION StatusMessage_Reference = 3
	// Status applies to the breakpoint and is related to its condition.
	StatusMessage_BREAKPOINT_CONDITION StatusMessage_Reference = 4
	// Status applies to the breakpoint and is related to its expressions.
	StatusMessage_BREAKPOINT_EXPRESSION StatusMessage_Reference = 7
	// Status applies to the entire variable.
	StatusMessage_VARIABLE_NAME StatusMessage_Reference = 5
	// Status applies to variable value (variable name is valid).
	StatusMessage_VARIABLE_VALUE StatusMessage_Reference = 6
)

var StatusMessage_Reference_name = map[int32]string{
	0: "UNSPECIFIED",
	3: "BREAKPOINT_SOURCE_LOCATION",
	4: "BREAKPOINT_CONDITION",
	7: "BREAKPOINT_EXPRESSION",
	5: "VARIABLE_NAME",
	6: "VARIABLE_VALUE",
}
var StatusMessage_Reference_value = map[string]int32{
	"UNSPECIFIED":                0,
	"BREAKPOINT_SOURCE_LOCATION": 3,
	"BREAKPOINT_CONDITION":       4,
	"BREAKPOINT_EXPRESSION":      7,
	"VARIABLE_NAME":              5,
	"VARIABLE_VALUE":             6,
}

func (x StatusMessage_Reference) String() string {
	return proto.EnumName(StatusMessage_Reference_name, int32(x))
}
func (StatusMessage_Reference) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

// Actions that can be taken when a breakpoint hits.
// Agents should reject breakpoints with unsupported or unknown action values.
type Breakpoint_Action int32

const (
	// Capture stack frame and variables and update the breakpoint.
	// The data is only captured once. After that the breakpoint is set
	// in a final state.
	Breakpoint_CAPTURE Breakpoint_Action = 0
	// Log each breakpoint hit. The breakpoint remains active until
	// deleted or expired.
	Breakpoint_LOG Breakpoint_Action = 1
)

var Breakpoint_Action_name = map[int32]string{
	0: "CAPTURE",
	1: "LOG",
}
var Breakpoint_Action_value = map[string]int32{
	"CAPTURE": 0,
	"LOG":     1,
}

func (x Breakpoint_Action) String() string {
	return proto.EnumName(Breakpoint_Action_name, int32(x))
}
func (Breakpoint_Action) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{5, 0} }

// Log severity levels.
type Breakpoint_LogLevel int32

const (
	// Information log message.
	Breakpoint_INFO Breakpoint_LogLevel = 0
	// Warning log message.
	Breakpoint_WARNING Breakpoint_LogLevel = 1
	// Error log message.
	Breakpoint_ERROR Breakpoint_LogLevel = 2
)

var Breakpoint_LogLevel_name = map[int32]string{
	0: "INFO",
	1: "WARNING",
	2: "ERROR",
}
var Breakpoint_LogLevel_value = map[string]int32{
	"INFO":    0,
	"WARNING": 1,
	"ERROR":   2,
}

func (x Breakpoint_LogLevel) String() string {
	return proto.EnumName(Breakpoint_LogLevel_name, int32(x))
}
func (Breakpoint_LogLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{5, 1} }

// Represents a message with parameters.
type FormatMessage struct {
	// Format template for the message. The `format` uses placeholders `$0`,
	// `$1`, etc. to reference parameters. `$$` can be used to denote the `$`
	// character.
	//
	// Examples:
	//
	// *   `Failed to load '$0' which helps debug $1 the first time it
	//     is loaded.  Again, $0 is very important.`
	// *   `Please pay $$10 to use $0 instead of $1.`
	Format string `protobuf:"bytes,1,opt,name=format" json:"format,omitempty"`
	// Optional parameters to be embedded into the message.
	Parameters []string `protobuf:"bytes,2,rep,name=parameters" json:"parameters,omitempty"`
}

func (m *FormatMessage) Reset()                    { *m = FormatMessage{} }
func (m *FormatMessage) String() string            { return proto.CompactTextString(m) }
func (*FormatMessage) ProtoMessage()               {}
func (*FormatMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *FormatMessage) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *FormatMessage) GetParameters() []string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

// Represents a contextual status message.
// The message can indicate an error or informational status, and refer to
// specific parts of the containing object.
// For example, the `Breakpoint.status` field can indicate an error referring
// to the `BREAKPOINT_SOURCE_LOCATION` with the message `Location not found`.
type StatusMessage struct {
	// Distinguishes errors from informational messages.
	IsError bool `protobuf:"varint,1,opt,name=is_error,json=isError" json:"is_error,omitempty"`
	// Reference to which the message applies.
	RefersTo StatusMessage_Reference `protobuf:"varint,2,opt,name=refers_to,json=refersTo,enum=google.devtools.clouddebugger.v2.StatusMessage_Reference" json:"refers_to,omitempty"`
	// Status message text.
	Description *FormatMessage `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *StatusMessage) Reset()                    { *m = StatusMessage{} }
func (m *StatusMessage) String() string            { return proto.CompactTextString(m) }
func (*StatusMessage) ProtoMessage()               {}
func (*StatusMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *StatusMessage) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

func (m *StatusMessage) GetRefersTo() StatusMessage_Reference {
	if m != nil {
		return m.RefersTo
	}
	return StatusMessage_UNSPECIFIED
}

func (m *StatusMessage) GetDescription() *FormatMessage {
	if m != nil {
		return m.Description
	}
	return nil
}

// Represents a location in the source code.
type SourceLocation struct {
	// Path to the source file within the source context of the target binary.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// Line inside the file. The first line in the file has the value `1`.
	Line int32 `protobuf:"varint,2,opt,name=line" json:"line,omitempty"`
}

func (m *SourceLocation) Reset()                    { *m = SourceLocation{} }
func (m *SourceLocation) String() string            { return proto.CompactTextString(m) }
func (*SourceLocation) ProtoMessage()               {}
func (*SourceLocation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *SourceLocation) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *SourceLocation) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

// Represents a variable or an argument possibly of a compound object type.
// Note how the following variables are represented:
//
// 1) A simple variable:
//
//     int x = 5
//
//     { name: "x", value: "5", type: "int" }  // Captured variable
//
// 2) A compound object:
//
//     struct T {
//         int m1;
//         int m2;
//     };
//     T x = { 3, 7 };
//
//     {  // Captured variable
//         name: "x",
//         type: "T",
//         members { name: "m1", value: "3", type: "int" },
//         members { name: "m2", value: "7", type: "int" }
//     }
//
// 3) A pointer where the pointee was captured:
//
//     T x = { 3, 7 };
//     T* p = &x;
//
//     {   // Captured variable
//         name: "p",
//         type: "T*",
//         value: "0x00500500",
//         members { name: "m1", value: "3", type: "int" },
//         members { name: "m2", value: "7", type: "int" }
//     }
//
// 4) A pointer where the pointee was not captured:
//
//     T* p = new T;
//
//     {   // Captured variable
//         name: "p",
//         type: "T*",
//         value: "0x00400400"
//         status { is_error: true, description { format: "unavailable" } }
//     }
//
// The status should describe the reason for the missing value,
// such as `<optimized out>`, `<inaccessible>`, `<pointers limit reached>`.
//
// Note that a null pointer should not have members.
//
// 5) An unnamed value:
//
//     int* p = new int(7);
//
//     {   // Captured variable
//         name: "p",
//         value: "0x00500500",
//         type: "int*",
//         members { value: "7", type: "int" } }
//
// 6) An unnamed pointer where the pointee was not captured:
//
//     int* p = new int(7);
//     int** pp = &p;
//
//     {  // Captured variable
//         name: "pp",
//         value: "0x00500500",
//         type: "int**",
//         members {
//             value: "0x00400400",
//             type: "int*"
//             status {
//                 is_error: true,
//                 description: { format: "unavailable" } }
//             }
//         }
//     }
//
// To optimize computation, memory and network traffic, variables that
// repeat in the output multiple times can be stored once in a shared
// variable table and be referenced using the `var_table_index` field.  The
// variables stored in the shared table are nameless and are essentially
// a partition of the complete variable. To reconstruct the complete
// variable, merge the referencing variable with the referenced variable.
//
// When using the shared variable table, the following variables:
//
//     T x = { 3, 7 };
//     T* p = &x;
//     T& r = x;
//
//     { name: "x", var_table_index: 3, type: "T" }  // Captured variables
//     { name: "p", value "0x00500500", type="T*", var_table_index: 3 }
//     { name: "r", type="T&", var_table_index: 3 }
//
//     {  // Shared variable table entry #3:
//         members { name: "m1", value: "3", type: "int" },
//         members { name: "m2", value: "7", type: "int" }
//     }
//
// Note that the pointer address is stored with the referencing variable
// and not with the referenced variable. This allows the referenced variable
// to be shared between pointers and references.
//
// The type field is optional. The debugger agent may or may not support it.
type Variable struct {
	// Name of the variable, if any.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Simple value of the variable.
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	// Variable type (e.g. `MyClass`). If the variable is split with
	// `var_table_index`, `type` goes next to `value`. The interpretation of
	// a type is agent specific. It is recommended to include the dynamic type
	// rather than a static type of an object.
	Type string `protobuf:"bytes,6,opt,name=type" json:"type,omitempty"`
	// Members contained or pointed to by the variable.
	Members []*Variable `protobuf:"bytes,3,rep,name=members" json:"members,omitempty"`
	// Reference to a variable in the shared variable table. More than
	// one variable can reference the same variable in the table. The
	// `var_table_index` field is an index into `variable_table` in Breakpoint.
	VarTableIndex *google_protobuf2.Int32Value `protobuf:"bytes,4,opt,name=var_table_index,json=varTableIndex" json:"var_table_index,omitempty"`
	// Status associated with the variable. This field will usually stay
	// unset. A status of a single variable only applies to that variable or
	// expression. The rest of breakpoint data still remains valid. Variables
	// might be reported in error state even when breakpoint is not in final
	// state.
	//
	// The message may refer to variable name with `refers_to` set to
	// `VARIABLE_NAME`. Alternatively `refers_to` will be set to `VARIABLE_VALUE`.
	// In either case variable value and members will be unset.
	//
	// Example of error message applied to name: `Invalid expression syntax`.
	//
	// Example of information message applied to value: `Not captured`.
	//
	// Examples of error message applied to value:
	//
	// *   `Malformed string`,
	// *   `Field f not found in class C`
	// *   `Null pointer dereference`
	Status *StatusMessage `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
}

func (m *Variable) Reset()                    { *m = Variable{} }
func (m *Variable) String() string            { return proto.CompactTextString(m) }
func (*Variable) ProtoMessage()               {}
func (*Variable) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *Variable) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Variable) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Variable) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Variable) GetMembers() []*Variable {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Variable) GetVarTableIndex() *google_protobuf2.Int32Value {
	if m != nil {
		return m.VarTableIndex
	}
	return nil
}

func (m *Variable) GetStatus() *StatusMessage {
	if m != nil {
		return m.Status
	}
	return nil
}

// Represents a stack frame context.
type StackFrame struct {
	// Demangled function name at the call site.
	Function string `protobuf:"bytes,1,opt,name=function" json:"function,omitempty"`
	// Source location of the call site.
	Location *SourceLocation `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
	// Set of arguments passed to this function.
	// Note that this might not be populated for all stack frames.
	Arguments []*Variable `protobuf:"bytes,3,rep,name=arguments" json:"arguments,omitempty"`
	// Set of local variables at the stack frame location.
	// Note that this might not be populated for all stack frames.
	Locals []*Variable `protobuf:"bytes,4,rep,name=locals" json:"locals,omitempty"`
}

func (m *StackFrame) Reset()                    { *m = StackFrame{} }
func (m *StackFrame) String() string            { return proto.CompactTextString(m) }
func (*StackFrame) ProtoMessage()               {}
func (*StackFrame) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *StackFrame) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *StackFrame) GetLocation() *SourceLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *StackFrame) GetArguments() []*Variable {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *StackFrame) GetLocals() []*Variable {
	if m != nil {
		return m.Locals
	}
	return nil
}

// Represents the breakpoint specification, status and results.
type Breakpoint struct {
	// Breakpoint identifier, unique in the scope of the debuggee.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Action that the agent should perform when the code at the
	// breakpoint location is hit.
	Action Breakpoint_Action `protobuf:"varint,13,opt,name=action,enum=google.devtools.clouddebugger.v2.Breakpoint_Action" json:"action,omitempty"`
	// Breakpoint source location.
	Location *SourceLocation `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
	// Condition that triggers the breakpoint.
	// The condition is a compound boolean expression composed using expressions
	// in a programming language at the source location.
	Condition string `protobuf:"bytes,3,opt,name=condition" json:"condition,omitempty"`
	// List of read-only expressions to evaluate at the breakpoint location.
	// The expressions are composed using expressions in the programming language
	// at the source location. If the breakpoint action is `LOG`, the evaluated
	// expressions are included in log statements.
	Expressions []string `protobuf:"bytes,4,rep,name=expressions" json:"expressions,omitempty"`
	// Only relevant when action is `LOG`. Defines the message to log when
	// the breakpoint hits. The message may include parameter placeholders `$0`,
	// `$1`, etc. These placeholders are replaced with the evaluated value
	// of the appropriate expression. Expressions not referenced in
	// `log_message_format` are not logged.
	//
	// Example: `Message received, id = $0, count = $1` with
	// `expressions` = `[ message.id, message.count ]`.
	LogMessageFormat string `protobuf:"bytes,14,opt,name=log_message_format,json=logMessageFormat" json:"log_message_format,omitempty"`
	// Indicates the severity of the log. Only relevant when action is `LOG`.
	LogLevel Breakpoint_LogLevel `protobuf:"varint,15,opt,name=log_level,json=logLevel,enum=google.devtools.clouddebugger.v2.Breakpoint_LogLevel" json:"log_level,omitempty"`
	// When true, indicates that this is a final result and the
	// breakpoint state will not change from here on.
	IsFinalState bool `protobuf:"varint,5,opt,name=is_final_state,json=isFinalState" json:"is_final_state,omitempty"`
	// Time this breakpoint was created by the server in seconds resolution.
	CreateTime *google_protobuf1.Timestamp `protobuf:"bytes,11,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	// Time this breakpoint was finalized as seen by the server in seconds
	// resolution.
	FinalTime *google_protobuf1.Timestamp `protobuf:"bytes,12,opt,name=final_time,json=finalTime" json:"final_time,omitempty"`
	// E-mail address of the user that created this breakpoint
	UserEmail string `protobuf:"bytes,16,opt,name=user_email,json=userEmail" json:"user_email,omitempty"`
	// Breakpoint status.
	//
	// The status includes an error flag and a human readable message.
	// This field is usually unset. The message can be either
	// informational or an error message. Regardless, clients should always
	// display the text message back to the user.
	//
	// Error status indicates complete failure of the breakpoint.
	//
	// Example (non-final state): `Still loading symbols...`
	//
	// Examples (final state):
	//
	// *   `Invalid line number` referring to location
	// *   `Field f not found in class C` referring to condition
	Status *StatusMessage `protobuf:"bytes,10,opt,name=status" json:"status,omitempty"`
	// The stack at breakpoint time.
	StackFrames []*StackFrame `protobuf:"bytes,7,rep,name=stack_frames,json=stackFrames" json:"stack_frames,omitempty"`
	// Values of evaluated expressions at breakpoint time.
	// The evaluated expressions appear in exactly the same order they
	// are listed in the `expressions` field.
	// The `name` field holds the original expression text, the `value` or
	// `members` field holds the result of the evaluated expression.
	// If the expression cannot be evaluated, the `status` inside the `Variable`
	// will indicate an error and contain the error text.
	EvaluatedExpressions []*Variable `protobuf:"bytes,8,rep,name=evaluated_expressions,json=evaluatedExpressions" json:"evaluated_expressions,omitempty"`
	// The `variable_table` exists to aid with computation, memory and network
	// traffic optimization.  It enables storing a variable once and reference
	// it from multiple variables, including variables stored in the
	// `variable_table` itself.
	// For example, the same `this` object, which may appear at many levels of
	// the stack, can have all of its data stored once in this table.  The
	// stack frame variables then would hold only a reference to it.
	//
	// The variable `var_table_index` field is an index into this repeated field.
	// The stored objects are nameless and get their name from the referencing
	// variable. The effective variable is a merge of the referencing variable
	// and the referenced variable.
	VariableTable []*Variable `protobuf:"bytes,9,rep,name=variable_table,json=variableTable" json:"variable_table,omitempty"`
	// A set of custom breakpoint properties, populated by the agent, to be
	// displayed to the user.
	Labels map[string]string `protobuf:"bytes,17,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Breakpoint) Reset()                    { *m = Breakpoint{} }
func (m *Breakpoint) String() string            { return proto.CompactTextString(m) }
func (*Breakpoint) ProtoMessage()               {}
func (*Breakpoint) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *Breakpoint) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Breakpoint) GetAction() Breakpoint_Action {
	if m != nil {
		return m.Action
	}
	return Breakpoint_CAPTURE
}

func (m *Breakpoint) GetLocation() *SourceLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Breakpoint) GetCondition() string {
	if m != nil {
		return m.Condition
	}
	return ""
}

func (m *Breakpoint) GetExpressions() []string {
	if m != nil {
		return m.Expressions
	}
	return nil
}

func (m *Breakpoint) GetLogMessageFormat() string {
	if m != nil {
		return m.LogMessageFormat
	}
	return ""
}

func (m *Breakpoint) GetLogLevel() Breakpoint_LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return Breakpoint_INFO
}

func (m *Breakpoint) GetIsFinalState() bool {
	if m != nil {
		return m.IsFinalState
	}
	return false
}

func (m *Breakpoint) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Breakpoint) GetFinalTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.FinalTime
	}
	return nil
}

func (m *Breakpoint) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func (m *Breakpoint) GetStatus() *StatusMessage {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Breakpoint) GetStackFrames() []*StackFrame {
	if m != nil {
		return m.StackFrames
	}
	return nil
}

func (m *Breakpoint) GetEvaluatedExpressions() []*Variable {
	if m != nil {
		return m.EvaluatedExpressions
	}
	return nil
}

func (m *Breakpoint) GetVariableTable() []*Variable {
	if m != nil {
		return m.VariableTable
	}
	return nil
}

func (m *Breakpoint) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// Represents the application to debug. The application may include one or more
// replicated processes executing the same code. Each of these processes is
// attached with a debugger agent, carrying out the debugging commands.
// The agents attached to the same debuggee are identified by using exactly the
// same field values when registering.
type Debuggee struct {
	// Unique identifier for the debuggee generated by the controller service.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Project the debuggee is associated with.
	// Use the project number when registering a Google Cloud Platform project.
	Project string `protobuf:"bytes,2,opt,name=project" json:"project,omitempty"`
	// Debuggee uniquifier within the project.
	// Any string that identifies the application within the project can be used.
	// Including environment and version or build IDs is recommended.
	Uniquifier string `protobuf:"bytes,3,opt,name=uniquifier" json:"uniquifier,omitempty"`
	// Human readable description of the debuggee.
	// Including a human-readable project name, environment name and version
	// information is recommended.
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// If set to `true`, indicates that the debuggee is considered as inactive by
	// the Controller service.
	IsInactive bool `protobuf:"varint,5,opt,name=is_inactive,json=isInactive" json:"is_inactive,omitempty"`
	// Version ID of the agent release. The version ID is structured as
	// following: `domain/type/vmajor.minor` (for example
	// `google.com/gcp-java/v1.1`).
	AgentVersion string `protobuf:"bytes,6,opt,name=agent_version,json=agentVersion" json:"agent_version,omitempty"`
	// If set to `true`, indicates that the agent should disable itself and
	// detach from the debuggee.
	IsDisabled bool `protobuf:"varint,7,opt,name=is_disabled,json=isDisabled" json:"is_disabled,omitempty"`
	// Human readable message to be displayed to the user about this debuggee.
	// Absence of this field indicates no status. The message can be either
	// informational or an error status.
	Status *StatusMessage `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
	// References to the locations and revisions of the source code used in the
	// deployed application.
	//
	// NOTE: This field is deprecated. Consumers should use
	// `ext_source_contexts` if it is not empty. Debug agents should populate
	// both this field and `ext_source_contexts`.
	SourceContexts []*google_devtools_source_v1.SourceContext `protobuf:"bytes,9,rep,name=source_contexts,json=sourceContexts" json:"source_contexts,omitempty"`
	// References to the locations and revisions of the source code used in the
	// deployed application.
	//
	// Contexts describing a remote repo related to the source code
	// have a `category` label of `remote_repo`. Source snapshot source
	// contexts have a `category` of `snapshot`.
	ExtSourceContexts []*google_devtools_source_v1.ExtendedSourceContext `protobuf:"bytes,13,rep,name=ext_source_contexts,json=extSourceContexts" json:"ext_source_contexts,omitempty"`
	// A set of custom debuggee properties, populated by the agent, to be
	// displayed to the user.
	Labels map[string]string `protobuf:"bytes,11,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Debuggee) Reset()                    { *m = Debuggee{} }
func (m *Debuggee) String() string            { return proto.CompactTextString(m) }
func (*Debuggee) ProtoMessage()               {}
func (*Debuggee) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *Debuggee) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Debuggee) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Debuggee) GetUniquifier() string {
	if m != nil {
		return m.Uniquifier
	}
	return ""
}

func (m *Debuggee) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Debuggee) GetIsInactive() bool {
	if m != nil {
		return m.IsInactive
	}
	return false
}

func (m *Debuggee) GetAgentVersion() string {
	if m != nil {
		return m.AgentVersion
	}
	return ""
}

func (m *Debuggee) GetIsDisabled() bool {
	if m != nil {
		return m.IsDisabled
	}
	return false
}

func (m *Debuggee) GetStatus() *StatusMessage {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Debuggee) GetSourceContexts() []*google_devtools_source_v1.SourceContext {
	if m != nil {
		return m.SourceContexts
	}
	return nil
}

func (m *Debuggee) GetExtSourceContexts() []*google_devtools_source_v1.ExtendedSourceContext {
	if m != nil {
		return m.ExtSourceContexts
	}
	return nil
}

func (m *Debuggee) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*FormatMessage)(nil), "google.devtools.clouddebugger.v2.FormatMessage")
	proto.RegisterType((*StatusMessage)(nil), "google.devtools.clouddebugger.v2.StatusMessage")
	proto.RegisterType((*SourceLocation)(nil), "google.devtools.clouddebugger.v2.SourceLocation")
	proto.RegisterType((*Variable)(nil), "google.devtools.clouddebugger.v2.Variable")
	proto.RegisterType((*StackFrame)(nil), "google.devtools.clouddebugger.v2.StackFrame")
	proto.RegisterType((*Breakpoint)(nil), "google.devtools.clouddebugger.v2.Breakpoint")
	proto.RegisterType((*Debuggee)(nil), "google.devtools.clouddebugger.v2.Debuggee")
	proto.RegisterEnum("google.devtools.clouddebugger.v2.StatusMessage_Reference", StatusMessage_Reference_name, StatusMessage_Reference_value)
	proto.RegisterEnum("google.devtools.clouddebugger.v2.Breakpoint_Action", Breakpoint_Action_name, Breakpoint_Action_value)
	proto.RegisterEnum("google.devtools.clouddebugger.v2.Breakpoint_LogLevel", Breakpoint_LogLevel_name, Breakpoint_LogLevel_value)
}

func init() { proto.RegisterFile("google/devtools/clouddebugger/v2/data.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x57, 0x4d, 0x73, 0xdb, 0xb6,
	0x16, 0x8d, 0x3e, 0x2c, 0x91, 0x97, 0x96, 0xac, 0xe0, 0x25, 0x6f, 0x18, 0xbf, 0x3c, 0x47, 0xa3,
	0x97, 0x85, 0xe7, 0x35, 0x23, 0x25, 0xca, 0xb4, 0x93, 0x34, 0x2b, 0x59, 0xa6, 0x5d, 0x4d, 0x14,
	0x49, 0x81, 0x6c, 0xb7, 0xd3, 0x0d, 0x0b, 0x8b, 0x10, 0xcb, 0x86, 0x22, 0x59, 0x00, 0x52, 0x9d,
	0x6d, 0xa7, 0x3f, 0xa2, 0x8b, 0xfe, 0xb7, 0xee, 0xba, 0xee, 0x4f, 0xe8, 0x00, 0x04, 0x15, 0x2a,
	0x69, 0xab, 0xb8, 0xc9, 0xee, 0xe2, 0xe0, 0xdc, 0x03, 0xea, 0xe2, 0xdc, 0x4b, 0x0a, 0x3e, 0xf1,
	0xe3, 0xd8, 0x0f, 0x69, 0xc7, 0xa3, 0x2b, 0x11, 0xc7, 0x21, 0xef, 0xcc, 0xc2, 0x78, 0xe9, 0x79,
	0xf4, 0x72, 0xe9, 0xfb, 0x94, 0x75, 0x56, 0xdd, 0x8e, 0x47, 0x04, 0x69, 0x27, 0x2c, 0x16, 0x31,
	0x6a, 0xa6, 0xe4, 0x76, 0x46, 0x6e, 0x6f, 0x90, 0xdb, 0xab, 0xee, 0xfe, 0x5d, 0x2d, 0x47, 0x92,
	0xa0, 0x43, 0xa2, 0x28, 0x16, 0x44, 0x04, 0x71, 0xc4, 0xd3, 0xfc, 0xfd, 0xf6, 0xdb, 0x87, 0xf1,
	0x78, 0xc9, 0x66, 0xb4, 0xb3, 0x7a, 0xa4, 0x23, 0x77, 0x16, 0x47, 0x82, 0x5e, 0x09, 0xcd, 0xbf,
	0xa7, 0xf9, 0x6a, 0x75, 0xb9, 0x9c, 0x77, 0x44, 0xb0, 0xa0, 0x5c, 0x90, 0x45, 0xa2, 0x09, 0x07,
	0x6f, 0x13, 0x7e, 0x60, 0x24, 0x49, 0x28, 0xd3, 0x07, 0xb6, 0x4e, 0xa1, 0x76, 0x12, 0xb3, 0x05,
	0x11, 0x2f, 0x28, 0xe7, 0xc4, 0xa7, 0xe8, 0xdf, 0x50, 0x99, 0x2b, 0xc0, 0x2e, 0x34, 0x0b, 0x87,
	0x26, 0xd6, 0x2b, 0x74, 0x00, 0x90, 0x10, 0x46, 0x16, 0x54, 0x50, 0xc6, 0xed, 0x62, 0xb3, 0x74,
	0x68, 0xe2, 0x1c, 0xd2, 0xfa, 0xbd, 0x08, 0xb5, 0xa9, 0x20, 0x62, 0xc9, 0x33, 0xa5, 0x3b, 0x60,
	0x04, 0xdc, 0xa5, 0x8c, 0xc5, 0x4c, 0x69, 0x19, 0xb8, 0x1a, 0x70, 0x47, 0x2e, 0xd1, 0x05, 0x98,
	0x8c, 0xce, 0x29, 0xe3, 0xae, 0x88, 0xed, 0x62, 0xb3, 0x70, 0x58, 0xef, 0x3e, 0x6d, 0x6f, 0x2b,
	0x5d, 0x7b, 0x43, 0xbe, 0x8d, 0xa5, 0x00, 0x8d, 0x66, 0x14, 0x1b, 0xa9, 0xd6, 0x59, 0x8c, 0x5e,
	0x82, 0xe5, 0x51, 0x3e, 0x63, 0x41, 0x22, 0x8b, 0x6a, 0x97, 0x9a, 0x85, 0x43, 0xab, 0xdb, 0xd9,
	0xae, 0xbc, 0x51, 0x02, 0x9c, 0xd7, 0x68, 0xfd, 0x5c, 0x00, 0x73, 0x7d, 0x14, 0xda, 0x03, 0xeb,
	0x7c, 0x34, 0x9d, 0x38, 0xfd, 0xc1, 0xc9, 0xc0, 0x39, 0x6e, 0xdc, 0x40, 0x07, 0xb0, 0x7f, 0x84,
	0x9d, 0xde, 0xf3, 0xc9, 0x78, 0x30, 0x3a, 0x73, 0xa7, 0xe3, 0x73, 0xdc, 0x77, 0xdc, 0xe1, 0xb8,
	0xdf, 0x3b, 0x1b, 0x8c, 0x47, 0x8d, 0x12, 0xb2, 0xe1, 0x56, 0x6e, 0xbf, 0x3f, 0x1e, 0x1d, 0x0f,
	0xd4, 0x4e, 0x19, 0xdd, 0x81, 0xdb, 0xb9, 0x1d, 0xe7, 0xab, 0x09, 0x76, 0xa6, 0x53, 0xb9, 0x55,
	0x45, 0x37, 0xa1, 0x76, 0xd1, 0xc3, 0x83, 0xde, 0xd1, 0xd0, 0x71, 0x47, 0xbd, 0x17, 0x4e, 0x63,
	0x07, 0x21, 0xa8, 0xaf, 0xa1, 0x8b, 0xde, 0xf0, 0xdc, 0x69, 0x54, 0x5a, 0x4f, 0xa0, 0x3e, 0x55,
	0xa6, 0x18, 0xc6, 0x33, 0xe5, 0x22, 0x84, 0xa0, 0x9c, 0x10, 0xf1, 0xad, 0xbe, 0x3a, 0x15, 0x4b,
	0x2c, 0x0c, 0x22, 0xaa, 0xca, 0xbc, 0x83, 0x55, 0xdc, 0xfa, 0xa5, 0x08, 0xc6, 0x05, 0x61, 0x01,
	0xb9, 0x0c, 0xa9, 0x24, 0x44, 0x64, 0x41, 0xb3, 0x24, 0x19, 0xa3, 0x5b, 0xb0, 0xb3, 0x22, 0xe1,
	0x32, 0xcd, 0x32, 0x71, 0xba, 0x90, 0x4c, 0xf1, 0x3a, 0xa1, 0x76, 0x25, 0x65, 0xca, 0x18, 0x1d,
	0x43, 0x75, 0x41, 0x17, 0x97, 0xd2, 0x14, 0xa5, 0x66, 0xe9, 0xd0, 0xea, 0xfe, 0x7f, 0x7b, 0xb9,
	0xb3, 0xa3, 0x71, 0x96, 0x8a, 0xfa, 0xb0, 0xb7, 0x22, 0xcc, 0x15, 0x12, 0x75, 0x83, 0xc8, 0xa3,
	0x57, 0x76, 0x59, 0x5d, 0xde, 0x7f, 0x32, 0xb5, 0xcc, 0xc0, 0xed, 0x41, 0x24, 0x1e, 0x77, 0x2f,
	0xe4, 0xf3, 0xe0, 0xda, 0x8a, 0xb0, 0x33, 0x99, 0x32, 0x90, 0x19, 0xe8, 0x14, 0x2a, 0x5c, 0x59,
	0xc4, 0xde, 0x79, 0xdf, 0x8b, 0xdf, 0xb0, 0x14, 0xd6, 0xe9, 0xad, 0x9f, 0x8a, 0x00, 0x53, 0x41,
	0x66, 0xaf, 0x4e, 0xa4, 0xbd, 0xd1, 0x3e, 0x18, 0xf3, 0x65, 0x34, 0x53, 0x96, 0x4a, 0x8b, 0xb4,
	0x5e, 0xa3, 0x21, 0x18, 0xa1, 0xae, 0xbe, 0xaa, 0x95, 0xd5, 0x7d, 0xf8, 0x1e, 0xa7, 0x6e, 0xdc,
	0x1a, 0x5e, 0x2b, 0xa0, 0x2f, 0xc0, 0x24, 0xcc, 0x5f, 0x2e, 0x68, 0x24, 0xfe, 0x49, 0x39, 0xdf,
	0x24, 0xa3, 0x23, 0xa8, 0x48, 0xd5, 0x90, 0xdb, 0xe5, 0x6b, 0xcb, 0xe8, 0xcc, 0xd6, 0xaf, 0x06,
	0xc0, 0x11, 0xa3, 0xe4, 0x55, 0x12, 0x07, 0x91, 0x40, 0x75, 0x28, 0x06, 0x9e, 0x2e, 0x40, 0x31,
	0xf0, 0xd0, 0x73, 0xa8, 0x90, 0xb4, 0x28, 0x35, 0xd5, 0xc1, 0x8f, 0xb7, 0x1f, 0xf1, 0x46, 0xad,
	0xdd, 0x53, 0xa9, 0x58, 0x4b, 0x7c, 0xe4, 0x3a, 0xde, 0x05, 0x73, 0x16, 0x47, 0x5e, 0xb0, 0x9e,
	0x02, 0x26, 0x7e, 0x03, 0xa0, 0x26, 0x58, 0xf4, 0x2a, 0x61, 0x94, 0x73, 0x39, 0x79, 0x55, 0x81,
	0x4c, 0x9c, 0x87, 0xd0, 0x03, 0x40, 0x61, 0xec, 0xbb, 0x8b, 0xd4, 0x17, 0xae, 0x1e, 0x88, 0x75,
	0x25, 0xd4, 0x08, 0x63, 0x5f, 0x1b, 0x26, 0x1d, 0x1b, 0x08, 0x83, 0x29, 0xd9, 0x21, 0x5d, 0xd1,
	0xd0, 0xde, 0x53, 0xb5, 0xf8, 0xf4, 0x5a, 0xb5, 0x18, 0xc6, 0xfe, 0x50, 0x26, 0xcb, 0x5f, 0x90,
	0x46, 0xe8, 0x3e, 0xd4, 0x03, 0xee, 0xce, 0x83, 0x88, 0x84, 0xae, 0x74, 0x25, 0x55, 0x9e, 0x36,
	0xf0, 0x6e, 0xc0, 0x4f, 0x24, 0x28, 0x8d, 0x4b, 0xd1, 0x33, 0xb0, 0x66, 0x8c, 0x12, 0x41, 0x5d,
	0x39, 0xf7, 0x6d, 0x4b, 0x15, 0x6e, 0xff, 0x9d, 0x96, 0x39, 0xcb, 0x5e, 0x0a, 0x18, 0x52, 0xba,
	0x04, 0xd0, 0x53, 0x80, 0x54, 0x5f, 0xe5, 0xee, 0x6e, 0xcd, 0x35, 0x15, 0x5b, 0xa5, 0xfe, 0x17,
	0x60, 0xc9, 0x29, 0x73, 0xe9, 0x82, 0x04, 0xa1, 0xdd, 0x48, 0x0b, 0x2c, 0x11, 0x47, 0x02, 0xb9,
	0x46, 0x84, 0x0f, 0x6a, 0x44, 0x34, 0x86, 0x5d, 0x2e, 0xfb, 0xd0, 0x9d, 0xcb, 0x46, 0xe4, 0x76,
	0x55, 0x79, 0xf9, 0xc1, 0x7b, 0xc9, 0xe9, 0xee, 0xc5, 0x16, 0x5f, 0xc7, 0x1c, 0xb9, 0x70, 0x9b,
	0xca, 0x59, 0x46, 0x04, 0xf5, 0xdc, 0xbc, 0x09, 0x8c, 0x6b, 0x77, 0xc9, 0xad, 0xb5, 0x90, 0x93,
	0x73, 0xce, 0x4b, 0xa8, 0xaf, 0x34, 0x23, 0x9d, 0x66, 0xb6, 0x79, 0x6d, 0xe5, 0x5a, 0xa6, 0xa0,
	0x66, 0x1b, 0x9a, 0x40, 0x25, 0x24, 0x97, 0x34, 0xe4, 0xf6, 0x4d, 0x25, 0xf5, 0xe4, 0x7a, 0xde,
	0x52, 0xa9, 0x4e, 0x24, 0xd8, 0x6b, 0xac, 0x75, 0xf6, 0x9f, 0x82, 0x95, 0x83, 0x51, 0x03, 0x4a,
	0xaf, 0xe8, 0x6b, 0xdd, 0xd9, 0x32, 0xfc, 0xf3, 0xf1, 0xff, 0x79, 0xf1, 0x49, 0xa1, 0x75, 0x00,
	0x95, 0xb4, 0x73, 0x91, 0x05, 0xd5, 0x7e, 0x6f, 0x72, 0x76, 0x8e, 0x9d, 0xc6, 0x0d, 0x54, 0x85,
	0xd2, 0x70, 0x7c, 0xda, 0x28, 0xb4, 0x1e, 0x80, 0x91, 0xb9, 0x19, 0x19, 0x50, 0x1e, 0x8c, 0x4e,
	0xc6, 0x8d, 0x1b, 0x92, 0xfb, 0x65, 0x0f, 0x8f, 0x06, 0xa3, 0xd3, 0x46, 0x01, 0x99, 0xb0, 0xe3,
	0x60, 0x3c, 0xc6, 0x8d, 0x62, 0xeb, 0xb7, 0x32, 0x18, 0xc7, 0xe9, 0x73, 0xd3, 0x77, 0xe6, 0x8b,
	0x0d, 0xd5, 0x84, 0xc5, 0xdf, 0xd1, 0x99, 0xd0, 0x8f, 0x91, 0x2d, 0xe5, 0xb7, 0xc8, 0x32, 0x0a,
	0xbe, 0x5f, 0x06, 0xf3, 0x80, 0x32, 0xdd, 0xdf, 0x39, 0x44, 0x36, 0x78, 0xfe, 0x33, 0xa0, 0xac,
	0x08, 0x79, 0x08, 0xdd, 0x03, 0x2b, 0xe0, 0x6e, 0x10, 0xc9, 0xe9, 0xb3, 0xca, 0x7a, 0x0b, 0x02,
	0x3e, 0xd0, 0x08, 0xfa, 0x1f, 0xd4, 0x88, 0x4f, 0x23, 0xe1, 0xae, 0x28, 0x93, 0x37, 0xab, 0xdf,
	0x79, 0xbb, 0x0a, 0xbc, 0x48, 0x31, 0xad, 0xe2, 0x05, 0x5c, 0xde, 0x93, 0x67, 0x57, 0x33, 0x95,
	0x63, 0x8d, 0xe4, 0x1a, 0xc1, 0xf8, 0xb0, 0x46, 0x78, 0x09, 0x7b, 0x9b, 0xdf, 0x7f, 0x5c, 0xfb,
	0xea, 0xf0, 0x1d, 0xc5, 0x94, 0xd7, 0x5e, 0x3d, 0xd2, 0xe3, 0xb1, 0x9f, 0x26, 0xe0, 0x3a, 0xcf,
	0x2f, 0x39, 0xfa, 0x06, 0xfe, 0x45, 0xaf, 0x84, 0xfb, 0xb6, 0x6c, 0x4d, 0xc9, 0x3e, 0xfc, 0x1b,
	0x59, 0xe7, 0x4a, 0xd0, 0xc8, 0xa3, 0xde, 0xa6, 0xfc, 0x4d, 0x7a, 0x25, 0xa6, 0x9b, 0x27, 0x8c,
	0xd6, 0xc6, 0xb5, 0x94, 0xe8, 0x67, 0xdb, 0x7f, 0x7d, 0x66, 0x86, 0x8f, 0x6c, 0xdb, 0xa3, 0x1f,
	0x0b, 0x70, 0x7f, 0x16, 0x2f, 0xb6, 0x3e, 0xc0, 0x91, 0x79, 0x4c, 0x04, 0x99, 0xc8, 0xe1, 0x37,
	0x29, 0x7c, 0xfd, 0x42, 0xd3, 0xfd, 0x38, 0x24, 0x91, 0xdf, 0x8e, 0x99, 0xdf, 0xf1, 0x69, 0xa4,
	0x46, 0x63, 0x27, 0xdd, 0x22, 0x49, 0xc0, 0xff, 0xfa, 0x9f, 0xc1, 0xb3, 0x0d, 0xe0, 0xb2, 0xa2,
	0x32, 0x1f, 0xff, 0x11, 0x00, 0x00, 0xff, 0xff, 0x56, 0xc4, 0x6b, 0x67, 0x52, 0x0c, 0x00, 0x00,
}