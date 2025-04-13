import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { AuthProvider } from './contexts/AuthContext';
import ProtectedRoute from './components/ProtectedRoute';
import Navbar from './components/Navbar';
import HomePage from './pages/HomePage';
import LoginPage from './pages/LoginPage';
import SignupPage from './pages/SignupPage';
import VenueListPage from './pages/VenueListPage';
import VenueDetailPage from './pages/VenueDetailPage';
import BookingPage from './pages/BookingPage';
import DashboardPage from './pages/DashboardPage';
import RegisterVenuePage from './pages/RegisterVenuePage';
import AddPackagePage from './pages/AddPackagePage';
import './App.css';

function App() {
  return (
      <AuthProvider>
        <BrowserRouter>
          <div className="flex flex-col min-h-screen">
            <Navbar />
            <main className="flex-grow">
              <Routes>
                <Route path="/" element={<HomePage />} />
                <Route path="/login" element={<LoginPage />} />
                <Route path="/signup" element={<SignupPage />} />
                <Route path="/venues" element={<VenueListPage />} />
                <Route path="/venues/:id" element={<VenueDetailPage />} />
                <Route path="/booking/:venueId/:packageId" element={
                  <ProtectedRoute>
                    <BookingPage />
                  </ProtectedRoute>
                } />
                <Route path="/dashboard" element={
                  <ProtectedRoute>
                    <DashboardPage />
                  </ProtectedRoute>
                } />
                <Route path="/register-venue" element={
                  <ProtectedRoute>
                    <RegisterVenuePage />
                  </ProtectedRoute>
                } />
                <Route path="/add-package/:venueName" element={
                  <ProtectedRoute>
                    <AddPackagePage />
                  </ProtectedRoute>
                } />
              </Routes>
            </main>
            <footer className="bg-gray-800 text-white p-4 text-center">
              <p>© 2025 Meeras Venue Booking. All rights reserved.</p>
            </footer>
          </div>
        </BrowserRouter>
      </AuthProvider>
  );
}

export default App;