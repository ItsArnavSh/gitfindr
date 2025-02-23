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

def import_redis_data(file_path="redis_backup.json"):
    with open(file_path, "r") as f:
        data = json.load(f)

    for key, value in data.items():
        if isinstance(value, str):  # String
            redis_client.set(key, value)
        elif isinstance(value, list) and all(isinstance(i, str) for i in value):  # List
            redis_client.delete(key)  # Ensure list is empty before pushing
            redis_client.rpush(key, *value)
        elif isinstance(value, list) and all(isinstance(i, list) and len(i) == 2 for i in value):  # Sorted Set
            redis_client.zadd(key, dict(value))
        elif isinstance(value, dict):  # Hash
            redis_client.hset(key, mapping=value)
        elif isinstance(value, list):  # Set
            redis_client.sadd(key, *value)

    print("Redis data imported successfully!")

# Run the import function
import_redis_data()
