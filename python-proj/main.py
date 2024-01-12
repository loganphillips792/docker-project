import time
import logging

def main():
    logging.basicConfig(level=logging.INFO)
    count = 0
    while True:
        time.sleep(5)
        count += 1
        logging.info(f"hello world, iteration {count}")

main()
