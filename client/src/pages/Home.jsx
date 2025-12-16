import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import axiosConfig from "../api/axiosConfig.js";
import Navbar from '../components/common/Navbar';
import Footer from '../components/common/Footer';
import MovieCard from '../components/movies/MovieCard';
import './Home.css';

const Home = () => {
    const [movies, setMovies] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const [loggedIn, setLoggedIn] = useState(false);
    const navigate = useNavigate();

    // Pagination and Genre state
    const [currentPage, setCurrentPage] = useState(1);
    const [selectedGenre, setSelectedGenre] = useState('');
    const itemsPerPage = 14;

    useEffect(() => {
        const token = localStorage.getItem('access_token');
        if (token) {
            setLoggedIn(true);
            fetchMovies();
        } else {
            navigate('/login');
        }
    }, [navigate, currentPage, selectedGenre]);

    const fetchMovies = async () => {
        try {
            setLoading(true);
            const genreParam = selectedGenre === 'All' ? '' : selectedGenre;
            const query = `?page=${currentPage}&size=${itemsPerPage}&search=&genre=${genreParam}`;

            const response = await axiosConfig.get(`/movies${query}`);

            if (response.data && response.data.data) {
                setMovies(response.data.data);
            } else {
                setMovies([]);
            }
        } catch (error) {
            console.error("Error fetching movies", error);
            setError("Failed to load movies. Please try again.");
        } finally {
            setLoading(false);
        }
    };

    const handleGenreClick = (genre) => {
        setSelectedGenre(genre);
        setCurrentPage(1); // Reset to first page on genre change
    };

    const handleNextPage = () => {
        setCurrentPage(prev => prev + 1);
    };

    const handlePrevPage = () => {
        setCurrentPage(prev => (prev > 1 ? prev - 1 : 1));
    };

    if (!loggedIn) return null;

    return (
        <div className="home-container">
            <Navbar />

            <div className="movies-container">
                <div className="section-header">
                    <h2>Featured Movies</h2>
                    {/* Category Filter Bar */}
                    <div className="filter-bar">
                        {['All', 'Comedy', 'Drama', 'Western', 'Fantasy', 'Thriller', 'Sci-Fi', 'Action', 'Mystery', 'Crime'].map((genre) => (
                            <button
                                key={genre}
                                className={`filter-btn ${selectedGenre === (genre === 'All' ? '' : genre) || (genre === 'All' && selectedGenre === '') ? 'active' : ''}`}
                                onClick={() => handleGenreClick(genre === 'All' ? '' : genre)}
                            >
                                {genre}
                            </button>
                        ))}
                    </div>
                </div>

                {error && <div className="alert alert-danger">{error}</div>}

                {loading ? (
                    <div className="loading-container">Loading amazing content...</div>
                ) : (
                    <>
                        {movies && movies.length > 0 ? (
                            <>
                                <div className="movies-grid">
                                    {movies.map((movie) => (
                                        <MovieCard key={movie._id} movie={movie} />
                                    ))}
                                </div>
                                <div className="pagination-controls">
                                    <button
                                        onClick={handlePrevPage}
                                        disabled={currentPage === 1}
                                        className="pagination-btn"
                                    >
                                        Previous
                                    </button>
                                    <span className="page-indicator">Page {currentPage}</span>
                                    <button
                                        onClick={handleNextPage}
                                        disabled={movies.length < itemsPerPage}
                                        className="pagination-btn"
                                    >
                                        Next
                                    </button>
                                </div>
                            </>
                        ) : (
                            <div className="empty-state">
                                <p>No movies found. Check back later!</p>
                            </div>
                        )}
                    </>
                )}
            </div>

            <Footer />
        </div>
    );
};

export default Home;