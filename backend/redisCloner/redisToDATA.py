import redis
import json

# Redis connection details
REDIS_ADDR = "172.17.0.1"
REDIS_PORT = 6379
REDIS_PASSWORD = "kingKong"

# Connect to Redis
redis_client = redis.StrictRedis(
    host=REDIS_ADDR,
    port=REDIS_PORT,
    password=REDIS_PASSWORD,
    decode_responses=True
)

def export_redis_data(file_path="redis_backup.json"):
    data = {}

    for key in redis_client.keys("*"):
        key_type = redis_client.type(key)

        if key_type == "string":
            data[key] = redis_client.get(key)
        elif key_type == "list":
            data[key] = redis_client.lrange(key, 0, -1)
        elif key_type == "set":
            data[key] = list(redis_client.smembers(key))
        elif key_type == "zset":
            data[key] = redis_client.zrange(key, 0, -1, withscores=True)
        elif key_type == "hash":
            data[key] = redis_client.hgetall(key)

    with open(file_path, "w") as f:
        json.dump(data, f, indent=4)

    print(f"Redis data exported to {file_path}")

# Run the export function
export_redis_data()
