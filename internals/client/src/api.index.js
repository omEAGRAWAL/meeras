import axios from 'axios';

const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api';

const api = axios.create({
    baseURL: API_URL,
    headers: {
        'Content-Type': 'application/json'
    }
});

// Add authorization header to requests when token is available
api.interceptors.request.use((config) => {
    const token = localStorage.getItem('token');
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

// Authentication
export const signup = (userData) => api.post('/signup', userData);
export const login = (credentials) => api.post('/login', credentials);

// Venues
export const getAllVenues = () => api.get('/getallvenues');
export const registerVenue = (venueData) => api.post('/registervenue', venueData);
export const getVenueDetails = (id) => api.get(`/venue/${id}`);

// Packages
export const addPackage = (venueName, packageData) => api.post(`/package/${venueName}`, packageData);
export const updatePackage = (venueName, packageId, packageData) => api.put(`/updatepackage/${venueName}/${packageId}`, packageData);

// Bookings
export const createBooking = (venueId, packageId, bookingData) => api.post(`/booking/${venueId}/${packageId}`, bookingData);
export const getUserBookings = () => api.get('/user/bookings');

// Image Upload
export const uploadImage = (formData) => {
    return api.post('/upload/image', formData, {
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
};

export default api;