import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
    resolve: {
        alias: {
            '@': '/src',
            '@wails': '/wailsjs',
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
