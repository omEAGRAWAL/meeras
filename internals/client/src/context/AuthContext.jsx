import React from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';

const ProtectedRoute = ({ children }) => {
    const { currentUser } = useAuth();

    if (!currentUser) {
        return <Navigate to="/login" />;
    }

    return children;
};

export default ProtectedRoute;

// File: src/components/Navbar.jsx
import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';

const Navbar = () => {
    const { currentUser, logout, isVenueManager } = useAuth();
    const navigate = useNavigate();
    const [isMenuOpen, setIsMenuOpen] = useState(false);

    const handleLogout = () => {
        logout();
        navigate('/login');
    };

    return (
        <nav className="bg-indigo-600 text-white shadow-lg">
            <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <div className="flex justify-between h-16">
                    <div className="flex items-center">
                        <Link to="/" className="flex-shrink-0 flex items-center">
                            <span className="text-xl font-bold">Meeras</span>
                        </Link>
                        <div className="hidden md:ml-6 md:flex md:space-x-8">
                            <Link to="/" className="px-3 py-2 rounded-md text-sm font-medium hover:bg-indigo-500">
                                Home
                            </Link>
                            <Link to="/venues" className="px-3 py-2 rounded-md text-sm font-medium hover:bg-indigo-500">
                                Venues
                            </Link>
                            {isVenueManager && (
                                <Link to="/register-venue" className="px-3 py-2 rounded-md text-sm font-medium hover:bg-indigo-500">
                                    Register Venue
                                </Link>
                            )}
                        </div>
                    </div>
                    <div className="hidden md:ml-6 md:flex md:items-center">
                        {currentUser ? (
                            <div className="flex items-center space-x-4">
                                <Link to="/dashboard" className="px-3 py-2 rounded-md text-sm font-medium hover:bg-indigo-500">
                                    Dashboard
                                </Link>
                                <button
                                    onClick={handleLogout}
                                    className="px-3 py-2 rounded-md text-sm font-medium bg-indigo-700 hover:bg-indigo-800"
                                >
                                    Logout
                                </button>
                                <div className="flex items-center ml-3">
                                    <div className="text-sm font-medium">
                                        {currentUser.name}
                                    </div>
                                </div>
                            </div>
                        ) : (
                            <div className="flex items-center space-x-4">
                                <Link to="/login" className="px-3 py-2 rounded-md text-sm font-medium hover:bg-indigo-500">
                                    Login
                                </Link>
                                <Link to="/signup" className="px-3 py-2 rounded-md text-sm font-medium bg-indigo-700 hover:bg-indigo-800">
                                    Sign Up
                                </Link>
                            </div>
                        )}
                    </div>
                    <div className="flex items-center md:hidden">
                        <button
                            onClick={() => setIsMenuOpen(!isMenuOpen)}
                            className="inline-flex items-center justify-center p-2 rounded-md text-white hover:bg-indigo-500"
                        >
                            <svg
                                className="h-6 w-6"
                                xmlns="http://www.w3.org/2000/svg"
                                fill="none"
                                viewBox="0 0 24 24"
                                stroke="currentColor"
                            >
                                {isMenuOpen ? (
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                                ) : (
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16M4 18h16" />
                                )}
                            </svg>
                        </button>
                    </div>
                </div>
            </div>
            {isMenuOpen && (
                <div className="md:hidden">
                    <div className="px-2 pt-2 pb-3 space-y-1 sm:px-3">
                        <Link
                            to="/"
                            className="block px-3 py-2 rounded-md text-base font-medium hover:bg-indigo-500"
                            onClick={() => setIsMenuOpen(false)}
                        >
                            Home
                        </Link>
                        <Link
                            to="/venues"
                            className="block px-3 py-2 rounded-md text-base font-medium hover:bg-indigo-500"
                            onClick={() => setIsMenuOpen(false)}
                        >
                            Venues
                        </Link>
                        {isVenueManager && (
                            <Link
                                to="/register-venue"
                                className="block px-3 py-2 rounded-md text-base font-medium hover:bg-indigo-500"
                                onClick={() => setIsMenuOpen(false)}
                            >
                                Register Venue
                            </Link>
                        )}
                        {currentUser ? (
                            <>
                                <Link
                                    to="/dashboard"
                                    className="block px-3 py-2 rounded-md text-base font-medium hover:bg-indigo-500"
                                    onClick={() => setIsMenuOpen(false)}
                                >
                                    Dashboard
                                </Link>
                                <button
                                    onClick={() => {
                                        handleLogout();
                                        setIsMenuOpen(false);
                                    }}
                                    className="block w-full text-left px-3 py-2 rounded-md text-base font-medium bg-indigo-700 hover:bg-indigo-800"
                                >
                                    Logout
                                </button>
                            </>
                        ) : (
                            <>
                                <Link
                                    to="/login"
                                    className="block px-3 py-2 rounded-md text-base font-medium hover:bg-indigo-500"
                                    onClick={() => setIsMenuOpen(false)}
                                >
                                    Login
                                </Link>
                                <Link
                                    to="/signup"
                                    className="block px-3 py-2 rounded-md text-base font-medium bg-indigo-700 hover:bg-indigo-800"
                                    onClick={() => setIsMenuOpen(false)}
                                >
                                    Sign Up
                                </Link>
                            </>
                        )}
                    </div>
                </div>
            )}
        </nav>
    );
};

export default Navbar;
