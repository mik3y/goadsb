// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: main.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SBS1Message_MessageType int32

const (
	SBS1Message_UNKNOWN          SBS1Message_MessageType = 0
	SBS1Message_SELECTION_CHANGE SBS1Message_MessageType = 1 // "SEL"
	SBS1Message_NEW_ID           SBS1Message_MessageType = 2 // "ID"
	SBS1Message_NEW_AIRCRAFT     SBS1Message_MessageType = 3 // "AIR"
	SBS1Message_STATUS_AIRCRAFT  SBS1Message_MessageType = 4 // "STA"
	SBS1Message_CLICK            SBS1Message_MessageType = 5 // "CLK"
	SBS1Message_TRANSMISSION     SBS1Message_MessageType = 6 // "MSG"
)

// Enum value maps for SBS1Message_MessageType.
var (
	SBS1Message_MessageType_name = map[int32]string{
		0: "UNKNOWN",
		1: "SELECTION_CHANGE",
		2: "NEW_ID",
		3: "NEW_AIRCRAFT",
		4: "STATUS_AIRCRAFT",
		5: "CLICK",
		6: "TRANSMISSION",
	}
	SBS1Message_MessageType_value = map[string]int32{
		"UNKNOWN":          0,
		"SELECTION_CHANGE": 1,
		"NEW_ID":           2,
		"NEW_AIRCRAFT":     3,
		"STATUS_AIRCRAFT":  4,
		"CLICK":            5,
		"TRANSMISSION":     6,
	}
)

func (x SBS1Message_MessageType) Enum() *SBS1Message_MessageType {
	p := new(SBS1Message_MessageType)
	*p = x
	return p
}

func (x SBS1Message_MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SBS1Message_MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_main_proto_enumTypes[0].Descriptor()
}

func (SBS1Message_MessageType) Type() protoreflect.EnumType {
	return &file_main_proto_enumTypes[0]
}

func (x SBS1Message_MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SBS1Message_MessageType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SBS1Message_MessageType(num)
	return nil
}

// Deprecated: Use SBS1Message_MessageType.Descriptor instead.
func (SBS1Message_MessageType) EnumDescriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{0, 0}
}

type SBS1Message_TransmissionType int32

const (
	SBS1Message_RESERVED              SBS1Message_TransmissionType = 0 // Not used
	SBS1Message_ES_IDENT_AND_CATEGORY SBS1Message_TransmissionType = 1 // ES identification and category
	SBS1Message_ES_SURFACE_POS        SBS1Message_TransmissionType = 2 // ES surface position message
	SBS1Message_ES_AIRBORNE_POS       SBS1Message_TransmissionType = 3 // ES airborne position message
	SBS1Message_ES_AIRBORNE_VEL       SBS1Message_TransmissionType = 4 // ES airborne velocity message
	SBS1Message_SURVEILLANCE_ALT      SBS1Message_TransmissionType = 5 // Surveillance alt message
	SBS1Message_SURVEILLANCE_ID       SBS1Message_TransmissionType = 6 // Surveillance ID message
	SBS1Message_AIR_TO_AIR            SBS1Message_TransmissionType = 7 // Air-to-air message
	SBS1Message_ALL_CALL_REPLY        SBS1Message_TransmissionType = 8 // All call reply
)

// Enum value maps for SBS1Message_TransmissionType.
var (
	SBS1Message_TransmissionType_name = map[int32]string{
		0: "RESERVED",
		1: "ES_IDENT_AND_CATEGORY",
		2: "ES_SURFACE_POS",
		3: "ES_AIRBORNE_POS",
		4: "ES_AIRBORNE_VEL",
		5: "SURVEILLANCE_ALT",
		6: "SURVEILLANCE_ID",
		7: "AIR_TO_AIR",
		8: "ALL_CALL_REPLY",
	}
	SBS1Message_TransmissionType_value = map[string]int32{
		"RESERVED":              0,
		"ES_IDENT_AND_CATEGORY": 1,
		"ES_SURFACE_POS":        2,
		"ES_AIRBORNE_POS":       3,
		"ES_AIRBORNE_VEL":       4,
		"SURVEILLANCE_ALT":      5,
		"SURVEILLANCE_ID":       6,
		"AIR_TO_AIR":            7,
		"ALL_CALL_REPLY":        8,
	}
)

