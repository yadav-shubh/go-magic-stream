import './MovieCard.css';

const MovieCard = ({ movie }) => {
    const { title, poster_path, ranking, genre } = movie;

    // Determine badge color based on ranking value
    const getBadgeClass = (value) => {
        if (value <= 2) return 'badge-excellent';
        if (value <= 3) return 'badge-okay';
        return 'badge-bad';
    };

    return (
        <div className="movie-card">
            <div className="poster-wrapper">
                <img src={poster_path} alt={title} loading="lazy" />
                <div className="rating-badge">
                    <span className={`badge ${getBadgeClass(ranking.ranking_value)}`}>
                        {ranking.ranking_name}
                    </span>
                </div>
            </div>
            <div className="movie-info">
                <h3>{title}</h3>
                <div className="genre-tags">
                    {genre.map((g, index) => (
                        <span key={index} className="genre-tag">{g.genre_name}</span>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default MovieCard;
