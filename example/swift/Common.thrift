namespace java com.seaWind0112.basic
namespace swift TRCommon

struct TCommon {}

service Common {
  	list<i16> getAges(1: i64 id),
  	list<string> getNames()
}
