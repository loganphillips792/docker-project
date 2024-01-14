# https://www.reddit.com/r/htmx/comments/16onteo/i_implemented_the_htmx_examples_in_flask/

import time
import logging
from flask import Flask

app = Flask(__name__)
logging.basicConfig(level=logging.INFO)
count = 0

@app.route('/')
def hello_world():
    global count
    count += 1
    logging.info(f"hello world, iteration {count}")
    return {'hello': 'world world'}

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8080)

# main()
