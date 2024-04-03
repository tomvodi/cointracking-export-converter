import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
    resolve: {
        alias: {
            '@': path.resolve(__dirname, '/src'),
            '@wails': path.resolve(__dirname, '/wailsjs'),
        }
    },
    plugins: [vue()],
    test: {
        globals: true,
        environment: "jsdom",
        server: {
            deps: {
                inline: ['vuetify']
            }
        }
    }
})
