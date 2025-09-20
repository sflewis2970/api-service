
<!--

    Licensed to the Apache Software Foundation (ASF) under one
    or more contributor license agreements.  See the NOTICE file
    distributed with this work for additional information
    regarding copyright ownership.  The ASF licenses this file
    to you under the Apache License, Version 2.0 (the
    "License"); you may not use this file except in compliance
    with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing,
    software distributed under the License is distributed on an
    "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
    KIND, either express or implied.  See the License for the
    specific language governing permissions and limitations
    under the License.

-->

# api-service

API-Serice is a tool for processing different APIs.
Currently there are only two APIs available for processing
Get, "/api/v1/trivia/getquestion" -- request trivia question
get request returns "questionid", "category", "question", and "choices" (one of the choices isd the correct answer)
Post, "/api/v1/trivia/submitanswer" -- post trivia question answer
send "questionid" and "response" in the post request


### Requirements
  * Git
  * Docker 4.19.0 (is the version being used)
  * VCisual Studio Code (or any editor that supports Golang add-ons or extensions)
  * Golang 1.25
  * Redia Server 7.0 or higher
  
### Download API-Service

Download the source with any tool of your choice (git is recommended).

### Building API-Service

Build service image
```
$ docker build -t api-service .
```
Make sure the current directory includes the Dockerfile file. The Dockerfile includes all the instructions for building the api-service image.

### Running API-Service
The compose.yml already has all the settings to pull from Docker Hub. If changes are made then re-build the api-service image then run the folloeing command:
```
$ docker-compose up
```

Run api-service using the docker images:
```
$ docker-compose up
```
Make sure the current directory includes the compose.yml file. The compose.yml includes all the images and parameters needed to pull and start images.
Once the images start the following hould display:
Redis Startup output:
<img width="1179" height="427" alt="Redis Startup" src="https://github.com/user-attachments/assets/79616081-956b-43f1-8d50-83933c6cda6b" />

trivia Service Strtup:
<img width="984" height="453" alt="Ttrivia-Service startup" src="https://github.com/user-attachments/assets/7e9f1ad2-3e53-4bc8-8bf7-5d9d17e081c0" />

Get Request:
the get request will return "questionid" which is needed for th post request. The get request alsop returns "category", "question", and "choices". One of the choices is th correct answer.
<img width="918" height="950" alt="Ttrivia-Service request   response" src="https://github.com/user-attachments/assets/e1ea69b7-9352-42fe-be5c-6536f1f271ca" />

Post Request & Response:
When sending the post request include the "questionid" included in the get request response. Once the answer is processed the question is deleted from the DB (whether the client provides right answer or not).
<img width="981" height="400" alt="Ttrivia-Service post response   response logs" src="https://github.com/user-attachments/assets/53c77120-a6e5-4dd2-8c7c-1e9ee563286e" />

Post Request (wrong answer) & Response:
<img width="970" height="495" alt="Ttrivia-Service post response   response bad answer logs" src="https://github.com/user-attachments/assets/a2c26ff3-6cdc-4705-a8e8-aaca58c7bfc4" />

Post Request (not found) & RFesponse:
<img width="914" height="900" alt="Ttrivia-Service post response   response -- not found" src="https://github.com/user-attachments/assets/179d9461-3549-4297-b451-3ca7b9d37056" />
<img width="981" height="282" alt="Ttrivia-Service post response   response -- not found logs" src="https://github.com/user-attachments/assets/80279a8a-e3b9-4ab9-848a-be19d8119ac2" />

### Reporting Bugs

Bugs should be reported to https://github.com/sflewis2970/api-service/issues

### Full History

