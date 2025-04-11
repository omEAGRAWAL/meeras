import React, { useState } from "react";

const RegisterVenueForm = () => {
    const [form, setForm] = useState({
        name: "",
        location: "",
        rating: "",
        description: "",
        map_url: "",
        packages: [{ name: "", price: "", duration: "", details: "" }],
        image: null,
    });
    const [status, setStatus] = useState("");
    const BASE_URL = "http://localhost:8080"; // replace with actual base url

    const handleChange = (e) => {
        setForm({ ...form, [e.target.name]: e.target.value });
    };

    const handlePackageChange = (index, field, value) => {
        const updatedPackages = [...form.packages];
        updatedPackages[index][field] = value;
        setForm({ ...form, packages: updatedPackages });
    };

    const handleImageChange = (e) => {
        setForm({ ...form, image: e.target.files[0] });
    };

    const addPackage = () => {
        setForm({
            ...form,
            packages: [...form.packages, { name: "", price: "", duration: "", details: "" }],
        });
    };

    const uploadImage = async () => {
        const formData = new FormData();
        formData.append("image", form.image);
        const res = await fetch(`${BASE_URL}/api/upload/image`, {
            method: "POST",
            body: formData,
        });
        const data = await res.json();
        if (!res.ok) throw new Error(data.error || "Image upload failed");
        return data.url;
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setStatus("Uploading image...");
        try {
            const imageUrl = await uploadImage();
            const venueData = { ...form, image_url: imageUrl };
            delete venueData.image;

            const res = await fetch(`${BASE_URL}/api/registervenue`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(venueData),
            });

            const result = await res.json();
            if (!res.ok) throw new Error(result.error || "Venue registration failed");

            setStatus("✅ Venue registered successfully!");
        } catch (err) {
            setStatus("❌ Error: " + err.message);
        }
    };

    return (
        <div className="max-w-3xl mx-auto p-6 bg-white rounded-xl shadow-md mt-8">
            <h2 className="text-2xl font-bold mb-4">Register a New Venue</h2>
            <form onSubmit={handleSubmit} className="space-y-4">
                <input type="text" name="name" placeholder="Venue Name" className="input" value={form.name} onChange={handleChange} required />
                <input type="text" name="location" placeholder="Location" className="input" value={form.location} onChange={handleChange} required />
                <input type="text" name="rating" placeholder="Rating (e.g. 4.7)" className="input" value={form.rating} onChange={handleChange} required />
                <textarea name="description" placeholder="Description" className="input" value={form.description} onChange={handleChange} required></textarea>
                <input type="text" name="map_url" placeholder="Map URL" className="input" value={form.map_url} onChange={handleChange} required />

                <label className="block font-medium">Upload Venue Image</label>
                <input type="file" accept="image/*" onChange={handleImageChange} required />

                <h3 className="text-lg font-semibold mt-4">Packages</h3>
                {form.packages.map((pkg, idx) => (
                    <div key={idx} className="border p-3 rounded-md mb-2">
                        <input type="text" placeholder="Package Name" className="input" value={pkg.name} onChange={(e) => handlePackageChange(idx, "name", e.target.value)} required />
                        <input type="number" placeholder="Price" className="input" value={pkg.price} onChange={(e) => handlePackageChange(idx, "price", e.target.value)} required />
                        <input type="text" placeholder="Duration" className="input" value={pkg.duration} onChange={(e) => handlePackageChange(idx, "duration", e.target.value)} required />
                        <input type="text" placeholder="Details" className="input" value={pkg.details} onChange={(e) => handlePackageChange(idx, "details", e.target.value)} required />
                    </div>
                ))}

                <button type="button" onClick={addPackage} className="px-3 py-1 bg-blue-500 text-white rounded">Add Package</button>

                <button type="submit" className="px-4 py-2 bg-green-600 text-white rounded mt-4">Register Venue</button>
                <p className="mt-2 text-sm text-gray-600">{status}</p>
            </form>
        </div>
    );
};

export default RegisterVenueForm;

// TailwindCSS: Add this in your global styles or component
// .input { @apply w-full p-2 border rounded mt-2; }
