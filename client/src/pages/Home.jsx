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

    useEffect(() => {
        const token = localStorage.getItem('access_token');
        if (token) {
            setLoggedIn(true);
            fetchMovies();
        } else {
            navigate('/login');
        }
    }, [navigate]);

    const fetchMovies = async () => {
        try {
            setLoading(true);
            const response = await axiosConfig.get("/movies");
            // API matches: { data: [...] } structure
            // But user showed: { data: [ ... ] } in the response body example.
            // Axios response object has .data property which is the body.
            // So result is response.data.data
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

    if (!loggedIn) return null;

    return (
        <div className="home-container">
            <Navbar />

            <div className="movies-container">
                <div className="section-header">
                    <h2>Featured Movies</h2>
                    {/* Category Filter Bar */}
                    <div className="filter-bar">
                        <button className="filter-btn active">All</button>
                        <button className="filter-btn">Action</button>
                        <button className="filter-btn">Comedy</button>
                        <button className="filter-btn">Drama</button>
                        <button className="filter-btn">Sci-Fi</button>
                    </div>
                </div>

                {error && <div className="alert alert-danger">{error}</div>}

                {loading ? (
                    <div className="loading-container">Loading amazing content...</div>
                ) : (
                    <>
                        {movies && movies.length > 0 ? (
                            <div className="movies-grid">
                                {movies.map((movie) => (
                                    <MovieCard key={movie._id} movie={movie} />
                                ))}
                            </div>
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