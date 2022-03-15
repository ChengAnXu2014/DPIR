// structure of whole file, fields aligned to block size。
type full_file struct{
  fileHead    file_head      // fileHead holds some global metadata.
  root        array_table/map_tabl/list_table/struct_table/string/integer/float/complex
  // field root is root of the whole data structure, it can be a array table, a map table, a list table, a struct table, a string, a integer, a float-point number, or a complex number.

  // followed by subsequent tables(if root is a table).
}
//--------------------------------------------------









//*******************************************************************
type file_head struct{
  blockSizeLog        uint8    // specified blockSize：blockSize=1 << blockSizeLog
  fileHeadSize        uint8    // size of file_head, block unit.
  majorVersion        uint8    // major version.
  minorVersion        uint8    // minor version.
  generalEncodeFlag   byte    // specified UTF and ASCII encoding, can be combined with '|'.
  charSetFlag         byte    // specified char set.
  encodeFormatFlag    byte    // specified encodings of char set specified by charSetFlag, can be combined with '|'.
  rootTypeFlag        byte    // specified type of root
  sonKeyTypeFlag      byte    // (when rootTypeFlag is TYPE_MAP)specified keyType of all entrys in root table.
  sonValueTypeFlag    byte    // (when rootTypeFlag is TYPE_ARRAY or TYPE_MAP)specified valueType of all entrys in root table.
  reserveFileHead     [n]byte  // reserve space, used to align file_head to block size.
}
// generalEncodeFlag:
const(
  ASCII      =  1<<iota
  UTF8
  UTF16LE
  UTF16BE
  UTF32LE
  UTF32BE
)


// charSetFlag:
const(
  CHINESE      = 1+iota        // Chinese char set.
  JAPANESE                    //  Japanese char set。
  // and so on
)


// encodeFormatFlag:
const(
  FLAG1    =  1<<iota  
  FLAG2
  FLAG3
  FLAG4
  FLAG5
  FLAG6
  FLAG7
  FLAG8
)
// For example:
// When charSetFlag is CHINESE, FLAG1 represent GBK, FLAG2 represent GB2312, FLAG3 represent GB18030 ...
// When charSetFlag is JAPANESE, FLAG1 represent Shift_JIS, FLAG2 represent EUC_JP, FLAG3 represent ISO-2022-JP ...





// typeFlag: for rootTypeFlag，keyTypeFlag，valueTypeFlag，sonKeyTypeFlag and sonValueTypeFlag.
const(
  TYPE_ARRAY      =  iota     // array type
  TYPE_MAP                    // map type
  TYPE_LIST                   // list type
  TYPE_STRUCT                 // struct type
  TYPE_STRING                 // string type
  TYPE_INT8                   // and so on...
  TYPE_INT16
  TYPE_INT32
  TYPE_INT64
  TYPE_INT128
  TYPE_INT256
  TYPE_UINT8
  TYPE_UINT16
  TYPE_UINT32
  TYPE_UINT64
  TYPE_UINT128
  TYPE_UINT256
  TYPE_FLOAT16
  TYPE_FLOAT32
  TYPE_FLOAT64
  TYPE_FLOAT128
  TYPE_FLOAT256
  TYPE_FLOAT512
  TYPE_COMPLEX32
  TYPE_COMPLEX64
  TYPE_COMPLEX128
  TYPE_COMPLEX256
  TYPE_COMPLEX512
  TYPE_COMPLEX1024
)
//----------------------------------------------------------------------










//*******************************************************************
// array table：
type array_table  struct{
  entryNum            uint16   // number of entrys in current table.
  sonKeyTypeFlag      byte     // this field only exist when value of all entrys in current table is map_table(specified by father entry's sonValueType), this field specified keyType of son table's all entrys.
  sonValueTypeFlag    byte     // this field only exist when value of all entrys in current table is map_table or array_table, this field specified valueType of son table's all entrys.
  entrys              [entryNum]array_entry    // entrys of current table.
  reserveTable        [n]byte  // reserved space, used to align current table to block size.
}


// array entry:
type  array_entry    struct{
  value  [valueSize]byte
  // according to father entry's sonValueTypeFlag：
  // if it's TYPE_ARRAY, TYPE_MAP, TYPE_LIST or TYPE_STRUCT，then value is uint16 type，specified son table's offset, relative to current table's end, block unit.
  // if it's TYPE_STRING, then value is uint16 type, specified the index of string in *StrFile*.
  // if it's other type(e.g. int64, complex128), then value is the real value of this entry.
}
//---------------------------------------------------------------------------------------









