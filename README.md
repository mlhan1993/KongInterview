# Service API

Tony Mei's implementation of Kong's interview question

# Running the api
#### Bringing up the database and the server
```bash
docker-compose up
```
#### Sending request to the api

Import the `Insomnia_KongInterview_API.json` https://docs.insomnia.rest/insomnia/import-export-data

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