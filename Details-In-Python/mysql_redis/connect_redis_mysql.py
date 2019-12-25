import pymysql
import redis

REDIS_CONF = {'host': '0.0.0.0', 'port': 566, 'password': '123', 'timeout': 3}
MYSQL_CONF = {
    'host': '0.0.0.0',
    'user': 'me',
    'password': '321',
    'db': 'user_score',
    'port': 566}


def save_to_redis():
    r = redis.Redis(**REDIS_CONF, decode_responses=True)
    pip = r.pipeline()
    main_key_user = "user"
    main_key_score = "score"
    # init new keys to redis
    for item in user_dict:
        new_key = item[0]
        value = item[1]
        pip.hset(main_key_user, new_key, value)
    for item in score_dict:
        new_key = item[0]
        value = item[1]
        pip.hset(main_key_score, new_key, value)
    output = pip.execute()
    print(output)


def save_user_score(model):
    user_f = [('0', '100'), ('1', '99')]
    sql = "insert into table (user, score) values(%s, %s) on duplicate key update feature= %s"
    params = [(us[0], us[1], us[1]) for uf in user_s]
    db = pymysql.connect(**MYSQL_CONF)
    try:
        cursor = db.cursor()
        cursor.executemany(sql, params)
        db.commit()
        cursor.close()
    finally:
        db.close()
