# SU.ED Fitness Management System
A web-based fitness center management system developed as a Senior Project for the Bachelor of Science in Information Technology program at Silpakorn University.
This system manages memberships, health questionnaires, fitness visits, QR Code check-ins, dashboards, and news announcements.

---

## Features

### ser Features
- Register and login with JWT authentication
- Apply for fitness memberships
- Fill in personal information and health questionnaires (PAR-Q)
- View membership details and status
- Display personal QR Code for check-in

### Admin Features
- Manage users (Create, Edit, Delete)
- Approve or reject memberships
- Record fitness visits
- Scan QR Code for automatic check-in
- View dashboard statistics
- Manage news and announcements

### Dashboard
- Total active members
- Daily fitness visits
- Membership statistics (daily, monthly, yearly)
- Visit statistics
- Recent visit history
- System activity logs

##Technologies Used

###Frontend
- React.js
- React Bootstrap
- Recharts
- React Router
- js-cookie

###Backend
- Golang
- Gin Framework
- GORM
- JWT Authentication
###Database
- PostgreSQL

###Cloud Services
- Cloudinary (Image Upload)

###Other Libraries
- html5-qrcode
- Casbin (Authorization)

โครงสร้างโปรแกรม Application server
1.my_backend ตัว App_server(Backend)
2.my_project ตัว Web_client(Frontend)
3.my_database ตัว Database 
