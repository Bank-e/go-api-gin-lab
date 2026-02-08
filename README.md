# Student Management API (Gin Edition)
> **Architecture:** Layered Architecture | **Framework:** Gin Gonic | **Database:** SQLite

---

## Overview
โปรเจกต์นี้เป็นการปรับปรุงระบบจัดการข้อมูลนักเรียน (Refactor) จาก Lab 2 ให้มาอยู่ในรูปแบบ **Layered Architecture** เพื่อแยกหน้าที่ความรับผิดชอบของโค้ดให้ชัดเจน และเปลี่ยนมาใช้ **Gin Framework** ซึ่งเป็น Web Framework ที่มีประสิทธิภาพสูงในภาษา Go เพื่อการจัดการเส้นทาง (Routing) และข้อมูล JSON ที่ง่ายขึ้น

> ## Lab 2 
> [https://github.com/Bank-e/go_lab2/tree/main]
---

## Features (Refactored)
* ✅ **Gin Framework:** ใช้ Framework ในการจัดการ HTTP Request และ JSON Binding ที่รวดเร็ว
* ✅ **Layered Architecture:** แบ่งส่วนการทำงานเป็น Handler, Service, Repository และ Model
* ✅ **Structured Code:** แยกไฟล์ตามหน้าที่เพื่อความง่ายในการบำรุงรักษาและขยายระบบในอนาคต
* ✅ **SQLite Integration:** เชื่อมต่อฐานข้อมูล SQLite ผ่าน Repository Layer

---

##  Tech Stack
| Component | Technology |
| :--- | :--- |
| **Backend** | ![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) |
| **Framework** | ![Gin](https://img.shields.io/badge/Gin-008080?style=for-the-badge&logo=go&logoColor=white) |
| **Database** | ![SQLite](https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white) |

---

##  Layered Architecture
การทำงานถูกแบ่งออกเป็นชั้นต่างๆ ดังนี้:

| Layer | Responsibility |
| :--- | :--- |
| **Handler** | รับ HTTP Request, เรียกใช้ Service และส่ง JSON Response กลับ |
| **Service** | จัดการ Business Logic และประสานงานระหว่าง Handler และ Repository |
| **Repository** | ติดต่อกับฐานข้อมูลโดยตรงผ่านคำสั่ง SQL |
| **Model** | กำหนดโครงสร้างข้อมูล (Struct) ที่ใช้ในระบบ |

---

## Project Structure
```bash
go-api-gin/
├─ config/        #  การตั้งค่าฐานข้อมูล (Database Connection)
├─ handlers/      #  ส่วนรับ Request และส่ง Response (Gin)
├─ models/        #  โครงสร้างข้อมูล (Data Models)
├─ repositories/  #  ส่วนจัดการคำสั่ง SQL (Database Access)
├─ services/      #  ส่วนจัดการตรรกะทางธุรกิจ (Business Logic)
├─ main.go        #  จุดเริ่มต้นของโปรแกรม (Entry Point)
└─ students.db    #  ไฟล์ฐานข้อมูล SQLite
