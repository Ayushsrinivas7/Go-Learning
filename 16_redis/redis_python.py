import csv
import redis


rdb = redis.Redis(host='localhost', port=6379, db=0, decode_responses=True)


BATCH_SIZE = 1000  

def load_pincodes_from_csv(filename):
    try:
        with open(filename, newline='') as csvfile:
            reader = csv.reader(csvfile)
            batch = {}
            for i, row in enumerate(reader):
                if i == 0:
                    continue  

                pincode = row[0]
                batch[pincode] = "1"

                if len(batch) >= BATCH_SIZE:
                    try:
                        rdb.hset("air_mode_pincode_map", mapping=batch)
                       
                    except Exception as e:
                        print(" Redis error during batch insert: {e}")
                    batch = {}

           
            if batch:
                try:
                    rdb.hset("air_mode_pincode_map", mapping=batch)
                    print(f" Inserted final batch of {len(batch)} pincodes.")
                except Exception as e:
                    print(f" Redis error in final batch: {e}")

    except FileNotFoundError:
        print(" pincode.csv not found")
        exit()

    print(" All pincodes loaded into Redis Hash!")



def get_pincode(pincode):
    try:
        value = rdb.hget("pincode_data", pincode)
        print(f"🔎 Get Pincode {pincode}: {value}")
        return value
    except Exception as e:
        print(f" Redis error: {e}")
        return None

def set_pincode(pincode):
    try:
        rdb.hset("pincode_data", pincode, "1")
        print(f" Pincode {pincode} added")
    except Exception as e:
        print(f" Redis error: {e}")

def check_pincode(pincode):
    try:
        exists = rdb.hexists("pincode_data", pincode)
        print(f" Does pincode {pincode} exist? {exists}")
        return exists
    except Exception as e:
        print(f" Redis error: {e}")
        return False


if __name__ == "__main__":
    load_pincodes_from_csv("pincode.csv")  
    get_pincode("110000")
    set_pincode("2101ee")
    check_pincode("2101ee")
