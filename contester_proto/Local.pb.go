// Code generated by protoc-gen-go from "Local.proto"
// DO NOT EDIT!

package contester_proto

import proto "code.google.com/p/goprotobuf/proto"
import "encoding/json"
import "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type BinaryTypeResponse_Win32BinaryType int32

const (
	BinaryTypeResponse_SCS_32BIT_BINARY BinaryTypeResponse_Win32BinaryType = 0
	BinaryTypeResponse_SCS_DOS_BINARY   BinaryTypeResponse_Win32BinaryType = 1
	BinaryTypeResponse_SCS_WOW_BINARY   BinaryTypeResponse_Win32BinaryType = 2
	BinaryTypeResponse_SCS_PIF_BINARY   BinaryTypeResponse_Win32BinaryType = 3
	BinaryTypeResponse_SCS_POSIX_BINARY BinaryTypeResponse_Win32BinaryType = 4
	BinaryTypeResponse_SCS_OS216_BINARY BinaryTypeResponse_Win32BinaryType = 5
	BinaryTypeResponse_SCS_64BIT_BINARY BinaryTypeResponse_Win32BinaryType = 6
)

var BinaryTypeResponse_Win32BinaryType_name = map[int32]string{
	0: "SCS_32BIT_BINARY",
	1: "SCS_DOS_BINARY",
	2: "SCS_WOW_BINARY",
	3: "SCS_PIF_BINARY",
	4: "SCS_POSIX_BINARY",
	5: "SCS_OS216_BINARY",
	6: "SCS_64BIT_BINARY",
}
var BinaryTypeResponse_Win32BinaryType_value = map[string]int32{
	"SCS_32BIT_BINARY": 0,
	"SCS_DOS_BINARY":   1,
	"SCS_WOW_BINARY":   2,
	"SCS_PIF_BINARY":   3,
	"SCS_POSIX_BINARY": 4,
	"SCS_OS216_BINARY": 5,
	"SCS_64BIT_BINARY": 6,
}

func (x BinaryTypeResponse_Win32BinaryType) Enum() *BinaryTypeResponse_Win32BinaryType {
	p := new(BinaryTypeResponse_Win32BinaryType)
	*p = x
	return p
}
func (x BinaryTypeResponse_Win32BinaryType) String() string {
	return proto.EnumName(BinaryTypeResponse_Win32BinaryType_name, int32(x))
}
func (x BinaryTypeResponse_Win32BinaryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *BinaryTypeResponse_Win32BinaryType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BinaryTypeResponse_Win32BinaryType_value, data, "BinaryTypeResponse_Win32BinaryType")
	if err != nil {
		return err
	}
	*x = BinaryTypeResponse_Win32BinaryType(value)
	return nil
}

