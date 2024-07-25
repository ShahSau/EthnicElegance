<p align="center">
  <h1 align="center">EthnicElegance</h1>
</p>
<p align="center">
    <em>Ecommerce site backend using golang</em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/license/ShahSau/EthnicElegance?style=flat&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/ShahSau/EthnicElegance?style=flat&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/ShahSau/EthnicElegance?style=flat&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/ShahSau/EthnicElegance?style=flat&color=0080ff" alt="repo-language-count">
<p>
<p align="center">
		<em>Developed with the software and tools below.</em>
</p>
<p align="center">
	<img src="https://img.shields.io/badge/YAML-CB171E.svg?style=flat&logo=YAML&logoColor=white" alt="YAML">
	<img src="https://img.shields.io/badge/JSON-000000.svg?style=flat&logo=JSON&logoColor=white" alt="JSON">
	<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=flat&logo=Go&logoColor=white" alt="Go">
	<img src="https://img.shields.io/badge/JWT-000000?style=flat&logo=Go&logoColor=white" alt="JWT">;
        <img src="https://img.shields.io/badge/Gin-black?style=flat&logo=Go&logoColor=white" alt="Gin-go">
</p>
<hr>

## 🔗 Quick Links

> - [📍 Overview](#-overview)
> - [📦 Features](#-features)
> - [📂 Repository Structure](#-repository-structure)
> - [🧩 Modules](#-modules)
> - [🚀 Getting Started](#-getting-started)
>   - [⚙️ Installation](#️-installation)
>   - [🤖 Running EthnicElegance](#-running-EthnicElegance)
>   - [🧪 Tests](#-tests)
> - [🛠 Project Roadmap](#-project-roadmap)
> - [🤝 Contributing](#-contributing)
> - [📄 License](#-license)
> - [👏 Acknowledgments](#-acknowledgments)

---

## 📍 Overview

Welcome to the eCommerce Backend project! This repository contains a robust and scalable backend system built using Go and the Gin framework, designed to support a comprehensive eCommerce platform. The system is structured to provide secure authentication, user management, and administrative functionalities, ensuring a seamless and efficient operation of an online store.

---

## 📦 Features
This eCommerce backend, built with Go and Gin, provides robust and scalable functionalities to support a full-fledged online store. The backend includes the following key features:

<h6>JWT Authentication</h6>
- Secure User Authentication: Implements JWT (JSON Web Token) for secure user authentication and authorization.
- Token Generation: Generates tokens upon successful login, ensuring secure and stateless user sessions.
- Token Validation: Validates tokens for each request to ensure that only authenticated users can access protected routes.

<h6>User Section</h6>
- User Registration: Allows new users to register with the platform by providing necessary details.
- User Login: Enables registered users to log in and receive a JWT for session management.
- Profile Management: Lets users view and update their profile information.
- Order Management: Users can place new orders, view their order history, and track order statuses.
- Product Browsing: Users can browse through the catalog of available products, view product details, and search for specific items.

<h6>Admin Section</h6>
- Admin Authentication: Admin-specific authentication to access the admin dashboard and perform administrative tasks.
- User Management: Admins can view, edit, and delete user accounts, as well as manage user permissions.
- Product Management: Admins can add, edit, and remove products from the catalog, including updating product details and stock levels.
- Order Management: Admins can view all orders, update order statuses, and manage the order fulfillment process.

---

## 📂 Repository Structure

```sh
└── EthnicElegance/
    ├── EthnicElegance
    ├── LICENSE
    ├── README.md
    ├── constant
    │   └── constant.go
    ├── controller
    │   ├── adminController.go
    │   ├── healthCheckController.go
    │   ├── productController.go
    │   └── userController.go
    ├── database
    │   └── connection.go
    ├── docs
    │   ├── docs.go
    │   ├── swagger.json
    │   └── swagger.yaml
    ├── go.mod
    ├── go.sum
    ├── helper
    │   └── helper.go
    ├── main.go
    ├── router
    │   ├── router.go
    │   └── routes.go
    └── types
        └── user-type.go
```

---

## 🧩 Modules

<details closed><summary>.</summary>
| File                                                                     | Summary                             |
| ---                                                                      | ---                                 |
| [main.go](https://github.com/ShahSau/EthnicElegance/blob/master/main.go) | entry point of the go project |

</details>

<details closed><summary>router</summary>

| File                                                                                | Summary                                      |
| ---                                                                                 | ---                                          |
| [routes.go](https://github.com/ShahSau/EthnicElegance/blob/master/router/routes.go) | All the routes of the project                |
| [router.go](https://github.com/ShahSau/EthnicElegance/blob/master/router/router.go) | Grouping, and creating common function for the routes |

</details>

<details closed><summary>constant</summary>

| File                                                                                      | Summary                                          |
| ---                                                                                       | ---                                              |
| [constant.go](https://github.com/ShahSau/EthnicElegance/blob/master/constant/constant.go) | variables like collection name, erroe messages of the whole project |

</details>

<details closed><summary>types</summary>

| File                                                                                     | Summary                                        |
| ---                                                                                      | ---                                            |
| [user-type.go](https://github.com/ShahSau/EthnicElegance/blob/master/types/user-type.go) |All the types of the whole project |

</details>

<details closed><summary>controller</summary>

| File                                                                                                                  | Summary                                                         |
| ---                                                                                                                   | ---                                                             |
| [productController.go](https://github.com/ShahSau/EthnicElegance/blob/master/controller/productController.go)         | Product routes controllers    |
| [adminController.go](https://github.com/ShahSau/EthnicElegance/blob/master/controller/adminController.go)             | Admin routes controllers     |
| [userController.go](https://github.com/ShahSau/EthnicElegance/blob/master/controller/userController.go)               | users routes controllers     |
| [healthCheckController.go](https://github.com/ShahSau/EthnicElegance/blob/master/controller/healthCheckController.go) | health check route controller |

</details>

<details closed><summary>database</summary>

| File                                                                                          | Summary                                            |
| ---                                                                                           | ---                                                |
| [connection.go](https://github.com/ShahSau/EthnicElegance/blob/master/database/connection.go) | Functions to connect and get collletion from database |

</details>

<details closed><summary>helper</summary>

| File                                                                                | Summary                                      |
| ---                                                                                 | ---                                          |
| [helper.go](https://github.com/ShahSau/EthnicElegance/blob/master/helper/helper.go) | Helper functions |

</details>

---

## 🚀 Getting Started

***Requirements***

Ensure you have the following dependencies installed on your system:

* **Go**: `version 1.22.5`

### ⚙️ Installation

1. Clone the EthnicElegance repository:

```sh
git clone https://github.com/ShahSau/EthnicElegance
```

2. Change to the project directory:

```sh
cd EthnicElegance
```

3. Install the dependencies:

```sh
go build -o backend
```

### 🤖 Running EthnicElegance

Use the following command to run EthnicElegance:

```sh
./backend
```

### 🧪 Tests

To execute tests, run:

```sh
go test
```

---

## 🛠 Project Roadmap

- [ ] `► Email Verification`
- [ ] `► OTP`
- [ ] `► API Testing`
- [ ] `► Docker`
- [ ] `► CI/CD`

---

## 🤝 Contributing

Contributions are welcome! Here are several ways you can contribute:

- **[Submit Pull Requests](https://github.com/ShahSau/EthnicElegance/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.
- **[Join the Discussions](https://github.com/ShahSau/EthnicElegance/discussions)**: Share your insights, provide feedback, or ask questions.
- **[Report Issues](https://github.com/ShahSau/EthnicElegance/issues)**: Submit bugs found or log feature requests for Ethnicelegance.

<details closed>
    <summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your GitHub account.
2. **Clone Locally**: Clone the forked repository to your local machine using a Git client.
   ```sh
   git clone https://github.com/ShahSau/EthnicElegance
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to GitHub**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.

Once your PR is reviewed and approved, it will be merged into the main branch.

</details>

---
## 📄 License

This project is protected under the [MIT License](https://choosealicense.com/licenses/mit) License. For more details, refer to the [LICENSE](https://choosealicense.com/licenses/mit) file.

---
