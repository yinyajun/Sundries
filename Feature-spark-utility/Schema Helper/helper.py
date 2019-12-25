from pyspark.sql.types import *
from pyspark import SparkContext


def _ignore_brackets_split(s, separator):
    from pyspark.sql.types import _BRACKETS
    parts = []
    buf = ""
    level = 0
    for c in s:
        if c in _BRACKETS.keys():
            level += 1
            buf += c
        elif c in _BRACKETS.values():
            if level == 0:
                raise ValueError("Brackets are not correctly paired: %s" % s)
            level -= 1
            buf += c
        elif c == separator and level > 0:
            buf += c
        elif c == separator:
            parts.append(buf)
            buf = ""
        else:
            buf += c
    if len(buf) == 0:
        raise ValueError("The %s cannot be the last char: %s" % (separator, s))
    parts.append(buf)
    return parts


def _parse_struct_fields_string(s):
    parts = _ignore_brackets_split(s, ",")
    fields = []
    for part in parts:
        name_and_type = _ignore_brackets_split(part, ":")
        if len(name_and_type) != 2:
            raise ValueError("The strcut field string format is: 'field_name:field_type', " +
                             "but got: %s" % part)
        field_name = name_and_type[0].strip()
        field_type = _parse_datatype_string(name_and_type[1])
        fields.append(StructField(field_name, field_type))
    return StructType(fields)


def _parse_datatype_string(s):
    s = s.strip()
    if s.startswith("array<"):
        if s[-1] != ">":
            raise ValueError("'>' should be the last char, but got: %s" % s)
        return ArrayType(_parse_datatype_string(s[6:-1]))
    elif s.startswith("map<"):
        if s[-1] != ">":
            raise ValueError("'>' should be the last char, but got: %s" % s)
        parts = _ignore_brackets_split(s[4:-1], ",")
        if len(parts) != 2:
            raise ValueError("The map type string format is: 'map<key_type,value_type>', " +
                             "but got: %s" % s)
        kt = _parse_datatype_string(parts[0])
        vt = _parse_datatype_string(parts[1])
        return MapType(kt, vt)
    elif s.startswith("struct<"):
        if s[-1] != ">":
            raise ValueError("'>' should be the last char, but got: %s" % s)
        return _parse_struct_fields_string(s[7:-1])
    elif ":" in s:
        return _parse_struct_fields_string(s)
    else:
        return _parse_basic_datatype_string(s)


def _parse_basic_datatype_string(s):
    from pyspark.sql.types import _all_atomic_types, _FIXED_DECIMAL
    if s in _all_atomic_types.keys():
        return _all_atomic_types[s]()
    elif s == "int":
        return IntegerType()
    elif _FIXED_DECIMAL.match(s):
        m = _FIXED_DECIMAL.match(s)
        return DecimalType(int(m.group(1)), int(m.group(2)))
    else:
        raise ValueError("Could not parse datatype: %s" % s)


class StructCollect(StructType):
    def __init__(self, original_df):
        super(StructCollect, self).__init__()
        from copy import deepcopy
        self.schema = deepcopy(original_df.schema)  # StructType
        self.fields = self.schema.fields  # StructField的集合
        self.names = self.schema.names

    def remove(self, name):
        assert name in self.names
        idx = self.names.index(name)
        self.names.pop(idx)
        self.fields.pop(idx)
        return self

    @staticmethod
    def _parse_schema(string):
        try:
            from pyspark.sql.types import _parse_datatype_string
            return _parse_datatype_string(string)
        except ImportError:
            return globals()["_parse_datatype_string"](string)

    def append(self, string):
        schema = self._parse_schema(string)
        self._merge(schema)
        return self

    def _merge(self, schema):
        # 注意没有检查name是否重复
        from copy import deepcopy
        schema = deepcopy(schema)
        fields = schema.fields
        for field in fields:
            self.add(field)

    def merge(self, df):
        self._merge(df.schema)
        return self

    def get(self, name):
        assert name in self.names
        idx = self.names.index(name)
        return self.fields[idx]


import unittest


class TestStructCollect(unittest.TestCase):
    def test_init(self):
        sc = SparkContext._active_spark_context
        cols = ['user', 'scores']
        df = sc.parallelize([['123', [5, 4]], ['fs', []], ['fsd', [2, 3, 4]]]).toDF(cols)
        schema = StructCollect(df)
        for i in range(len(cols)):
            self.assertEqual(cols[i], schema.names[i])
        return schema

    def test_get(self):
        schema = self.test_init()
        self.assertTrue(isinstance(schema.get("user"), StructField))
        self.assertEqual(schema.get('user').name, "user")

    def test_merge(self):
        sc = SparkContext._active_spark_context
        df = sc.parallelize([[123]]).toDF(["height"])
        schema = self.test_init()
        schema.merge(df)
        cols = ['user', 'scores', 'height']
        for i in range(len(cols)):
            self.assertEqual(cols[i], schema.names[i])

    def test_append(self):
        string = "items:array<int>, education:string"
        schema = self.test_init()
        schema.append(string)
        self.assertTrue("items" in schema.names)
        self.assertTrue("education" in schema.names)


if __name__ == '__main__':
    unittest.main()