type LocalEnvironment struct {
	Empty            *bool                        `protobuf:"varint,1,opt,name=empty" json:"empty,omitempty"`
	Variable         []*LocalEnvironment_Variable `protobuf:"bytes,2,rep,name=variable" json:"variable,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (this *LocalEnvironment) Reset()         { *this = LocalEnvironment{} }
func (this *LocalEnvironment) String() string { return proto.CompactTextString(this) }
func (*LocalEnvironment) ProtoMessage()       {}

func (this *LocalEnvironment) GetEmpty() bool {
	if this != nil && this.Empty != nil {
		return *this.Empty
	}
	return false
}

type LocalEnvironment_Variable struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Expand           *bool   `protobuf:"varint,3,opt,name=expand" json:"expand,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *LocalEnvironment_Variable) Reset()         { *this = LocalEnvironment_Variable{} }
func (this *LocalEnvironment_Variable) String() string { return proto.CompactTextString(this) }
func (*LocalEnvironment_Variable) ProtoMessage()       {}

func (this *LocalEnvironment_Variable) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

func (this *LocalEnvironment_Variable) GetValue() string {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return ""
}

func (this *LocalEnvironment_Variable) GetExpand() bool {
	if this != nil && this.Expand != nil {
		return *this.Expand
	}
	return false
}

type LocalExecutionParameters struct {
	ApplicationName       *string             `protobuf:"bytes,1,opt,name=application_name" json:"application_name,omitempty"`
	CommandLine           *string             `protobuf:"bytes,2,opt,name=command_line" json:"command_line,omitempty"`
	CurrentDirectory      *string             `protobuf:"bytes,3,opt,name=current_directory" json:"current_directory,omitempty"`
	TimeLimitMicros       *uint64             `protobuf:"varint,4,opt,name=time_limit_micros" json:"time_limit_micros,omitempty"`
	MemoryLimit           *uint64             `protobuf:"varint,5,opt,name=memory_limit" json:"memory_limit,omitempty"`
	CheckIdleness         *bool               `protobuf:"varint,6,opt,name=check_idleness" json:"check_idleness,omitempty"`
	Environment           *LocalEnvironment   `protobuf:"bytes,7,opt,name=environment" json:"environment,omitempty"`
	RestrictUi            *bool               `protobuf:"varint,8,opt,name=restrict_ui" json:"restrict_ui,omitempty"`
	NoJob                 *bool               `protobuf:"varint,9,opt,name=no_job" json:"no_job,omitempty"`
	ProcessLimit          *uint32             `protobuf:"varint,10,opt,name=process_limit" json:"process_limit,omitempty"`
	TimeLimitHardMicros   *uint64             `protobuf:"varint,15,opt,name=time_limit_hard_micros" json:"time_limit_hard_micros,omitempty"`
	StdIn                 *RedirectParameters `protobuf:"bytes,12,opt,name=std_in" json:"std_in,omitempty"`
	StdOut                *RedirectParameters `protobuf:"bytes,13,opt,name=std_out" json:"std_out,omitempty"`
	StdErr                *RedirectParameters `protobuf:"bytes,14,opt,name=std_err" json:"std_err,omitempty"`
	CommandLineParameters []string            `protobuf:"bytes,16,rep,name=command_line_parameters" json:"command_line_parameters,omitempty"`
	SandboxId             *string             `protobuf:"bytes,17,opt,name=sandbox_id" json:"sandbox_id,omitempty"`
	XXX_unrecognized      []byte              `json:"-"`
}

func (this *LocalExecutionParameters) Reset()         { *this = LocalExecutionParameters{} }
func (this *LocalExecutionParameters) String() string { return proto.CompactTextString(this) }
func (*LocalExecutionParameters) ProtoMessage()       {}

func (this *LocalExecutionParameters) GetApplicationName() string {
	if this != nil && this.ApplicationName != nil {
		return *this.ApplicationName
	}
	return ""
}

func (this *LocalExecutionParameters) GetCommandLine() string {
	if this != nil && this.CommandLine != nil {
		return *this.CommandLine
	}
	return ""
}

func (this *LocalExecutionParameters) GetCurrentDirectory() string {
	if this != nil && this.CurrentDirectory != nil {
		return *this.CurrentDirectory
	}
	return ""
}

func (this *LocalExecutionParameters) GetTimeLimitMicros() uint64 {
	if this != nil && this.TimeLimitMicros != nil {
		return *this.TimeLimitMicros
	}
	return 0
}

func (this *LocalExecutionParameters) GetMemoryLimit() uint64 {
	if this != nil && this.MemoryLimit != nil {
		return *this.MemoryLimit
	}
	return 0
}

func (this *LocalExecutionParameters) GetCheckIdleness() bool {
	if this != nil && this.CheckIdleness != nil {
		return *this.CheckIdleness
	}
	return false
}

func (this *LocalExecutionParameters) GetEnvironment() *LocalEnvironment {
	if this != nil {
		return this.Environment
	}
	return nil
}

func (this *LocalExecutionParameters) GetRestrictUi() bool {
	if this != nil && this.RestrictUi != nil {
		return *this.RestrictUi
	}
	return false
}

func (this *LocalExecutionParameters) GetNoJob() bool {
	if this != nil && this.NoJob != nil {
		return *this.NoJob
	}
	return false
}

func (this *LocalExecutionParameters) GetProcessLimit() uint32 {
	if this != nil && this.ProcessLimit != nil {
		return *this.ProcessLimit
	}
	return 0
}

func (this *LocalExecutionParameters) GetTimeLimitHardMicros() uint64 {
	if this != nil && this.TimeLimitHardMicros != nil {
		return *this.TimeLimitHardMicros
	}
	return 0
}

func (this *LocalExecutionParameters) GetStdIn() *RedirectParameters {
	if this != nil {
		return this.StdIn
	}
	return nil
}

func (this *LocalExecutionParameters) GetStdOut() *RedirectParameters {
	if this != nil {
		return this.StdOut
	}
	return nil
}

func (this *LocalExecutionParameters) GetStdErr() *RedirectParameters {
	if this != nil {
		return this.StdErr
	}
	return nil
}

func (this *LocalExecutionParameters) GetSandboxId() string {
	if this != nil && this.SandboxId != nil {
		return *this.SandboxId
	}
	return ""
}

type LocalExecutionResult struct {
	Flags            *ExecutionResultFlags `protobuf:"bytes,1,opt,name=flags" json:"flags,omitempty"`
	Time             *ExecutionResultTime  `protobuf:"bytes,2,opt,name=time" json:"time,omitempty"`
	Memory           *uint64               `protobuf:"varint,3,opt,name=memory" json:"memory,omitempty"`
	ReturnCode       *uint32               `protobuf:"varint,4,opt,name=return_code" json:"return_code,omitempty"`
	StdOut           *Blob                 `protobuf:"bytes,5,opt,name=std_out" json:"std_out,omitempty"`
	StdErr           *Blob                 `protobuf:"bytes,6,opt,name=std_err" json:"std_err,omitempty"`
	TotalProcesses   *uint64               `protobuf:"varint,7,opt,name=total_processes" json:"total_processes,omitempty"`
	KillSignal       *int32                `protobuf:"varint,8,opt,name=kill_signal" json:"kill_signal,omitempty"`
	StopSignal       *int32                `protobuf:"varint,9,opt,name=stop_signal" json:"stop_signal,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (this *LocalExecutionResult) Reset()         { *this = LocalExecutionResult{} }
func (this *LocalExecutionResult) String() string { return proto.CompactTextString(this) }
func (*LocalExecutionResult) ProtoMessage()       {}

func (this *LocalExecutionResult) GetFlags() *ExecutionResultFlags {
	if this != nil {
		return this.Flags
	}
	return nil
}

func (this *LocalExecutionResult) GetTime() *ExecutionResultTime {
	if this != nil {
		return this.Time
	}
	return nil
}

func (this *LocalExecutionResult) GetMemory() uint64 {
	if this != nil && this.Memory != nil {
		return *this.Memory
	}
	return 0
}

func (this *LocalExecutionResult) GetReturnCode() uint32 {
	if this != nil && this.ReturnCode != nil {
		return *this.ReturnCode
	}
	return 0
}

func (this *LocalExecutionResult) GetStdOut() *Blob {
	if this != nil {
		return this.StdOut
	}
	return nil
}

func (this *LocalExecutionResult) GetStdErr() *Blob {
	if this != nil {
		return this.StdErr
	}
	return nil
}

func (this *LocalExecutionResult) GetTotalProcesses() uint64 {
	if this != nil && this.TotalProcesses != nil {
		return *this.TotalProcesses
	}
	return 0
}

func (this *LocalExecutionResult) GetKillSignal() int32 {
	if this != nil && this.KillSignal != nil {
		return *this.KillSignal
	}
	return 0
}

func (this *LocalExecutionResult) GetStopSignal() int32 {
	if this != nil && this.StopSignal != nil {
		return *this.StopSignal
	}
	return 0
}

type LocalExecution struct {
	Parameters       *LocalExecutionParameters `protobuf:"bytes,1,req,name=parameters" json:"parameters,omitempty"`
	Result           *LocalExecutionResult     `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (this *LocalExecution) Reset()         { *this = LocalExecution{} }
func (this *LocalExecution) String() string { return proto.CompactTextString(this) }
func (*LocalExecution) ProtoMessage()       {}

func (this *LocalExecution) GetParameters() *LocalExecutionParameters {
	if this != nil {
		return this.Parameters
	}
	return nil
}

func (this *LocalExecution) GetResult() *LocalExecutionResult {
	if this != nil {
		return this.Result
	}
	return nil
}

type OwnerInfo struct {
	Uid              *uint32  `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Pathname         []string `protobuf:"bytes,2,rep,name=pathname" json:"pathname,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *OwnerInfo) Reset()         { *this = OwnerInfo{} }
func (this *OwnerInfo) String() string { return proto.CompactTextString(this) }
func (*OwnerInfo) ProtoMessage()       {}

func (this *OwnerInfo) GetUid() uint32 {
	if this != nil && this.Uid != nil {
		return *this.Uid
	}
	return 0
}

type BinaryTypeRequest struct {
	Pathname         *string `protobuf:"bytes,1,opt,name=pathname" json:"pathname,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *BinaryTypeRequest) Reset()         { *this = BinaryTypeRequest{} }
func (this *BinaryTypeRequest) String() string { return proto.CompactTextString(this) }
func (*BinaryTypeRequest) ProtoMessage()       {}

func (this *BinaryTypeRequest) GetPathname() string {
	if this != nil && this.Pathname != nil {
		return *this.Pathname
	}
	return ""
}

type BinaryTypeResponse struct {
	Failure          *bool                               `protobuf:"varint,1,opt,name=failure" json:"failure,omitempty"`
	Result           *BinaryTypeResponse_Win32BinaryType `protobuf:"varint,2,opt,name=result,enum=contester.proto.BinaryTypeResponse_Win32BinaryType" json:"result,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (this *BinaryTypeResponse) Reset()         { *this = BinaryTypeResponse{} }
func (this *BinaryTypeResponse) String() string { return proto.CompactTextString(this) }
func (*BinaryTypeResponse) ProtoMessage()       {}

func (this *BinaryTypeResponse) GetFailure() bool {
	if this != nil && this.Failure != nil {
		return *this.Failure
	}
	return false
}

func (this *BinaryTypeResponse) GetResult() BinaryTypeResponse_Win32BinaryType {
	if this != nil && this.Result != nil {
		return *this.Result
	}
	return 0
}

type ClearRequest struct {
	Sandbox          *string `protobuf:"bytes,1,opt,name=sandbox" json:"sandbox,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ClearRequest) Reset()         { *this = ClearRequest{} }
func (this *ClearRequest) String() string { return proto.CompactTextString(this) }
func (*ClearRequest) ProtoMessage()       {}

func (this *ClearRequest) GetSandbox() string {
	if this != nil && this.Sandbox != nil {
		return *this.Sandbox
	}
	return ""
}

type IdentifyRequest struct {
	ContesterId      *string `protobuf:"bytes,1,opt,name=contester_id" json:"contester_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *IdentifyRequest) Reset()         { *this = IdentifyRequest{} }
func (this *IdentifyRequest) String() string { return proto.CompactTextString(this) }
func (*IdentifyRequest) ProtoMessage()       {}

func (this *IdentifyRequest) GetContesterId() string {
	if this != nil && this.ContesterId != nil {
		return *this.ContesterId
	}
	return ""
}

type SandboxLocations struct {
	Compile          *string `protobuf:"bytes,1,opt,name=compile" json:"compile,omitempty"`
	Run              *string `protobuf:"bytes,2,opt,name=run" json:"run,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *SandboxLocations) Reset()         { *this = SandboxLocations{} }
func (this *SandboxLocations) String() string { return proto.CompactTextString(this) }
func (*SandboxLocations) ProtoMessage()       {}

func (this *SandboxLocations) GetCompile() string {
	if this != nil && this.Compile != nil {
		return *this.Compile
	}
	return ""
}

func (this *SandboxLocations) GetRun() string {
	if this != nil && this.Run != nil {
		return *this.Run
	}
	return ""
}

type IdentifyResponse struct {
	InvokerId        *string             `protobuf:"bytes,1,opt,name=invoker_id" json:"invoker_id,omitempty"`
	Sandboxes        []*SandboxLocations `protobuf:"bytes,2,rep,name=sandboxes" json:"sandboxes,omitempty"`
	Environment      *LocalEnvironment   `protobuf:"bytes,3,opt,name=environment" json:"environment,omitempty"`
	Platform         *string             `protobuf:"bytes,4,opt,name=platform" json:"platform,omitempty"`
	PathSeparator    *string             `protobuf:"bytes,5,opt,name=path_separator" json:"path_separator,omitempty"`
	Disks            []string            `protobuf:"bytes,6,rep,name=disks" json:"disks,omitempty"`
	ProgramFiles     []string            `protobuf:"bytes,7,rep,name=programFiles" json:"programFiles,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (this *IdentifyResponse) Reset()         { *this = IdentifyResponse{} }
func (this *IdentifyResponse) String() string { return proto.CompactTextString(this) }
func (*IdentifyResponse) ProtoMessage()       {}

func (this *IdentifyResponse) GetInvokerId() string {
	if this != nil && this.InvokerId != nil {
		return *this.InvokerId
	}
	return ""
}

func (this *IdentifyResponse) GetEnvironment() *LocalEnvironment {
	if this != nil {
		return this.Environment
	}
	return nil
}

func (this *IdentifyResponse) GetPlatform() string {
	if this != nil && this.Platform != nil {
		return *this.Platform
	}
	return ""
}

func (this *IdentifyResponse) GetPathSeparator() string {
	if this != nil && this.PathSeparator != nil {
		return *this.PathSeparator
	}
	return ""
}

type FileStat struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	IsDirectory      *bool   `protobuf:"varint,2,opt,name=is_directory" json:"is_directory,omitempty"`
	Size             *uint64 `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *FileStat) Reset()         { *this = FileStat{} }
func (this *FileStat) String() string { return proto.CompactTextString(this) }
func (*FileStat) ProtoMessage()       {}

func (this *FileStat) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

func (this *FileStat) GetIsDirectory() bool {
	if this != nil && this.IsDirectory != nil {
		return *this.IsDirectory
	}
	return false
}

func (this *FileStat) GetSize() uint64 {
	if this != nil && this.Size != nil {
		return *this.Size
	}
	return 0
}

type NameList struct {
	Name             []string `protobuf:"bytes,1,rep,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *NameList) Reset()         { *this = NameList{} }
func (this *NameList) String() string { return proto.CompactTextString(this) }
func (*NameList) ProtoMessage()       {}

type FileStats struct {
	Name             *string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Results          []*FileStat `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (this *FileStats) Reset()         { *this = FileStats{} }
func (this *FileStats) String() string { return proto.CompactTextString(this) }
func (*FileStats) ProtoMessage()       {}

func (this *FileStats) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

type FileStatsList struct {
	Results          []*FileStats `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (this *FileStatsList) Reset()         { *this = FileStatsList{} }
func (this *FileStatsList) String() string { return proto.CompactTextString(this) }
func (*FileStatsList) ProtoMessage()       {}

type PutRequest struct {
	Files            []*File `protobuf:"bytes,1,rep,name=files" json:"files,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *PutRequest) Reset()         { *this = PutRequest{} }
func (this *PutRequest) String() string { return proto.CompactTextString(this) }
func (*PutRequest) ProtoMessage()       {}

type FileContents struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Files            []*File `protobuf:"bytes,2,rep,name=files" json:"files,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *FileContents) Reset()         { *this = FileContents{} }
func (this *FileContents) String() string { return proto.CompactTextString(this) }
func (*FileContents) ProtoMessage()       {}

func (this *FileContents) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

type FileContentsList struct {
	Results          []*FileContents `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (this *FileContentsList) Reset()         { *this = FileContentsList{} }
func (this *FileContentsList) String() string { return proto.CompactTextString(this) }
func (*FileContentsList) ProtoMessage()       {}

type EmptyMessage struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *EmptyMessage) Reset()         { *this = EmptyMessage{} }
func (this *EmptyMessage) String() string { return proto.CompactTextString(this) }
func (*EmptyMessage) ProtoMessage()       {}

func init() {
	proto.RegisterEnum("contester.proto.BinaryTypeResponse_Win32BinaryType", BinaryTypeResponse_Win32BinaryType_name, BinaryTypeResponse_Win32BinaryType_value)
}
