## What are APIs
APIs (Application Programming Interfaces) are mediums of communication between two applications. APIs are 
written using REST (Representational State Transfer) architecture, and accessed via the CLI (Command LIne Interface).

There are different types of web services. The four most prominent include:

- Simple Object Access Protocol (SOAP)
- Universal Description, Discovery, and Integration (UDDI)
- Web Services Description Language (WSDL)
- Representational State Transfer (REST)

## SOAP API
SOAP was used in the early 2000s when XML was in vogue as a data format -- now JSON has taken over. Soap requests have 
three components:
- envelope
- header
- body

SOAP APIs are best used in fintech solutions as they are more secure.

## REST API
REST is more simplified and lightweight than SOAP. It is performant, scalable, portable, flexible, etc. 
Characteristics of RESTFul Services include:
- it uses client-server based architecture
- it is stateless; once a request is served, the server removes it from memory, hence it's a stateless operation
- it is cacheable; responses are cached for better throughput
- resources are represented properly
- there is freedom of implementing any web service pattern you want using rest.

REST API is made of three components:
- the REST verb
- header information
- body (optional)

The verbs include:
- GET: fetch resource from the server
- OPTIONS: fetch all REST operations available
- POST: create new resource
- PUT: update/replace resource
- PATCH: modify resource
- DELETE: delete resource

These verbs have status codes. They include:

- Success: 200 - 226
- Error: 400 - 499 for client and 500 - 599 for server
- Redirect: 300- 308

## Cross-Origin Resource Sharing (CORS)
CORS is the most important application of OPTIONS method. When a client requests OPTIONS method from a server, the 
server grants it access. Then the client can call a GET method to receive resource from the client, while the server 
responds with an OK.

## SPAs and REST APIs
Most modern JavaScript frameworks are MVC: Angular, React, Vue... They only implement one design patter: no requesting 
of web pages, only REST API usage. Developers follow a particular method of writing code while working with REST APIs:
- request the HTML templates with the browser in one go
- query the JSON REST API to fill a model (the data object)
- adjust the UI accordingly in JSON
- push back the changes to the server via API call