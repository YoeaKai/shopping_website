# Shopping_website

#### Shopping_website is a price comparison service that provides users with the cheapest product from different websites.

## About Projects

##### Shopping_website implements a gRPC server with Golang to establish a shopping website serving several users parallelized and streaming and equips:

* Plural workers crawl the product information using colly (framework for Golang) and the APIs opened by the target platform (like PChome) from different platforms concurrently.
* Database (MariaDB) establishes a cache mechanism.
* Structured logs make debugging more efficient.
* Interfaces increase scalability and achieve low coupling.
* Sleep time for rate limit avoids being mistaken for DDoS.
* Docker packages it to a container and I deployed it in AWS.

## Getting Started

This is an instruction on how we set up Shopping_website locally. Running follow these simple example steps.

### Prerequisites

1. [Download and install Go](https://go.dev/doc/install)

2. Install Docker

      #### **For Mac:**
      Install Homebrew
      ```
      /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
      ```
      Install Docker
      ```
      brew cask install docker
      ```
      #### **For Amazon Linux 2:**
      Apply pending updates using the yum command
      ```
      sudo yum update
      ```
      Search for Docker package
      ```
      sudo yum search docker
      ```

3. Install chrome (For momo crawler, needs to install chrome browser in OS or container)

