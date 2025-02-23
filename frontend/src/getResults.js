async function searchQuery(query) {
    try {
        const response = await fetch("http://localhost:3000/search", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ query }),
        });

        if (!response.ok) {
            throw new Error("Failed to fetch results");
        }

        const data = await response.json();
        console.log("Search Results:", data.results);
        
        return data.results; 
    } catch (error) {
        console.error("Error:", error);
        return [];
    }
}

