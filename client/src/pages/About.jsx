import React, { useEffect } from 'react';
import Navbar from '../components/common/Navbar';
import Footer from '../components/common/Footer';
import './About.css';

const features = [
    {
        title: "Unlimited Content",
        description: "Access a massive library of movies, TV shows, and exclusive originals. There's always something new to discover."
    },
    {
        title: "Watch Anywhere",
        description: "Stream seamlessly on your phone, tablet, laptop, or TV. Your content follows you wherever you go."
    },
    {
        title: "Premium Quality",
        description: "Experience your favorite shows in breathtaking 4K Ultra HD with immersive sound."
    }
];

const About = () => {
    useEffect(() => {
        document.title = "About Us - Magic Stream";
    }, []);

    return (
        <div className="about-page bg-light min-vh-100 d-flex flex-column">
            <header>
                <Navbar />
            </header>

            <main className="flex-grow-1">
                {/* Hero Section */}
                <section className="hero-section">
                    <div className="container position-relative">
                        <h1 className="display-3 fw-bold mb-3 text-gradient-gold">Our Mission</h1>
                        <p className="lead text-white-50 mx-auto">
                            At Magic Stream, we believe entertainment should be limitless.<br />We are dedicated to bringing you the world's best stories, anytime, anywhere.
                        </p>
                    </div>
                </section>

                {/* Features Section */}
                <section className="container py-5">
                    <div className="row g-4 text-center">
                        {features.map((feature, index) => (
                            <div className="col-md-4" key={index}>
                                <div className="card h-100 border-0 shadow-sm p-4">
                                    <div className="card-body">
                                        <h3 className="h4 fw-bold mb-3 text-primary">{feature.title}</h3>
                                        <p className="text-secondary">
                                            {feature.description}
                                        </p>
                                    </div>
                                </div>
                            </div>
                        ))}
                    </div>
                </section>

                {/* Team/Story Section */}
                <section className="container py-5 mb-5">
                    <div className="row align-items-center">
                        <div className="col-md-6 mb-4 mb-md-0">
                            <img
                                src="https://images.unsplash.com/photo-1522071820081-009f0129c71c?q=80&w=2070&auto=format&fit=crop"
                                alt="Our Team collaboration"
                                className="img-fluid rounded-3 shadow"
                                width="600"
                                height="400"
                                loading="lazy"
                            />
                        </div>
                        <div className="col-md-6 ps-md-5">
                            <h2 className="fw-bold mb-4 display-5">Who We Are</h2>
                            <p className="lead text-secondary mb-4">
                                We are a team of passionate storytellers and technologists united by a single goal: to redefine how the world consumes entertainment.
                            </p>
                            <p className="text-muted">
                                Founded in 2024, Magic Stream has grown from a small startup to a global platform. We innovate constantly to ensure the best viewing experience for our millions of subscribers.
                            </p>
                        </div>
                    </div>
                </section>

                {/* CTA */}
                <section className="bg-white py-5 text-center border-top">
                    <div className="container">
                        <h2 className="fw-bold mb-3">Ready to Watch?</h2>
                        <p className="text-secondary mb-4">Join Magic Stream today and start your journey.</p>
                    </div>
                </section>
            </main>

            <Footer />
        </div>
    );
};

export default About;
