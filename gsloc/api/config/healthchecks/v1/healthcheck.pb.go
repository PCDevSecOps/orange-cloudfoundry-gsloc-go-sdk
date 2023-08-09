// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: gsloc/api/config/healthchecks/v1/healthcheck.proto

package hcconf

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/orange-cloudfoundry/gsloc-go-sdk/gsloc/api/config/core/v1"
	_ "github.com/orange-cloudfoundry/gsloc-go-sdk/gsloc/type/matcher/v1"
	v11 "github.com/orange-cloudfoundry/gsloc-go-sdk/gsloc/type/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// HTTP request method.
type RequestMethod int32

const (
	RequestMethod_METHOD_UNSPECIFIED RequestMethod = 0
	RequestMethod_GET                RequestMethod = 1
	RequestMethod_HEAD               RequestMethod = 2
	RequestMethod_POST               RequestMethod = 3
	RequestMethod_PUT                RequestMethod = 4
	RequestMethod_DELETE             RequestMethod = 5
	RequestMethod_CONNECT            RequestMethod = 6
	RequestMethod_OPTIONS            RequestMethod = 7
	RequestMethod_TRACE              RequestMethod = 8
	RequestMethod_PATCH              RequestMethod = 9
)

// Enum value maps for RequestMethod.
var (
	RequestMethod_name = map[int32]string{
		0: "METHOD_UNSPECIFIED",
		1: "GET",
		2: "HEAD",
		3: "POST",
		4: "PUT",
		5: "DELETE",
		6: "CONNECT",
		7: "OPTIONS",
		8: "TRACE",
		9: "PATCH",
	}
	RequestMethod_value = map[string]int32{
		"METHOD_UNSPECIFIED": 0,
		"GET":                1,
		"HEAD":               2,
		"POST":               3,
		"PUT":                4,
		"DELETE":             5,
		"CONNECT":            6,
		"OPTIONS":            7,
		"TRACE":              8,
		"PATCH":              9,
	}
)

func (x RequestMethod) Enum() *RequestMethod {
	p := new(RequestMethod)
	*p = x
	return p
}

func (x RequestMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_gsloc_api_config_healthchecks_v1_healthcheck_proto_enumTypes[0].Descriptor()
}

func (RequestMethod) Type() protoreflect.EnumType {
	return &file_gsloc_api_config_healthchecks_v1_healthcheck_proto_enumTypes[0]
}

func (x RequestMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestMethod.Descriptor instead.
func (RequestMethod) EnumDescriptor() ([]byte, []int) {
	return file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDescGZIP(), []int{0}
}

// Health check configuration.
type HealthCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The time to wait for a health check response. If the timeout is reached the
	// health check attempt will be considered a failure.
	Timeout *durationpb.Duration `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// The interval between health checks.
	Interval *durationpb.Duration `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	// port specifies the port to use when performing health checks.
	Port uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	// Types that are assignable to HealthChecker:
	//
	//	*HealthCheck_HttpHealthCheck
	//	*HealthCheck_TcpHealthCheck
	//	*HealthCheck_GrpcHealthCheck
	//	*HealthCheck_NoHealthCheck
	HealthChecker isHealthCheck_HealthChecker `protobuf_oneof:"health_checker"`
	// Enable TLS for healthcheck.
	EnableTls bool `protobuf:"varint,8,opt,name=enable_tls,json=enableTls,proto3" json:"enable_tls,omitempty"`
}

func (x *HealthCheck) Reset() {
	*x = HealthCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheck) ProtoMessage() {}

func (x *HealthCheck) ProtoReflect() protoreflect.Message {
	mi := &file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheck.ProtoReflect.Descriptor instead.
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDescGZIP(), []int{0}
}

func (x *HealthCheck) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *HealthCheck) GetInterval() *durationpb.Duration {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *HealthCheck) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (m *HealthCheck) GetHealthChecker() isHealthCheck_HealthChecker {
	if m != nil {
		return m.HealthChecker
	}
	return nil
}

