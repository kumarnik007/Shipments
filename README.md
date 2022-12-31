# Shipments

> Description:
1. The web server is designed to handle HTTP GET /shipments and HTTP POST /shipment requests.
2. HTTP POST /shipment request body is expected to be a JSON object like below :
  ```json
  {
      "sender_country": "SE",
       "receiver_country": "IN",
       "weight": 100
  }
  ```
3. Go programming language has been used to develop the web server.
4. The server is designed to listen on port 8081.

> Sample HTTP Requests:
1. GET /shipments
  - GET /shipments
  - Host: localhost:8081
  - Sample URL : http://localhost:8081/shipments 
  - Sample Response :
  ```json
  {
      "shipments": [
        {
            "sender_country": "SE",
            "receiver_country": "IN",
            "weight": 100,
            "price": 5000,
            "currency": "SEK"
        }
      ]
  }
  ```
2. POST /shipment
  - POST / shipment
  - Host: localhost:8081
  - Content-Type: application/json
  - Sample URL : http://localhost:8081/shipment 
  - Sample Request :
  ```json
  {
      "sender_country": "SE",
      "receiver_country": "IN",
      "weight": 100
  }
  ```
  - Sample Response :
  ```json
  {
      "sender_country": "SE",
      "receiver_country": "IN",
      "weight": 100,
      "price": 5000,
      "currency": "SEK"
  }
  ```

> Assumptions/Limitations:
1. For POST /shipment API, the server returns a 400 BAD REQUEST with appropriate error message in the following scenarios:
  - Content-Type header of the request is anything other than application/json.
  - Request has no body.
  - Request body is not in the correct form.
  - Sender or receiver country value is not a valid code as per the wiki.
  - Weight value is less than 0 or greater than 1000.
2 If a shipment has same sender and receiver country but its outside EU, then it is considered as international shipment.
3. Used the package "github.com/biter777/countries" for country codes.
4. EU members are taken as mentioned in the wiki and stored in a JSON file.
5. Currently there is support for only 2 APIs :
  - POST /shipment - Send the shipment details (sender_country, receiver_country, weight of a shipment), the API returns all these details plus the calculated price of the package.
  - GET /shipments to get a list of all shipments sent (which the user sent a POST /shipment request)
6. There is no support for filtering or storing shipments as per the user (since the user is not authenticated), so all posted shipments show up when a GET /shipments API endpoint is invoked.
7. The details about the shipments are stored in a JSON file named storage.json in the root directory of the project and so any shipments in this file (written by POST /shipment) are listed by GET /shipments API.
8. Intended pricing information is stored in a JSON file and read from there.

