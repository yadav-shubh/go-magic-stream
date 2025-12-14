import { fetchAuthInfo } from '../../api/authService';
import { useNavigate } from 'react-router-dom';
import './Navbar.css';

const Navbar = () => {
    const navigate = useNavigate();

    const handleLogout = async () => {
        try {
            const response = await fetchAuthInfo();
            const { logout_url } = response.data;

            localStorage.removeItem('access_token');
            localStorage.removeItem('refresh_token');

            if (logout_url) {
                window.location.href = logout_url;
            } else {
                navigate('/login');
            }
        } catch (error) {
            console.error("Error during logout:", error);
            localStorage.removeItem('access_token');
            localStorage.removeItem('refresh_token');
            navigate('/login');
        }
    };

    return (
        <nav className="navbar">
            <div className="navbar-brand">
                <h1>Magic Stream</h1>
            </div>
            <div className="navbar-actions">
                <button className="btn-logout" onClick={handleLogout}>
                    Logout
                </button>
            </div>
        </nav>

    );
};

export default Navbar;