func (x *HealthCheck) GetHttpHealthCheck() *HttpHealthCheck {
	if x, ok := x.GetHealthChecker().(*HealthCheck_HttpHealthCheck); ok {
		return x.HttpHealthCheck
	}
	return nil
}

func (x *HealthCheck) GetTcpHealthCheck() *TcpHealthCheck {
	if x, ok := x.GetHealthChecker().(*HealthCheck_TcpHealthCheck); ok {
		return x.TcpHealthCheck
	}
	return nil
}

func (x *HealthCheck) GetGrpcHealthCheck() *GrpcHealthCheck {
	if x, ok := x.GetHealthChecker().(*HealthCheck_GrpcHealthCheck); ok {
		return x.GrpcHealthCheck
	}
	return nil
}

func (x *HealthCheck) GetNoHealthCheck() *NoHealthCheck {
	if x, ok := x.GetHealthChecker().(*HealthCheck_NoHealthCheck); ok {
		return x.NoHealthCheck
	}
	return nil
}

func (x *HealthCheck) GetEnableTls() bool {
	if x != nil {
		return x.EnableTls
	}
	return false
}

type isHealthCheck_HealthChecker interface {
	isHealthCheck_HealthChecker()
}

type HealthCheck_HttpHealthCheck struct {
	// HTTP health check.
	HttpHealthCheck *HttpHealthCheck `protobuf:"bytes,4,opt,name=http_health_check,json=httpHealthCheck,proto3,oneof"`
}

type HealthCheck_TcpHealthCheck struct {
	// TCP health check.
	TcpHealthCheck *TcpHealthCheck `protobuf:"bytes,5,opt,name=tcp_health_check,json=tcpHealthCheck,proto3,oneof"`
}

type HealthCheck_GrpcHealthCheck struct {
	// gRPC health check.
	GrpcHealthCheck *GrpcHealthCheck `protobuf:"bytes,6,opt,name=grpc_health_check,json=grpcHealthCheck,proto3,oneof"`
}

type HealthCheck_NoHealthCheck struct {
	// No health check.
	NoHealthCheck *NoHealthCheck `protobuf:"bytes,7,opt,name=no_health_check,json=noHealthCheck,proto3,oneof"`
}

func (*HealthCheck_HttpHealthCheck) isHealthCheck_HealthChecker() {}

func (*HealthCheck_TcpHealthCheck) isHealthCheck_HealthChecker() {}

func (*HealthCheck_GrpcHealthCheck) isHealthCheck_HealthChecker() {}

func (*HealthCheck_NoHealthCheck) isHealthCheck_HealthChecker() {}

// Describes the encoding of the payload bytes in the payload.
type HealthCheckPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*HealthCheckPayload_Text
	//	*HealthCheckPayload_Binary
	Payload isHealthCheckPayload_Payload `protobuf_oneof:"payload"`
}

func (x *HealthCheckPayload) Reset() {
	*x = HealthCheckPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckPayload) ProtoMessage() {}

func (x *HealthCheckPayload) ProtoReflect() protoreflect.Message {
	mi := &file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckPayload.ProtoReflect.Descriptor instead.
func (*HealthCheckPayload) Descriptor() ([]byte, []int) {
	return file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDescGZIP(), []int{1}
}

func (m *HealthCheckPayload) GetPayload() isHealthCheckPayload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *HealthCheckPayload) GetText() string {
	if x, ok := x.GetPayload().(*HealthCheckPayload_Text); ok {
		return x.Text
	}
	return ""
}

func (x *HealthCheckPayload) GetBinary() []byte {
	if x, ok := x.GetPayload().(*HealthCheckPayload_Binary); ok {
		return x.Binary
	}
	return nil
}

type isHealthCheckPayload_Payload interface {
	isHealthCheckPayload_Payload()
}

