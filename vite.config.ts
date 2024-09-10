import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
    plugins: [react()],
    server: {
        proxy: {
            // Proxy pour les requêtes vers /api
            "/api": {
                target: "http://localhost:8080", // URL du serveur backend
                changeOrigin: true, // Changer l'origine de l'en-tête de la requête
                rewrite: (path) => path.replace(/^\/api/, ""), // Réécrire le chemin
            },
        },
    },
});
