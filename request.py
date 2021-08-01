import requests
import time

from requests.exceptions import HTTPError

for x in range(100):
    with requests.Session() as session:
        for url in ['http://gowiki.abcd.com/', 'http://gowiki.abcd.com/oidc']:
            try:
                response = session.get(url)

                # If the response was successful, no Exception will be raised
                response.raise_for_status()
                print(response.json())
            except HTTPError as http_err:
                print(f'HTTP error occurred: {http_err}')  # Python 3.6
            except Exception as err:
                print(f'Other error occurred: {err}')  # Python 3.6
            else:
                print('Success!')
            
            if url == 'http://gowiki.abcd.com/':
                time.sleep(15)