package app

import (
	"fmt"
	"net/http"
)

func (server *Server) getIndex(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<head>
    <script src="https://unpkg.com/htmx.org@1.9.10"></script>
    <script src="https://cdn.tailwindcss.com"></script>
    <title>Weather</title>
</head>
<body class="bg-gray-100 flex items-center justify-center h-screen">
    <div class="w-full max-w-xs">
        <h1 class="mb-6 text-center text-4xl font-bold text-gray-800">Weather</h1>
        <form hx-get="/weather" hx-trigger="submit" hx-target="#weather" class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
            <div class="mb-4">
                <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" type="text" name="city" placeholder="City" required>
            </div>
            <div class="flex items-center justify-between">
                <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="submit">
                    Get weather
                </button>
            </div>
        </form>
        <div id="weather" class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4"></div>
    </div>
</body>
</html>
`

	_, err := fmt.Fprintf(w, html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