type HealthCheckPayload_Text struct {
	// Hex encoded payload. E.g., "000000FF".
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type HealthCheckPayload_Binary struct {
	// Binary payload.
	Binary []byte `protobuf:"bytes,2,opt,name=binary,proto3,oneof"`
}

func (*HealthCheckPayload_Text) isHealthCheckPayload_Payload() {}

func (*HealthCheckPayload_Binary) isHealthCheckPayload_Payload() {}

type HttpHealthCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The value of the host header in the HTTP health check request. If
	// left empty (default value), the name of the cluster this health check is associated
	// with will be used.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Specifies the HTTP path that will be requested during health checking. For example
	// */healthcheck*.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// HTTP specific payload.
	Send *HealthCheckPayload `protobuf:"bytes,3,opt,name=send,proto3" json:"send,omitempty"`
	// HTTP specific response.
	Receive *HealthCheckPayload `protobuf:"bytes,4,opt,name=receive,proto3" json:"receive,omitempty"`
	// Specifies a list of HTTP headers that should be added to each request that is sent to the
	// health checked cluster.
	RequestHeadersToAdd []*v1.HeaderValueOption `protobuf:"bytes,5,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	// Specifies a list of HTTP response statuses considered healthy. If provided, replaces default
	// 200-only policy - 200 must be included explicitly as needed. Ranges follow half-open
	// semantics of Int64Range. The start and end of each
	// range are required. Only statuses in the range [100, 600) are allowed.
	ExpectedStatuses *v11.Int64Range `protobuf:"bytes,6,opt,name=expected_statuses,json=expectedStatuses,proto3" json:"expected_statuses,omitempty"`
	// Use specified application protocol for health checks.
	CodecClientType v11.CodecClientType `protobuf:"varint,7,opt,name=codec_client_type,json=codecClientType,proto3,enum=gsloc.type.v1.CodecClientType" json:"codec_client_type,omitempty"`
	// HTTP Method that will be used for health checking, default is "GET".
	// GET, HEAD, POST, PUT, DELETE, OPTIONS, TRACE, PATCH methods are supported, but making request body is not supported.
	// CONNECT method is disallowed because it is not appropriate for health check request.
	// If a non-200 response is expected by the method, it needs to be set in expected_statuses.
	Method RequestMethod `protobuf:"varint,8,opt,name=method,proto3,enum=gsloc.api.config.healthchecks.v1.RequestMethod" json:"method,omitempty"`
}

func (x *HttpHealthCheck) Reset() {
	*x = HttpHealthCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpHealthCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpHealthCheck) ProtoMessage() {}

func (x *HttpHealthCheck) ProtoReflect() protoreflect.Message {
	mi := &file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpHealthCheck.ProtoReflect.Descriptor instead.
func (*HttpHealthCheck) Descriptor() ([]byte, []int) {
	return file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDescGZIP(), []int{2}
}

func (x *HttpHealthCheck) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *HttpHealthCheck) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *HttpHealthCheck) GetSend() *HealthCheckPayload {
	if x != nil {
		return x.Send
	}
	return nil
}

func (x *HttpHealthCheck) GetReceive() *HealthCheckPayload {
	if x != nil {
		return x.Receive
	}
	return nil
}

func (x *HttpHealthCheck) GetRequestHeadersToAdd() []*v1.HeaderValueOption {
	if x != nil {
		return x.RequestHeadersToAdd
	}
	return nil
}

func (x *HttpHealthCheck) GetExpectedStatuses() *v11.Int64Range {
	if x != nil {
		return x.ExpectedStatuses
	}
	return nil
}

func (x *HttpHealthCheck) GetCodecClientType() v11.CodecClientType {
	if x != nil {
		return x.CodecClientType
	}
	return v11.CodecClientType(0)
}

func (x *HttpHealthCheck) GetMethod() RequestMethod {
	if x != nil {
		return x.Method
	}
	return RequestMethod_METHOD_UNSPECIFIED
}

type TcpHealthCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Empty payloads imply a connect-only health check.
	Send *HealthCheckPayload `protobuf:"bytes,1,opt,name=send,proto3" json:"send,omitempty"`
	// When checking the response, “fuzzy” matching is performed such that each
	// binary block must be found, and in the order specified, but not
	// necessarily contiguous.
	Receive []*HealthCheckPayload `protobuf:"bytes,2,rep,name=receive,proto3" json:"receive,omitempty"`
}

func (x *TcpHealthCheck) Reset() {
	*x = TcpHealthCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TcpHealthCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TcpHealthCheck) ProtoMessage() {}

func (x *TcpHealthCheck) ProtoReflect() protoreflect.Message {
	mi := &file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TcpHealthCheck.ProtoReflect.Descriptor instead.
func (*TcpHealthCheck) Descriptor() ([]byte, []int) {
	return file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDescGZIP(), []int{3}
}

func (x *TcpHealthCheck) GetSend() *HealthCheckPayload {
	if x != nil {
		return x.Send
	}
	return nil
}

func (x *TcpHealthCheck) GetReceive() []*HealthCheckPayload {
	if x != nil {
		return x.Receive
	}
	return nil
}

// `grpc.health.v1.Health
// <https://github.com/grpc/grpc/blob/master/src/proto/grpc/health/v1/health.proto>`_-based
// healthcheck. See `gRPC doc <https://github.com/grpc/grpc/blob/master/doc/health-checking.md>`_
// for details.
type GrpcHealthCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An optional service name parameter which will be sent to gRPC service in
	// `grpc.health.v1.HealthCheckRequest
	// <https://github.com/grpc/grpc/blob/master/src/proto/grpc/health/v1/health.proto#L20>`_.
	// message. See `gRPC health-checking overview
	// <https://github.com/grpc/grpc/blob/master/doc/health-checking.md>`_ for more information.
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The value of the :authority header in the gRPC health check request. If
	// left empty (default value), the name of the cluster this health check is associated
	// with will be used. The authority header can be customized for a specific endpoint by setting
	// the HealthCheckConfig.hostname field.
	Authority string `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (x *GrpcHealthCheck) Reset() {
	*x = GrpcHealthCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcHealthCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcHealthCheck) ProtoMessage() {}

func (x *GrpcHealthCheck) ProtoReflect() protoreflect.Message {
	mi := &file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcHealthCheck.ProtoReflect.Descriptor instead.
func (*GrpcHealthCheck) Descriptor() ([]byte, []int) {
	return file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDescGZIP(), []int{4}
}

func (x *GrpcHealthCheck) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *GrpcHealthCheck) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

// No health check. This health check is always considered healthy.
// This is particularly useful for udp route which cannot be health checked.
type NoHealthCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoHealthCheck) Reset() {
	*x = NoHealthCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoHealthCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoHealthCheck) ProtoMessage() {}

func (x *NoHealthCheck) ProtoReflect() protoreflect.Message {
	mi := &file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoHealthCheck.ProtoReflect.Descriptor instead.
func (*NoHealthCheck) Descriptor() ([]byte, []int) {
	return file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDescGZIP(), []int{5}
}

var File_gsloc_api_config_healthchecks_v1_healthcheck_proto protoreflect.FileDescriptor

var file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67, 0x73, 0x6c, 0x6f, 0x63,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x04, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x3f, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0xaa, 0x01, 0x04, 0x08, 0x01, 0x2a, 0x00, 0x52, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x41, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0xaa, 0x01, 0x04, 0x08, 0x01, 0x2a, 0x00,
	0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20,
	0x00, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x5f, 0x0a, 0x11, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x5c, 0x0a, 0x10, 0x74, 0x63, 0x70, 0x5f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x63, 0x70, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x0e, 0x74, 0x63, 0x70, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x5f, 0x0a, 0x11, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x59, 0x0a, 0x0f, 0x6e, 0x6f, 0x5f, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2f, 0x2e, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x48, 0x00, 0x52, 0x0d, 0x6e, 0x6f, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x6c, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x6c,
	0x73, 0x42, 0x15, 0x0a, 0x0e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x5d, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1d,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a,
	0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x42, 0x0e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0xd2, 0x04, 0x0a, 0x0f, 0x48, 0x74, 0x74, 0x70,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x22, 0x0a, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x72, 0x09,
	0xc8, 0x01, 0x00, 0xd0, 0x01, 0x01, 0xc0, 0x01, 0x02, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xfa,
	0x42, 0x0a, 0x72, 0x08, 0x10, 0x01, 0xc8, 0x01, 0x00, 0xc0, 0x01, 0x02, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x48, 0x0a, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x4e, 0x0a, 0x07,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e,
	0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x12, 0x6b, 0x0a, 0x16,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f,
	0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67,
	0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x92, 0x01,
	0x03, 0x10, 0xe8, 0x07, 0x52, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x12, 0x46, 0x0a, 0x11, 0x65, 0x78, 0x70,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x10, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x12, 0x54, 0x0a, 0x11, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x67,
	0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x63, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x53, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x82, 0x01, 0x04,
	0x10, 0x01, 0x20, 0x06, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0xaa, 0x01, 0x0a,
	0x0e, 0x54, 0x63, 0x70, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x48, 0x0a, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e,
	0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x4e, 0x0a, 0x07, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x73, 0x6c,
	0x6f, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x22, 0x5f, 0x0a, 0x0f, 0x47, 0x72, 0x70,
	0x63, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x29, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xc8, 0x01, 0x00, 0xc0, 0x01, 0x02, 0x52,
	0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x0f, 0x0a, 0x0d, 0x4e, 0x6f,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2a, 0x89, 0x01, 0x0a, 0x0d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a,
	0x12, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x48, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x53, 0x54,
	0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10,
	0x07, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x52, 0x41, 0x43, 0x45, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05,
	0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x09, 0x42, 0x55, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x2f, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2d,
	0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x73, 0x6c, 0x6f, 0x63, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x68, 0x63, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDescOnce sync.Once
	file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDescData = file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDesc
)

func file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDescGZIP() []byte {
	file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDescOnce.Do(func() {
		file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDescData = protoimpl.X.CompressGZIP(file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDescData)
	})
	return file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDescData
}

var file_gsloc_api_config_healthchecks_v1_healthcheck_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gsloc_api_config_healthchecks_v1_healthcheck_proto_goTypes = []interface{}{
	(RequestMethod)(0),           // 0: gsloc.api.config.healthchecks.v1.RequestMethod
	(*HealthCheck)(nil),          // 1: gsloc.api.config.healthchecks.v1.HealthCheck
	(*HealthCheckPayload)(nil),   // 2: gsloc.api.config.healthchecks.v1.HealthCheckPayload
	(*HttpHealthCheck)(nil),      // 3: gsloc.api.config.healthchecks.v1.HttpHealthCheck
	(*TcpHealthCheck)(nil),       // 4: gsloc.api.config.healthchecks.v1.TcpHealthCheck
	(*GrpcHealthCheck)(nil),      // 5: gsloc.api.config.healthchecks.v1.GrpcHealthCheck
	(*NoHealthCheck)(nil),        // 6: gsloc.api.config.healthchecks.v1.NoHealthCheck
	(*durationpb.Duration)(nil),  // 7: google.protobuf.Duration
	(*v1.HeaderValueOption)(nil), // 8: gsloc.api.config.core.v1.HeaderValueOption
	(*v11.Int64Range)(nil),       // 9: gsloc.type.v1.Int64Range
	(v11.CodecClientType)(0),     // 10: gsloc.type.v1.CodecClientType
}
var file_gsloc_api_config_healthchecks_v1_healthcheck_proto_depIdxs = []int32{
	7,  // 0: gsloc.api.config.healthchecks.v1.HealthCheck.timeout:type_name -> google.protobuf.Duration
	7,  // 1: gsloc.api.config.healthchecks.v1.HealthCheck.interval:type_name -> google.protobuf.Duration
	3,  // 2: gsloc.api.config.healthchecks.v1.HealthCheck.http_health_check:type_name -> gsloc.api.config.healthchecks.v1.HttpHealthCheck
	4,  // 3: gsloc.api.config.healthchecks.v1.HealthCheck.tcp_health_check:type_name -> gsloc.api.config.healthchecks.v1.TcpHealthCheck
	5,  // 4: gsloc.api.config.healthchecks.v1.HealthCheck.grpc_health_check:type_name -> gsloc.api.config.healthchecks.v1.GrpcHealthCheck
	6,  // 5: gsloc.api.config.healthchecks.v1.HealthCheck.no_health_check:type_name -> gsloc.api.config.healthchecks.v1.NoHealthCheck
	2,  // 6: gsloc.api.config.healthchecks.v1.HttpHealthCheck.send:type_name -> gsloc.api.config.healthchecks.v1.HealthCheckPayload
	2,  // 7: gsloc.api.config.healthchecks.v1.HttpHealthCheck.receive:type_name -> gsloc.api.config.healthchecks.v1.HealthCheckPayload
	8,  // 8: gsloc.api.config.healthchecks.v1.HttpHealthCheck.request_headers_to_add:type_name -> gsloc.api.config.core.v1.HeaderValueOption
	9,  // 9: gsloc.api.config.healthchecks.v1.HttpHealthCheck.expected_statuses:type_name -> gsloc.type.v1.Int64Range
	10, // 10: gsloc.api.config.healthchecks.v1.HttpHealthCheck.codec_client_type:type_name -> gsloc.type.v1.CodecClientType
	0,  // 11: gsloc.api.config.healthchecks.v1.HttpHealthCheck.method:type_name -> gsloc.api.config.healthchecks.v1.RequestMethod
	2,  // 12: gsloc.api.config.healthchecks.v1.TcpHealthCheck.send:type_name -> gsloc.api.config.healthchecks.v1.HealthCheckPayload
	2,  // 13: gsloc.api.config.healthchecks.v1.TcpHealthCheck.receive:type_name -> gsloc.api.config.healthchecks.v1.HealthCheckPayload
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_gsloc_api_config_healthchecks_v1_healthcheck_proto_init() }
func file_gsloc_api_config_healthchecks_v1_healthcheck_proto_init() {
	if File_gsloc_api_config_healthchecks_v1_healthcheck_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpHealthCheck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TcpHealthCheck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcHealthCheck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoHealthCheck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*HealthCheck_HttpHealthCheck)(nil),
		(*HealthCheck_TcpHealthCheck)(nil),
		(*HealthCheck_GrpcHealthCheck)(nil),
		(*HealthCheck_NoHealthCheck)(nil),
	}
	file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*HealthCheckPayload_Text)(nil),
		(*HealthCheckPayload_Binary)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gsloc_api_config_healthchecks_v1_healthcheck_proto_goTypes,
		DependencyIndexes: file_gsloc_api_config_healthchecks_v1_healthcheck_proto_depIdxs,
		EnumInfos:         file_gsloc_api_config_healthchecks_v1_healthcheck_proto_enumTypes,
		MessageInfos:      file_gsloc_api_config_healthchecks_v1_healthcheck_proto_msgTypes,
	}.Build()
	File_gsloc_api_config_healthchecks_v1_healthcheck_proto = out.File
	file_gsloc_api_config_healthchecks_v1_healthcheck_proto_rawDesc = nil
	file_gsloc_api_config_healthchecks_v1_healthcheck_proto_goTypes = nil
	file_gsloc_api_config_healthchecks_v1_healthcheck_proto_depIdxs = nil
}
