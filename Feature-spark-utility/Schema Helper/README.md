# DataFrame Schema Helper
[中文博客Chinese Blog](https://yinyajun.github.io/CS-Data/spark_02/)

You need to provide schema if inferring fails. To this,  you should do like this to create schema:

```python
from pyspark.sql.types import StructType, StructField, IntegerType, StringType, ArrayType

schema = StructType([
    StructField("name", StringType(), False), 
    StructField("age", IntegerType(), False), 
    StructField("height", IntegerType(), False)，
    StructField("scores", ArrayType(IntergerType()))
])

df = rdd.toDF(schema)
```

If this dataframe has more than 20 or 30 columns, you will be mad. You can use this helper to  get schema conveniently.

```python
from pyspark.sql.types import *
from helper import StructCollect

cols = ['user', 'scores']
df = sc.parallelize([['123', [5, 4]], ['fs', []], ['fsd', [2, 3, 4]]]).toDF(cols)

schema = StructCollect(df)
print(schema.names)

# get a field
print(schema.get("user"))

# merge two dataframe schema
new_df = sc.parallelize([[123]]).toDF(["height"])
schema = schema.merge(new_df)
print(schema)

# most convient way
schema = schema.append("items:array<int>, education:string")
print(schema)
```