func (x SBS1Message_TransmissionType) Enum() *SBS1Message_TransmissionType {
	p := new(SBS1Message_TransmissionType)
	*p = x
	return p
}

func (x SBS1Message_TransmissionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SBS1Message_TransmissionType) Descriptor() protoreflect.EnumDescriptor {
	return file_main_proto_enumTypes[1].Descriptor()
}

func (SBS1Message_TransmissionType) Type() protoreflect.EnumType {
	return &file_main_proto_enumTypes[1]
}

func (x SBS1Message_TransmissionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SBS1Message_TransmissionType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SBS1Message_TransmissionType(num)
	return nil
}

// Deprecated: Use SBS1Message_TransmissionType.Descriptor instead.
func (SBS1Message_TransmissionType) EnumDescriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{0, 1}
}

// A parsed SBS1 message.
type SBS1Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageType      *SBS1Message_MessageType      `protobuf:"varint,1,req,name=message_type,json=messageType,enum=proto.SBS1Message_MessageType" json:"message_type,omitempty"`
	TransmissionType *SBS1Message_TransmissionType `protobuf:"varint,2,req,name=transmission_type,json=transmissionType,enum=proto.SBS1Message_TransmissionType" json:"transmission_type,omitempty"`
	SessionId        *string                       `protobuf:"bytes,3,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	AircraftId       *string                       `protobuf:"bytes,4,opt,name=aircraft_id,json=aircraftId" json:"aircraft_id,omitempty"`
	HexIdent         *string                       `protobuf:"bytes,5,opt,name=hex_ident,json=hexIdent" json:"hex_ident,omitempty"`
	FlightId         *string                       `protobuf:"bytes,6,opt,name=flight_id,json=flightId" json:"flight_id,omitempty"`
	GeneratedDate    *string                       `protobuf:"bytes,7,opt,name=generated_date,json=generatedDate" json:"generated_date,omitempty"`
	GeneratedTime    *string                       `protobuf:"bytes,8,opt,name=generated_time,json=generatedTime" json:"generated_time,omitempty"`
	LoggedDate       *string                       `protobuf:"bytes,9,opt,name=logged_date,json=loggedDate" json:"logged_date,omitempty"`
	LoggedTime       *string                       `protobuf:"bytes,10,opt,name=logged_time,json=loggedTime" json:"logged_time,omitempty"`
	Callsign         *string                       `protobuf:"bytes,11,opt,name=callsign" json:"callsign,omitempty"`
	Altitude         *int32                        `protobuf:"varint,12,opt,name=altitude" json:"altitude,omitempty"`
	GroundSpeed      *int32                        `protobuf:"varint,13,opt,name=ground_speed,json=groundSpeed" json:"ground_speed,omitempty"`
	Track            *int32                        `protobuf:"varint,14,opt,name=track" json:"track,omitempty"`
	Lat              *float32                      `protobuf:"fixed32,15,opt,name=lat" json:"lat,omitempty"`
	Lng              *float32                      `protobuf:"fixed32,16,opt,name=lng" json:"lng,omitempty"`
	VerticalRate     *int32                        `protobuf:"varint,17,opt,name=vertical_rate,json=verticalRate" json:"vertical_rate,omitempty"`
	Squawk           *string                       `protobuf:"bytes,18,opt,name=squawk" json:"squawk,omitempty"`
	Alert            *bool                         `protobuf:"varint,19,opt,name=alert" json:"alert,omitempty"`
	Emergency        *bool                         `protobuf:"varint,20,opt,name=emergency" json:"emergency,omitempty"`
	Spi              *bool                         `protobuf:"varint,21,opt,name=spi" json:"spi,omitempty"`
	IsOnGround       *bool                         `protobuf:"varint,22,opt,name=is_on_ground,json=isOnGround" json:"is_on_ground,omitempty"`
}

func (x *SBS1Message) Reset() {
	*x = SBS1Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SBS1Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SBS1Message) ProtoMessage() {}

func (x *SBS1Message) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SBS1Message.ProtoReflect.Descriptor instead.
func (*SBS1Message) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{0}
}

func (x *SBS1Message) GetMessageType() SBS1Message_MessageType {
	if x != nil && x.MessageType != nil {
		return *x.MessageType
	}
	return SBS1Message_UNKNOWN
}

func (x *SBS1Message) GetTransmissionType() SBS1Message_TransmissionType {
	if x != nil && x.TransmissionType != nil {
		return *x.TransmissionType
	}
	return SBS1Message_RESERVED
}

func (x *SBS1Message) GetSessionId() string {
	if x != nil && x.SessionId != nil {
		return *x.SessionId
	}
	return ""
}

func (x *SBS1Message) GetAircraftId() string {
	if x != nil && x.AircraftId != nil {
		return *x.AircraftId
	}
	return ""
}

func (x *SBS1Message) GetHexIdent() string {
	if x != nil && x.HexIdent != nil {
		return *x.HexIdent
	}
	return ""
}

func (x *SBS1Message) GetFlightId() string {
	if x != nil && x.FlightId != nil {
		return *x.FlightId
	}
	return ""
}

func (x *SBS1Message) GetGeneratedDate() string {
	if x != nil && x.GeneratedDate != nil {
		return *x.GeneratedDate
	}
	return ""
}

func (x *SBS1Message) GetGeneratedTime() string {
	if x != nil && x.GeneratedTime != nil {
		return *x.GeneratedTime
	}
	return ""
}

func (x *SBS1Message) GetLoggedDate() string {
	if x != nil && x.LoggedDate != nil {
		return *x.LoggedDate
	}
	return ""
}

func (x *SBS1Message) GetLoggedTime() string {
	if x != nil && x.LoggedTime != nil {
		return *x.LoggedTime
	}
	return ""
}

func (x *SBS1Message) GetCallsign() string {
	if x != nil && x.Callsign != nil {
		return *x.Callsign
	}
	return ""
}

func (x *SBS1Message) GetAltitude() int32 {
	if x != nil && x.Altitude != nil {
		return *x.Altitude
	}
	return 0
}

func (x *SBS1Message) GetGroundSpeed() int32 {
	if x != nil && x.GroundSpeed != nil {
		return *x.GroundSpeed
	}
	return 0
}

func (x *SBS1Message) GetTrack() int32 {
	if x != nil && x.Track != nil {
		return *x.Track
	}
	return 0
}

func (x *SBS1Message) GetLat() float32 {
	if x != nil && x.Lat != nil {
		return *x.Lat
	}
	return 0
}

func (x *SBS1Message) GetLng() float32 {
	if x != nil && x.Lng != nil {
		return *x.Lng
	}
	return 0
}

func (x *SBS1Message) GetVerticalRate() int32 {
	if x != nil && x.VerticalRate != nil {
		return *x.VerticalRate
	}
	return 0
}

func (x *SBS1Message) GetSquawk() string {
	if x != nil && x.Squawk != nil {
		return *x.Squawk
	}
	return ""
}

func (x *SBS1Message) GetAlert() bool {
	if x != nil && x.Alert != nil {
		return *x.Alert
	}
	return false
}

func (x *SBS1Message) GetEmergency() bool {
	if x != nil && x.Emergency != nil {
		return *x.Emergency
	}
	return false
}

func (x *SBS1Message) GetSpi() bool {
	if x != nil && x.Spi != nil {
		return *x.Spi
	}
	return false
}

func (x *SBS1Message) GetIsOnGround() bool {
	if x != nil && x.IsOnGround != nil {
		return *x.IsOnGround
	}
	return false
}

// A summarized "full telemetry" message.
type PositionUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IcaoId       *string  `protobuf:"bytes,1,req,name=icao_id,json=icaoId" json:"icao_id,omitempty"`
	Timestamp    *int64   `protobuf:"varint,2,req,name=timestamp" json:"timestamp,omitempty"`
	Callsign     *string  `protobuf:"bytes,3,opt,name=callsign" json:"callsign,omitempty"`
	Lat          *float32 `protobuf:"fixed32,4,opt,name=lat" json:"lat,omitempty"`
	Lng          *float32 `protobuf:"fixed32,5,opt,name=lng" json:"lng,omitempty"`
	Altitude     *int32   `protobuf:"varint,6,opt,name=altitude" json:"altitude,omitempty"`
	GroundSpeed  *int32   `protobuf:"varint,7,opt,name=ground_speed,json=groundSpeed" json:"ground_speed,omitempty"`
	Track        *int32   `protobuf:"varint,8,opt,name=track" json:"track,omitempty"`
	VerticalRate *int32   `protobuf:"varint,9,opt,name=vertical_rate,json=verticalRate" json:"vertical_rate,omitempty"`
	Squawk       *string  `protobuf:"bytes,10,opt,name=squawk" json:"squawk,omitempty"`
}

func (x *PositionUpdate) Reset() {
	*x = PositionUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionUpdate) ProtoMessage() {}

func (x *PositionUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionUpdate.ProtoReflect.Descriptor instead.
func (*PositionUpdate) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{1}
}

func (x *PositionUpdate) GetIcaoId() string {
	if x != nil && x.IcaoId != nil {
		return *x.IcaoId
	}
	return ""
}

func (x *PositionUpdate) GetTimestamp() int64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *PositionUpdate) GetCallsign() string {
	if x != nil && x.Callsign != nil {
		return *x.Callsign
	}
	return ""
}

func (x *PositionUpdate) GetLat() float32 {
	if x != nil && x.Lat != nil {
		return *x.Lat
	}
	return 0
}

func (x *PositionUpdate) GetLng() float32 {
	if x != nil && x.Lng != nil {
		return *x.Lng
	}
	return 0
}

func (x *PositionUpdate) GetAltitude() int32 {
	if x != nil && x.Altitude != nil {
		return *x.Altitude
	}
	return 0
}

func (x *PositionUpdate) GetGroundSpeed() int32 {
	if x != nil && x.GroundSpeed != nil {
		return *x.GroundSpeed
	}
	return 0
}

func (x *PositionUpdate) GetTrack() int32 {
	if x != nil && x.Track != nil {
		return *x.Track
	}
	return 0
}

func (x *PositionUpdate) GetVerticalRate() int32 {
	if x != nil && x.VerticalRate != nil {
		return *x.VerticalRate
	}
	return 0
}

func (x *PositionUpdate) GetSquawk() string {
	if x != nil && x.Squawk != nil {
		return *x.Squawk
	}
	return ""
}

var File_main_proto protoreflect.FileDescriptor

var file_main_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x08, 0x0a, 0x0b, 0x53, 0x42, 0x53, 0x31, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x42, 0x53, 0x31, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x50, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x0e, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x42, 0x53, 0x31, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x69, 0x72, 0x63, 0x72,
	0x61, 0x66, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x69,
	0x72, 0x63, 0x72, 0x61, 0x66, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x65, 0x78, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x65, 0x78,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x74, 0x69,
	0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x71, 0x75, 0x61, 0x77, 0x6b, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x71,
	0x75, 0x61, 0x77, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6d,
	0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65,
	0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x70, 0x69, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x73, 0x70, 0x69, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x73,
	0x5f, 0x6f, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x69, 0x73, 0x4f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x80, 0x01, 0x0a,
	0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x45, 0x4c,
	0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x4e, 0x45, 0x57, 0x5f, 0x49, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x4e,
	0x45, 0x57, 0x5f, 0x41, 0x49, 0x52, 0x43, 0x52, 0x41, 0x46, 0x54, 0x10, 0x03, 0x12, 0x13, 0x0a,
	0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x49, 0x52, 0x43, 0x52, 0x41, 0x46, 0x54,
	0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x10, 0x05, 0x12, 0x10, 0x0a,
	0x0c, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x22,
	0xc8, 0x01, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x53, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x5f, 0x41,
	0x4e, 0x44, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x10, 0x01, 0x12, 0x12, 0x0a,
	0x0e, 0x45, 0x53, 0x5f, 0x53, 0x55, 0x52, 0x46, 0x41, 0x43, 0x45, 0x5f, 0x50, 0x4f, 0x53, 0x10,
	0x02, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x53, 0x5f, 0x41, 0x49, 0x52, 0x42, 0x4f, 0x52, 0x4e, 0x45,
	0x5f, 0x50, 0x4f, 0x53, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x53, 0x5f, 0x41, 0x49, 0x52,
	0x42, 0x4f, 0x52, 0x4e, 0x45, 0x5f, 0x56, 0x45, 0x4c, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x53,
	0x55, 0x52, 0x56, 0x45, 0x49, 0x4c, 0x4c, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x41, 0x4c, 0x54, 0x10,
	0x05, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x55, 0x52, 0x56, 0x45, 0x49, 0x4c, 0x4c, 0x41, 0x4e, 0x43,
	0x45, 0x5f, 0x49, 0x44, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x49, 0x52, 0x5f, 0x54, 0x4f,
	0x5f, 0x41, 0x49, 0x52, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x4c, 0x4c, 0x5f, 0x43, 0x41,
	0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x10, 0x08, 0x22, 0x99, 0x02, 0x0a, 0x0e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x69, 0x63, 0x61, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06,
	0x69, 0x63, 0x61, 0x6f, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c,
	0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x03, 0x6c, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x71, 0x75, 0x61, 0x77, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x71, 0x75, 0x61, 0x77, 0x6b, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x6b, 0x33, 0x79, 0x2f, 0x67, 0x6f, 0x61, 0x64, 0x73,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_main_proto_rawDescOnce sync.Once
	file_main_proto_rawDescData = file_main_proto_rawDesc
)

func file_main_proto_rawDescGZIP() []byte {
	file_main_proto_rawDescOnce.Do(func() {
		file_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_main_proto_rawDescData)
	})
	return file_main_proto_rawDescData
}

var file_main_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_main_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_main_proto_goTypes = []interface{}{
	(SBS1Message_MessageType)(0),      // 0: proto.SBS1Message.MessageType
	(SBS1Message_TransmissionType)(0), // 1: proto.SBS1Message.TransmissionType
	(*SBS1Message)(nil),               // 2: proto.SBS1Message
	(*PositionUpdate)(nil),            // 3: proto.PositionUpdate
}
var file_main_proto_depIdxs = []int32{
	0, // 0: proto.SBS1Message.message_type:type_name -> proto.SBS1Message.MessageType
	1, // 1: proto.SBS1Message.transmission_type:type_name -> proto.SBS1Message.TransmissionType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_main_proto_init() }
func file_main_proto_init() {
	if File_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SBS1Message); i {
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
		file_main_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionUpdate); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_main_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_main_proto_goTypes,
		DependencyIndexes: file_main_proto_depIdxs,
		EnumInfos:         file_main_proto_enumTypes,
		MessageInfos:      file_main_proto_msgTypes,
	}.Build()
	File_main_proto = out.File
	file_main_proto_rawDesc = nil
	file_main_proto_goTypes = nil
	file_main_proto_depIdxs = nil
}