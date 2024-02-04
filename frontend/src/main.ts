import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import '@mdi/font/css/materialdesignicons.css'

import 'vuetify/styles'
import {createVuetify} from 'vuetify'
import {aliases, mdi} from 'vuetify/iconsets/mdi'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import {createPinia, PiniaVuePlugin} from "pinia";
import ExportView from "./components/ExportView.vue";
import SettingsView from "./components/SettingsView.vue";

import {createRouter, createWebHashHistory} from 'vue-router'

const routes = [
    {
        path: '/export',
        name: 'export',
        component: ExportView
    },
    {
        path: '/',
        redirect: '/export'
    },
    {
        path: '/settings',
        name: 'settings',
        component: SettingsView
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

const pinia = createPinia()

const vuetify = createVuetify({
    components,
    directives,
    theme: {
        defaultTheme: 'light'
    },
    icons: {
        defaultSet: 'mdi',
        aliases,
        sets: {
            mdi,
        },
    },
})

const app = createApp(App)
app.use(router)
app.use(vuetify)
app.use(pinia)
app.use(PiniaVuePlugin)
app.mount('#app')
