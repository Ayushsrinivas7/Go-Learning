import csv
import redis
import os


rdb = redis.Redis(host='localhost', port=6379, db=0, decode_responses=True)


BATCH_SIZE = 1000
FOLDER_PATH = "./pincodeFolder"  # Change this to your folder path
REDIS_HASH = "air_mode_pincode_map-1"

# 3. Load All CSVs in Folder
def load_pincodes_from_folder(folder_path):
    batch = {}

    for filename in os.listdir(folder_path):
        if filename.endswith(".csv"):
            filepath = os.path.join(folder_path, filename)
            print(f" {filepath}")
            with open(filepath, newline='') as csvfile:
                reader = csv.reader(csvfile)
                for i, row in enumerate(reader):
                    if i == 0:
                        continue  
                    
                    pincode = row[0].strip()
                    

                    batch[pincode] = "1"

                    if len(batch) >= BATCH_SIZE:
                        try:
                            rdb.hset(REDIS_HASH, mapping=batch)
                          
                        except Exception as e:
                            print(f" Redis error during batch insert: {e}")
                        batch.clear()

    # Insert final batch
    if batch:
        try:
            rdb.hset(REDIS_HASH, mapping=batch)
            print(f"✅ Inserted final batch of {len(batch)} pincodes.")
        except Exception as e:
            print(f" Redis error in final batch: {e}")

    print(" All pincodes loaded from folder into Redis!")



def get_pincode(pincode):
    try:
        value = rdb.hget("air_mode_pincode_map-1", pincode)
        print(f"🔎 Get Pincode {pincode}: {value}")
        return value
    except Exception as e:
        print(f" Redis error: {e}")
        return None

def set_pincode(pincode):
    try:
        rdb.hset("air_mode_pincode_map-1", pincode, "1")
        print(f" Pincode {pincode} added")
    except Exception as e:
        print(f" Redis error: {e}")

def check_pincode(pincode):
    try:
        exists = rdb.hexists("air_mode_pincode_map-1", pincode)
        print(f" Does pincode {pincode} exist? {exists}")
        return exists
    except Exception as e:
        print(f" Redis error: {e}")
        return False


# 4. Run
if __name__ == "__main__":
    load_pincodes_from_folder(FOLDER_PATH)
    get_pincode("110000")
    set_pincode("2101ee")
    check_pincode("2101ee")
    
