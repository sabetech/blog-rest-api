# Blog - RESTful API

## Background Information
### What is REST?

REST, coined by Roy Fielding in 2000, is an acronym for Representational State Transfer. You can read more from the source in Roy Fielding's famous [Dissertation](https://www.ics.uci.edu/~fielding/pubs/dissertation/fielding_dissertation.pdf)

### REST is an Architecture

When a Web Service or API conforms to the REST architecture it is said to be a RESTful API.

So what makes a Web API RESTful?  There are six guidelines or principles:

- #### Uniform Interface.
    Resources should be uniquely identified through a single URL and be manipulated using the methods of the network protocol: GET, PUT, POST, DELETE
- #### Client - Server
    Should be a clear separation between the client and the server.  The server should not know or care what the client is and how it uses the API.
- #### Stateless
    Each request to the server should contain all information to complete the request.  The client should handle any session state
- #### Cacheable
    Server response should say if the resource is cacheable: if so, client could reuse the response
- #### Layered System
    Responses and calls go through different systems.  The server nor client can assume they are communicating end point or an intermediary.
- #### Code on Demand (optional)
    REST APIs generally send static data, such as JSON, but in some cases the response can be executable code.

### What is a resource?

Roy Fielding put it best:
> The  key  abstraction  of  information  in  REST  is  a  resource.  Any information  that  can  be named can be a resource: a document or image, a temporal service (e.g. “today’s weather in Los Angeles”), a collection of other resources, a non-virtual object (e.g. a person), and so  on.

For instance, using the Twitter api you can retreive a tweet or using OpenWeather's API you can get weather data.  Both of these are examples of what a resource is.

### Endpoints

The URL that has the resource is known as the endpoint.  
Example: http://web.site/api/book/123, where this resource would be a book with the ID of 123 and could return JSON, XML, or even HTML with more data about this resource.

At these endpoints we can manipulate resources using common HTTP methods:
- POST
- GET
- PUT/PATCH
- DELETE

These methods are used in what is commonly known as CRUD: Create, Read, Update, and Delete.
- POST :: **Create**  Using the HTTP POST Method we can create a new resource.
- GET :: **Read** We can retrieve and read the resource with a GET request
- PUT/PATCH :: **Update**  Using PUT/PATCH, update an existing resource.
- DELETE :: **Delete**  Delete a resource.