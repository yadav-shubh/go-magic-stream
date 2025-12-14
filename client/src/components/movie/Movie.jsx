const Movie = ({ movie }) => {
    return (
        <div className="col-md-4 mb-4">
            <div style={{ position: "relative" }}>
                <img src={movie.poster_path} alt={movie.title}
                    className="card-img-top"
                    style={{
                        objectFit: "cover",
                        width: "100%",
                        height: "250px"
                    }} />
            </div>
            <div className="card h-100 shadow-sm">
                <div className="card-body d-flex flex-column">
                    <h5 className="card-title">{movie.title}</h5>
                    <p className="card-text mb-2">{movie.imdb_id}</p>
                </div>
            </div>
        </div>
    );
};

export default Movie;