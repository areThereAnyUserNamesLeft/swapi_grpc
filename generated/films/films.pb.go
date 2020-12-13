// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: films.proto

package films

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type FilmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PK []int32 `protobuf:"varint,1,rep,packed,name=PK,proto3" json:"PK,omitempty"`
}

func (x *FilmRequest) Reset() {
	*x = FilmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilmRequest) ProtoMessage() {}

func (x *FilmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilmRequest.ProtoReflect.Descriptor instead.
func (*FilmRequest) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{0}
}

func (x *FilmRequest) GetPK() []int32 {
	if x != nil {
		return x.PK
	}
	return nil
}

type FilmsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Films []*Film `protobuf:"bytes,1,rep,name=Films,proto3" json:"Films,omitempty"`
}

func (x *FilmsResponse) Reset() {
	*x = FilmsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilmsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilmsResponse) ProtoMessage() {}

func (x *FilmsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilmsResponse.ProtoReflect.Descriptor instead.
func (*FilmsResponse) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{1}
}

func (x *FilmsResponse) GetFilms() []*Film {
	if x != nil {
		return x.Films
	}
	return nil
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchText string `protobuf:"bytes,1,opt,name=searchText,proto3" json:"searchText,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{2}
}

func (x *SearchRequest) GetSearchText() string {
	if x != nil {
		return x.SearchText
	}
	return ""
}

type Film struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Edited       *timestamp.Timestamp `protobuf:"bytes,1,opt,name=Edited,proto3" json:"Edited,omitempty"`
	Producer     string               `protobuf:"bytes,2,opt,name=Producer,proto3" json:"Producer,omitempty"`
	Title        string               `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	Created      *timestamp.Timestamp `protobuf:"bytes,4,opt,name=Created,proto3" json:"Created,omitempty"`
	EpisodeID    int32                `protobuf:"varint,5,opt,name=EpisodeID,proto3" json:"EpisodeID,omitempty"`
	Director     string               `protobuf:"bytes,6,opt,name=Director,proto3" json:"Director,omitempty"`
	ReleaseDate  string               `protobuf:"bytes,7,opt,name=ReleaseDate,proto3" json:"ReleaseDate,omitempty"`
	OpeningCrawl string               `protobuf:"bytes,8,opt,name=OpeningCrawl,proto3" json:"OpeningCrawl,omitempty"`
	Characters   []int32              `protobuf:"varint,9,rep,packed,name=Characters,proto3" json:"Characters,omitempty"`
	Species      []int32              `protobuf:"varint,10,rep,packed,name=Species,proto3" json:"Species,omitempty"`
	Starships    []int32              `protobuf:"varint,11,rep,packed,name=Starships,proto3" json:"Starships,omitempty"`
	Vehicles     []int32              `protobuf:"varint,12,rep,packed,name=Vehicles,proto3" json:"Vehicles,omitempty"`
	Planets      []int32              `protobuf:"varint,13,rep,packed,name=Planets,proto3" json:"Planets,omitempty"`
	PK           int32                `protobuf:"varint,14,opt,name=PK,proto3" json:"PK,omitempty"`
}

func (x *Film) Reset() {
	*x = Film{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Film) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Film) ProtoMessage() {}

func (x *Film) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Film.ProtoReflect.Descriptor instead.
func (*Film) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{3}
}

func (x *Film) GetEdited() *timestamp.Timestamp {
	if x != nil {
		return x.Edited
	}
	return nil
}

func (x *Film) GetProducer() string {
	if x != nil {
		return x.Producer
	}
	return ""
}

func (x *Film) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Film) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Film) GetEpisodeID() int32 {
	if x != nil {
		return x.EpisodeID
	}
	return 0
}

func (x *Film) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *Film) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

func (x *Film) GetOpeningCrawl() string {
	if x != nil {
		return x.OpeningCrawl
	}
	return ""
}

func (x *Film) GetCharacters() []int32 {
	if x != nil {
		return x.Characters
	}
	return nil
}

func (x *Film) GetSpecies() []int32 {
	if x != nil {
		return x.Species
	}
	return nil
}

func (x *Film) GetStarships() []int32 {
	if x != nil {
		return x.Starships
	}
	return nil
}

func (x *Film) GetVehicles() []int32 {
	if x != nil {
		return x.Vehicles
	}
	return nil
}

func (x *Film) GetPlanets() []int32 {
	if x != nil {
		return x.Planets
	}
	return nil
}

func (x *Film) GetPK() int32 {
	if x != nil {
		return x.PK
	}
	return 0
}

var File_films_proto protoreflect.FileDescriptor

var file_films_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66,
	0x69, 0x6c, 0x6d, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x50, 0x4b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x02, 0x50, 0x4b, 0x22, 0x32, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x46, 0x69, 0x6c,
	0x6d, 0x52, 0x05, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x22, 0x2f, 0x0a, 0x0d, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74, 0x22, 0xc0, 0x03, 0x0a, 0x04, 0x46, 0x69,
	0x6c, 0x6d, 0x12, 0x32, 0x0a, 0x06, 0x45, 0x64, 0x69, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06,
	0x45, 0x64, 0x69, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x4f, 0x70,
	0x65, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x12, 0x1e,
	0x0a, 0x0a, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x0a, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x07, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x53, 0x74, 0x61,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x07, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x50, 0x4b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x50, 0x4b, 0x32, 0xbc, 0x01, 0x0a,
	0x0c, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a,
	0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x12, 0x2e, 0x66,
	0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x12, 0x12, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x69, 0x6c,
	0x6d, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x39, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x12,
	0x14, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_films_proto_rawDescOnce sync.Once
	file_films_proto_rawDescData = file_films_proto_rawDesc
)

