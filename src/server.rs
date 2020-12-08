use tonic::{transport::Server, Request, Response, Status};
    use films::films_server::{Films,FilmsServer};
    use films::{FilmRequest, FilmsResponse};
    use films::{Film};
    mod films; 

    // defining a struct for our service
    #[derive(Default)]
    pub struct MyFilms {}

    // implementing rpc for service defined in .proto
    #[tonic::async_trait]
    impl Films for MyFilms {
    // our rpc impelemented as function
        async fn request_film(&self,request:Request<FilmRequest>)->Result<Response<FilmsResponse>,Status>{
    // returning a response as SayResponse message as defined in .proto
            println!("Got a request: {:?}", request);
            Ok(Response::new(FilmsResponse{
    // reading data from request which is awrapper around our SayRequest message defined in .proto
                 films: [Film{title: "words".to_string()}].to_vec(),
            }))
        }
    }

    #[tokio::main]
    async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // defining address for our service
        let addr = "[::1]:50051".parse().unwrap();
    // creating a service
        let films = MyFilms::default();
        println!("Server listening on {}", addr);
    // adding our service to our server.
        Server::builder()
            .add_service(FilmsServer::new(films))
            .serve(addr)
            .await?;
        Ok(())
    }
