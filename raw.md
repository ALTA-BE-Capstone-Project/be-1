# mrClinic

## Tech Stack

[![Go.Dev reference](https://img.shields.io/badge/-echo%20-05122A?style=flat&logo=go)](https://github.com/labstack/echo)
[![Go.Dev reference](https://img.shields.io/badge/-gorm%20-05122A?style=flat&logo=go)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/-aws--s3%20-05122A?style=flat&logo=Amazon%20Aws)](https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/go/example_code/s3)
[![Go.Dev reference](https://img.shields.io/badge/-google--calendar%20-05122A?style=flat&logo=Google%20Calendar)](https://github.com/googleapis/google-api-go-client/blob/main/examples/calendar.go)

## About The Project

<p>
  Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce erat ipsum, auctor eu condimentum quis, sagittis in risus. Nulla facilisi. Mauris et malesuada lorem. Nam eget accumsan velit, id aliquet magna. Sed quis ornare nibh. Cras sodales elit ex, ut tempus mi porta id. Donec finibus pulvinar est, sed tempor eros lacinia aliquet. Maecenas et dui sit amet ipsum luctus eleifend sed at nisl. Nam et ligula ac massa feugiat porttitor non eget ipsum. Sed velit mi, pellentesque id velit non, commodo mollis turpis. Praesent volutpat, mi sed ultricies finibus, nisi dolor eleifend augue, congue luctus ipsum diam eget lacus. In eget nulla diam. Curabitur auctor tortor lectus, in dictum turpis aliquet id. Ut sed pellentesque ligula. Aliquam erat volutpat. Pellentesque in nulla ac tellus consequat accumsan.
</p>

<p>
Maecenas vel justo venenatis, rhoncus mauris vel, elementum augue. Aenean volutpat est quis mi pharetra pharetra. Sed accumsan molestie mattis. Vestibulum id erat risus. Nunc placerat et arcu quis euismod. Fusce accumsan nibh nisi, id interdum sapien lacinia quis. Fusce eu enim hendrerit, laoreet lectus sed, egestas est. Aliquam bibendum vestibulum nibh, facilisis maximus magna convallis a.
</p>

[OPEN API](https://app.swaggerhub.com/apis/faliqadlan/mrClinic/1.0.0)

<details>
<summary>ERD</summary>
</details>
<details>
<summary>&nbsp;Patient</summary>
| Feature User | Endpoint | Query Param | Request Body | JWT Token | Fungsi |
| ------------ | ---------| ----------- | ------------ | --------- | ------ |
</details>

| Feature Doctor | Endpoint | Query Param | Request Body | JWT Token | Fungsi |
| ------------ | ---------| ----------- | ------------ | --------- | ------ |
| POST   | /login   | -   | indentity & | 

  
| Feature User | Endpoint | Query Param | Request Body | JWT Token | Fungsi |
| ------------ | ---------| ----------- | ------------ | --------- | ------ |
| POST         | /users/login | - | identity & password | NO | login user with identity & passwords |
| POST         | /users/register  | - | - | NO |  | 
| POST         | /users/avatar  | - | avatar | YES | upload avatar for user profile |
| GET          | /users/profile | - | - | YES | get current user profile |
| PUT          | /users/ | - | password | YES | update current user profile |
