<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->


<!-- PROJECT LOGO -->
<br />
<div align="center">
    <img src="" alt="Logo">

  <h3 align="center">Managers</h3>

  <p align="center">
    Managers
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
      <ul>
        <li><a href="#api-documentation">API Documentation</a></li>
      </ul>
      <ul>
        <li><a href="#microservice-diagram">Microservice Diagram</a></li>
      </ul>
      <ul>
        <li><a href="#sequence-diagram">Sequence Diagram</a></li>
      </ul>
    </li>
    <li>
        <a href="#developement">Developement</a>
        <ul>
            <li><a href="#credential">Credential</a></li>
        </ul>
        <ul>
            <li><a href="#required-stuff">Required Stuff</a></li>
        </ul>
        <ul>
            <li><a href="#required-step">Required Step</a></li>
        </ul>
        <ul>
            <li><a href="#usage">Usage</a></li>
        </ul>
    </li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This project is a test project

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

This the tech stack that we used for developing Postgre service

* [![Fiber][fiber]][fiber-url]
* [![Gorm][gorm]][gorm-url]
* [![Golang][golang]][golang-url]
* [![GoValidator][govalidator]][go-validator]
* [![Postgre][postgre]][postgre-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### API Documentation

This is API Documentation [API Documentation](https://documenter.getpostman.com/view/11975231/2s8YekSvNC)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Microservice Diagram

<div align="center">
    <img src="https://i.ibb.co/b3n2RgH/microservice-diagram-drawio.png" alt="Logo">
</div>

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Sequence Diagram

<div align="center">
    <img src="https://i.ibb.co/SdGQvn2/j-PEz-Ji-Cm4-CVt-FCMz-G5-Mr-Om-VKWOW5-YKg-K6-Jdsj-LWu-Tk-Ip-Y-3f-YLq-RAqh41qstq-Ny-X5t-P9-Zrlb-MLSmx.png" alt="Logo">
</div>

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- Developement -->
## Developement

<br />

### Credential:

- admin:

email: admin@gmail.com
<br/>
password: adminpassword
<br/>


- user

email: user@gmail.com
<br/>
password: userpassword
<br/>
<br/>



### Required Stuff:

- golang >= 1.19 -> [golang](https://go.dev/)
- kubernetes latest
- minikube

### Required Step:

Please enable Ingress during apply the k8s configuration
Please check DB_HOST in ``authentication-service.yaml`` and ``user-service.yaml`` use local ip instead of FQDN after apply k8s configuration
After running, in terminal type ``minikube ip`` then point to:

unix: 
- /etc/hosts
  managers.com : the ip

### Usage:

Command to run environment :
```bash
./run.sh
```
<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- Contributors -->
## Contributors

Havis Iqbal Syahrunizar - [@hirasakavizu](https://twitter.com/hirasakavizu) - havisikkubaru@gmail.com

Github: [https://github.com/vizucode](https://github.com/vizucode)

[![linkedinhavis][linkedinhavis-shield]][linkedinhavis-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[linkedinhavis-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedinhavis-url]: https://www.linkedin.com/in/havis-iqbal/


[fiber]: https://img.shields.io/badge/fiber-gray?style=for-the-badge&logo=go&logoColor=00ADD8
[fiber-url]: https://gofiber.io/

[gorm]: https://img.shields.io/badge/gorm-gray?style=for-the-badge&logo=go&logoColor=00ADD8
[gorm-url]: https://gorm.io/

[govalidator]: https://img.shields.io/badge/go_validator-gray?style=for-the-badge&logo=go&logoColor=00ADD8
[go-validator]: https://github.com/go-playground/validator

[golang]: https://img.shields.io/badge/golang-gray?style=for-the-badge&logo=go&logoColor=00ADD8
[golang-url]: https://go.dev/

[postgre]: https://img.shields.io/badge/Postgre-gray?style=for-the-badge&logo=postgresql&logoColor=00ADD8
[postgre-url]: https://www.postgresql.org/
