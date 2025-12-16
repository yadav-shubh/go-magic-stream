import { useState } from 'react';
import axiosConfig from "../api/axiosConfig";
import Navbar from '../components/common/Navbar';
import Footer from '../components/common/Footer';
import './Contact.css';

const Contact = () => {
    const [formData, setFormData] = useState({
        name: '',
        email: '',
        phone: '',
        query: ''
    });

    const [errors, setErrors] = useState({});
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [submitSuccess, setSubmitSuccess] = useState(false);

    const validate = () => {
        const newErrors = {};

        // Name is mandatory
        if (!formData.name.trim()) {
            newErrors.name = 'Name is required';
        }

        // Either email or phone is mandatory
        if (!formData.email.trim() && !formData.phone.trim()) {
            newErrors.contact = 'Please provide either Email or Phone number';
        }

        // Basic Email validation if provided
        if (formData.email && !/\S+@\S+\.\S+/.test(formData.email)) {
            newErrors.email = 'Please enter a valid email address';
        }

        setErrors(newErrors);
        return Object.keys(newErrors).length === 0;
    };

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData(prev => ({
            ...prev,
            [name]: value
        }));
        // Clear errors when user types
        if (errors[name]) {
            setErrors(prev => ({ ...prev, [name]: null }));
        }
        if (errors.contact && (name === 'email' || name === 'phone')) {
            setErrors(prev => ({ ...prev, contact: null }));
        }
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        if (!validate()) return;

        setIsSubmitting(true);

        try {
            await axiosConfig.post('/contacts', formData);
            // On success, show success message and clear form
            setSubmitSuccess(true);
            setFormData({ name: '', email: '', phone: '', query: '' });
            setTimeout(() => setSubmitSuccess(false), 5000); // Hide success message after 5s
        } catch (error) {
            console.error("Error sending contact message", error);
            // Optionally set a global error state here if UI supports it, 
            // but for now we just log it and stop loading.
            setErrors(prev => ({ ...prev, contact: "Failed to send message. Please try again later." }));
        } finally {
            setIsSubmitting(false);
        }
    };

    return (
        <div className="contact-page">
            <Navbar />

            <div className="contact-container">
                <div className="contact-card">
                    <div className="contact-header">
                        <h1>Get in Touch</h1>
                        <p>We'd love to hear from you. Send us a message and we'll respond as soon as possible.</p>
                    </div>

                    {submitSuccess ? (
                        <div className="success-message">
                            <div className="success-icon">✓</div>
                            <h3>Message Sent!</h3>
                            <p>Thank you for reaching out. We will get back to you shortly.</p>
                            <button onClick={() => setSubmitSuccess(false)} className="submit-btn">Send Another Message</button>
                        </div>
                    ) : (
                        <form onSubmit={handleSubmit} className="contact-form">
                            <div className="form-group">
                                <label htmlFor="name">Name <span className="required">*</span></label>
                                <input
                                    type="text"
                                    id="name"
                                    name="name"
                                    value={formData.name}
                                    onChange={handleChange}
                                    placeholder="Your Name"
                                    className={errors.name ? 'error-input' : ''}
                                />
                                {errors.name && <span className="error-text">{errors.name}</span>}
                            </div>

                            <div className="form-row">
                                <div className="form-group">
                                    <label htmlFor="email">Email</label>
                                    <input
                                        type="email"
                                        id="email"
                                        name="email"
                                        value={formData.email}
                                        onChange={handleChange}
                                        placeholder="john@example.com"
                                        className={errors.email || errors.contact ? 'error-input' : ''}
                                    />
                                    {errors.email && <span className="error-text">{errors.email}</span>}
                                </div>

                                <div className="form-group">
                                    <label htmlFor="phone">Phone</label>
                                    <input
                                        type="tel"
                                        id="phone"
                                        name="phone"
                                        value={formData.phone}
                                        onChange={handleChange}
                                        placeholder="+1 (555) 000-0000"
                                        className={errors.contact ? 'error-input' : ''}
                                    />
                                </div>
                            </div>

                            {errors.contact && <div className="error-text global-error">{errors.contact}</div>}

                            <div className="form-group">
                                <label htmlFor="query">Message</label>
                                <textarea
                                    id="query"
                                    name="query"
                                    value={formData.query}
                                    onChange={handleChange}
                                    placeholder="How can we help you?"
                                    rows="4"
                                ></textarea>
                            </div>

                            <button type="submit" className="submit-btn" disabled={isSubmitting}>
                                {isSubmitting ? 'Sending...' : 'Send Message'}
                            </button>
                        </form>
                    )}
                </div>
            </div>

            <Footer />
        </div>
    );
};

export default Contact;
