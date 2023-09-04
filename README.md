# Spotify Records

<section name"summary" id="summary">

## Summary

<ul>
 <a href="#desc"><li>Description</li></a>
 <a href="#prototype"><li>Prototype</li></a>
 <a href="#author"><li>Author</li></a>
 <a href="#how-init"><li>How init the api?</li></a>
</ul>

</section>

<section name"desc" id="desc">

## Description:

``Spotify Records`` it is a wiki about all the artists on the Spotify platform, there, the user can connect with his Spotify Music account and after that, the most diverse information about his favorite artists will be shown. On the site, it is also possible to stay on top of news about your favorite artists, in addition to closely monitoring their upcoming presentations and releases. Also, exclusive information of the artist content will be shown

The back-end of the project will be a REST API that will provide the database data for the front-end. The code is made of ``golang`` and using hexagonal architecture for better security and code organization. 

The front-end will be developed in ``React`` with some possible libraries to improve the performance and operation of the application

</section>
 
<section name"prototype" id="prototype">

## Diagrams:

#### Entity-Relashionship Model:

![SR Backend - Diagrama ER de banco de dados (pé de galinha)](https://github.com/emanuelvsz/spotify-records-backend/assets/84058517/e6c0f658-0771-42d0-b4e0-a9bb97c2d8d7)

## Prototype:
 
#### Home page
<img src="https://user-images.githubusercontent.com/84058517/226115550-bb4d1f5b-7513-4dd5-ba9b-ef4421cfd424.png" width=560/>

#### About the artist page

<img src="https://user-images.githubusercontent.com/84058517/226115645-7ef8dfa4-8609-4d6e-aa05-56b3b9d3b8aa.png" width=560/>

<a href="https://www.figma.com/file/MObQo3CpTAPbX2fPYZ6BeI/Spotify-Records?node-id=0%3A1&t=9kqFiWsBJ8W3ne76-1">click here to see the design in a better resolution</a>

<section id="how-init"></section>

### How init?

#### <section id="what-are-the-necessary-tools"></section>Necessary tools

- IDE of your choice
- Ubuntu Terminal <a href="https://learn.microsoft.com/en-us/windows/wsl/install">how to install?</a>
- Golang <a href="https://go.dev/doc/">how to install?</a>
- Migrate <a href="https://pkg.go.dev/github.com/golang-migrate/migrate/v4">how to install?</a>
- Sqlc <a href="https://docs.sqlc.dev/en/latest/overview/install.html">how to install?</a>
- Docker <a href="https://docs.docker.com/engine/install">how to install?</a>
- Swagger <a href="https://github.com/swaggo/swag">how to install?</a>

#### Clone the project 
 
```
git clone https://github.com/emanuelvsz/spotify-records-backend
cd spotify-records-backend
```

### <section id="como-rodar-o-projeto"></section>How run the project?

1. Open your development environment (<a href="https://code.visualstudio.com/download">VS Code</a>, <a href="https://www.jetbrains.com/go/promo/?source=google&medium=cpc&campaign=10156130867&term=goland&content=438684701890&gad=1&gclid=Cj0KCQjwmN2iBhCrARIsAG_G2i7Tsx5AtYXU8TlqWbu6rqD6AO_C6sJs4C8plkJPbA0HNRWExrQFFmgaAhIlEALw_wcB">Goland</a> e etc.)
2. Open a ``terminal linux``
3. Verify if the dependencies listed up are sucessful installed
4. Run ``docker``
5. Open the directory: ``cd spotify-records-backend``
6. Inside the project, execute the command ``cd config/docker`` to enter the directory where the docker settings are running
7. Run the command: ``docker compose rm -sf && docker compose up --build`` to run the spotify records container. Volte para o caminho padrão do projeto com o comando ``cd ../..``
9. Execute o comando ``go mod tidy``
10. Run the command ``bash -c "cd src/app/api && swag init -g ../../main.go --output ./docs --dir ./endpoints/handlers"
`` to generate the API swagger documentation
11. Run the server ``go run ./src/app/api``

### <section id="how-to-access-route-documentation"></section>How to access route documentation?
 
After running the server locally, the route documentation can be accessed at the following address: `http://localhost:8000/api/docs/index.html#/`.

<section name"author" id="author">

## Author

<a href="https://github.com/emanuelvsz">@emanuelvsz</a>

</section>
