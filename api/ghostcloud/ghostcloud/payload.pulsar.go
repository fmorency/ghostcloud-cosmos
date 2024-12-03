// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package ghostcloud

import (
	fmt "fmt"
	io "io"
	reflect "reflect"
	sync "sync"

	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

var (
	md_Payload         protoreflect.MessageDescriptor
	fd_Payload_dataset protoreflect.FieldDescriptor
	fd_Payload_archive protoreflect.FieldDescriptor
)

func init() {
	file_ghostcloud_ghostcloud_payload_proto_init()
	md_Payload = File_ghostcloud_ghostcloud_payload_proto.Messages().ByName("Payload")
	fd_Payload_dataset = md_Payload.Fields().ByName("dataset")
	fd_Payload_archive = md_Payload.Fields().ByName("archive")
}

var _ protoreflect.Message = (*fastReflection_Payload)(nil)

type fastReflection_Payload Payload

func (x *Payload) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Payload)(x)
}

func (x *Payload) slowProtoReflect() protoreflect.Message {
	mi := &file_ghostcloud_ghostcloud_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Payload_messageType fastReflection_Payload_messageType
var _ protoreflect.MessageType = fastReflection_Payload_messageType{}

type fastReflection_Payload_messageType struct{}

func (x fastReflection_Payload_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Payload)(nil)
}
func (x fastReflection_Payload_messageType) New() protoreflect.Message {
	return new(fastReflection_Payload)
}
func (x fastReflection_Payload_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Payload
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Payload) Descriptor() protoreflect.MessageDescriptor {
	return md_Payload
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Payload) Type() protoreflect.MessageType {
	return _fastReflection_Payload_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Payload) New() protoreflect.Message {
	return new(fastReflection_Payload)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Payload) Interface() protoreflect.ProtoMessage {
	return (*Payload)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Payload) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.PayloadOption != nil {
		switch o := x.PayloadOption.(type) {
		case *Payload_Dataset:
			v := o.Dataset
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_Payload_dataset, value) {
				return
			}
		case *Payload_Archive:
			v := o.Archive
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_Payload_archive, value) {
				return
			}
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Payload) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "ghostcloud.ghostcloud.Payload.dataset":
		if x.PayloadOption == nil {
			return false
		} else if _, ok := x.PayloadOption.(*Payload_Dataset); ok {
			return true
		} else {
			return false
		}
	case "ghostcloud.ghostcloud.Payload.archive":
		if x.PayloadOption == nil {
			return false
		} else if _, ok := x.PayloadOption.(*Payload_Archive); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ghostcloud.ghostcloud.Payload"))
		}
		panic(fmt.Errorf("message ghostcloud.ghostcloud.Payload does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Payload) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "ghostcloud.ghostcloud.Payload.dataset":
		x.PayloadOption = nil
	case "ghostcloud.ghostcloud.Payload.archive":
		x.PayloadOption = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ghostcloud.ghostcloud.Payload"))
		}
		panic(fmt.Errorf("message ghostcloud.ghostcloud.Payload does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Payload) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "ghostcloud.ghostcloud.Payload.dataset":
		if x.PayloadOption == nil {
			return protoreflect.ValueOfMessage((*Dataset)(nil).ProtoReflect())
		} else if v, ok := x.PayloadOption.(*Payload_Dataset); ok {
			return protoreflect.ValueOfMessage(v.Dataset.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*Dataset)(nil).ProtoReflect())
		}
	case "ghostcloud.ghostcloud.Payload.archive":
		if x.PayloadOption == nil {
			return protoreflect.ValueOfMessage((*Archive)(nil).ProtoReflect())
		} else if v, ok := x.PayloadOption.(*Payload_Archive); ok {
			return protoreflect.ValueOfMessage(v.Archive.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*Archive)(nil).ProtoReflect())
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ghostcloud.ghostcloud.Payload"))
		}
		panic(fmt.Errorf("message ghostcloud.ghostcloud.Payload does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Payload) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "ghostcloud.ghostcloud.Payload.dataset":
		cv := value.Message().Interface().(*Dataset)
		x.PayloadOption = &Payload_Dataset{Dataset: cv}
	case "ghostcloud.ghostcloud.Payload.archive":
		cv := value.Message().Interface().(*Archive)
		x.PayloadOption = &Payload_Archive{Archive: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ghostcloud.ghostcloud.Payload"))
		}
		panic(fmt.Errorf("message ghostcloud.ghostcloud.Payload does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Payload) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "ghostcloud.ghostcloud.Payload.dataset":
		if x.PayloadOption == nil {
			value := &Dataset{}
			oneofValue := &Payload_Dataset{Dataset: value}
			x.PayloadOption = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.PayloadOption.(type) {
		case *Payload_Dataset:
			return protoreflect.ValueOfMessage(m.Dataset.ProtoReflect())
		default:
			value := &Dataset{}
			oneofValue := &Payload_Dataset{Dataset: value}
			x.PayloadOption = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "ghostcloud.ghostcloud.Payload.archive":
		if x.PayloadOption == nil {
			value := &Archive{}
			oneofValue := &Payload_Archive{Archive: value}
			x.PayloadOption = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.PayloadOption.(type) {
		case *Payload_Archive:
			return protoreflect.ValueOfMessage(m.Archive.ProtoReflect())
		default:
			value := &Archive{}
			oneofValue := &Payload_Archive{Archive: value}
			x.PayloadOption = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ghostcloud.ghostcloud.Payload"))
		}
		panic(fmt.Errorf("message ghostcloud.ghostcloud.Payload does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Payload) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "ghostcloud.ghostcloud.Payload.dataset":
		value := &Dataset{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "ghostcloud.ghostcloud.Payload.archive":
		value := &Archive{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ghostcloud.ghostcloud.Payload"))
		}
		panic(fmt.Errorf("message ghostcloud.ghostcloud.Payload does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Payload) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "ghostcloud.ghostcloud.Payload.payload_option":
		if x.PayloadOption == nil {
			return nil
		}
		switch x.PayloadOption.(type) {
		case *Payload_Dataset:
			return x.Descriptor().Fields().ByName("dataset")
		case *Payload_Archive:
			return x.Descriptor().Fields().ByName("archive")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in ghostcloud.ghostcloud.Payload", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Payload) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Payload) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Payload) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Payload) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Payload)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		switch x := x.PayloadOption.(type) {
		case *Payload_Dataset:
			if x == nil {
				break
			}
			l = options.Size(x.Dataset)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Payload_Archive:
			if x == nil {
				break
			}
			l = options.Size(x.Archive)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Payload)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		switch x := x.PayloadOption.(type) {
		case *Payload_Dataset:
			encoded, err := options.Marshal(x.Dataset)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		case *Payload_Archive:
			encoded, err := options.Marshal(x.Archive)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Payload)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Payload: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Payload: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Dataset", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v := &Dataset{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.PayloadOption = &Payload_Dataset{v}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Archive", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v := &Archive{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.PayloadOption = &Payload_Archive{v}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: ghostcloud/ghostcloud/payload.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to PayloadOption:
	//
	//	*Payload_Dataset
	//	*Payload_Archive
	PayloadOption isPayload_PayloadOption `protobuf_oneof:"payload_option"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ghostcloud_ghostcloud_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_ghostcloud_ghostcloud_payload_proto_rawDescGZIP(), []int{0}
}

func (x *Payload) GetPayloadOption() isPayload_PayloadOption {
	if x != nil {
		return x.PayloadOption
	}
	return nil
}

func (x *Payload) GetDataset() *Dataset {
	if x, ok := x.GetPayloadOption().(*Payload_Dataset); ok {
		return x.Dataset
	}
	return nil
}

func (x *Payload) GetArchive() *Archive {
	if x, ok := x.GetPayloadOption().(*Payload_Archive); ok {
		return x.Archive
	}
	return nil
}

type isPayload_PayloadOption interface {
	isPayload_PayloadOption()
}

type Payload_Dataset struct {
	Dataset *Dataset `protobuf:"bytes,1,opt,name=dataset,proto3,oneof"`
}

type Payload_Archive struct {
	Archive *Archive `protobuf:"bytes,2,opt,name=archive,proto3,oneof"`
}

func (*Payload_Dataset) isPayload_PayloadOption() {}

func (*Payload_Archive) isPayload_PayloadOption() {}

var File_ghostcloud_ghostcloud_payload_proto protoreflect.FileDescriptor

var file_ghostcloud_ghostcloud_payload_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x68, 0x6f,
	0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x1a, 0x23, 0x67, 0x68,
	0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x23, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x68,
	0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x3a, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x48, 0x00, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x3a,
	0x0a, 0x07, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x68, 0x6f,
	0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x48,
	0x00, 0x52, 0x07, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xda, 0x01, 0x0a,
	0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x0c, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x66, 0x74, 0x65, 0x64, 0x69, 0x6e, 0x69,
	0x74, 0x2f, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x68, 0x6f, 0x73,
	0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0xa2, 0x02, 0x03, 0x47, 0x47, 0x58, 0xaa, 0x02, 0x15, 0x47,
	0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0xca, 0x02, 0x15, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0xe2, 0x02, 0x21, 0x47,
	0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x16, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x47,
	0x68, 0x6f, 0x73, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ghostcloud_ghostcloud_payload_proto_rawDescOnce sync.Once
	file_ghostcloud_ghostcloud_payload_proto_rawDescData = file_ghostcloud_ghostcloud_payload_proto_rawDesc
)

func file_ghostcloud_ghostcloud_payload_proto_rawDescGZIP() []byte {
	file_ghostcloud_ghostcloud_payload_proto_rawDescOnce.Do(func() {
		file_ghostcloud_ghostcloud_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_ghostcloud_ghostcloud_payload_proto_rawDescData)
	})
	return file_ghostcloud_ghostcloud_payload_proto_rawDescData
}

var file_ghostcloud_ghostcloud_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ghostcloud_ghostcloud_payload_proto_goTypes = []interface{}{
	(*Payload)(nil), // 0: ghostcloud.ghostcloud.Payload
	(*Dataset)(nil), // 1: ghostcloud.ghostcloud.Dataset
	(*Archive)(nil), // 2: ghostcloud.ghostcloud.Archive
}
var file_ghostcloud_ghostcloud_payload_proto_depIdxs = []int32{
	1, // 0: ghostcloud.ghostcloud.Payload.dataset:type_name -> ghostcloud.ghostcloud.Dataset
	2, // 1: ghostcloud.ghostcloud.Payload.archive:type_name -> ghostcloud.ghostcloud.Archive
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ghostcloud_ghostcloud_payload_proto_init() }
func file_ghostcloud_ghostcloud_payload_proto_init() {
	if File_ghostcloud_ghostcloud_payload_proto != nil {
		return
	}
	file_ghostcloud_ghostcloud_archive_proto_init()
	file_ghostcloud_ghostcloud_dataset_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ghostcloud_ghostcloud_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
	file_ghostcloud_ghostcloud_payload_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Payload_Dataset)(nil),
		(*Payload_Archive)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ghostcloud_ghostcloud_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ghostcloud_ghostcloud_payload_proto_goTypes,
		DependencyIndexes: file_ghostcloud_ghostcloud_payload_proto_depIdxs,
		MessageInfos:      file_ghostcloud_ghostcloud_payload_proto_msgTypes,
	}.Build()
	File_ghostcloud_ghostcloud_payload_proto = out.File
	file_ghostcloud_ghostcloud_payload_proto_rawDesc = nil
	file_ghostcloud_ghostcloud_payload_proto_goTypes = nil
	file_ghostcloud_ghostcloud_payload_proto_depIdxs = nil
}
