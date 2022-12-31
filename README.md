# Shipments

> Description:
1. The web server is designed to handle HTTP GET /shipments and HTTP POST /shipment requests.
2. HTTP POST /shipment request body is expected to be a JSON object like below :
  {
      "sender_country": "SE",
       "receiver_country": "IN",
       "weight": 100
  } 
3. Go programming language has been used to develop the web server.
4. The server is designed to listen on port 8081.
