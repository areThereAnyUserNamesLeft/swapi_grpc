fn main() -> Result<(), Box<dyn std::error::Error>> {
    // compiling protos using path on build time
    tonic_build::compile_protos(
        "protos/github.com/areThereAnyUserNamesLeft/swapi_contract/films.proto",
    )?;
    Ok(())
}
