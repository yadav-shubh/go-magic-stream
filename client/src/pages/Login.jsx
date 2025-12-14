import { useEffect, useState } from 'react';
import { fetchAuthInfo } from '../api/authService';
import './Login.css';

const Login = () => {
    const [loginUrl, setLoginUrl] = useState(null);
    const [error, setError] = useState(null);

    useEffect(() => {
        const getAuthInfo = async () => {
            try {
                const response = await fetchAuthInfo();
                setLoginUrl(response.data.login_url);
            } catch (err) {
                console.error("Failed to fetch auth info", err);
                setError("Unable to initialize login. Please try again later.");
            }
        };
        getAuthInfo();
    }, []);

    const handleLogin = () => {
        if (loginUrl) {
            window.location.href = loginUrl;
        }
    };

    return (
        <div className="login-container">
            <div className="login-left">
                <div className="brand-content">
                    <h1>Magic Stream</h1>
                    <p>Unlimited movies, TV shows, and more.</p>
                </div>
                <div className="hero-overlay"></div>
            </div>

            <div className="login-right">
                <div className="login-form-container">
                    <h2>Welcome Back</h2>
                    <p className="login-subtext">Please sign in to continue enjoying your favorite content.</p>

                    {error && <div className="error-message">{error}</div>}

                    <div className="login-actions">
                        <button
                            className="btn-login-kinde"
                            onClick={handleLogin}
                            disabled={!loginUrl}
                        >
                            {loginUrl ? 'Sign In with SSO' : 'Loading...'}
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Login;
