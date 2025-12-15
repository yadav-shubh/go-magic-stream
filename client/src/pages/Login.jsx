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
            window.location.href = loginUrl.toString();
        }
    };

    return (
        <div className="container-fluid p-0 overflow-hidden">
            <div className="row g-0 min-vh-100">
                {/* Left Side - Hero/Brand */}
                <div className="col-md-7 login-left d-flex flex-column justify-content-center align-items-center text-center text-white position-relative p-5">
                    <div className="brand-content position-relative" style={{ zIndex: 2 }}>
                        <h1 className="display-1 fw-bold mb-3">Magic Stream</h1>
                        <p className="lead fs-4 text-white-50">Unlimited movies, TV shows, and more.</p>
                    </div>
                    <div className="hero-overlay"></div>
                </div>

                {/* Right Side - Login Form */}
                <div className="col-md-5 login-right d-flex justify-content-center align-items-center bg-light p-4 p-md-5">
                    <div className="login-form-container w-100" style={{ maxWidth: '450px' }}>
                        <h2 className="mb-2 fw-bold text-dark">Welcome Back</h2>
                        <p className="text-secondary mb-4">Please sign in to continue enjoying your favorite content.</p>

                        {error && <div className="alert alert-danger" role="alert">{error}</div>}

                        <div className="d-grid gap-2">
                            <button
                                className="btn btn-primary btn-lg btn-login-kinde border-0 py-3"
                                onClick={handleLogin}
                                disabled={!loginUrl}
                            >
                                {loginUrl ? 'Sign In with SSO' : 'Loading...'}
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Login;
