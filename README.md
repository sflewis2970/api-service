
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

API-Serice is a tool for processing dfferent APIs.
Currently there are only two APIs available for processing
Get, "/api/v1/trivia/getquestion" -- request trivia question
Post, "/api/v1/trivia/submitanswer" -- post trivia question answer

### Build status

### Requirements
  * Git
  * Docker
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
<img width="918" height="950" alt="Ttrivia-Service request   response" src="https://github.com/user-attachments/assets/e1ea69b7-9352-42fe-be5c-6536f1f271ca" />

Post Request & Response:
<img width="981" height="400" alt="Ttrivia-Service post response   response logs" src="https://github.com/user-attachments/assets/53c77120-a6e5-4dd2-8c7c-1e9ee563286e" />

Post Request (wrong answer) & RFesponse:
<img width="970" height="495" alt="Ttrivia-Service post response   response bad answer logs" src="https://github.com/user-attachments/assets/a2c26ff3-6cdc-4705-a8e8-aaca58c7bfc4" />

Post Request (not found) & RFesponse:
<img width="914" height="900" alt="Ttrivia-Service post response   response -- not found" src="https://github.com/user-attachments/assets/179d9461-3549-4297-b451-3ca7b9d37056" />
<img width="981" height="282" alt="Ttrivia-Service post response   response -- not found logs" src="https://github.com/user-attachments/assets/80279a8a-e3b9-4ab9-848a-be19d8119ac2" />



### Download

Developer builds can be downloaded: https://builds.apache.org/job/incubator-netbeans-linux.

Convenience binary of released source artifacts: https://netbeans.apache.org/download/index.html.

### Reporting Bugs

Bugs should be reported to https://issues.apache.org/jira/projects/NETBEANS/issues/

### Full History

The origins of the code in this repository are older than its Apache existence.
As such significant part of the history (before the code was donated to Apache)
is kept in an independent repository. To fully understand the code
you may want to merge the modern and ancient versions together:

```bash
$ git clone https://github.com/apache/incubator-netbeans.git
$ cd incubator-netbeans
$ git log uihandler/arch.xml
```

This gives you just few log entries including the initial checkin and
change of the file headers to Apache. But then the magic comes:

```bash
$ git remote add emilian https://github.com/emilianbold/netbeans-releases.git
$ git fetch emilian # this takes a while, the history is huge!
$ git replace 6daa72c98 32042637 # the 1st donation
$ git replace 6035076ee 32042637 # the 2nd donation
```

When you search the log, or use the blame tool, the full history is available:

```bash
$ git log uihandler/arch.xml
$ git blame uihandler/arch.xml
```

Many thanks to Emilian Bold who converted the ancient history to his
[Git repository](https://github.com/emilianbold/netbeans-releases)
and made the magic possible!