func file_films_proto_rawDescGZIP() []byte {
	file_films_proto_rawDescOnce.Do(func() {
		file_films_proto_rawDescData = protoimpl.X.CompressGZIP(file_films_proto_rawDescData)
	})
	return file_films_proto_rawDescData
}

var file_films_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_films_proto_goTypes = []interface{}{
	(*FilmRequest)(nil),         // 0: films.filmRequest
	(*FilmsResponse)(nil),       // 1: films.filmsResponse
	(*SearchRequest)(nil),       // 2: films.searchRequest
	(*Film)(nil),                // 3: films.Film
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_films_proto_depIdxs = []int32{
	3, // 0: films.filmsResponse.Films:type_name -> films.Film
	4, // 1: films.Film.Edited:type_name -> google.protobuf.Timestamp
	4, // 2: films.Film.Created:type_name -> google.protobuf.Timestamp
	0, // 3: films.filmsService.requestFilm:input_type -> films.filmRequest
	0, // 4: films.filmsService.requestFilms:input_type -> films.filmRequest
	2, // 5: films.filmsService.searchFilms:input_type -> films.searchRequest
	1, // 6: films.filmsService.requestFilm:output_type -> films.filmsResponse
	1, // 7: films.filmsService.requestFilms:output_type -> films.filmsResponse
	1, // 8: films.filmsService.searchFilms:output_type -> films.filmsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_films_proto_init() }
func file_films_proto_init() {
	if File_films_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_films_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilmRequest); i {
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
		file_films_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilmsResponse); i {
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
		file_films_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_films_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Film); i {
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
			RawDescriptor: file_films_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_films_proto_goTypes,
		DependencyIndexes: file_films_proto_depIdxs,
		MessageInfos:      file_films_proto_msgTypes,
	}.Build()
	File_films_proto = out.File
	file_films_proto_rawDesc = nil
	file_films_proto_goTypes = nil
	file_films_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FilmsServiceClient is the client API for FilmsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FilmsServiceClient interface {
	RequestFilm(ctx context.Context, in *FilmRequest, opts ...grpc.CallOption) (*FilmsResponse, error)
	RequestFilms(ctx context.Context, in *FilmRequest, opts ...grpc.CallOption) (*FilmsResponse, error)
	SearchFilms(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*FilmsResponse, error)
}

type filmsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFilmsServiceClient(cc grpc.ClientConnInterface) FilmsServiceClient {
	return &filmsServiceClient{cc}
}

func (c *filmsServiceClient) RequestFilm(ctx context.Context, in *FilmRequest, opts ...grpc.CallOption) (*FilmsResponse, error) {
	out := new(FilmsResponse)
	err := c.cc.Invoke(ctx, "/films.filmsService/requestFilm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsServiceClient) RequestFilms(ctx context.Context, in *FilmRequest, opts ...grpc.CallOption) (*FilmsResponse, error) {
	out := new(FilmsResponse)
	err := c.cc.Invoke(ctx, "/films.filmsService/requestFilms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmsServiceClient) SearchFilms(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*FilmsResponse, error) {
	out := new(FilmsResponse)
	err := c.cc.Invoke(ctx, "/films.filmsService/searchFilms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilmsServiceServer is the server API for FilmsService service.
type FilmsServiceServer interface {
	RequestFilm(context.Context, *FilmRequest) (*FilmsResponse, error)
	RequestFilms(context.Context, *FilmRequest) (*FilmsResponse, error)
	SearchFilms(context.Context, *SearchRequest) (*FilmsResponse, error)
}

// UnimplementedFilmsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFilmsServiceServer struct {
}

func (*UnimplementedFilmsServiceServer) RequestFilm(context.Context, *FilmRequest) (*FilmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestFilm not implemented")
}
func (*UnimplementedFilmsServiceServer) RequestFilms(context.Context, *FilmRequest) (*FilmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestFilms not implemented")
}
func (*UnimplementedFilmsServiceServer) SearchFilms(context.Context, *SearchRequest) (*FilmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchFilms not implemented")
}

func RegisterFilmsServiceServer(s *grpc.Server, srv FilmsServiceServer) {
	s.RegisterService(&_FilmsService_serviceDesc, srv)
}

func _FilmsService_RequestFilm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServiceServer).RequestFilm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/films.filmsService/RequestFilm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServiceServer).RequestFilm(ctx, req.(*FilmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmsService_RequestFilms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServiceServer).RequestFilms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/films.filmsService/RequestFilms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServiceServer).RequestFilms(ctx, req.(*FilmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmsService_SearchFilms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmsServiceServer).SearchFilms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/films.filmsService/SearchFilms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmsServiceServer).SearchFilms(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FilmsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "films.filmsService",
	HandlerType: (*FilmsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "requestFilm",
			Handler:    _FilmsService_RequestFilm_Handler,
		},
		{
			MethodName: "requestFilms",
			Handler:    _FilmsService_RequestFilms_Handler,
		},
		{
			MethodName: "searchFilms",
			Handler:    _FilmsService_SearchFilms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "films.proto",
}