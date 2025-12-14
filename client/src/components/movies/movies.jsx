import "../movie/Movie"

const Movies = ({ movies, message }) => {
    return (
        <div className="container mt-4">
            <div className="row">
                {movies && movies.length > 0 ? (
                    movies.map(movie => (
                        <Movie key={movie._id} movie={movie} />
                    ))
                ) : (
                    <p>{message}</p>
                )}
            </div>
        </div>
    );
};

export default Movies;