//*******************************************************************************************
// map table:
type  map_table  struct{
  entryNum           uint16            // specified nuber of entrys in current table.
  sonKeyTypeFlag     byte              // this field only exist when value of all entrys in current table is map_table(specified by father entry's sonValueType), this field specified keyType of son table's all entrys.
  sonValueTypeFlag   byte              // this field only exist when value of all entrys in current table is map_table or array_table, this field specified valueType of son table's all entrys.
  entrys             [entryNum]map_entry      // entrys of current table.
  reserveTable       [n]byte            // reserved space, used to align current table to block size.
}


// map entry
type  map_entry  struct{
  key      [keySize]byte          // current entry's key, type and size is specified by father entry's sonKeyTypeFlag; 
  // if father entry's sonKeyTypeFlag is TYPE_STRING, then key is uint16 type, specified the index of string in *StrFile*.

  value    [valueSize]byte        // current entry's value, according to father entry's sonValueTypeFlag:
  // if it's TYPE_ARRAY, TYPE_MAP, TYPE_LIST or TYPE_STRUCT，then value is uint16 type，specified son table's offset, relative to current table's end, block unit.
  // if it's TYPE_STRING, then value is uint16 type, specified the index of string in *StrFile*.
  // if it's other type(e.g. int64, complex128), then value is the real value of this entry.
}
//-------------------------------------------------------------------------------------









//*************************************************************************************
// list table:
type list_table  struct{
  entryNum      uint16      // specified nuber of entrys in current table.
  entrys        [entryNum]list_entry  // entrys in current table.
  exValues      [n]byte    // real values for types beyond size of 32bit(e.g. int64, complex128).
  reserveTable  [n]byte    // reserved space, used to align current table to block size.
}


// list entry:
type  list_entry    struct{
  valueTypeFlag     byte      // specified valueType of current entry.
  sonKeyTypeFlag    byte      // this field only exist when valueTypeFlag is TYPE_MAP, this field specified keyType of son table's all entrys.
  sonValueTypeFlag  byte      // this field only exist when valueTypeFlag is TYPE_MAP or TYPE_ARRAY, this field specified valueType of son table's all entrys.
  value             [n]byte    // the "value" of current entry.
  reserveEntry      [n]byte    // reserved space, used to align current entry to 5 bytes.
  // according to valueTypeFlag：
  // if it's TYPE_ARRAY, TYPE_MAP, TYPE_LIST or TYPE_STRUCT，then value is uint16 type，specified son table's offset, relative to current table's end, block unit.
  // if it's TYPE_STRING, then value is uint16 type, specified the index of string in *StrFile*.
  // if it's number type below size of 32bit(int8-32, uint8-32, float32), then value is the real value of current entry.
  // if it's number type beyond size of 32bit, then value is uint32 type, specified the offset of real value, relative to start of exValues, byte unit.
}
//---------------------------------------------------------------------------------------









//*******************************************************************************************
// struct table:
type  struct_table  struct{
  entryNum      byte    // number of entrys in current table.
  entrys        [entryNum]struct_entry  // entrys of current table.
  exValues      [n]byte  // real values for types beyond size of 32bit(e.g. int64, complex128).
  reserveTable  [n]byte  // reserved space, used to align current table to block size.
}


// struct entry：
type  struct_entry  struct{
  valueTypeFlag     byte      // specified valueType of current entry.
  key               uint16    // index of string in *StrFile*.
  sonKeyTypeFlag    byte      // this field only exist when valueTypeFlag is TYPE_MAP, this field specified keyType of son table's all entrys.
  sonValueTypeFlag  byte      // this field only exist when valueTypeFlag is TYPE_MAP or TYPE_ARRAY, this field specified valueType of son table's all entrys.
  value             [n]byte    // the "value" of current entry.
  reserveEntry      [n]byte    // reserved space, used to align current entry to 7 bytes.
  // according to valueTypeFlag：
  // if it's TYPE_ARRAY, TYPE_MAP, TYPE_LIST or TYPE_STRUCT，then value is uint16 type，specified son table's offset, relative to current table's end, block unit.
  // if it's TYPE_STRING, then value is uint16 type, specified the index of string in *StrFile*.
  // if it's number type below size of 32bit(int8-32, uint8-32, float32), then value is the real value of current entry.
  // if it's number type beyond size of 32bit, then value is uint32 type, specified the offset of real value, relative to start of exValues, byte unit.
}