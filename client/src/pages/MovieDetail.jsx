import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import axiosConfig from "../api/axiosConfig";
import Navbar from '../components/common/Navbar';
import Footer from '../components/common/Footer';
import './MovieDetail.css';

const MovieDetail = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const [movie, setMovie] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchMovieDetail = async () => {
            try {
                const response = await axiosConfig.get(`/movies/${id}`);
                // API response: { data: [ { ... } ], ... }
                if (response.data && response.data.data && response.data.data.length > 0) {
                    setMovie(response.data.data[0]);
                } else {
                    setError("Movie not found.");
                }
            } catch (err) {
                console.error("Error fetching movie details", err);
                setError("Failed to load movie details.");
            } finally {
                setLoading(false);
            }
        };

        if (id) {
            fetchMovieDetail();
        }
    }, [id]);

    const getBadgeClass = (value) => {
        if (!value) return '';
        if (value <= 2) return 'badge-excellent';
        if (value <= 3) return 'badge-okay';
        return 'badge-bad';
    };

    if (loading) return <div className="loading-container">Loading details...</div>;
    if (error) return <div className="error-container">{error} <button onClick={() => navigate('/')}>Go Home</button></div>;
    if (!movie) return null;

    return (
        <div className="detail-page">
            <Navbar />

            <div className="movie-detail-container">
                <button className="back-btn" onClick={() => navigate('/')}>
                    ← Back to Movies
                </button>

                <div className="detail-content">
                    {/* Left Column: Poster & Trailer */}
                    <div className="media-column">
                        <div className="detail-poster">
                            <img src={movie.poster_path} alt={movie.title} />
                        </div>

                        {movie.youtube_id && (
                            <div className="trailer-container">
                                <h3>Official Trailer</h3>
                                <div className="iframe-wrapper">
                                    <iframe
                                        src={`https://www.youtube.com/embed/${movie.youtube_id}`}
                                        title="YouTube video player"
                                        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                                        allowFullScreen
                                    ></iframe>
                                </div>
                            </div>
                        )}
                    </div>

                    {/* Right Column: Info */}
                    <div className="info-column">
                        <h1 className="movie-title">{movie.title}</h1>

                        <div className="meta-row">
                            {movie.ranking && (
                                <span className={`detail-badge ${getBadgeClass(movie.ranking.ranking_value)}`}>
                                    {movie.ranking.ranking_name}
                                </span>
                            )}
                            <div className="detail-genres">
                                {movie.genre && movie.genre.map((g) => (
                                    <span key={g.genre_id} className="detail-genre-tag">
                                        {g.genre_name}
                                    </span>
                                ))}
                            </div>
                        </div>

                        <div className="review-section">
                            <h3>Admin Review</h3>
                            <p className="review-text">"{movie.admin_review}"</p>
                        </div>

                        <div className="movie-info-raw">
                            <p><span className="label">IMDb ID:</span> {movie.imdb_id}</p>
                            {/* Assuming created_at might not be relevant to show user, but ok to have if needed */}
                        </div>
                    </div>
                </div>
            </div>

            <Footer />
        </div>
    );
};

export default MovieDetail;
