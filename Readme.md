# Phase 1
# How to convert a monolith to  microservice

First we will Design a monlith architecture then , switch to microservice and deploy it 
using containers and kubernet . This will be a hand written code 
and gradually we will improve the project

# Build an E-commerce Application in Go – Monolithic Architecture

We'll build a small but realistic e-commerce backend using Go.

The objective is to understand how a monolithic application is structured before we break it into microservices later.

We'll start with three business modules:
1. User: Registration, login and user information.
2. Catalog: Product listing, product details and inventory.
3. Shopping & Orders: Cart, checkout and purchasing products.

# 1. What architecture are we building?
   A monolith means all business modules run inside one deployable application.

For example, User, Catalog and Order are separate logical modules,
but they share the same application process and database. 

           | user | 
client---> |Catalog|   ----> | Postgres | 
           |Order|   


All three modules run within one Go application.

What makes this a monolith?
- One Go application.
- One deployment.
- One process.
- All three modules run together.
- Modules communicate through direct function calls.
- No network communication between modules.



Important limitations of our first implementation
This is a learning prototype, not production-ready:
Current implementation	Production improvement
In-memory storage	PostgreSQL
One global mutex	Database transactions and appropriate concurrency control
No authentication	JWT/session-based authentication
No payment integration	Payment gateway
Basic error handling	Structured errors and logging
No cart	Dedicated cart functionality


# The global mutex currently serializes HTTP operations to protect our in-memory data. We will replace this approach as we introduce proper persistence and concurrency control.




Look at this line in main.go:
orders := order.NewService(users, products)
We are injecting the User and Catalog services into the Order service.
There is no HTTP call between these modules.
This is one of the biggest differences between a monolith and microservices.
In a microservices architecture, the Order Service might communicate with the User Service through HTTP or another messaging mechanism.



