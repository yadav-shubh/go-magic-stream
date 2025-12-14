import { useEffect, useRef } from 'react';
import { useSearchParams, useNavigate } from 'react-router-dom';
import { authenticateUser } from '../api/authService';

const SignInCallback = () => {
    const [searchParams] = useSearchParams();
    const navigate = useNavigate();
    const codeProcessed = useRef(false);

    useEffect(() => {
        const code = searchParams.get('code');

        if (code && !codeProcessed.current) {
            codeProcessed.current = true;

            authenticateUser(code)
                .then((response) => {
                    console.log("Authentication successful:", response.data);
                    const { access_token, refresh_token } = response.data.data[0];

                    localStorage.setItem('access_token', access_token);
                    localStorage.setItem('refresh_token', refresh_token);

                    navigate('/');
                })
                .catch((error) => {
                    console.error("Authentication failed:", error);
                    // Handle error (e.g., show message or redirect to login)
                    navigate('/'); // Fallback for now
                });
        } else if (!code) {
            console.warn("No code found in URL");
            navigate('/');
        }
    }, [searchParams, navigate]);

    return (
        <div className="d-flex justify-content-center align-items-center vh-100">
            <div className="text-center">
                <h2>Authenticating...</h2>
                <div className="spinner-border text-primary" role="status">
                    <span className="visually-hidden">Loading...</span>
                </div>
            </div>
        </div>
    );
};

export default SignInCallback;
