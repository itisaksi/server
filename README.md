A simple Go web server that accepts POST requests with JSON body in the format: 
{"name": "Maksim", "age": 24}. The server saves the received data in memory and returns the list 
of all saved users in response to a GET request to the endpoint "/users".
