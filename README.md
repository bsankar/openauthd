
  

<!-- ABOUT THE PROJECT -->
## About The Project

OpenAuthD is a developer friendly OAuth2.0 Auth Provider to be used in Dev/Test environments. It is developed with two main goals - be light weight and easy to setup, so that developers can easily test their apps against an OAuth2.0 Auth provider without setting up complex authentication systems and wait for Enterprise IT to provide access.

Why its called OpenAuthD? 
Ultimate aim of the project is to make OpenAuthD as a decentralized trust architecture platform for enterprise developers.

<p align="right">(<a href="#top">back to top</a>)</p>



### Built With

This section should list any major frameworks/libraries used to bootstrap your project. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.

* Golang
* Supports MS SQL Server Express as backend
* bcrypt
* jwt
* Golang Templates using Booststrap


<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

Below are the Prerequisites and Install instructions

### Prerequisites

Requires latest Golang version and MS Sql server 
 

### Installation

Below are the steps to run OpenAuthD by building from source. Development of setup file is in progress.

1. Clone the repo
2. Setup the database server and tables
3. Setup your own certificate
4.  Setup below environment variables
	1. DBUSERNAME
	2. DBPASSWORD
	3. TOKENSECRET
5. Build and run using`
   ```cli	
   go run .
   ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

The application runs on port 4443 and the endpoints are TLS based.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Resource Owner Password Credentials Grant
- [x] Implicit Grant
- [x] JWT
- [ ] Authorization Code Grant
- [ ] Client Credentials Grant
- [ ] OIDC Support
- [ ] MySQL and Other DBs support
- [ ] Docker and Kubernetes Support
- [ ] Web UI Based User import support
- [ ] Blockchain based Decentralized ID (DID)
  



<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion or would like to contribute in any way that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the Apache 2.0 License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Your Name - [@balajesankar](https://twitter.com/balajesankar) - balajesankar@outlook.com

Project Link: [https://github.com/bsankar/openauthd/](https://github.com/bsankar/openauthd/)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

* [OAuth 2.0 RFC 6749](https://tools.ietf.org/html/rfc6749)
* golang
