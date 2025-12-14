import './Footer.css';

const Footer = () => {
    return (
        <footer className="footer">
            <div className="footer-content">
                <div className="footer-brand">
                    <h2>Magic Stream</h2>
                    <p>Unlimited movies, TV shows, and more.</p>
                </div>
                <div className="footer-links">
                    <div className="link-column">
                        <h4>Company</h4>
                        <a href="#">About Us</a>
                        <a href="#">Careers</a>
                        <a href="#">Contact</a>
                    </div>
                    <div className="link-column">
                        <h4>Support</h4>
                        <a href="#">Help Center</a>
                        <a href="#">Terms of Service</a>
                        <a href="#">Privacy Policy</a>
                    </div>
                    <div className="link-column">
                        <h4>Social</h4>
                        <a href="#">Twitter</a>
                        <a href="#">Instagram</a>
                        <a href="#">Facebook</a>
                    </div>
                </div>
            </div>
            <div className="footer-bottom">
                <p>&copy; {new Date().getFullYear()} Magic Stream. All rights reserved.</p>
            </div>
        </footer>
    );
};

export default Footer;
