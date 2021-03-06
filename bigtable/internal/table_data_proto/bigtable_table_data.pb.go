// Code generated by protoc-gen-go.
// source: google.golang.org/cloud/bigtable/internal/table_data_proto/bigtable_table_data.proto
// DO NOT EDIT!

/*
Package google_bigtable_admin_table_v1 is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/cloud/bigtable/internal/table_data_proto/bigtable_table_data.proto

It has these top-level messages:
	Table
	ColumnFamily
	GcRule
*/
package google_bigtable_admin_table_v1

import proto "github.com/golang/protobuf/proto"
import google_protobuf "google.golang.org/cloud/bigtable/internal/duration_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Table_TimestampGranularity int32

const (
	Table_MILLIS Table_TimestampGranularity = 0
)

var Table_TimestampGranularity_name = map[int32]string{
	0: "MILLIS",
}
var Table_TimestampGranularity_value = map[string]int32{
	"MILLIS": 0,
}

func (x Table_TimestampGranularity) String() string {
	return proto.EnumName(Table_TimestampGranularity_name, int32(x))
}

// A collection of user data indexed by row, column, and timestamp.
// Each table is served using the resources of its parent cluster.
type Table struct {
	// A unique identifier of the form
	// <cluster_name>/tables/[_a-zA-Z0-9][-_.a-zA-Z0-9]*
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The column families configured for this table, mapped by column family id.
	ColumnFamilies map[string]*ColumnFamily `protobuf:"bytes,3,rep,name=column_families" json:"column_families,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The granularity (e.g. MILLIS, MICROS) at which timestamps are stored in
	// this table. Timestamps not matching the granularity will be rejected.
	// Cannot be changed once the table is created.
	Granularity Table_TimestampGranularity `protobuf:"varint,4,opt,name=granularity,enum=google.bigtable.admin.table.v1.Table_TimestampGranularity" json:"granularity,omitempty"`
}

func (m *Table) Reset()         { *m = Table{} }
func (m *Table) String() string { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()    {}

func (m *Table) GetColumnFamilies() map[string]*ColumnFamily {
	if m != nil {
		return m.ColumnFamilies
	}
	return nil
}

// A set of columns within a table which share a common configuration.
type ColumnFamily struct {
	// A unique identifier of the form <table_name>/families/[-_.a-zA-Z0-9]+
	// The last segment is the same as the "name" field in
	// google.bigtable.v1.Family.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Garbage collection expression specified by the following grammar:
	//   GC = EXPR
	//      | "" ;
	//   EXPR = EXPR, "||", EXPR              (* lowest precedence *)
	//        | EXPR, "&&", EXPR
	//        | "(", EXPR, ")"                (* highest precedence *)
	//        | PROP ;
	//   PROP = "version() >", NUM32
	//        | "age() >", NUM64, [ UNIT ] ;
	//   NUM32 = non-zero-digit { digit } ;    (* # NUM32 <= 2^32 - 1 *)
	//   NUM64 = non-zero-digit { digit } ;    (* # NUM64 <= 2^63 - 1 *)
	//   UNIT =  "d" | "h" | "m"  (* d=days, h=hours, m=minutes, else micros *)
	// GC expressions can be up to 500 characters in length
	//
	// The different types of PROP are defined as follows:
	//   version() - cell index, counting from most recent and starting at 1
	//   age() - age of the cell (current time minus cell timestamp)
	//
	// Example: "version() > 3 || (age() > 3d && version() > 1)"
	//   drop cells beyond the most recent three, and drop cells older than three
	//   days unless they're the most recent cell in the row/column
	//
	// Garbage collection executes opportunistically in the background, and so
	// it's possible for reads to return a cell even if it matches the active GC
	// expression for its family.
	GcExpression string `protobuf:"bytes,2,opt,name=gc_expression" json:"gc_expression,omitempty"`
	// Garbage collection rule specified as a protobuf.
	// Supercedes `gc_expression`.
	// Garbage collection executes opportunistically in the background, and so
	// it's possible for reads to return a cell even if it matches the active GC
	// expression for its family.
	GcRule *GcRule `protobuf:"bytes,3,opt,name=gc_rule" json:"gc_rule,omitempty"`
}

func (m *ColumnFamily) Reset()         { *m = ColumnFamily{} }
func (m *ColumnFamily) String() string { return proto.CompactTextString(m) }
func (*ColumnFamily) ProtoMessage()    {}

func (m *ColumnFamily) GetGcRule() *GcRule {
	if m != nil {
		return m.GcRule
	}
	return nil
}

// Rule for determining which cells to delete during garbage collection.
type GcRule struct {
	// Delete all cells in a column except the most recent N.
	MaxNumVersions int32 `protobuf:"varint,1,opt,name=max_num_versions" json:"max_num_versions,omitempty"`
	// Delete cells in a column older than the given age.
	MaxAge *google_protobuf.Duration `protobuf:"bytes,2,opt,name=max_age" json:"max_age,omitempty"`
	// Delete cells that would be deleted by every nested rule.
	Intersection *GcRule_Intersection `protobuf:"bytes,3,opt,name=intersection" json:"intersection,omitempty"`
	// Delete cells that would be deleted by any nested rule.
	Union *GcRule_Union `protobuf:"bytes,4,opt,name=union" json:"union,omitempty"`
}

func (m *GcRule) Reset()         { *m = GcRule{} }
func (m *GcRule) String() string { return proto.CompactTextString(m) }
func (*GcRule) ProtoMessage()    {}

func (m *GcRule) GetMaxAge() *google_protobuf.Duration {
	if m != nil {
		return m.MaxAge
	}
	return nil
}

func (m *GcRule) GetIntersection() *GcRule_Intersection {
	if m != nil {
		return m.Intersection
	}
	return nil
}

func (m *GcRule) GetUnion() *GcRule_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

// A GcRule which deletes cells matching all of the given rules.
type GcRule_Intersection struct {
	// Only delete cells which would be deleted by every element of `rules`.
	Rules []*GcRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *GcRule_Intersection) Reset()         { *m = GcRule_Intersection{} }
func (m *GcRule_Intersection) String() string { return proto.CompactTextString(m) }
func (*GcRule_Intersection) ProtoMessage()    {}

func (m *GcRule_Intersection) GetRules() []*GcRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// A GcRule which deletes cells matching any of the given rules.
type GcRule_Union struct {
	// Delete cells which would be deleted by any element of `rules`.
	Rules []*GcRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *GcRule_Union) Reset()         { *m = GcRule_Union{} }
func (m *GcRule_Union) String() string { return proto.CompactTextString(m) }
func (*GcRule_Union) ProtoMessage()    {}

func (m *GcRule_Union) GetRules() []*GcRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.bigtable.admin.table.v1.Table_TimestampGranularity", Table_TimestampGranularity_name, Table_TimestampGranularity_value)
}
