namespace java com.seaWind0112.basic
namespace swift TRShared

include "Common.thrift"

struct TBasic {}

service Shared {
	TBasic getBasic(1: i32 key, 2: i64 id),
  	list<TBasic> getBasics(1: i32 key, 2: i64 id, 3: list<i64> int64s),
  	list<i64> getInt64s(1: i64 id, 2: list<i64> int64s),
  	list<Common.TCommon> getCommons(1: i32 key, 2: i64 id)
}
