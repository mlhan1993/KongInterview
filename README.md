# Service API

Tony Mei's implementation of Kong's interview question

# Running the api
#### Bringing up the database and the server
```bash
docker-compose up
```
#### Sending request to the api

Import the `Insomnia_KongInterview_API.json` https://docs.insomnia.rest/insomnia/import-export-data

# Design Considerations

### Database

- Using mysql as a db for the following reasons
  - join operation is needed
  - low write and high read
  - Data volume likely low and less need for scaling 
- Schema design consideration
  - service - version is a 1 to many relationship, therefore version should contain service id
  - version's serviceID reference service's id to avoid version without services
  - unique key on version tag and service id to avoid duplication. 
- Query consideration
  - While it's true that using `LIKE` could reduce performance, the number of services are likely not extremely high, therefore it should be ok for the first iteration
  - Filtering first then join will increase performance

### API
- using gorilla as a library for better middleware syntax
- Introduced some middlewares
  - adding request id to request for better tracking / logging
  - adding logging of the request itself for easier debugging in the future in case something is wrong. 
- Separate end point for versions and services as standard REST api practice for different resources
- Used a logging library for json fmt logging, which allow easier parsing and query if a logging pipeline exists
- Versions for handlers to support future changes
- Interfaces between handlers and db layer to allow plug and play in case implementation changes
- Using environment variables for db connection properties for better security and deployment configurability
  - Comparing to a config file, i think using an environment variable is slightly more secure, and given 
- Introduce special error types for convent api response 

### Local testing and containerization
- single docker-compose file for easy local testing
- separate Dockerfile so db and app can be build separately
- Insomnia json file for easy local testing. 


# Assumptions
- I'm able to talk with the stakeholders to make small changes regarding api route names and certain small updates on the request body/response body
- The number of services is not huge
- DB read is much more frequent than write
- It's ok to pre-populate some data for dev purposes. 
- Versions have "version tags" and they are string formatted.
- 

# API documentation

### /v1/service
This endpoint retrieves all available services based on filter and pagination provided.

##### Method: POST

##### Request body format:
```json
{
  "sortOrder": string,
  "filter": string,
  "numPerPage": uint,
  "pageNumber": uint
}

```
where:

- sortOrder (optional): specifies the order in which the results should be sorted.
- filter (optional): specifies a string to filter the results by.
- numPerPage (optional): specifies the number of results to return per page. Must be used together with pageNumber
- pageNumber (optional): specifies the page number of the results to return. Must be used together with numPerPage

##### Response body format:

```json
{
  "total": uint,
  "services": [
    {
    "name": string,
    "id": string,
    "description": string,
    "numVersions": int
    },
  ...
  ]
}
```
where:

- total: the total number of services returned.
- services: an array of objects, each representing a service, with the following fields:
  - name: the name of the service.
  - id: the ID of the service.
  - description: a short description of the service.
  - numVersions: the number of available versions for the service.

### /v1/version
  This endpoint retrieves version information about a specific service.

#### Method: POST
#### Request body format:

```json
{
  "serviceID": uint,
  "sortOrder": string,
  "numPerPage": uint,
  "pageNumber": uint
}
```

where:

- serviceID: the ID of the service to retrieve versions for.
- sortOrder (optional): specifies the order in which the results should be sorted. Either asc or desc
- numPerPage (optional): specifies the number of results to return per page.
- pageNumber (optional): specifies the page number of the results to return.

#### Response body format:

```json 
{
  "total": uint,
  "versions": [
    {
      "id": string,
      "tag": string,
      "serviceID": uint,
      "dateCreated": string
    },
    ...
  ]
}
```
where:

- total: the total number of versions available
- versions: an array of objects, each representing a service version, with the following fields:
  - id: the ID of the version
  - tag: version tag.
  - serviceID: the ID of the service the version belongs to.
  - dateCreated: the date and time the version was created.