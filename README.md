Online Library System - Test Report

📌 Overview

This document provides a test report for the Online Library System, detailing the test execution results, observations, and areas for improvement.

📅 Date: February 24, 2025

⏱ Total Execution Time: 0.969s

📂 Module: online-library-system/controllers

✅ Passed Tests

1️⃣ User Authentication

TestSignup – Successfully created a new user (John Doe)

TestLogin – Successfully authenticated user and generated a token

2️⃣ Books Management

TestAddBook – Successfully added a new book

TestGetBooks – Successfully retrieved all books

TestGetBook – Successfully retrieved a specific book (ISBN: 3333333333)

TestRemoveBook – Successfully removed a book (ISBN: 4444444444)

3️⃣ Library Management

TestGetLibraries – Successfully retrieved all libraries

TestCreateLibrary – Successfully created a new library (New Library)

TestCreateDuplicateLibrary – Prevented duplicate library creation (Conflict 409)

TestDeleteLibrary – Successfully deleted library (ID: 1)

TestDeleteNonExistentLibrary – Handled non-existent library deletion (ID: 999)

4️⃣ Requests & Users Management

TestCreateIssueRequest – Successfully created a book issue request (Reader ID: 1)

TestCreateAdmin – Successfully created an admin user

TestGetUser – Successfully retrieved user details (ID: 1)

TestGetUsers – Successfully retrieved all users

TestGetAdmins – Successfully retrieved all admins

TestUpdateUser – Successfully updated user information (ID: 1)

TestDeleteUser – Successfully deleted user (ID: 1)
