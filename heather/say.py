import requests
text = "Hi I'm Heather"
requests.get('http://localhost:8080/say?text=%s' % text)

